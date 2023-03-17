package helpers

import "github.com/dimassfeb-09/MyLibraryApp-BE.git/entity/response"

func ToWebResponse(status string, code int, msg string, data any) *response.WebResponse {
	return &response.WebResponse{
		Status: status,
		Code:   code,
		Msg:    msg,
		Data:   data,
	}
}
