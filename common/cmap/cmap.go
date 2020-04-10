package cmap

import (
	"encoding/json"
	"sort"
	"strings"
)

type H map[string]interface{}


// 设置参数
func (h H) Set(key string, value interface{}) {
	h[key] = value
}
func (h H)Get(key string) string {
	value, ok := h[key]
	if !ok {
		return ""
	}
	v, ok := value.(string)
	if !ok {
		return convertToString(value)
	}
	return v
}
func convertToString(v interface{}) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}
func (h H) Sort() string{
	var (
		buf strings.Builder
		keys []string
	)
	for k:=range h{
		keys=append(keys, k)
	}
	sort.Strings(keys)
	for i,k:=range keys{
		if v := h.Get(k); v != "" {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
			if i+1<len(keys){
				buf.WriteByte('&')
			}
		}
	}
	return buf.String()
}