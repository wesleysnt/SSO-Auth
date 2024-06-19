package schemas

// HTTP status codes as registered with IANA.
// See: https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
type ApiStatusCode int

const (
	StatusContinue           ApiStatusCode = 100 // RFC 9110, 15.2.1
	StatusSwitchingProtocols ApiStatusCode = 101 // RFC 9110, 15.2.2
	StatusProcessing         ApiStatusCode = 102 // RFC 2518, 10.1
	StatusEarlyHints         ApiStatusCode = 103 // RFC 8297

	StatusOK                   ApiStatusCode = 200 // RFC 9110, 15.3.1
	StatusCreated              ApiStatusCode = 201 // RFC 9110, 15.3.2
	StatusAccepted             ApiStatusCode = 202 // RFC 9110, 15.3.3
	StatusNonAuthoritativeInfo ApiStatusCode = 203 // RFC 9110, 15.3.4
	StatusNoContent            ApiStatusCode = 204 // RFC 9110, 15.3.5
	StatusResetContent         ApiStatusCode = 205 // RFC 9110, 15.3.6
	StatusPartialContent       ApiStatusCode = 206 // RFC 9110, 15.3.7
	StatusMultiStatus          ApiStatusCode = 207 // RFC 4918, 11.1
	StatusAlreadyReported      ApiStatusCode = 208 // RFC 5842, 7.1
	StatusIMUsed               ApiStatusCode = 226 // RFC 3229, 10.4.1

	StatusMultipleChoices   ApiStatusCode = 300 // RFC 9110, 15.4.1
	StatusMovedPermanently  ApiStatusCode = 301 // RFC 9110, 15.4.2
	StatusFound             ApiStatusCode = 302 // RFC 9110, 15.4.3
	StatusSeeOther          ApiStatusCode = 303 // RFC 9110, 15.4.4
	StatusNotModified       ApiStatusCode = 304 // RFC 9110, 15.4.5
	StatusUseProxy          ApiStatusCode = 305 // RFC 9110, 15.4.6
	_                       ApiStatusCode = 306 // RFC 9110, 15.4.7 (Unused)
	StatusTemporaryRedirect ApiStatusCode = 307 // RFC 9110, 15.4.8
	StatusPermanentRedirect ApiStatusCode = 308 // RFC 9110, 15.4.9

	StatusBadRequest                   ApiStatusCode = 400 // RFC 9110, 15.5.1
	StatusUnauthorized                 ApiStatusCode = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired              ApiStatusCode = 402 // RFC 9110, 15.5.3
	StatusForbidden                    ApiStatusCode = 403 // RFC 9110, 15.5.4
	StatusNotFound                     ApiStatusCode = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed             ApiStatusCode = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable                ApiStatusCode = 406 // RFC 9110, 15.5.7
	StatusProxyAuthRequired            ApiStatusCode = 407 // RFC 9110, 15.5.8
	StatusRequestTimeout               ApiStatusCode = 408 // RFC 9110, 15.5.9
	StatusConflict                     ApiStatusCode = 409 // RFC 9110, 15.5.10
	StatusGone                         ApiStatusCode = 410 // RFC 9110, 15.5.11
	StatusLengthRequired               ApiStatusCode = 411 // RFC 9110, 15.5.12
	StatusPreconditionFailed           ApiStatusCode = 412 // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge        ApiStatusCode = 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong            ApiStatusCode = 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType         ApiStatusCode = 415 // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable ApiStatusCode = 416 // RFC 9110, 15.5.17
	StatusExpectationFailed            ApiStatusCode = 417 // RFC 9110, 15.5.18
	StatusTeapot                       ApiStatusCode = 418 // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest           ApiStatusCode = 421 // RFC 9110, 15.5.20
	StatusUnprocessableEntity          ApiStatusCode = 422 // RFC 9110, 15.5.21
	StatusLocked                       ApiStatusCode = 423 // RFC 4918, 11.3
	StatusFailedDependency             ApiStatusCode = 424 // RFC 4918, 11.4
	StatusTooEarly                     ApiStatusCode = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired              ApiStatusCode = 426 // RFC 9110, 15.5.22
	StatusPreconditionRequired         ApiStatusCode = 428 // RFC 6585, 3
	StatusTooManyRequests              ApiStatusCode = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge  ApiStatusCode = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons   ApiStatusCode = 451 // RFC 7725, 3

	StatusInternalServerError           ApiStatusCode = 500 // RFC 9110, 15.6.1
	StatusNotImplemented                ApiStatusCode = 501 // RFC 9110, 15.6.2
	StatusBadGateway                    ApiStatusCode = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            ApiStatusCode = 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                ApiStatusCode = 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       ApiStatusCode = 505 // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         ApiStatusCode = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           ApiStatusCode = 507 // RFC 4918, 11.5
	StatusLoopDetected                  ApiStatusCode = 508 // RFC 5842, 7.2
	StatusNotExtended                   ApiStatusCode = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired ApiStatusCode = 511 // RFC 6585, 6
)

var StatusTextMap = map[ApiStatusCode]string{
	StatusContinue:                      "Continue",
	StatusSwitchingProtocols:            "Switching Protocols",
	StatusProcessing:                    "Processing",
	StatusEarlyHints:                    "Early Hints",
	StatusOK:                            "OK",
	StatusCreated:                       "Created",
	StatusAccepted:                      "Accepted",
	StatusNonAuthoritativeInfo:          "Non-Authoritative Information",
	StatusNoContent:                     "No Content",
	StatusResetContent:                  "Reset Content",
	StatusPartialContent:                "Partial Content",
	StatusMultiStatus:                   "Multi-Status",
	StatusAlreadyReported:               "Already Reported",
	StatusIMUsed:                        "IM Used",
	StatusMultipleChoices:               "Multiple Choices",
	StatusMovedPermanently:              "Moved Permanently",
	StatusFound:                         "Found",
	StatusSeeOther:                      "See Other",
	StatusNotModified:                   "Not Modified",
	StatusUseProxy:                      "Use Proxy",
	StatusTemporaryRedirect:             "Temporary Redirect",
	StatusPermanentRedirect:             "Permanent Redirect",
	StatusBadRequest:                    "Bad Request",
	StatusUnauthorized:                  "Unauthorized",
	StatusPaymentRequired:               "Payment Required",
	StatusForbidden:                     "Forbidden",
	StatusNotFound:                      "Not Found",
	StatusMethodNotAllowed:              "Method Not Allowed",
	StatusNotAcceptable:                 "Not Acceptable",
	StatusProxyAuthRequired:             "Proxy Authentication Required",
	StatusRequestTimeout:                "Request Timeout",
	StatusConflict:                      "Conflict",
	StatusGone:                          "Gone",
	StatusLengthRequired:                "Length Required",
	StatusPreconditionFailed:            "Precondition Failed",
	StatusRequestEntityTooLarge:         "Request Entity Too Large",
	StatusRequestURITooLong:             "Request URI Too Long",
	StatusUnsupportedMediaType:          "Unsupported Media Type",
	StatusRequestedRangeNotSatisfiable:  "Requested Range Not Satisfiable",
	StatusExpectationFailed:             "Expectation Failed",
	StatusTeapot:                        "I'm a teapot",
	StatusMisdirectedRequest:            "Misdirected Request",
	StatusUnprocessableEntity:           "Unprocessable Entity",
	StatusLocked:                        "Locked",
	StatusFailedDependency:              "Failed Dependency",
	StatusTooEarly:                      "Too Early",
	StatusUpgradeRequired:               "Upgrade Required",
	StatusPreconditionRequired:          "Precondition Required",
	StatusTooManyRequests:               "Too Many Requests",
	StatusRequestHeaderFieldsTooLarge:   "Request Header Fields Too Large",
	StatusUnavailableForLegalReasons:    "Unavailable For Legal Reasons",
	StatusInternalServerError:           "Internal Server Error",
	StatusNotImplemented:                "Not Implemented",
	StatusBadGateway:                    "Bad Gateway",
	StatusServiceUnavailable:            "Service Unavailable",
	StatusGatewayTimeout:                "Gateway Timeout",
	StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	StatusInsufficientStorage:           "Insufficient Storage",
	StatusLoopDetected:                  "Loop Detected",
	StatusNotExtended:                   "Not Extended",
	StatusNetworkAuthenticationRequired: "Network Authentication Required",
}
