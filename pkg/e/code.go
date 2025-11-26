package e

// ErrorCode枚举
const (
	SUCCESS  = 200
	CREATED  = 201
	ACCEPTED = 202

	ERROR_BAD_REQUEST  = 400
	ERROR_UNAUTHORIZED = 401
	ERROR_FORBIDDEN    = 403
	ERROR_NOT_FOUND    = 404
	ERROR_CONFLICT     = 409

	ERROR_INTERNAL = 500
)

var MsgFlags = map[int]string{
	SUCCESS:            "OK",
	CREATED:            "Created",
	ACCEPTED:           "Accepted",
	ERROR_BAD_REQUEST:  "Bad Request",
	ERROR_UNAUTHORIZED: "Unauthorized",
	ERROR_FORBIDDEN:    "Forbidden",
	ERROR_NOT_FOUND:    "Not Found",
	ERROR_CONFLICT:     "Conflict",
	ERROR_INTERNAL:     "Internal Server Error",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg

	}
	return MsgFlags[ERROR_INTERNAL]
}
