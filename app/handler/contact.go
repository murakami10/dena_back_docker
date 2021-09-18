package handler

import (
	"net/http"
	"fmt"
	"dena-hackathon21/api_model"
	"dena-hackathon21/repository"
	"github.com/labstack/echo"
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
			return c.String(http.StatusInternalServerError, fmt.Sprintf("POST /api/contact Error: %s", err.Error()))
		}
	}
	return c.String(http.StatusCreated, "Created")
}

func (ch *ContactHandler) Get(c echo.Context) error {
	//TODO 後でjwt使った関数に置き換える
	var user_id uint64 = 1

	// 受信ユーザ、受信メッセージ、受信日時を取得
	contactItems, err := ch.contactRepository.GetReceivedContact(c.Request().Context(), user_id)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("GET /api/contact Error: %s", err.Error()))
	}

	var response := &api_model.GetContactResponse{
	}

	return c.JSON(http.StatusOK, "OK")
}