package swagext

import (
	"fmt"
	"github.com/go-openapi/spec"
	"github.com/swaggo/swag"
	"strings"
)

const (
	namesKey        = "x-enum-varnames"
	schemesTemplate = `"schemes": {{ marshal .Schemes }},`
)

func init() {
	swagger := swag.GetSwagger(Name).(*swag.Spec)

	doc := new(spec.Swagger)
	_ = doc.UnmarshalJSON([]byte(
		strings.Replace(swagger.SwaggerTemplate, schemesTemplate, "", 1)))

	for k, v := range doc.Definitions {
		if v.Enum == nil {
			continue
		}
		names, ok := v.VendorExtensible.Extensions.GetStringSlice(namesKey)
		if !ok {
			continue
		}

		v.Type = spec.StringOrArray{"string"}
		for i := range v.Enum {
			if strings.Contains(names[i], "_") {
				v.Enum[i] = strings.SplitN(names[i], "_", 2)[1]
			}
		}
		delete(v.VendorExtensible.Extensions, namesKey)

		doc.Definitions[k] = v
	}

	bytes, _ := doc.MarshalJSON()
	swagger.SwaggerTemplate = strings.Replace(
		fmt.Sprintf("%s", bytes), "{", fmt.Sprintf("{%v", schemesTemplate), 1)
}
