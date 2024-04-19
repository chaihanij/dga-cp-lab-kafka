package errors

const (
	SuccessCode                = "10000"
	ErrCodeInternalServerError = "E50000"
	ErrCodeBadRequest          = "E40000"
	ErrCodeUnauthorized        = "E40100"
	ErrCodeForbidden           = "E40300"
	ErrCodeNotFound            = "E40400"
	ErrCodeConflict            = "E40900"
	ErrCodeUnprocessableEntity = "E42200"
	ErrCodeTooManyRequests     = "E42900"
)

func ErrorCodeToMessage(code string) string {
	switch code {
	case SuccessCode:
		return "Success"
	case ErrCodeInternalServerError:
		return "Internal server error"
	case ErrCodeBadRequest:
		return "Bad request"
	case ErrCodeUnauthorized:
		return "Unauthorized"
	case ErrCodeForbidden:
		return "Forbidden"
	case ErrCodeNotFound:
		return "Not found"
	case ErrCodeConflict:
		return "Conflict"
	case ErrCodeUnprocessableEntity:
		return "Unprocessable entity"
	case ErrCodeTooManyRequests:
		return "Too many requests"
	default:
		return "Unknown error"
	}
}
