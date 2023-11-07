package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func Any2Int64Vector(v any) []int64 {
	{
		l, ok := v.([]int64)
		if ok {
			return l
		}
	}
	{
		l, ok := v.([]int)
		if ok {
			r := make([]int64, len(l))
			for i, v := range l {
				r[i] = int64(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]int32)
		if ok {
			r := make([]int64, len(l))
			for i, v := range l {
				r[i] = int64(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]int16)
		if ok {
			r := make([]int64, len(l))
			for i, v := range l {
				r[i] = int64(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]int8)
		if ok {
			r := make([]int64, len(l))
			for i, v := range l {
				r[i] = int64(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]uint64)
		if ok {
			r := make([]int64, len(l))
			for i, v := range l {
				r[i] = int64(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]uint)
		if ok {
			r := make([]int64, len(l))
			for i, v := range l {
				r[i] = int64(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]uint32)
		if ok {
			r := make([]int64, len(l))
			for i, v := range l {
				r[i] = int64(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]uint16)
		if ok {
			r := make([]int64, len(l))
			for i, v := range l {
				r[i] = int64(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]uint8)
		if ok {
			r := make([]int64, len(l))
			for i, v := range l {
				r[i] = int64(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]float64)
		if ok {
			r := make([]int64, len(l))
			for i, v := range l {
				r[i] = int64(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]float32)
		if ok {
			r := make([]int64, len(l))
			for i, v := range l {
				r[i] = int64(v)
			}
			return r
		}
	}
	{
		l := v.([]any)
		r := make([]int64, len(l))
		for i, v := range l {
			r[i] = Any2Int64(v)
		}
	}
	{
		l := v.([]string)
		r := make([]int64, len(l))
		for i, v := range l {
			if strings.HasPrefix(v, "0x") {
				r[i] = HexToInt64(v)
			} else {
				v, _ := strconv.ParseInt(v, 10, 64)
				r[i] = v
			}
		}
	}
	panic("any2Int64Vector: unknown type")
}

func Any2Int64(v any) int64 {
	switch v.(type) {
	case int64:
		return v.(int64)
	case int:
		return int64(v.(int))
	case int32:
		return int64(v.(int32))
	case int16:
		return int64(v.(int16))
	case int8:
		return int64(v.(int8))
	case uint64:
		return int64(v.(uint64))
	case uint:
		return int64(v.(uint))
	case uint32:
		return int64(v.(uint32))
	case uint16:
		return int64(v.(uint16))
	case uint8:
		return int64(v.(uint8))
	case float64:
		return int64(v.(float64))
	case float32:
		return int64(v.(float32))
	case string:
		if strings.HasPrefix(v.(string), "0x") {
			return HexToInt64(v.(string))
		} else {
			v, _ := strconv.ParseInt(v.(string), 10, 64)
			return v
		}
	}
	panic(fmt.Sprintf("any2Int64: unknown type %T", v))
	return 0
}

func Any2Int32Vector(v any) []int32 {
	{
		l, ok := v.([]int32)
		if ok {
			return l
		}
	}
	{
		l, ok := v.([]int)
		if ok {
			r := make([]int32, len(l))
			for i, v := range l {
				r[i] = int32(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]int64)
		if ok {
			r := make([]int32, len(l))
			for i, v := range l {
				r[i] = int32(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]int16)
		if ok {
			r := make([]int32, len(l))
			for i, v := range l {
				r[i] = int32(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]int8)
		if ok {
			r := make([]int32, len(l))
			for i, v := range l {
				r[i] = int32(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]uint32)
		if ok {
			r := make([]int32, len(l))
			for i, v := range l {
				r[i] = int32(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]uint)
		if ok {
			r := make([]int32, len(l))
			for i, v := range l {
				r[i] = int32(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]uint32)
		if ok {
			r := make([]int32, len(l))
			for i, v := range l {
				r[i] = int32(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]uint16)
		if ok {
			r := make([]int32, len(l))
			for i, v := range l {
				r[i] = int32(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]uint8)
		if ok {
			r := make([]int32, len(l))
			for i, v := range l {
				r[i] = int32(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]float64)
		if ok {
			r := make([]int32, len(l))
			for i, v := range l {
				r[i] = int32(v)
			}
			return r
		}
	}
	{
		l, ok := v.([]float32)
		if ok {
			r := make([]int32, len(l))
			for i, v := range l {
				r[i] = int32(v)
			}
			return r
		}
	}
	{
		l := v.([]any)
		r := make([]int32, len(l))
		for i, v := range l {
			r[i] = int32(Any2Int64(v))
		}
	}
	{
		l := v.([]string)
		r := make([]int32, len(l))
		for i, v := range l {
			if strings.HasPrefix(v, "0x") {
				r[i] = int32(HexToInt64(v))
			} else {
				v, _ := strconv.ParseInt(v, 10, 64)
				r[i] = int32(v)
			}
		}
	}
	panic("any2Int64Vector: unknown type")
}

func Uint32ToHex(u uint32) string {
	s := fmt.Sprintf("%08x", u)
	// 去掉前面的0
	for strings.HasPrefix(s, "0") {
		s = strings.TrimPrefix(s, "0")
	}
	return s
}

func HexToUint32(hex string) uint32 {
	u, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		panic(err)
	}
	return uint32(u)
}

func HexToInt64(h string) int64 {
	u, err := strconv.ParseUint(h, 16, 64)
	if err != nil {
		panic(err)
	}
	return int64(u)
}

func Any2Float64(v any) float64 {
	//数字类型和string类型
	switch v.(type) {
	case float64:
		return v.(float64)
	case float32:
		return float64(v.(float32))
	case int64:
		return float64(v.(int64))
	case int:
		return float64(v.(int))
	case int32:
		return float64(v.(int32))
	case int16:
		return float64(v.(int16))
	case int8:
		return float64(v.(int8))
	case uint64:
		return float64(v.(uint64))
	case uint:
		return float64(v.(uint))
	case uint32:
		return float64(v.(uint32))
	case uint16:
		return float64(v.(uint16))
	case uint8:
		return float64(v.(uint8))
	case string:
		if strings.HasPrefix(v.(string), "0x") {
			return float64(HexToInt64(v.(string)))
		} else {
			v, _ := strconv.ParseFloat(v.(string), 64)
			return v
		}
	}
	panic(fmt.Sprintf("any2Float64: unknown type %T", v))
}

func Any2Float64Vector(v any) []float64 {
	if l, ok := v.([]float64); ok {
		return l
	}
	if l, ok := v.([]float32); ok {
		r := make([]float64, len(l))
		for i, v := range l {
			r[i] = float64(v)
		}
		return r
	}
	if l, ok := v.([]any); ok {
		r := make([]float64, len(l))
		for i, v := range l {
			r[i] = Any2Float64(v)
		}
		return r
	}
	if l, ok := v.([]string); ok {
		r := make([]float64, len(l))
		for i, v := range l {
			if strings.HasPrefix(v, "0x") {
				r[i] = float64(HexToInt64(v))
			} else {
				v, _ := strconv.ParseFloat(v, 64)
				r[i] = v
			}
		}
		return r
	}
	panic("any2Float64Vector: unknown type")
}
