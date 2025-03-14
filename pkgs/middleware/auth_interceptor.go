package middleware

import (
	"context"
	"encoding/json"
	"goparking/pkgs/token"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthInterceptor struct {
	ignoredMethods []string
	token          token.IMarker
}

func NewAuthInterceptor(ignoredMethods []string, token token.IMarker) *AuthInterceptor {
	return &AuthInterceptor{
		ignoredMethods: ignoredMethods,
		token:          token,
	}
}

func (ai *AuthInterceptor) authorize(ctx context.Context) (context.Context, string, error) {
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(m["token"]) == 0 {
		return ctx, "", status.New(codes.Unauthenticated, "missing token").Err()
	}

	payload, err := ai.token.ValidateToken(m["token"][0])
	if err != nil {
		return ctx, "", status.New(codes.Unauthenticated, "unauthorized").Err()
	}

	var meta map[string]interface{}
	b, err := json.Marshal(payload)
	if err != nil {
		return ctx, "", status.New(codes.Unauthenticated, "unauthorized").Err()
	} else {
		if err := json.Unmarshal(b, &meta); err != nil {
			log.Println("Error while unmarshalling auth data", err)
		}
	}

	return ctx, payload.ID, nil
}

func (ai *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		for _, m := range ai.ignoredMethods {
			if info.FullMethod == m {
				return handler(ctx, req)
			}
		}

		ctx, userID, err := ai.authorize(ctx)
		if err != nil {
			return nil, status.New(codes.Internal, "Loi").Err()
		}

		ctx = context.WithValue(ctx, "userId", userID)

		return handler(ctx, req)
	}
}
