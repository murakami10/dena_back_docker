package handler

import (
	"dena-hackathon21/api_model"
	"dena-hackathon21/repository"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ContactHandler struct {
	contactRepository *repository.ContactRepository
}

func NewContactHandler(contactRepository *repository.ContactRepository) *ContactHandler {
	return &ContactHandler{
		contactRepository: contactRepository,
	}
}

func (ch *ContactHandler) Send(c echo.Context) error {
	//TODO 後でjwt使った関数に置き換える
	var sender_id uint64 = 1

	req := new(api_model.SendContactRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	for _, receiver_id := range req.RequestUseIDList {
		err := ch.contactRepository.SendContact(c.Request().Context(), sender_id, receiver_id, req.Message)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("POST /contact Error: %s", err.Error()))
		}
	}
	return c.String(http.StatusCreated, "OK")
}
