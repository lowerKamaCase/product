package req

import (
	"net/http"

	"github.com/lowerKamaCase/product/pkg/res"
)

func HandleBody[T any](w *http.ResponseWriter, req *http.Request) (*T, error) {
	body, err := Decode[T](req.Body)
	if err != nil {
		res.Json(*w, err.Error(), http.StatusPaymentRequired)
		return nil, err
	}

	err = IsValid(body)
	if err != nil {
		res.Json(*w, err.Error(), http.StatusPaymentRequired)
		return nil, err
	}

	return &body, nil

}
