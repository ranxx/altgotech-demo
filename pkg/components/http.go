package components

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ranxx/altgotech-demo/pkg/errors"
)

var encode = json.Marshal
var contentType = "application/json; charset=utf-8"

// EncodeHTTPGenericResponse ....
func EncodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if response == nil {
		return responseToServer(ctx, errors.Response{Msg: "success", Data: struct{}{}}, w)
	}
	return responseToServer(ctx, errors.Response{Msg: "success", Data: response}, w)
}

// ServerErrorEncoder ...
func ServerErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	resp := errors.Convert2Response(err)
	_ = responseToServer(ctx, resp, w)
}

func responseToServer(ctx context.Context, resp errors.Response, w http.ResponseWriter) error {
	data, err := []byte{}, (error)(nil)
	if resp.Data == nil {
		resp.Data = struct{}{}
	}

	if len(data) > 0 {
		data = []byte(fmt.Sprintf(`{"code":%d,"msg":"%s","data":%s}`, resp.Code, resp.Msg, string(data)))
	} else {
		data, err = encode(resp)
	}

	if err != nil {
		return err
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(data)))
	w.Header().Set("Content-Type", contentType)
	if resp.HTTPStatus > 0 {
		w.WriteHeader(resp.HTTPStatus)
	}
	_, err = w.Write(data)
	return err
}
