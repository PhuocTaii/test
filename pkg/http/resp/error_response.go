package resp

import (
	"errors"
)

type ErrorResp struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Log        string `json:"-"`
	RootErr    error  `json:"-"`
}

// RootError returns the root error of the ResError.
//
// It does not take any parameters.
// It returns an error.
func (e *ErrorResp) RootError() error {
	var err *ErrorResp

	if errors.As(e.RootErr, &err) {
		return err.RootError()
	}

	return e.RootErr
}

// Error returns the error message of the ResError.
//
// It returns a string.
func (e *ErrorResp) Error() string {
	return e.RootError().Error()
}

// NewErrorResp creates a new custom error with the given code, root error, message, and key.
//
// Parameters:
//
//   - code: The error code.
//
//   - root: The root error.
//
//   - msg: The error message.
//
// Returns:
//   - *ResError: The new custom error.
func NewErrorResp(statusCode int, root error, msg string) *ErrorResp {
	if root != nil {
		return &ErrorResp{
			StatusCode: statusCode,
			RootErr:    root,
			Message:    msg,
			Log:        root.Error(),
		}
	}

	return &ErrorResp{
		StatusCode: statusCode,
		RootErr:    errors.New(msg),
		Message:    msg,
		Log:        msg,
	}
}
