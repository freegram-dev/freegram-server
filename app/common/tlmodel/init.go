package tlmodel

import (
	"github.com/freegram-dev/freegram-server/app/common/embed"
	"github.com/freegram-dev/freegram-server/app/common/utils"
	"github.com/freegram-dev/freegram-server/app/common/xlog"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	entries, err := embed.TLFiles.ReadDir("tlfiles")
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		filename := entry.Name()
		file, err := embed.TLFiles.ReadFile("tlfiles/" + filename)
		if err != nil {
			panic(err)
		}
		parseTLFile(filename, file)
	}
}

func parseTLFile(filename string, file []byte) {
	// filename: layer158.tl
	layer := int32(0)
	// 提取出layer
	pattern := `layer(\d+?)\.tl`
	reg := regexp.MustCompile(pattern)
	match := reg.FindStringSubmatch(filename)
	if len(match) > 1 {
		layerTmp, _ := strconv.Atoi(match[1])
		layer = int32(layerTmp)
	}
	//xlog.Debugf("parseTLFile: %s, layer: %d", filename, layer)
	// 提取出所有的type
	// type是---functions---之前的数据
	content := string(file)
	split := strings.Split(content, "---functions---")
	if len(split) < 2 {
		xlog.Fatalf("parseTLFile: %s, split len < 2", filename)
	}
	typeContent := split[0]
	//xlog.Debugf("parseTLFile: %s, typeContent: %s", filename, typeContent)
	// 逐行解析
	lines := strings.Split(typeContent, "\n")
	for _, line := range lines {
		// 如果是//开头的，跳过
		if strings.HasPrefix(line, "//") {
			continue
		}
		// 如果是空行，跳过
		if strings.TrimSpace(line) == "" {
			continue
		}
		_, ok := SubTypeMap[layer]
		if !ok {
			SubTypeMap[layer] = make(map[string]*SubType)
		}
		subType := parseLine(layer, line)
		SubTypeMap[layer][subType.Crc32Hex] = subType
	}
	fixTypeMap()
	functionContent := split[1]
	//xlog.Debugf("parseTLFile: %s, functionContent: %s", filename, functionContent)
	// 逐行解析
	lines = strings.Split(functionContent, "\n")
	for _, line := range lines {
		// 如果是//开头的，跳过
		if strings.HasPrefix(line, "//") {
			continue
		}
		// 如果是空行，跳过
		if strings.TrimSpace(line) == "" {
			continue
		}
		_, ok := SubTypeMap[layer]
		if !ok {
			SubTypeMap[layer] = make(map[string]*SubType)
		}
		subType := parseLine(layer, line)
		SubTypeMap[layer][subType.Crc32Hex] = subType
	}
}

func parseLine(layer int32, line string) *SubType {
	//boolFalse#bc799737 = Bool; #前表字类型，#后表crc32值，=后表类型
	//boolTrue#997275b5 = Bool;
	//vector#1cb5c415 {t:Type} # [ t ] = Vector t; {t:Type}表示泛型，[ t ]表示数组
	//error#c4b9f9bb code:int text:string = Error; code:int text:string表示参数
	//inputMediaUploadedPhoto#1e287d04 flags:# spoiler:flags.2?true file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia; flags:#表示flags参数，?表示可选参数
	crc32Hex := utils.Uint32ToHex(CalcCrc32(line))

	if crc32Hex == VectorCrc32Hex {
		// vector#1cb5c415 {t:Type} # [ t ] = Vector t;
		return &SubType{
			SubTypeName: "vector",
			Crc32Hex:    crc32Hex,
			TypeName:    "Vector",
		}
	}
	// 把#和后面的crc32值替换成#crc32Hex
	line = regexp.MustCompile("#[a-f0-9]{1,8}").ReplaceAllString(line, "#"+crc32Hex)
	subTypeName := ""
	// 提取出subTypeName
	pattern := `([a-zA-Z0-9.]+?)#`
	reg := regexp.MustCompile(pattern)
	match := reg.FindStringSubmatch(line)
	if len(match) > 1 {
		subTypeName = match[1]
	} else {
		xlog.Fatalf("subTypeName line: %s, match len < 2", line)
	}
	typeName := ""
	// 提取出typeName
	pattern = `= ([a-zA-Z0-9.<>]+?);`
	reg = regexp.MustCompile(pattern)
	match = reg.FindStringSubmatch(line)
	if len(match) > 1 {
		typeName = match[1]
	} else {
		xlog.Fatalf("typeName line: %s, match len < 2", line)
	}
	paramString := ""
	// 提取出paramString
	pattern = `#[a-f0-9]{1,8} (.*?) =`
	reg = regexp.MustCompile(pattern)
	match = reg.FindStringSubmatch(line)
	if len(match) > 1 {
		paramString = match[1]
	}
	if paramString != "" {
		//xlog.Debugf("parseTLFile: %s, subTypeName: %s, crc32Hex: %s, typeName: %s, paramString: %s", filename, subTypeName, crc32Hex, typeName, paramString)
	}
	// 解析paramString
	params := make([]*TypeParam, 0)
	if paramString != "" {
		paramSplit := strings.Split(paramString, " ")
		for _, param := range paramSplit {
			//xlog.Debugf("parseTLFile: %s, param: %s", filename, param)
			// 解析kv
			kvSplit := strings.Split(param, ":")
			if len(kvSplit) != 2 {
				xlog.Fatalf("param: %s, kvSplit len != 2", param)
			}
			paramName := kvSplit[0]
			paramTypeName := kvSplit[1]
			isFlags := false
			isOptional := false
			var flagName = ""
			var flagIndex = 0
			isVector := false
			if paramTypeName == "#" {
				// 需要计算flags
				isFlags = true
				flagName = paramName
			} else {
				if strings.HasPrefix(paramTypeName, "flags") {
					// optional
					isOptional = true
					// ?后面是真正的类型
					paramTypeNameSplit := strings.Split(paramTypeName, "?")
					paramTypeName = paramTypeNameSplit[1]
					flagName = paramTypeNameSplit[0]
					//第几个索引
					flagIndex, _ = strconv.Atoi(strings.Split(flagName, ".")[1])
					flagName = strings.Split(flagName, ".")[0]
					//xlog.Debugf("paramName: %s, paramTypeName: %s", paramName, paramTypeName)
					//paramTypeName是否是Vector
					if strings.HasPrefix(paramTypeName, "Vector<") {
						isVector = true
						paramTypeName = strings.TrimPrefix(paramTypeName, "Vector<")
						paramTypeName = strings.TrimSuffix(paramTypeName, ">")
					}
				} else {
					//paramTypeName是否是Vector
					if strings.HasPrefix(paramTypeName, "Vector<") {
						isVector = true
						paramTypeName = strings.TrimPrefix(paramTypeName, "Vector<")
						paramTypeName = strings.TrimSuffix(paramTypeName, ">")
					}
				}
			}
			params = append(params, &TypeParam{
				ParamName:     paramName,
				ParamTypeName: paramTypeName,
				IsFlags:       isFlags,
				IsVector:      isVector,
				Optional: &Optional{
					IsOptional: isOptional,
					FlagsName:  flagName,
					FlagIndex:  flagIndex,
				},
			})
		}
	}
	return &SubType{
		SubTypeName: subTypeName,
		Crc32Hex:    crc32Hex,
		Params:      params,
		TypeName:    typeName,
	}
}

func fixTypeMap() {
	for layer, typeMap := range SubTypeMap {
		{
			if _, ok := TypeMap[layer]; !ok {
				TypeMap[layer] = make(map[string]*Type)
			}
			subTypesMap := make(map[string][]*SubType)
			for _, typ := range typeMap {
				if _, ok := subTypesMap[typ.TypeName]; !ok {
					subTypesMap[typ.TypeName] = make([]*SubType, 0)
				}
				subTypesMap[typ.TypeName] = append(subTypesMap[typ.TypeName], typ)
			}
			for typeName, subTypes := range subTypesMap {
				TypeMap[layer][typeName] = &Type{
					SubTypes: subTypes,
					TypeName: typeName,
				}
			}
		}
		{
			if _, ok := SubTypeCrc32Map[layer]; !ok {
				SubTypeCrc32Map[layer] = make(map[string]string)
			}
			for _, typ := range typeMap {
				SubTypeCrc32Map[layer][typ.TypeName] = typ.Crc32Hex
			}
		}
	}
}
