package tlmodel

import "fmt"

var (
	ErrLayerNotFound         = fmt.Errorf("layer not found")
	ErrCrc32HexNotFound      = fmt.Errorf("crc32Hex not found")
	ErrTypeNameNotFound      = fmt.Errorf("typeName not found")
	ErrInputTypeNotTLSubType = fmt.Errorf("input type not TLSubType")
	ErrInputTypeNil          = fmt.Errorf("input type is nil")

	DecodeErrDataTooShort      = fmt.Errorf("decode: data too short")
	DecodeErrCrc32HexNotMatch  = fmt.Errorf("decode: crc32Hex not match")
	DecodeErrUnknownVectorSize = fmt.Errorf("decode: unknown vector size")
)
