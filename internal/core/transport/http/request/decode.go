package core_http_request

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	core_errors "github.com/i-dontneedaname/study-todoapp/internal/core/errors"
)

type validatable interface {
	Validate() error
}

var requestValidator = validator.New()

func DecodeAndValidateRequest(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("decode json: %v: %w", err, core_errors.ErrInvalidArg)
	}

	v, ok := dest.(validatable)
	if ok {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("request validation: %v: %w", err, core_errors.ErrInvalidArg)
		}
	} else {
		if err := requestValidator.Struct(dest); err != nil {
			return fmt.Errorf("request validation: %v: %w", err, core_errors.ErrInvalidArg)
		}
	}

	return nil
}
