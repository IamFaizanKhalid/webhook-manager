package output

import (
	"github.com/IamFaizanKhalid/webhook-api/internal/utils"
	"net/http"
)

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e ErrResponse) Error() string {
	return e.Err.Error()
}

func ErrInvalidRequest(err error) error {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     utils.StatusText(http.StatusBadRequest),
		ErrorText:      err.Error(),
	}
}

func ErrInternalError(err error) error {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     utils.StatusText(http.StatusInternalServerError),
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{
	HTTPStatusCode: http.StatusNotFound,
	StatusText:     utils.StatusText(http.StatusNotFound),
}
