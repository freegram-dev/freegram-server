package tlmodel

var SubTypeMap = make(map[int32]map[string]*SubType)    // layer -> crc32Hex -> SubType
var TypeMap = make(map[int32]map[string]*Type)          // layer -> typeName -> Type
var SubTypeCrc32Map = make(map[int32]map[string]string) // layer -> typeName -> crc32Hex

type SubType struct {
	SubTypeName string
	Crc32Hex    string
	Params      []*TypeParam
	TypeName    string
}

type Type struct {
	SubTypes []*SubType
	TypeName string
}

type TypeParam struct {
	ParamName     string
	ParamTypeName string
	Optional      *Optional
	IsFlags       bool
	IsVector      bool
}

type Optional struct {
	IsOptional bool
	FlagsName  string // flags / flags2 / flags3 ...
	FlagIndex  int    // 0 / 1 / 2 ...
}

type TLSubType struct {
	Line string
	Args any
}

type TLObject struct {
	SubTypeName string
	Crc32Hex    string
	Values      map[string]any
}

func (o *TLObject) GetInt64(key string) (int64, bool) {
	v, ok := o.Values[key]
	if !ok {
		return 0, false
	}
	if v == nil {
		return 0, false
	}
	return v.(int64), true
}

func (o *TLObject) GetInt64Slice(key string) []int64 {
	v, ok := o.Values[key]
	if !ok {
		return nil
	}
	return v.([]int64)
}

func (o *TLObject) GetInt32(key string) (int32, bool) {
	v, ok := o.Values[key]
	if !ok {
		return 0, false
	}
	if v == nil {
		return 0, false
	}
	return v.(int32), true
}

func (o *TLObject) GetInt32Slice(key string) []int32 {
	v, ok := o.Values[key]
	if !ok {
		return nil
	}
	return v.([]int32)
}

func (o *TLObject) GetUInt32(key string) (uint32, bool) {
	v, ok := o.Values[key]
	if !ok {
		return 0, false
	}
	if v == nil {
		return 0, false
	}
	return v.(uint32), true
}

func (o *TLObject) GetUInt32Slice(key string) []uint32 {
	v, ok := o.Values[key]
	if !ok {
		return nil
	}
	return v.([]uint32)
}

func (o *TLObject) GetUInt64(key string) (uint64, bool) {
	v, ok := o.Values[key]
	if !ok {
		return 0, false
	}
	if v == nil {
		return 0, false
	}
	return v.(uint64), true
}

func (o *TLObject) GetUInt64Slice(key string) []uint64 {
	v, ok := o.Values[key]
	if !ok {
		return nil
	}
	return v.([]uint64)
}

func (o *TLObject) GetString(key string) (string, bool) {
	v, ok := o.Values[key]
	if !ok {
		return "", false
	}
	if v == nil {
		return "", false
	}
	return v.(string), true
}

func (o *TLObject) GetStringSlice(key string) []string {
	v, ok := o.Values[key]
	if !ok {
		return nil
	}
	return v.([]string)
}

func (o *TLObject) GetBytes(key string) []byte {
	v, ok := o.Values[key]
	if !ok {
		return nil
	}
	return v.([]byte)
}

func (o *TLObject) GetBytesSlice(key string) [][]byte {
	v, ok := o.Values[key]
	if !ok {
		return nil
	}
	return v.([][]byte)
}

func (o *TLObject) GetBool(key string) bool {
	v, ok := o.Values[key]
	if !ok {
		return false
	}
	return v.(bool)
}

func (o *TLObject) GetFloat64(key string) (float64, bool) {
	v, ok := o.Values[key]
	if !ok {
		return 0, false
	}
	if v == nil {
		return 0, false
	}
	return v.(float64), true
}

func (o *TLObject) GetFloat64Slice(key string) []float64 {
	v, ok := o.Values[key]
	if !ok {
		return nil
	}
	return v.([]float64)
}

func (o *TLObject) GetTLObject(key string) (*TLObject, bool) {
	v, ok := o.Values[key]
	if !ok {
		return nil, false
	}
	if v == nil {
		return nil, false
	}
	return v.(*TLObject), true
}

func (o *TLObject) GetTLObjectSlice(key string) []*TLObject {
	v, ok := o.Values[key]
	if !ok {
		return nil
	}
	return v.([]*TLObject)
}
