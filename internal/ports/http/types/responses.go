package types

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/render"
)

var (
	_ render.Renderer = Response{}

	errWrongResponseType = errors.New("wrong response type")

	InternalError = errResponse(HttpError{
		Code:        http.StatusInternalServerError,
		Error:       "Internal Server Error",
		Description: "Something went wrong, contact the administration",
	})
)

type Response struct {
	Ok     bool       `json:"ok"`
	Result any        `json:"result,omitempty"`
	Error  *HttpError `json:"error,omitempty"`
}

type HttpError struct {
	Code        int    `json:"error_code"`
	Error       string `json:"error"`
	Description string `json:"description"`
}

func (r Response) Render(_ http.ResponseWriter, _ *http.Request) error {
	if !r.Ok && (r.Result != nil || r.Error == nil) {
		return errWrongResponseType
	}

	return nil
}

func errResponse(err HttpError) Response {
	return Response{
		Ok:     false,
		Result: nil,
		Error:  &err,
	}
}

type ErrInvalidEmail struct {
	Email string
}

func (e ErrInvalidEmail) Error() string {
	return "invalid email: " + e.Email
}

type UserResponse struct {
	ID            uint64    `json:"id"`
	Name          string    `json:"name"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	BitcoinAmount float64   `json:"bitcoin_amount"`
	UsdBalance    float64   `json:"usd_balance"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
