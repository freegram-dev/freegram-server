package tlmodel

import (
	"encoding/json"
	"github.com/freegram-dev/freegram-server/app/common/utils"
	"strconv"
	"strings"
	"testing"
)

func marshal(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func TestTlModel(t *testing.T) {
	//fmt.Printf("%v", marshal(SubTypeMap))
	//fmt.Printf("%v", marshal(TypeMap))
	//fmt.Printf("%v", marshal(SubTypeCrc32Map))
	//line := parseLine(158, "inputMediaUploadedPhoto#1e287d04 flags:# spoiler:flags.2?true file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;")
	//t.Logf("%v", marshal(line))
	if false {
		b, _ := NewEncoder(158).Encode("boolFalse#bc799737 = Bool;", nil)
		t.Logf("%v", b)
	}
	if false {
		b, _ := NewEncoder(158).Encode(
			"inputPeerChat#35a95cb9 chat_id:long = InputPeer;",
			[]any{123456789},
		)
		t.Logf("%v", b)
	}
	if false {
		b, _ := NewEncoder(158).Encode(
			"inputPeerUserFromMessage#a87b0a1c peer:InputPeer msg_id:int user_id:long = InputPeer;",
			[]any{TLSubType{
				Line: "inputPeerChat#35a95cb9 chat_id:long = InputPeer;",
				Args: []any{123456789},
			}, 123456789, 123456789},
		)
		t.Logf("%v", b)
	}
	if false {
		b, _ := NewEncoder(158).Encode(
			"inputMediaUploadedPhoto#1e287d04 flags:# spoiler:flags.2?true file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;",
			[]any{
				nil,
				nil,
				TLSubType{
					Line: "inputFile#f52ff27f id:long parts:int name:string md5_checksum:string = InputFile;",
					Args: []any{123456789, 123456789, "123456789", "123456789"},
				},
				[]TLSubType{TLSubType{
					Line: "inputDocument#1abfb575 id:long access_hash:long file_reference:bytes = InputDocument;",
					Args: []any{123456789, 123456789, []byte("123456789")},
				}},
				nil,
			},
		)
		t.Logf("%v", b)
	}
	if false {
		b, _ := NewEncoder(158).Encode(
			"photos.uploadProfilePhoto#388a3b5 flags:# "+
				"fallback:flags.3?true "+
				"bot:flags.5?InputUser "+
				"file:flags.0?InputFile "+
				"video:flags.1?InputFile "+
				"video_start_ts:flags.2?double "+
				"video_emoji_markup:flags.4?VideoSize = photos.Photo;",
			[]any{
				nil,
				true,
				nil,
				nil,
				nil,
				123456789.12345,
				nil,
			},
		)
		t.Logf("%v", b)
	}
	if true {
		b, _ := NewEncoder(158).Encode(
			"0x388a3b5",
			map[string]any{
				"fallback":       true,
				"video_start_ts": 123456789.12345,
				"bot": TLSubType{
					//Line: "inputUser#f21158c6 user_id:long access_hash:long = InputUser;",
					Line: "0xf21158c6",
					Args: []any{
						123456789,
						123456789123456789,
					},
				},
			},
		)
		t.Logf("%v", b)
		decoder := NewDecoder(158, b)
		o := decoder.Decode()
		e := decoder.Error()
		if e != nil {
			t.Fatalf("decode error: %v", e)
		}
		t.Logf("%v", marshal(o))
		flags, _ := o.GetUInt32("flags")
		t.Logf("flags: %v", flags)
		video_start_ts, _ := o.GetFloat64("video_start_ts")
		t.Logf("video_start_ts: %s", strconv.FormatFloat(video_start_ts, 'f', -1, 64))
		fallback := o.GetBool("fallback")
		t.Logf("fallback: %v", fallback)
		file, _ := o.GetTLObject("file")
		t.Logf("file: %v", marshal(file))
		bot, _ := o.GetTLObject("bot")
		t.Logf("bot: %v", marshal(bot))
	}
	if true {
		//photos.uploadProfilePhoto#93c9a51
		//flags:#
		//fallback:flags.3?true
		//file:flags.0?InputFile
		//video:flags.1?InputFile
		//video_start_ts:flags.2?double
		//video_emoji_markup:flags.4?VideoSize
		//= photos.Photo;
		b, _ := NewEncoder(155).Encode(
			"0x93c9a51",
			map[string]any{
				"fallback":       true,
				"video_start_ts": 123456789.12345,
				"video_emoji_markup": TLSubType{
					//Line: "videoSize#de33b094
					//flags:#
					//type:string
					//w:int
					//h:int
					//size:int
					//video_start_ts:flags.0?double
					//= VideoSize;",
					Line: "0xde33b094",
					Args: map[string]any{
						"type":           "type",
						"w":              123,
						"h":              123,
						"size":           123,
						"video_start_ts": nil,
					},
				},
			},
		)
		t.Logf("%v", b)
		decoder := NewDecoder(155, b)
		o := decoder.Decode()
		e := decoder.Error()
		if e != nil {
			t.Fatalf("decode error: %v", e)
		}
		t.Logf("%v", marshal(o))
		flags, _ := o.GetUInt32("flags")
		t.Logf("flags: %v", flags)
		video_start_ts, _ := o.GetFloat64("video_start_ts")
		t.Logf("video_start_ts: %s", strconv.FormatFloat(video_start_ts, 'f', -1, 64))
		fallback := o.GetBool("fallback")
		t.Logf("fallback: %v", fallback)
		file, _ := o.GetTLObject("file")
		t.Logf("file: %v", marshal(file))
		video_emoji_markup, _ := o.GetTLObject("video_emoji_markup")
		t.Logf("video_emoji_markup: %v", marshal(video_emoji_markup))
	}
}

func TestCalcCrc32(t *testing.T) {
	lineString := `
boolFalse#bc799737 = Bool;
boolTrue#997275b5 = Bool;

true#3fedd339 = True;

vector#1cb5c415 {t:Type} # [ t ] = Vector t;

error#c4b9f9bb code:int text:string = Error;

null#56730bcc = Null;

inputPeerEmpty#7f3b18ea = InputPeer;
inputPeerSelf#7da07ec9 = InputPeer;
inputPeerChat#35a95cb9 chat_id:long = InputPeer;
inputPeerUser#dde8a54c user_id:long access_hash:long = InputPeer;
inputPeerChannel#27bcbbfc channel_id:long access_hash:long = InputPeer;
inputPeerUserFromMessage#a87b0a1c peer:InputPeer msg_id:int user_id:long = InputPeer;
inputPeerChannelFromMessage#bd2a0840 peer:InputPeer msg_id:int channel_id:long = InputPeer;

inputMediaUploadedPhoto#1e287d04 flags:# spoiler:flags.2?true file:InputFile stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
inputMediaPhoto#b3ba0635 flags:# spoiler:flags.1?true id:InputPhoto ttl_seconds:flags.0?int = InputMedia;
photos.uploadProfilePhoto#388a3b5 flags:# fallback:flags.3?true bot:flags.5?InputUser file:flags.0?InputFile video:flags.1?InputFile video_start_ts:flags.2?double video_emoji_markup:flags.4?VideoSize = photos.Photo;
`
	for _, line := range strings.Split(lineString, "\n") {
		if strings.HasPrefix(line, "//") {
			continue
		}
		if strings.TrimSpace(line) == "" {
			continue
		}
		crc32 := CalcCrc32(line)
		hex := utils.Uint32ToHex(crc32)
		t.Logf("hex: %s, line: %s, crc32: %d", hex, line, crc32)
	}
}
