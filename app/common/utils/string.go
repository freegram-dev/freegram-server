package utils

import (
	"encoding/json"
	"strconv"
)

func Any2StringVector(v any) []string {
	{
		if l, ok := v.([]string); ok {
			return l
		}
	}
	{
		if l, ok := v.([][]byte); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = string(v)
			}
			return r
		}
	}
	{
		if l, ok := v.([][]rune); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = string(v)
			}
			return r
		}
	}
	{
		if l, ok := v.([]int); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = strconv.FormatInt(int64(v), 10)
			}
			return r
		}
	}
	{
		if l, ok := v.([]int8); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = strconv.FormatInt(int64(v), 10)
			}
			return r
		}
	}
	{
		if l, ok := v.([]int16); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = strconv.FormatInt(int64(v), 10)
			}
			return r
		}
	}
	{
		if l, ok := v.([]int32); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = strconv.FormatInt(int64(v), 10)
			}
			return r
		}
	}
	{
		if l, ok := v.([]int64); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = strconv.FormatInt(v, 10)
			}
			return r
		}
	}
	{
		if l, ok := v.([]uint); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = strconv.FormatUint(uint64(v), 10)
			}
			return r
		}
	}
	{
		if l, ok := v.([]uint8); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = strconv.FormatUint(uint64(v), 10)
			}
			return r
		}
	}
	{
		if l, ok := v.([]uint16); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = strconv.FormatUint(uint64(v), 10)
			}
			return r
		}
	}
	{
		if l, ok := v.([]uint32); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = strconv.FormatUint(uint64(v), 10)
			}
			return r
		}
	}
	{
		if l, ok := v.([]uint64); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = strconv.FormatUint(v, 10)
			}
			return r
		}
	}
	{
		if l, ok := v.([]float32); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = strconv.FormatFloat(float64(v), 'f', -1, 32)
			}
			return r
		}
	}
	{
		if l, ok := v.([]float64); ok {
			r := make([]string, len(l))
			for i, v := range l {
				r[i] = strconv.FormatFloat(v, 'f', -1, 64)
			}
			return r
		}
	}
	{
		l := v.([]any)
		r := make([]string, len(l))
		for i, v := range l {
			r[i] = Any2String(v)
		}
	}
	panic("Any2StringVector: unknown type")
}

func Any2String(i any) string {
	switch v := i.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case []rune:
		return string(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		b, _ := json.Marshal(v)
		return string(b)
	}
}

func Any2Bytes(i any) []byte {
	switch v := i.(type) {
	case string:
		return []byte(v)
	case []byte:
		return v
	}
	panic("Any2Any2Bytes: unknown type")
}

func Any2BytesVector(i any) [][]byte {
	{
		if l, ok := i.([][]byte); ok {
			return l
		}
	}
	{
		if l, ok := i.([]string); ok {
			r := make([][]byte, len(l))
			for i, v := range l {
				r[i] = []byte(v)
			}
			return r
		}
	}
	panic("Any2BytesVector: unknown type")
}
