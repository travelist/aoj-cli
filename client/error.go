package client

import (
	"fmt"
	"github.com/travelist/aoj-cli/client/response"
	"net/http"
)

type AOJErrorCode string

// AOJ API Error
type AOJClientError struct {
	// HTTP Status code
	statusCode int

	// Actual API response
	response response.ErrorResponse
}

func (e *AOJClientError) Error() string {
	return fmt.Sprintf("aoj api client error: status_code=%d", e.statusCode)
}

//func (m *AOJClientError) String() string { return m.value }

func newAOJClientError(res *http.Response) *AOJClientError {
	var result = AOJClientError{statusCode: res.StatusCode}
	var body = response.ErrorResponse{}
	if e := decodeBody(res, &body); e != nil {
		return &result
	}

	result.response = body
	return &result
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
