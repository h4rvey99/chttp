package chttp

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type request struct {
	c      *client
	url    string
	method string

	meta   http.Header
	params url.Values
}

func toString(v interface{}) string {
	switch v := v.(type) {
	case string:
		return v
	case int8:
		return strconv.Itoa(int(v))
	case int16:
		return strconv.Itoa(int(v))
	case int32:
		return strconv.Itoa(int(v))
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	}
	return fmt.Sprintf("%v", v)
}

func (r *request) Meta(k string, v interface{}) *request {
	r.meta.Add(k, toString(v))
	return r
}

func (r *request) Param(k string, v interface{}) *request {
	r.params.Add(k, toString(v))
	return r
}
