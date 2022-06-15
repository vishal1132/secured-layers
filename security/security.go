package security

import (
	"reflect"
	"strings"
)

type SecurityCompliance interface {
	EncryptRegisteredFields(key interface{})
	DecryptRegisteredFields()
}

type security interface {
	RegisterFields(encryptingType interface{}, fields []string)
	SecurityCompliance
}

type securityCompliant struct {
	Fields map[string]map[string]struct{}
}

func New() security {
	return &securityCompliant{
		Fields: make(map[string]map[string]struct{}),
	}
}

func (s *securityCompliant) RegisterFields(encrypted interface{}, fields []string) {
	refType := reflect.TypeOf(encrypted)
	typeKey := refType.String()
	_, ok := s.Fields[typeKey]
	if !ok {
		s.Fields[typeKey] = make(map[string]struct{})
	}
	for _, v := range fields {
		s.Fields[typeKey][v] = struct{}{}
	}
}

func (s *securityCompliant) EncryptRegisteredFields(key interface{}) {
	refType := reflect.TypeOf(key)
	typeKey := refType.String()
	typeKey = strings.TrimPrefix(typeKey, "*")
	v, ok := s.Fields[typeKey]
	if !ok {
		return
	}
	e := reflect.Indirect(reflect.ValueOf(key))
	for i := 0; i < e.NumField(); i++ {
		varName := e.Type().Field(i).Name
		if _, ok := v[varName]; ok {
			e.Field(i).Set(reflect.ValueOf(encrypt()))
		}
	}
}

// dummy function
func encrypt(args ...string) string {
	return "obfuscated"
}

func (s *securityCompliant) DecryptRegisteredFields() {

}
