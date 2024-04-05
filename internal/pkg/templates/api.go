package templates

const (
	ApiModel = `
syntax = "proto3";
package api;
option go_package = "{{ .repository }}/api";

message {{ .Singular }} {
  // @inject_tag: json:"id,omitempty"
  string ID = 1;
}

message {{ .Plural }} {
  // @inject_tag: json:"elements,omitempty"
  repeated {{ .Singular }} Elements = 1;
  // @inject_tag: json:"count"
  int64 Count = 2;
}
`
	ApiService = `
syntax = "proto3";
package api;
option go_package = "{{ .repository }}/api";
import "api/src/misc.proto";
import "api/src/{{ .singular }}_model.proto";

service {{ .Singular }}Service {
  rpc Create ({{ .Singular }}Create) returns ({{ .Singular }});
  rpc Update ({{ .Singular }}Update) returns ({{ .Singular }});
  rpc Return ({{ .Singular }}Return) returns ({{ .Singular }});
  rpc Delete ({{ .Singular }}Delete) returns (Void);
  rpc List ({{ .Singular }}List) returns ({{ .Plural }});
}

message {{ .Singular }}Create {
  // @inject_tag: json:"query"
  string Query = 1;
}

message {{ .Singular }}Update {
  // @inject_tag: json:"{{ .singular }}_id"
  string {{ .Singular }}ID = 1;
  // @inject_tag: json:"query"
  string Query = 2;
}

message {{ .Singular }}Return {
  // @inject_tag: json:"{{ .singular }}_id"
  string {{ .Singular }}ID = 1;
  // @inject_tag: json:"query"
  string Query = 2;
}

message {{ .Singular }}Delete {
  // @inject_tag: json:"{{ .singular }}_id"
  string {{ .Singular }}ID = 1;
}

message {{ .Singular }}List {
  // @inject_tag: json:"query"
  string Query = 1;
}
`

	ApiEnum = `//region enum {{ .Enum }} methods
func ({{ .Enum }}) Values() []string {
	return mapext.Values({{ .Enum }}_name)
}
func ({{ .Enum }}) InRange(v interface{}) bool {
	_, ok := {{ .Enum }}_value[v.({{ .Enum }}).String()]
	return ok
}
func (x *{{ .Enum }}) Scan(value interface{}) error {
	*x = {{ .Enum }}({{ .Enum }}_value[value.(string)])
	return nil
}
func (x {{ .Enum }}) Value() (driver.Value, error) {
	return x.String(), nil
}
func (x *{{ .Enum }}) UnmarshalJSON(b []byte) error {
	var str string
	if err := json.Unmarshal(b, &str); err != nil {
		return err
	}
	*x = {{ .Enum }}({{ .Enum }}_value[str])
	return nil
}
func (x {{ .Enum }}) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
//endregion`
)
