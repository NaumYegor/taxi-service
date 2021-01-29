package requests

type GetAvailableOrdersResponseAttributes struct {
	ID   int32  `json:"id"`
	Info string `json:"info"`
}

type GetAvailableOrdersResponseList []GetAvailableOrdersResponseAttributes
