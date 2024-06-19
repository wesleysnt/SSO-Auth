package helpers

import (
	"net/http"
	"sso-auth/app/schemas"

	"github.com/gofiber/fiber/v2"
)

func CatchErrorResponseApi(r *schemas.ResponseApiError) *schemas.SetResponseApiError {
	switch r.Status {
	case schemas.ApiErrorBadRequest:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusBadRequest,
			Message:    r.Message,
		}
	case schemas.ApiErrorForbidden:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusForbidden,
			Message:    r.Message,
		}
	case schemas.ApiErrorNotFound:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusNotFound,
			Message:    r.Message,
		}
	case schemas.ApiErrorUnprocessAble:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    r.Message,
		}
	default:
		return &schemas.SetResponseApiError{
			StatusCode: http.StatusInternalServerError,
			Message:    r.Message,
		}
	}
}

func ResponseApi(ctx *fiber.Ctx, msg string, status string, statusCode int, data any, errors any) error {
	details := schemas.DetailResponse{
		StatusCode: statusCode,
		Path:       ctx.Request().URI().String(),
		Method:     string(ctx.Request().Header.Method()),
		Status:     status,
	}

	if statusCode >= 400 {
		return ctx.Status(statusCode).JSON(schemas.ResponseApi{
			Valid:   false,
			Message: msg,
			Data:    data,
			Errors:  errors,
			Details: details,
		})
	}

	return ctx.Status(statusCode).JSON(schemas.ResponseApi{
		Valid:   true,
		Message: msg,
		Data:    data,
		Errors:  errors,
		Details: details,
	})
}

func ResponseApiCreated(ctx *fiber.Ctx, msg string, data any) error {
	return ResponseApi(ctx, msg, "success_created", http.StatusCreated, data, nil)
}

func ResponseApiOk(ctx *fiber.Ctx, msg string, data any) error {
	return ResponseApi(ctx, msg, "success_ok", http.StatusOK, data, nil)
}

func ResponseApiUnauthorized(ctx *fiber.Ctx, msg string) error {
	return ResponseApi(ctx, msg, "error_unauthorized", http.StatusUnauthorized, nil, nil)
}

func ResponseApiForbidden(ctx *fiber.Ctx, msg string) error {
	return ResponseApi(ctx, msg, "error_forbidden", http.StatusForbidden, nil, nil)
}

func ResponseApiBadRequest(ctx *fiber.Ctx, msg string, errors any) error {
	return ResponseApi(ctx, msg, "error_bad_request", http.StatusBadRequest, nil, errors)
}

func ResponseApiError(ctx *fiber.Ctx, msg string, statusCode int, errors any) error {
	return ResponseApi(ctx, msg, "error_api", statusCode, nil, errors)
}
