package tlmodel

import (
	"hash/crc32"
	"regexp"
	"strings"
)

func CalcCrc32(line string) uint32 {
	if strings.Contains(line, "1cb5c415") {
		// vector#1cb5c415 {t:Type} # [ t ] = Vector t;
		return 0x1cb5c415
	}
	line = regexp.MustCompile("[a-zA-Z0-9_]+\\:flags\\.[0-9]+\\?true").ReplaceAllString(line, "")
	line = strings.ReplaceAll(line, "<", " ")
	line = strings.ReplaceAll(line, ">", " ")
	line = strings.ReplaceAll(line, "  ", " ")
	line = regexp.MustCompile("^ ").ReplaceAllString(line, "")
	line = regexp.MustCompile(" $").ReplaceAllString(line, "")
	line = strings.ReplaceAll(line, ":bytes", ":string")
	line = strings.ReplaceAll(line, "?bytes", "?string")
	line = strings.ReplaceAll(line, "{", "?string")
	line = strings.ReplaceAll(line, "}", "?string")
	line = strings.ReplaceAll(line, ";", "")
	// 删掉#和后面的crc32值
	line = regexp.MustCompile("#[a-f0-9]{1,8}").ReplaceAllString(line, "")
	ieee := crc32.ChecksumIEEE([]byte(line))
	return ieee
}
