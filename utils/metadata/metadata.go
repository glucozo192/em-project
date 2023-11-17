package metadata

import (
	"context"
	"github.com/glu/shopvui/utils/authenticate"

	"google.golang.org/grpc/metadata"
)

const (
	MDUserIDKey     = "user_id"
	MDUserNameKey   = "user_name"
	MDIpKey         = "ip"
	MDTokenKey      = "token"
	MDRolesKey      = "roles"
	MDXForwardedFor = "x-forwarded-for"
)

func ImportIpToCtx(ip string) metadata.MD {
	md := metadata.Pairs(MDIpKey, ip)

	return md
}

func ExtractIpFromCtx(ctx context.Context) (*authenticate.Payload, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, false
	}
	values := md.Get(MDIpKey)
	if len(values) < 1 {
		return nil, false
	}
	return &authenticate.Payload{
		Ip: values[0],
	}, true
}

func ImportUserInfoToCtx(payload *authenticate.Payload) metadata.MD {
	md := metadata.Pairs(MDUserIDKey, payload.UserID)
	md.Append(MDUserNameKey, payload.UserName)
	md.Append(MDTokenKey, payload.Token)
	md.Append(MDRolesKey, payload.Roles...)
	return md
}

func ExtractUserInfoFromCtx(ctx context.Context) (*authenticate.Payload, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, false
	}

	vals := []string{}

	for _, key := range []string{MDUserIDKey, MDUserNameKey, MDIpKey, MDTokenKey, MDXForwardedFor} {
		values := md.Get(key)
		if len(values) < 1 {
			vals = append(vals, "")
		} else {
			vals = append(vals, values[0])
		}
	}

	ip := vals[2]
	if vals[2] == "" {
		ip = vals[4]
	}

	return &authenticate.Payload{
		UserID:   vals[0],
		UserName: vals[1],
		Ip:       ip,
		Token:    vals[3],
		Roles:    md.Get(MDRolesKey),
	}, true
}

func InjectIncomingCtxToOutgoingCtx(ctx context.Context) context.Context {
	md, _ := metadata.FromIncomingContext(ctx)
	return metadata.NewOutgoingContext(ctx, md)
}
