package authorize

import (
	"context"
	"encoding/json"
	"github.com/mohsensamiei/gopher/pkg/metadataext"
)

func ToContext(ctx context.Context, claim *Claims) context.Context {
	bin, _ := json.Marshal(claim)
	return metadataext.SetValue(ctx, CurrentClaims, string(bin))
}

func FromContext(ctx context.Context) *Claims {
	str, ok := metadataext.GetValue(ctx, CurrentClaims)
	if !ok {
		return nil
	}
	claim := new(Claims)
	_ = json.Unmarshal([]byte(str), claim)
	return claim
}
