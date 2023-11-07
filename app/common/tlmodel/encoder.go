package tlmodel

import (
	"encoding/binary"
	"github.com/freegram-dev/freegram-server/app/common/utils"
	"github.com/freegram-dev/freegram-server/app/common/xlog"
	"math"
	"strings"
)

type Encoder struct {
	buf   []byte
	Layer int32
}

func NewEncoder(layer int32) *Encoder {
	return &Encoder{buf: make([]byte, 0), Layer: layer}
}

func (e *Encoder) UInt(v uint32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], v)
}

func (e *Encoder) Long(v int64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], uint64(v))
}

func (e *Encoder) Int(v int32) {
	e.buf = append(e.buf, 0, 0, 0, 0)
	binary.LittleEndian.PutUint32(e.buf[len(e.buf)-4:], uint32(v))
}

func (e *Encoder) Double(v float64) {
	e.buf = append(e.buf, 0, 0, 0, 0, 0, 0, 0, 0)
	binary.LittleEndian.PutUint64(e.buf[len(e.buf)-8:], math.Float64bits(v))
}

func (e *Encoder) String(v string) {
	var res []byte
	s := []byte(v)
	size := len(s)
	if size < 254 {
		nl := 1 + size + (4-(size+1)%4)&3
		res = make([]byte, nl)
		res[0] = byte(size)
		copy(res[1:], s)
	} else {
		nl := 4 + size + (4-size%4)&3
		res = make([]byte, nl)
		binary.LittleEndian.PutUint32(res, uint32(size<<8|254))
		copy(res[4:], s)
	}
	e.buf = append(e.buf, res...)
}

func (e *Encoder) Bytes(s []byte) {
	var res []byte
	size := len(s)
	if size < 254 {
		nl := 1 + size + (4-(size+1)%4)&3
		res = make([]byte, nl)
		res[0] = byte(size)
		copy(res[1:], s)
	} else {
		nl := 4 + size + (4-size%4)&3
		res = make([]byte, nl)
		binary.LittleEndian.PutUint32(res, uint32(size<<8|254))
		copy(res[4:], s)
	}
	e.buf = append(e.buf, res...)
}

func (e *Encoder) VectorFlag(length uint32) {
	// 写入VectorCrc32Hex
	e.UInt(VectorCrc32)
	// 写入长度
	e.UInt(length)
}

func (e *Encoder) LongVector(v []int64) {
	e.VectorFlag(uint32(len(v)))
	for _, i := range v {
		e.Long(i)
	}
}

func (e *Encoder) VectorDouble(v []float64) {
	e.VectorFlag(uint32(len(v)))
	for _, i := range v {
		e.Double(i)
	}
}

func (e *Encoder) IntVector(v []int32) {
	e.VectorFlag(uint32(len(v)))
	for _, i := range v {
		e.Int(i)
	}
}

func (e *Encoder) StringVector(v []string) {
	e.VectorFlag(uint32(len(v)))
	for _, i := range v {
		e.String(i)
	}
}

func (e *Encoder) BytesVector(v [][]byte) {
	e.VectorFlag(uint32(len(v)))
	for _, i := range v {
		e.Bytes(i)
	}
}

// Encode 编码
// layer: 层级
// line: 一行，或crc32Hex
// Values: 参数，可以是[]any，也可以是map[string]any
func (e *Encoder) Encode(
	line string,
	values any,
) ([]byte, error) {
	var layer = e.Layer
	var subType *SubType
	// 判断line是不是crc32Hex
	// 如果以0x开头，就是crc32Hex
	if strings.HasPrefix(line, "0x") {
		h := strings.TrimPrefix(line, "0x")
		sm, ok := SubTypeMap[layer]
		if !ok {
			xlog.Errorf("Encode: layer %d not found", layer)
			return nil, ErrLayerNotFound
		}
		subType, ok = sm[h]
		if !ok {
			xlog.Errorf("Encode: layer %d, crc32Hex %s not found", layer, h)
			return nil, ErrCrc32HexNotFound
		}
	} else {
		subType = parseLine(layer, line)
	}
	var args []any
	switch values.(type) {
	case []any:
		args = values.([]any)
	case map[string]any:
		args = make([]any, len(subType.Params), len(subType.Params))
		m := values.(map[string]any)
		for i, param := range subType.Params {
			var ok bool
			args[i], ok = m[param.ParamName]
			if !ok {
				args[i] = nil
			}
		}
	}
	crc32Value := utils.HexToUint32(subType.Crc32Hex)
	var writeFunctions = make([]func() error, len(subType.Params)+1, len(subType.Params)+1)
	writeFunctions[0] = func() error {
		e.UInt(crc32Value)
		return nil
	}
	for i, param := range subType.Params {
		// go的bug，闭包的坑，需要copy一份
		i := i
		param := param
		val := args[i]
		switch param.ParamTypeName {
		case "long":
			writeFunctions[i+1] = func() error {
				if val == nil {
					// 如果是optional，return nil，否则报错
					if param.Optional.IsOptional {
						return nil
					} else if param.IsVector {
						writeFunctions[i+1] = func() error {
							e.VectorFlag(0)
							return nil
						}
					} else {
						return ErrInputTypeNil
					}
				}
				// 是不是Vector
				if param.IsVector {
					e.LongVector(utils.Any2Int64Vector(val))
				} else {
					e.Long(utils.Any2Int64(val))
				}
				return nil
			}
		case "int":
			writeFunctions[i+1] = func() error {
				if val == nil {
					// 如果是optional，return nil，否则报错
					if param.Optional.IsOptional {
						return nil
					} else if param.IsVector {
						writeFunctions[i+1] = func() error {
							e.VectorFlag(0)
							return nil
						}
					} else {
						return ErrInputTypeNil
					}
				}
				// 是不是Vector
				if param.IsVector {
					e.IntVector(utils.Any2Int32Vector(val))
				} else {
					e.Int(int32(utils.Any2Int64(val)))
				}
				return nil
			}
		case "double":
			writeFunctions[i+1] = func() error {
				if val == nil {
					// 如果是optional，return nil，否则报错
					if param.Optional.IsOptional {
						return nil
					} else if param.IsVector {
						writeFunctions[i+1] = func() error {
							e.VectorFlag(0)
							return nil
						}
					} else {
						return ErrInputTypeNil
					}
				}
				// 是不是Vector
				if param.IsVector {
					e.VectorDouble(utils.Any2Float64Vector(val))
				} else {
					d := utils.Any2Float64(val)
					e.Double(d)
				}
				return nil
			}
		case "true":
			writeFunctions[i+1] = func() error {
				if param.IsVector {
					panic("true must not be vector")
				}
				return nil
			}
		case "#":
			// 要计算flag
			writeFunctions[i+1] = func() error {
				flagName := param.Optional.FlagsName
				flagValue := uint32(0)
				for j, p := range subType.Params {
					if p.Optional.FlagsName == flagName && p.Optional.IsOptional {
						// 判断值是否存在
						if args[j] != nil {
							flagValue |= 1 << p.Optional.FlagIndex
						}
					}
				}
				e.UInt(flagValue)
				return nil
			}
		case "string":
			writeFunctions[i+1] = func() error {
				if val == nil {
					// 如果是optional，return nil，否则报错
					if param.Optional.IsOptional {
						return nil
					} else if param.IsVector {
						writeFunctions[i+1] = func() error {
							e.VectorFlag(0)
							return nil
						}
					} else {
						return ErrInputTypeNil
					}
				}
				// 是不是Vector
				if param.IsVector {
					e.StringVector(utils.Any2StringVector(val))
				} else {
					e.String(utils.Any2String(val))
				}
				return nil
			}
		case "bytes":
			writeFunctions[i+1] = func() error {
				if val == nil {
					// 如果是optional，return nil，否则报错
					if param.Optional.IsOptional {
						return nil
					} else if param.IsVector {
						writeFunctions[i+1] = func() error {
							e.VectorFlag(0)
							return nil
						}
					} else {
						return ErrInputTypeNil
					}
				}
				// 是不是Vector
				if param.IsVector {
					e.BytesVector(utils.Any2BytesVector(val))
				} else {
					e.Bytes(utils.Any2Bytes(val))
				}
				return nil
			}
		default:
			typeMap, ok := TypeMap[layer]
			if !ok {
				xlog.Errorf("Encode: layer %d not found", layer)
				return nil, ErrLayerNotFound
			}
			typeName := param.ParamTypeName
			_, ok = typeMap[typeName]
			if !ok {
				xlog.Errorf("Encode: layer %d, typeName %s not found", layer, typeName)
				return nil, ErrTypeNameNotFound
			}
			if val == nil {
				// 如果是optional，或是vector，return nil，否则报错
				if param.Optional.IsOptional {
					writeFunctions[i+1] = func() error {
						return nil
					}
				} else if param.IsVector {
					writeFunctions[i+1] = func() error {
						e.VectorFlag(0)
						return nil
					}
				} else {
					return nil, ErrInputTypeNil
				}
			} else {
				// 是不是Vector
				if param.IsVector {
					tlSubTypes, ok := val.([]TLSubType)
					if !ok {
						xlog.Errorf("Encode: layer %d, typeName %s, args[%d] not []TLSubType, value: %v", layer, typeName, i, val)
						return nil, ErrInputTypeNotTLSubType
					}
					writeFunctions[i+1] = func() error {
						e.VectorFlag(uint32(len(tlSubTypes)))
						for _, tlSubType := range tlSubTypes {
							_, err := e.Encode(tlSubType.Line, tlSubType.Args)
							if err != nil {
								return err
							}
						}
						return nil
					}
				} else {
					tlSubType, ok := val.(TLSubType)
					if !ok {
						xlog.Errorf("Encode: layer %d, typeName %s, args[%d] not TLSubType, value: %v, value == nil:%v", layer, typeName, i, val, val == nil)
						return nil, ErrInputTypeNotTLSubType
					}
					writeFunctions[i+1] = func() error {
						_, err := e.Encode(tlSubType.Line, tlSubType.Args)
						if err != nil {
							return err
						}
						return nil
					}
				}
			}
		}
	}
	for _, writeFunction := range writeFunctions {
		err := writeFunction()
		if err != nil {
			return nil, err
		}
	}
	return e.buf, nil
}
