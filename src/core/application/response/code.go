package response

type Code string

const (
	CodeSuccess       Code = "OK"
	CodeNotFound      Code = "NOT_FOUND"
	CodeConflict      Code = "CONFLICT"
	CodeBadRequest    Code = "BAD_REQUEST"
	CodeUnauthorized  Code = "UNAUTHORIZED"
	CodeDatabaseError Code = "DATABASE_ERROR"
	CodeGatewayError  Code = "GATEWAY_ERROR"
	CodeInvalidToken  Code = "INVALID_TOKEN"
	CodeExpiredToken  Code = "EXPIRED_TOKEN"
	CodeFileError     Code = "FILE_ERROR"
	CodeForbidden     Code = "FORBIDDEN"
	CodeAsyncError    Code = "ASYNC_ERROR"
)
