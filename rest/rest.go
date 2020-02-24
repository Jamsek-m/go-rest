package rest

import (
	"encoding/json"
	"net/http"

	"github.com/Jamsek-m/go-rest/errors"
	"github.com/Jamsek-m/go-rest/headers"
	"github.com/Jamsek-m/go-rest/media"
)

type JsonArgs struct {
	Data    interface{}
	Status  int
	Headers map[string]string
}

func Json(res http.ResponseWriter, args JsonArgs) {
	if args.Status == 0 {
		args.Status = http.StatusOK
	}

	if args.Headers != nil {
		if _, present := args.Headers[headers.CONTENT_TYPE]; !present {
			res.Header().Add(headers.CONTENT_TYPE, media.APPLICATION_JSON)
		}
		for headerName, headerValue := range args.Headers {
			res.Header().Add(headerName, headerValue)
		}
	} else {
		res.Header().Add(headers.CONTENT_TYPE, media.APPLICATION_JSON)
	}

	res.WriteHeader(args.Status)

	if args.Data != nil {
		if err := SerializeBody(res, args.Data); err != nil {
			println(err)
		}
	}
}

func HandleError(res http.ResponseWriter, err error) {
	res.Header().Add(headers.CONTENT_TYPE, media.APPLICATION_JSON)

	var statusCode int
	var responseBody errors.ErrorResponse
	switch e := err.(type) {
	case *errors.RestError:
		responseBody = errors.NewErrorResponseFromError(*e)
		statusCode = e.Status
	default:
		responseBody = errors.NewErrorResponse(e.Error(), http.StatusInternalServerError)
		statusCode = http.StatusInternalServerError
	}

	res.WriteHeader(statusCode)
	if err := SerializeBody(res, responseBody); err != nil {
		println(err)
	}
}

func ParseBody(req *http.Request, entity interface{}) error {
	err := json.NewDecoder(req.Body).Decode(entity)
	return err
}

func SerializeBody(res http.ResponseWriter, body interface{}) error {
	err := json.NewEncoder(res).Encode(body)
	return err
}
