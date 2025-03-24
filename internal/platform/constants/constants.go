package constants

const (
	ContentTypeJSON       = "application/json"
	ContentTypeText       = "text/plain"
	ContentTypeHTML       = "text/html"
	ContentTypeXML        = "application/xml"
	ContentTypeFormData   = "multipart/form-data"
	ContentTypeURLEncoded = "application/x-www-form-urlencoded"
	
	HeaderContentType     = "Content-Type"
	HeaderContentLength   = "Content-Length"
	HeaderAuthorization   = "Authorization"
	HeaderAccept          = "Accept"
	HeaderUserAgent       = "User-Agent"
	HeaderXForwardedFor   = "X-Forwarded-For"
	
	MethodGET     = "GET"
	MethodPOST    = "POST"
	MethodPUT     = "PUT"
	MethodDELETE  = "DELETE"
	MethodPATCH   = "PATCH"
	MethodOPTIONS = "OPTIONS"
	MethodHEAD    = "HEAD"
	
	StatusOK                  = 200
	StatusCreated             = 201
	StatusAccepted            = 202
	StatusNoContent           = 204
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusInternalServerError = 500
	StatusNotImplemented      = 501
	StatusBadGateway          = 502
	StatusServiceUnavailable  = 503
)

const (
	RouteHealth = "/health"
	RouteRoot   = "/"
	RouteVersion = "/version"
	
	MessageHealthOK = "OK"
	MessageNotFound = "Not Found"
	MessageError    = "Error"
)
