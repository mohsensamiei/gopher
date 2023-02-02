package errors

import (
	"encoding/json"
	"encoding/xml"
)

type Model struct {
	Code    uint32   `json:"code,omitempty" xml:"code,omitempty" yaml:"code,omitempty"`
	Message string   `json:"message,omitempty" xml:"message,omitempty" yaml:"message,omitempty"`
	Details []string `json:"details,omitempty" xml:"details,omitempty" yaml:"details,omitempty"`
}

func (e Error) Model() *Model {
	model := &Model{
		Code:    uint32(e.Code()),
		Message: e.Localize(),
	}
	model.Details = append([]string{e.Message()}, e.Details()...)
	return model
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
