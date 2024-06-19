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

	return helpers.ResponseApiCreated(c, "Client created successfullt", nil)
}
