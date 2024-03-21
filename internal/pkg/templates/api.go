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
)
