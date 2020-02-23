package rest

import (
	"encoding/json"
	"net/http"
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
		if _, present := args.Headers["Content-Type"]; !present {
			res.Header().Add("Content-Type", "application/json")
		}
		for headerName, headerValue := range args.Headers {
			res.Header().Add(headerName, headerValue)
		}
	} else {
		res.Header().Add("Content-Type", "application/json")
	}

	res.WriteHeader(args.Status)

	if args.Data != nil {
		if err := json.NewEncoder(res).Encode(args.Data); err != nil {
			// logger.Log.WithError(err).Error("Error encoding to json!")
			println(err)
		}
	}
}
