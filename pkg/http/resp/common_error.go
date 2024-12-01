package resp

import "net/http"

func ErrInternalServer(err error) *ErrorResp {
	return NewErrorResp(
		http.StatusInternalServerError,
		err,
		err.Error(),
	)
}

func ErrInvalidRequest(err error) *ErrorResp {
	return NewErrorResp(
		http.StatusBadRequest,
		err,
		err.Error(),
	)
}

func ErrMissingTokenInHeader(err error) *ErrorResp {
	return NewErrorResp(
		http.StatusUnauthorized,
		err,
		"missing token in header",
	)
}

func ErrMissingUserIdInHeader(err error) *ErrorResp {
	return NewErrorResp(
		http.StatusUnauthorized,
		err,
		"missing uid in header",
	)
}

func ErrFailedToParseUserId(err error) *ErrorResp {
	return NewErrorResp(
		http.StatusUnauthorized,
		err,
		"failed to parse uid",
	)
}

func ErrInvalidTokenFormat(err error) *ErrorResp {
	return NewErrorResp(
		http.StatusUnauthorized,
		err,
		"token is invalid signature",
	)
}

func ErrInvalidSignature(err error) *ErrorResp {
	return NewErrorResp(
		http.StatusUnauthorized,
		err,
		"token is invalid signature",
	)
}
