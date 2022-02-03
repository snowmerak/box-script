package golang

import (
	"strconv"
	"strings"

	"github.com/snowmerak/box-script/grammar"
)

func Value(typ *grammar.Type, value *grammar.Value) string {
	sb := strings.Builder{}
	sb.WriteString("{")
	if value.Array != nil {
		num := 0
		for i, v := range value.Array.Values {
			switch strings.Join(value.Array.Type.Name, ".") {
			case grammar.Int8:
				if v.Int != nil {
					n := int8(v.Int.Int.Int64())
					if v.Int.Sign != nil && *v.Int.Sign {
						n = -n
					}
					sb.WriteString(strconv.FormatInt(int64(n), 10))
				}
			case grammar.Int16:
				if v.Int != nil {
					n := int16(v.Int.Int.Int64())
					if v.Int.Sign != nil && *v.Int.Sign {
						n = -n
					}
					sb.WriteString(strconv.FormatInt(int64(n), 10))
				}
			case grammar.Int32:
				if v.Int != nil {
					n := int32(v.Int.Int.Int64())
					if v.Int.Sign != nil && *v.Int.Sign {
						n = -n
					}
					sb.WriteString(strconv.FormatInt(int64(n), 10))
				}
			case grammar.Int64:
				if v.Int != nil {
					n := int64(v.Int.Int.Int64())
					if v.Int.Sign != nil && *v.Int.Sign {
						n = -n
					}
					sb.WriteString(strconv.FormatInt(n, 10))
				}
			case grammar.Uint8:
				if v.Int != nil {
					n := uint8(v.Int.Int.Int64())
					sb.WriteString(strconv.FormatUint(uint64(n), 10))
				}
			case grammar.Uint16:
				if v.Int != nil {
					n := uint16(v.Int.Int.Int64())
					sb.WriteString(strconv.FormatUint(uint64(n), 10))
				}
			case grammar.Uint32:
				if v.Int != nil {
					n := uint32(v.Int.Int.Int64())
					sb.WriteString(strconv.FormatUint(uint64(n), 10))
				}
			case grammar.Uint64:
				if v.Int != nil {
					n := uint64(v.Int.Int.Int64())
					sb.WriteString(strconv.FormatUint(n, 10))
				}
			case grammar.Float32:
				if v.Float != nil {
					f, _ := v.Float.Float.Float64()
					n := float32(f)
					if v.Float.Sign != nil && *v.Float.Sign {
						n = -n
					}
					sb.WriteString(strconv.FormatFloat(float64(n), 'f', -1, 32))
				}
			case grammar.Float64:
				if v.Float != nil {
					f, _ := v.Float.Float.Float64()
					n := float64(f)
					if v.Float.Sign != nil && *v.Float.Sign {
						n = -n
					}
					sb.WriteString(strconv.FormatFloat(n, 'f', -1, 64))
				}
			case grammar.String:
				if v.String != nil {
					sb.WriteString(*v.String)
				}
			case grammar.Char:
				if v.Char != nil && 1 < len(*v.Char) && len(*v.Char) < 4 {
					sb.WriteString(*v.Char)
				}
			case grammar.Bool:
				if v.Bool != nil {
					sb.WriteString(strconv.FormatBool(*v.Bool))
				}
			default:
				continue
			}
			if i < len(value.Array.Values)-1 {
				sb.WriteString(", ")
			}
			num++
		}
		return "[" + strconv.FormatInt(int64(num), 10) + "]" + sb.String()
	}
	if value.Bool != nil {
		return strconv.FormatBool(*value.Bool)
	}
	if value.Char != nil && len(*value.Char) < 2 {
		return *value.Char
	}
	if value.String != nil {
		return *value.String
	}
	if value.Float != nil {
		f, _ := value.Float.Float.Float64()
		n := float64(f)
		if value.Float.Sign != nil && *value.Float.Sign {
			n = -n
		}
		if strings.Join(typ.Name, ".") == grammar.Float32 {
			return strconv.FormatFloat(n, 'f', -1, 32)
		}
		return strconv.FormatFloat(n, 'f', -1, 64)
	}
	if value.Int != nil {
		n := int64(value.Int.Int.Int64())
		if value.Int.Sign != nil && *value.Int.Sign {
			n = -n
		}
		if strings.Join(typ.Name, ".") == grammar.Int8 {
			v := int8(n)
			return strconv.FormatInt(int64(v), 10)
		}
		if strings.Join(typ.Name, ".") == grammar.Int16 {
			v := int16(n)
			return strconv.FormatInt(int64(v), 10)
		}
		if strings.Join(typ.Name, ".") == grammar.Int32 {
			v := int32(n)
			return strconv.FormatInt(int64(v), 10)
		}
		if strings.Join(typ.Name, ".") == grammar.Int64 {
			v := int64(n)
			return strconv.FormatInt(v, 10)
		}
		if strings.Join(typ.Name, ".") == grammar.Uint8 {
			v := uint8(n)
			return strconv.FormatUint(uint64(v), 10)
		}
		if strings.Join(typ.Name, ".") == grammar.Uint16 {
			v := uint16(n)
			return strconv.FormatUint(uint64(v), 10)
		}
		if strings.Join(typ.Name, ".") == grammar.Uint32 {
			v := uint32(n)
			return strconv.FormatUint(uint64(v), 10)
		}
		if strings.Join(typ.Name, ".") == grammar.Uint64 {
			v := uint64(n)
			return strconv.FormatUint(v, 10)
		}
		return value.Int.Int.String()
	}
	return ""
}
