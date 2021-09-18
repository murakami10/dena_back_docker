package api_model

type SendContactRequest struct {
	AddressList []uint64 `json:"address_list"`
	Message string `json:"message"`
}