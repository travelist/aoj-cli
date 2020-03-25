package client

import (
	"fmt"
	"github.com/travelist/aoj-cli/client/response"
	"net/http"
)

type AOJErrorCode string

// AOJ API Error
type AOJClientError struct {

	// Actual API response
	response *http.Response
}

func (e *AOJClientError) Error() string {
	if e.response == nil {
		return fmt.Sprintf("AOJ API Client Error: no response")
	}

	url := e.response.Request.URL
	method := e.response.Request.Method
	statusCode := e.response.StatusCode
	var body response.ErrorResponse
	decodeBody(e.response, &body)
	fmt.Println(body[0].Code)

	return fmt.Sprintf("AOJ API Client Error: url=%s method=%s status_code=%d", url, method, statusCode)
}

//func (m *AOJClientError) String() string { return m.value }

func newAOJClientError(res *http.Response) *AOJClientError {
	return &AOJClientError{response: res}
}

//func IsBadRequestError(e AOJClientError) bool {
//	return 400 <= e.statusCode && e.statusCode < 500
//}
//
//func IsInternalServerError(e AOJClientError) bool {
//	return 500 <= e.statusCode
//}

//const (
//	// API
//	ErrApiExecutionError AOJErrorCode = "API_EXECUTION_ERROR"
//	ErrJudgeQueueError   AOJErrorCode = "JUDGE_QUEUE_ERROR"
//
//	// Database
//	ErrDatabaseConnectionError AOJErrorCode = "DATABASE_CONNECTION_ERROR"
//	ErrDatabaseQueryError      AOJErrorCode = "DATABASE_QUERY_ERROR"
//
//	// Parameter
//	ErrValidationError       AOJErrorCode = "VALIDATION_ERROR"
//	ErrDuplicationError      AOJErrorCode = "DUPLICATION_ERROR"
//	ErrDataNotExistError     AOJErrorCode = "DATA_NOT_EXIST_ERROR"
//	ErrResourceNotExistError AOJErrorCode = "RESOURCE_NOT_EXIST_ERROR"
//	ErrMethodNotAllowedError AOJErrorCode = "METHOD_NOT_ALLOWED_ERROR"
//
//	// Security
//	ErrInvalidRefreshToken AOJErrorCode = "INVALID_REFRESH_TOKEN_ERROR"
//	ErrInvalidAccessToken  AOJErrorCode = "INVALID_ACCESS_TOKEN_ERROR"
//	ErrAccessDeniedError   AOJErrorCode = "ACCESS_DENIED_ERROR"
//	ErrUserNotFoundError   AOJErrorCode = "USER_NOT_FOUND_ERROR"
//
//	// Unknown
//	ErrUnknownError AOJErrorCode = "UNKNOWN_ERROR"
//)
