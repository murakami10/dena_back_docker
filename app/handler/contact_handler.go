package handler

import (
	"dena-hackathon21/api_model"
	"dena-hackathon21/repository"
	"dena-hackathon21/auth"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ContactHandler struct {
	contactRepository *repository.ContactRepository
	jwtHandler        *auth.JWTHandler
}

func NewContactHandler(contactRepository *repository.ContactRepository, jwtHandler *auth.JWTHandler) *ContactHandler {
	return &ContactHandler{
		contactRepository: contactRepository,
		jwtHandler:        jwtHandler,
	}
}

func (ch *ContactHandler) Send(c echo.Context) error {
	// ユーザIDを取得
	cookie, err := c.Cookie("token")
	if err != nil {
		return c.String(500, err.Error())
	}
	token := cookie.Value
	sender_id, err := ch.jwtHandler.GetUserIDFromToken(token)

	// リクエスト取得
	req := new(api_model.SendContactRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	for _, receiver_id := range req.RequestUseIDList {
    // コンタクトを登録
		err := ch.contactRepository.SendContact(c.Request().Context(), sender_id, receiver_id, req.Message)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("POST /api/contact Error: %s", err.Error()))
		}
    
		// ルームが無い場合は新規作成
		is_exist, err := ch.contactRepository.IsExistRoom(c.Request().Context(), sender_id, receiver_id)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("POST /api/contact Error: %s", err.Error()))
		}
		fmt.Println("is_exist:", is_exist)

		if !is_exist {
			err := ch.contactRepository.CreateRoom(c.Request().Context(), sender_id, receiver_id, req.Message)
			if err != nil {
				return c.String(http.StatusInternalServerError, fmt.Sprintf("POST /api/contact Error: %s", err.Error()))
			}
		}
	}
	return c.String(http.StatusCreated, "Created")
}

func (ch *ContactHandler) Get(c echo.Context) error {
	// ユーザIDを取得
	cookie, err := c.Cookie("token")
	if err != nil {
		return c.String(500, err.Error())
	}
	token := cookie.Value
	user_id, err := ch.jwtHandler.GetUserIDFromToken(token)

	// 受信ユーザ、受信メッセージ、その受信日時を取得
	receivedContactItemList, err := ch.contactRepository.GetReceivedContact(c.Request().Context(), user_id)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("GET /api/contact Error (receivedContactItemList): %s", err.Error()))
	}

	// ルームIDを取得
	for index, receivedContactItem := range receivedContactItemList {
		room_id, err := ch.contactRepository.GetRoomId(c.Request().Context(), user_id, receivedContactItem.SenderID)
		if err != nil {
			return c.String(http.StatusInternalServerError, fmt.Sprintf("GET /api/contact Error (room_id): %s", err.Error()))
		}
		receivedContactItemList[index].RoomID = room_id
	}

	// 受信ユーザ、最新のメッセージ、ルームID、受信日時を取得
	pastContactItemList, err := ch.contactRepository.GetPastContact(c.Request().Context(), user_id)
	if err != nil {
		return c.String(http.StatusInternalServerError, fmt.Sprintf("GET /api/contact Error (pastContactItemList): %s", err.Error()))
	}

	// レスポンスを作成
	var res api_model.GetContactResponse
	res.ReceivedContactList = receivedContactItemList
	res.PastMessageList = pastContactItemList

	return c.JSON(http.StatusOK, res)
}
