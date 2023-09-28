package http

const (
	Get    = "GET"
	Post   = "POST"
	DELETE = "DELETE"
)

const VersionsFor11 = "HTTP/1.1"

const (
	StatusSuccessCode             = "200"
	StatusRedirectCode            = "302"
	StatusBadRequestCode          = "400"
	StatusUnauthorizedCode        = "401"
	StatusNotFoundCode            = "404"
	StatusMethodNotAllowedCode    = "405"
	StatusInternalServerErrorCode = "500"
)

const (
	StatusReasonOk                  = "OK"
	StatusReasonRedirect            = "Found"
	StatusReasonBadRequest          = "Bad Request"
	StatusReasonUnauthorized        = "Unauthorized"
	StatusReasonNotFound            = "Not Found"
	StatusReasonMethodNotAllowed    = "Method Not Allowed"
	StatusReasonInternalServerError = "Internal Server Error"
)
