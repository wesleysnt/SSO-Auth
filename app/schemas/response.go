package schemas

type ApiStatusError string

const (
	ApiErrorNotFound       ApiStatusError = "API_ERROR_NOT_FOUND"
	ApiErrorInternalServer ApiStatusError = "API_ERROR_INTERNAL_SERVER"
	ApiErrorBadRequest     ApiStatusError = "API_ERROR_BAD_REQUEST"
	ApiErrorForbidden      ApiStatusError = "API_ERROR_FORBIDDEN"
	ApiErrorUnauthorized   ApiStatusError = "API_ERROR_UNAUTHORIZED"
	ApiErrorUnprocessAble  ApiStatusError = "API_ERROR_UNPROCESSABLE"
)

type ResponseApi struct {
	Details DetailResponse `json:"details"`
	Valid   bool           `json:"valid"`
	Data    any            `json:"data"`
	Errors  any            `json:"errors"`
	Message string         `json:"message"`
}
type DetailResponse struct {
	Path       string `json:"path"`
	Query      string `json:"query"`
	StatusCode int    `json:"status_code"`
	Method     string `json:"method"`
	Status     string `json:"status"`
}
type ResponseApiError struct {
	Status  ApiStatusError
	Message string
}

type ResponseApiErrorValidaton struct {
	StatusCode int    `json:"status_code"`
	Errors     any    `json:"errors"`
	Message    string `json:"message"`
}

type SetResponseApiError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (r *ResponseApiError) Error() string {
	return r.Message
}
