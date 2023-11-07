package tlmodel

import (
	"encoding/binary"
	"github.com/freegram-dev/freegram-server/app/common/utils"
	"math"
)

type Decoder struct {
	Layer int32
	buf   []byte
	off   int
	size  int
	err   error
}

func NewDecoder(layer int32, buf []byte) *Decoder {
	return &Decoder{
		Layer: layer,
		buf:   buf,
		off:   0,
		size:  len(buf),
		err:   nil,
	}
}

func (d *Decoder) Decode() *TLObject {
	o := &TLObject{
		SubTypeName: "",
		Crc32Hex:    "",
		Values:      map[string]any{},
	}
	if d.err != nil {
		return o
	}
	crc32 := d.Int()
	if d.err != nil {
		return o
	}
	// find subType
	sm, ok := SubTypeMap[d.Layer]
	if !ok {
		d.err = ErrLayerNotFound
		return o
	}
	subType, ok := sm[utils.Uint32ToHex(uint32(crc32))]
	if !ok {
		d.err = ErrCrc32HexNotFound
		return o
	}
	o.SubTypeName = subType.SubTypeName
	o.Crc32Hex = subType.Crc32Hex
	flagMap := map[string]uint32{}
	// decode params
	for i, param := range subType.Params {
		i := i
		_ = i
		param := param
		_ = param
		switch param.ParamTypeName {
		case "#":
			flagMap[param.Optional.FlagsName] = d.UInt()
			if d.err != nil {
				return o
			}
			o.Values[param.ParamName] = flagMap[param.Optional.FlagsName]
		case "long":
			if param.Optional.IsOptional {
				// 判断是否存在值
				if flagMap[param.Optional.FlagsName]&(1<<param.Optional.FlagIndex) == 0 {
					// 不存在值
					if param.IsVector {
						o.Values[param.ParamName] = make([]int64, 0)
					} else {
						o.Values[param.ParamName] = nil
					}
				} else {
					// 存在值
					// 是否是Vector
					if param.IsVector {
						o.Values[param.ParamName] = d.VectorLong()
					} else {
						o.Values[param.ParamName] = d.Long()
					}
					if d.err != nil {
						return o
					}
				}
			} else {
				// 存在值
				// 是否是Vector
				if param.IsVector {
					o.Values[param.ParamName] = d.VectorLong()
				} else {
					o.Values[param.ParamName] = d.Long()
				}
				if d.err != nil {
					return o
				}
			}
		case "int":
			if param.Optional.IsOptional {
				// 判断是否存在值
				if flagMap[param.Optional.FlagsName]&(1<<param.Optional.FlagIndex) == 0 {
					// 不存在值
					o.Values[param.ParamName] = nil
				} else {
					// 存在值
					// 是否是Vector
					if param.IsVector {
						o.Values[param.ParamName] = d.VectorInt()
					} else {
						o.Values[param.ParamName] = d.Int()
					}
				}
			} else {
				// 存在值
				// 是否是Vector
				if param.IsVector {
					o.Values[param.ParamName] = d.VectorInt()
				} else {
					o.Values[param.ParamName] = d.Int()
				}
			}
		case "double":
			if param.Optional.IsOptional {
				// 判断是否存在值
				if flagMap[param.Optional.FlagsName]&(1<<param.Optional.FlagIndex) == 0 {
					// 不存在值
					if param.IsVector {
						o.Values[param.ParamName] = make([]float64, 0)
					} else {
						o.Values[param.ParamName] = nil
					}
				} else {
					// 存在值
					// 是否是Vector
					if param.IsVector {
						o.Values[param.ParamName] = d.VectorDouble()
					} else {
						o.Values[param.ParamName] = d.Double()
					}
					if d.err != nil {
						return o
					}
				}
			} else {
				// 存在值
				// 是否是Vector
				if param.IsVector {
					o.Values[param.ParamName] = d.VectorDouble()
				} else {
					o.Values[param.ParamName] = d.Double()
				}
				if d.err != nil {
					return o
				}
			}
		case "string":
			if param.Optional.IsOptional {
				// 判断是否存在值
				if flagMap[param.Optional.FlagsName]&(1<<param.Optional.FlagIndex) == 0 {
					// 不存在值
					if param.IsVector {
						o.Values[param.ParamName] = make([]string, 0)
					} else {
						o.Values[param.ParamName] = nil
					}
				} else {
					// 存在值
					// 是否是Vector
					if param.IsVector {
						o.Values[param.ParamName] = d.VectorString()
					} else {
						o.Values[param.ParamName] = d.String()
					}
					if d.err != nil {
						return o
					}
				}
			} else {
				// 存在值
				// 是否是Vector
				if param.IsVector {
					o.Values[param.ParamName] = d.VectorString()
				} else {
					o.Values[param.ParamName] = d.String()
				}
				if d.err != nil {
					return o
				}
			}
		case "bytes":
			if param.Optional.IsOptional {
				// 判断是否存在值
				if flagMap[param.Optional.FlagsName]&(1<<param.Optional.FlagIndex) == 0 {
					// 不存在值
					if param.IsVector {
						o.Values[param.ParamName] = make([][]byte, 0)
					} else {
						o.Values[param.ParamName] = nil
					}
				} else {
					// 存在值
					// 是否是Vector
					if param.IsVector {
						o.Values[param.ParamName] = d.VectorBytes()
					} else {
						o.Values[param.ParamName] = d.String()
					}
					if d.err != nil {
						return o
					}
				}
			} else {
				// 存在值
				// 是否是Vector
				if param.IsVector {
					o.Values[param.ParamName] = d.VectorBytes()
				} else {
					o.Values[param.ParamName] = d.String()
				}
				if d.err != nil {
					return o
				}
			}
		case "true":
			if param.Optional.IsOptional {
				// 判断是否存在值
				if flagMap[param.Optional.FlagsName]&(1<<param.Optional.FlagIndex) == 0 {
					// 不存在值
					o.Values[param.ParamName] = false
				} else {
					// 存在值
					// 是否是Vector
					if param.IsVector {
						panic("true must not be vector")
					} else {
						o.Values[param.ParamName] = true
					}
				}
			} else {
				// 存在值
				// 是否是Vector
				if param.IsVector {
					panic("true must not be vector")
				} else {
					o.Values[param.ParamName] = true
				}
			}
		default:
			if param.Optional.IsOptional {
				// 判断是否存在值
				if flagMap[param.Optional.FlagsName]&(1<<param.Optional.FlagIndex) == 0 {
					// 不存在值
					if param.IsVector {
						o.Values[param.ParamName] = make([]*TLObject, 0)
					} else {
						o.Values[param.ParamName] = nil
					}
				} else {
					// 存在值，递归解析
					// 是否是Vector
					if param.IsVector {
						o.Values[param.ParamName] = d.VectorDecode()
					} else {
						o.Values[param.ParamName] = d.Decode()
					}
					if d.err != nil {
						return o
					}
				}
			} else {
				// 存在值，递归解析
				// 是否是Vector
				if param.IsVector {
					o.Values[param.ParamName] = d.VectorDecode()
				} else {
					o.Values[param.ParamName] = d.Decode()
				}
				if d.err != nil {
					return o
				}
			}
		}
	}
	return o
}

func (d *Decoder) GetOffset() int {
	return d.off
}

func (d *Decoder) GetSize() int {
	return d.size
}

func (d *Decoder) Error() error {
	return d.err
}

func (d *Decoder) SetError(err error) {
	d.err = err
}

func (d *Decoder) Long() int64 {
	if d.err != nil {
		return 0
	}
	if d.off+8 > d.size {
		d.err = DecodeErrDataTooShort
		return 0
	}
	x := int64(binary.LittleEndian.Uint64(d.buf[d.off : d.off+8]))
	d.off += 8
	return x
}

func (d *Decoder) Double() float64 {
	if d.err != nil {
		return 0
	}
	if d.off+8 > d.size {
		d.err = DecodeErrDataTooShort
		return 0
	}
	x := math.Float64frombits(binary.LittleEndian.Uint64(d.buf[d.off : d.off+8]))
	d.off += 8
	return x
}

func (d *Decoder) Int() int32 {
	if d.err != nil {
		return 0
	}
	if d.off+4 > d.size {
		d.err = DecodeErrDataTooShort
		return 0
	}
	x := int32(binary.LittleEndian.Uint32(d.buf[d.off : d.off+4]))
	d.off += 4
	return x
}

func (d *Decoder) UInt() uint32 {
	if d.err != nil {
		return 0
	}
	if d.off+4 > d.size {
		d.err = DecodeErrDataTooShort
		return 0
	}
	x := binary.LittleEndian.Uint32(d.buf[d.off : d.off+4])
	d.off += 4
	return x
}

func (d *Decoder) Bytes(size int) []byte {
	if d.err != nil {
		return nil
	}
	if d.off+size > d.size {
		d.err = DecodeErrDataTooShort
		return nil
	}
	x := d.buf[d.off : d.off+size]
	d.off += size
	return x
}

func (d *Decoder) String() string {
	if d.err != nil {
		return ""
	}
	var (
		size    int
		padding int
	)
	if d.off+1 > d.size {
		d.err = DecodeErrDataTooShort
		return ""
	}
	size = int(d.buf[d.off])
	d.off++
	padding = (4 - (size+1)%4) & 3
	if size == 254 {
		if d.off+3 > d.size {
			d.err = DecodeErrDataTooShort
			return ""
		}
		size = int(d.buf[d.off]) | int(d.buf[d.off+1])<<8 | int(d.buf[d.off+2])<<16
		d.off += 3
		padding = (4 - size%4) & 3
	}
	if d.off+size > d.size {
		d.err = DecodeErrDataTooShort
		return ""
	}
	x := make([]byte, size)
	copy(x, d.buf[d.off:d.off+size])
	d.off += size
	if d.off+padding > d.size {
		d.err = DecodeErrDataTooShort
		return ""
	}
	d.off += padding
	return string(x)
}

func (d *Decoder) VectorInt() []int32 {
	if d.Int() != VectorCrc32 {
		d.err = DecodeErrCrc32HexNotMatch
		return nil
	}
	size := d.Int()
	if d.err != nil {
		return nil
	}
	if size < 0 {
		d.err = DecodeErrUnknownVectorSize
		return nil
	}
	x := make([]int32, size)
	i := int32(0)
	for i < size {
		y := d.Int()
		if d.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (d *Decoder) VectorLong() []int64 {
	if d.Int() != VectorCrc32 {
		d.err = DecodeErrCrc32HexNotMatch
		return nil
	}
	size := d.Int()
	if d.err != nil {
		return nil
	}
	if size < 0 {
		d.err = DecodeErrUnknownVectorSize
		return nil
	}
	x := make([]int64, size)
	i := int32(0)
	for i < size {
		y := d.Long()
		if d.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (d *Decoder) VectorDouble() []float64 {
	if d.Int() != VectorCrc32 {
		d.err = DecodeErrCrc32HexNotMatch
		return nil
	}
	size := d.Int()
	if d.err != nil {
		return nil
	}
	if size < 0 {
		d.err = DecodeErrUnknownVectorSize
		return nil
	}
	x := make([]float64, size)
	i := int32(0)
	for i < size {
		y := d.Double()
		if d.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (d *Decoder) VectorString() []string {
	if d.Int() != VectorCrc32 {
		d.err = DecodeErrCrc32HexNotMatch
		return nil
	}
	size := d.Int()
	if d.err != nil {
		return nil
	}
	if size < 0 {
		d.err = DecodeErrUnknownVectorSize
		return nil
	}
	x := make([]string, size)
	i := int32(0)
	for i < size {
		y := d.String()
		if d.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (d *Decoder) VectorBytes() [][]byte {
	if d.Int() != VectorCrc32 {
		d.err = DecodeErrCrc32HexNotMatch
		return nil
	}
	size := d.Int()
	if d.err != nil {
		return nil
	}
	if size < 0 {
		d.err = DecodeErrUnknownVectorSize
		return nil
	}
	x := make([][]byte, size)
	i := int32(0)
	for i < size {
		y := []byte(d.String())
		if d.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}

func (d *Decoder) Bool() bool {
	if d.err != nil {
		return false
	}
	switch int(d.Int()) {
	case BoolFalseCrc32:
		return false
	case BoolTrueCrc32:
		return true
	}
	d.err = DecodeErrCrc32HexNotMatch
	return false
}

func (d *Decoder) VectorDecode() []*TLObject {
	if d.Int() != VectorCrc32 {
		d.err = DecodeErrCrc32HexNotMatch
		return nil
	}
	size := d.Int()
	if d.err != nil {
		return nil
	}
	if size < 0 {
		d.err = DecodeErrUnknownVectorSize
		return nil
	}
	x := make([]*TLObject, size)
	i := int32(0)
	for i < size {
		y := d.Decode()
		if d.err != nil {
			return nil
		}
		x[i] = y
		i++
	}
	return x
}
