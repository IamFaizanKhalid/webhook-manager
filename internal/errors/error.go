package errors

import (
	"github.com/IamFaizanKhalid/webhook-manager/internal/utils"
	"net/http"
)

// Response error type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type Response struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e Response) Error() string {
	return e.Err.Error()
}

func InvalidRequest(err error) error {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     utils.StatusText(http.StatusBadRequest),
		ErrorText:      err.Error(),
	}
}

func InternalError(err error) error {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     utils.StatusText(http.StatusInternalServerError),
		ErrorText:      err.Error(),
	}
}

func Conflict(err error) error {
	return &Response{
		Err:            err,
		HTTPStatusCode: http.StatusConflict,
		StatusText:     utils.StatusText(http.StatusConflict),
		ErrorText:      err.Error(),
	}
}

var NotFound = &Response{
	HTTPStatusCode: http.StatusNotFound,
	StatusText:     utils.StatusText(http.StatusNotFound),
}

var Unauthorized = &Response{
	HTTPStatusCode: http.StatusUnauthorized,
	StatusText:     utils.StatusText(http.StatusUnauthorized),
}
