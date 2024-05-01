package apierror

import "net/http"

type ErrorType string

const (
	Unauthorized               ErrorType = "UNAUTHORIZED"
	BadRequest                 ErrorType = "BAD_REQUEST"
	InternalServerError        ErrorType = "INTERNAL_SERVER_ERROR"
	APIKeyVerificationFailed   ErrorType = "API_KEY_VERIFICATION_FAILED"
	TokenVerificationFailed    ErrorType = "TOKEN_VERIFICATION_FAILED"
	AuthorizationHeaderMissing ErrorType = "AUTHORIZATION_HEADER_MISSING"
	NoFile                     ErrorType = "NO_FILE_UPLOADED"
	NoProjectName              ErrorType = "NO_PROJECT_NAME_PROVIDED"
)

type APIError struct {
	Reason       ErrorType `json:"reason"`
	ErrorMessage string    `json:"message"`
}

func NewAPIError(reason ErrorType, message string) APIError {
	return APIError{
		Reason:       reason,
		ErrorMessage: message,
	}
}

func (e APIError) Message() string {
	return e.ErrorMessage
}

func (e APIError) Status() int {
	switch e.Reason {
	case Unauthorized:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case InternalServerError:
		return http.StatusInternalServerError
	case APIKeyVerificationFailed:
		return http.StatusUnauthorized
	case TokenVerificationFailed:
		return http.StatusUnauthorized
	case AuthorizationHeaderMissing:
		return http.StatusUnauthorized
	case NoFile:
		return http.StatusBadRequest
	case NoProjectName:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
