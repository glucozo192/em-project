package metadata

import (
	"context"

	"github.com/glu/shopvui/utils/authenticate"

	"google.golang.org/grpc/metadata"
)

const (
	MDUserIDKey     = "user_id"
	MDTokenKey      = "token"
	MDXForwardedFor = "x-forwarded-for"
)

func ImportUserInfoToCtx(payload *authenticate.Payload) metadata.MD {
	md := metadata.Pairs(MDUserIDKey, payload.UserID)
	return md
}

func ExtractUserInfoFromCtx(ctx context.Context) (*authenticate.Payload, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, false
	}

	vals := []string{}

	for _, key := range []string{MDUserIDKey, MDTokenKey, MDXForwardedFor} {
		values := md.Get(key)
		if len(values) < 1 {
			vals = append(vals, "")
		} else {
			vals = append(vals, values[0])
		}
	}

	return &authenticate.Payload{
		UserID: vals[0],
	}, true
}

func InjectIncomingCtxToOutgoingCtx(ctx context.Context) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)
	return metadata.NewOutgoingContext(ctx, md)
}
