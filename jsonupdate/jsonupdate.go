package jsonupdate

import "github.com/tidwall/sjson"

type JsonUpdate struct {
	jsondata string
}

func (s *JsonUpdate) String() string {
	return s.jsondata
}

func (s *JsonUpdate) Set(key, value string) *JsonUpdate {
	s.jsondata, _ = sjson.Set(s.jsondata, key, value)
	return s
}

func NewJsonUpdate(jsondata string) *JsonUpdate {
	return &JsonUpdate{
		jsondata: jsondata,
	}
}
