package errors

import (
	"encoding/json"
	"encoding/xml"
)

type Validation struct {
	Field string `json:"field,omitempty" xml:"field,omitempty" yaml:"field,omitempty"`
	Tag   string `json:"tag,omitempty" xml:"tag,omitempty" yaml:"tag,omitempty"`
	Param string `json:"param,omitempty" xml:"param,omitempty" yaml:"param,omitempty"`
}

type Model struct {
	Code        uint32        `json:"code,omitempty" xml:"code,omitempty" yaml:"code,omitempty"`
	Message     string        `json:"message,omitempty" xml:"message,omitempty" yaml:"message,omitempty"`
	Details     []string      `json:"details,omitempty" xml:"details,omitempty" yaml:"details,omitempty"`
	Validations []*Validation `json:"validations,omitempty" xml:"validations,omitempty" yaml:"validations,omitempty"`
}

func (e Error) Model() *Model {
	return &Model{
		Code:        uint32(e.Code()),
		Message:     e.Localize(),
		Details:     append([]string{e.Message()}, e.Details()...),
		Validations: e.Validations(),
	}
}

func (e Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.Model())
}

func (e Error) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	return encoder.EncodeElement(e.Model(), start)
}

func (e Error) MarshalYAML() (any, error) {
	return e.Model(), nil
}
