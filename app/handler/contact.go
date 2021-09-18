package handler

import (
	"net/http"
	"fmt"
	"dena-hackathon21/api_model"
	"dena-hackathon21/repository"
	"dena-hackathon21/sql_handler"
	"github.com/labstack/echo"
)

type ContactHandler struct {
	sqlHandler *sql_handler.SQLHandler
}

func NewContactHandler(sqlHandler *sql_handler.SQLHandler) *ContactHandler {
	return &ContactHandler{
		sqlHandler: sqlHandler,
	}
}

func (ch *ContactHandler) Send(c echo.Context) error {
	// TODO 環境変数から取りたい
	sqlHandler, err := sql_handler.NewHandler("user:password@tcp(db:3306)/test_database")
	if err != nil {
		return c.String(500, fmt.Sprintf("db scan error: %s", err.Error()))
	}

	//TODO 後でjwt使った関数に置き換える
	var sender_id uint64 = 1

	req := new(api_model.SendContactRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	for _, receiver_id := range req.RequestUseIDList {
		contactRepository := repository.NewContactRepository(sqlHandler)
		err := contactRepository.SendContact(c.Request().Context(), sender_id, receiver_id, req.Message)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("POST /contact Error: %s", err.Error()))
		}
	}
	return c.String(http.StatusCreated, "OK")
}