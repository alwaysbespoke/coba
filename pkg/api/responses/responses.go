package responses

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrFailedToMarshalJson = errors.New("failed to marshal JSON")
var ErrFailedToWriteJson = errors.New("failed to write JSON")

func WriteJson(w http.ResponseWriter, resp interface{}, statusCode int) error {
	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		return ErrFailedToMarshalJson
	}

	w.WriteHeader(statusCode)

	if _, err := w.Write(jsonBytes); err != nil {
		return fmt.Errorf("%s: %w", ErrFailedToWriteJson.Error(), err)
	}

	return err
}
