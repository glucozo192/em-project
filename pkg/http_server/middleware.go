package http_server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/glu/shopvui/idl/pb"
	"github.com/glu/shopvui/utils/authenticate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/status"
)

type middlewareFunc func(http.Handler) http.Handler

var (
	Authorization = "Authorization"
	Bearer        = "Bearer"
	InfoKey       = "info_key"
	CookieKey     = "h5token"

	ignoredAPIs []string
)

func init() {
	ignoredAPIs = []string{
		pb.AllPathToMethodMap[pb.UserService_Login_API],
		pb.AllPathToMethodMap[pb.UserService_Register_API],
	}
}

type payloadKeys struct{}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, authorization")
		if r.Method != "OPTIONS" {
			h.ServeHTTP(w, r)
		}
	})
}

type Response struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
	Data    any      `json:"data"`
}

func ErrorResponse(w http.ResponseWriter, code int, err error) {
	resp := &Response{
		Code:    code,
		Message: err.Error(),
		Details: []string{},
	}

	jData, _ := json.Marshal(resp)

	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func DataResponse(w http.ResponseWriter, data any) {
	resp := &Response{
		Data: data,
	}

	jData, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func forwardErrorResponse(ctx context.Context, s *runtime.ServeMux, m runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	sta := status.Convert(err)
	errStr := sta.Message()
	log.Print(err)
	firstColonPos := strings.Index(errStr, ":")

	if firstColonPos > 0 {
		errStr = errStr[:firstColonPos]
	}

	runtime.DefaultHTTPErrorHandler(ctx, s, m, w, r, errors.New(errStr))
}

func authorized(
	authenticator authenticate.Authenticator,
) middlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			incomingPath := fmt.Sprintf("%s %s", r.Method, r.URL.Path)
			fmt.Println("incomingPath: ", incomingPath)

			if isCheckingPassed(ignoredAPIs, incomingPath) {
				h.ServeHTTP(w, r)
				return
			}

			authorization := r.Header.Get(Authorization)
			scheme, token, found := strings.Cut(authorization, " ")
			if !found {
				ErrorResponse(w, http.StatusUnauthorized, fmt.Errorf("invalid authorization token"))
				return
			}
			if !strings.EqualFold(scheme, Bearer) {
				ErrorResponse(w, http.StatusUnauthorized, fmt.Errorf("invalid authorization token"))
				return
			}
			fmt.Println("Verify(token): ", token)
			payload, err := authenticator.VerifyToken(token)
			fmt.Println("payload", payload)
			if err != nil {
				ErrorResponse(w, http.StatusUnauthorized, err)
				return
			}

			h.ServeHTTP(w, r.WithContext(context.WithValue(ctx, payloadKeys{}, payload)))
		})
	}
}
