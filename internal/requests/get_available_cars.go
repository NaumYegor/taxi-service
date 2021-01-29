package requests

type GetAvailableCarsResponseAttributes struct {
	ID     int32  `json:"id"`
	Driver int32  `json:"driver_id"`
	Model  string `json:"model"`
	Number string `json:"number"`
}

type GetAvailableCarsResponseList []GetAvailableCarsResponseAttributes
