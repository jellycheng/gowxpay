package gowxpay

import "strconv"

type MapParams map[string]string

// SetString map本来已经是引用类型了，所以不需要 *MapParams
func (p MapParams) SetString(k, s string) MapParams {
	p[k] = s
	return p
}

func (p MapParams) GetString(k string) string {
	s, _ := p[k]
	return s
}

func (p MapParams) SetInt64(k string, i int64) MapParams {
	p[k] = strconv.FormatInt(i, 10)
	return p
}

func (p MapParams) GetInt64(k string) int64 {
	i, _ := strconv.ParseInt(p.GetString(k), 10, 64)
	return i
}

// ContainsKey 判断key是否存在
func (p MapParams) ContainsKey(key string) bool {
	_, ok := p[key]
	return ok
}
