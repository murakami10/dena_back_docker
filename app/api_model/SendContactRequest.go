package api_model

type SendContactRequest struct {
	RequestUseIDList []uint64 `json:"request_user_id_list"`
	Message string `json:"message"`
}