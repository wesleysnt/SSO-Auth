package clienthandler

import (
	"sso-auth/app/helpers"
	"sso-auth/app/http/requests"
	"sso-auth/app/schemas"
	"sso-auth/app/services"
	"sso-auth/pkg"

	"github.com/gofiber/fiber/v2"
)

type ClientHandler struct {
	clientService *services.ClientService
}

func NewClientHandler() *ClientHandler {
	return &ClientHandler{
		clientService: services.NewClientService(),
	}
}

func (h *ClientHandler) Create(c *fiber.Ctx) error {
	data := requests.ClientRequest{}

	c.BodyParser(&data)
	pkg.NewValidator()
	err := pkg.Validate(data)

	if err != nil {
		return helpers.ResponseApiBadRequest(c, err.Error(), nil)
	}

	err = h.clientService.Create(&data)

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)
		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiCreated(c, "Client created successfully", nil)
}

func (h *ClientHandler) List(c *fiber.Ctx) error {
	page := c.Query("page")
	limit := c.Query("limit")
	sort := c.Query("sort")

	res, err := h.clientService.List(page, limit, sort)
	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)
		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiOk(c, "OK", &res)
}

func (h *ClientHandler) Detail(c *fiber.Ctx) error {
	clientId, _ := c.ParamsInt("id", 0)

	if clientId == 0 {
		return helpers.ResponseApiBadRequest(c, "Param id is required", nil)
	}

	res, err := h.clientService.Detail(uint(clientId))
	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)
		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiOk(c, "OK", res)
}

func (h *ClientHandler) Update(c *fiber.Ctx) error {
	clientId, _ := c.ParamsInt("id", 0)
	data := requests.ClientRequest{}

	c.BodyParser(&data)
	pkg.NewValidator()
	err := pkg.Validate(data)

	if err != nil {
		return helpers.ResponseApiBadRequest(c, err.Error(), nil)
	}

	if clientId == 0 {
		return helpers.ResponseApiBadRequest(c, "Param id is required", nil)
	}

	err = h.clientService.Update(uint(clientId), &data)

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)
		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiCreated(c, "Client updated successfully", nil)
}

func (h *ClientHandler) GenerateSecret(c *fiber.Ctx) error {
	clientId, _ := c.ParamsInt("id", 0)
	secret, error := h.clientService.GenerateSecret(uint(clientId))

	if error != nil {
		respErr := error.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)
		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiOk(c, "OK", secret)
}

func (h *ClientHandler) Delete(c *fiber.Ctx) error {
	clientId, _ := c.ParamsInt("id", 0)

	if clientId == 0 {
		return helpers.ResponseApiBadRequest(c, "Param id is required", nil)
	}

	err := h.clientService.Delete(uint(clientId))

	if err != nil {
		respErr := err.(*schemas.ResponseApiError)
		catchErr := helpers.CatchErrorResponseApi(respErr)
		return helpers.ResponseApiError(c, catchErr.Message, catchErr.StatusCode, nil)
	}

	return helpers.ResponseApiOk(c, "Client deleted successfully", nil)
}
