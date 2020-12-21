package errs

import (
	"github.com/mauwahid/kafman/internal/presenter/dto"
	"net/http"
	"strconv"
)

var RespBadRequest = dto.PubResponse{
	Status:  strconv.Itoa(http.StatusBadRequest),
	Message: "invalid parameter",
}

func Error(data interface{}, err error) (pubRes dto.PubResponse) {
	pubRes.Status = "99"
	pubRes.Message = err.Error()
	pubRes.Data = data
	return
}

func Success(key, data string) (pubRes dto.PubResponse) {
	pubRes.Status = "00"
	pubRes.Message = "success"
	pubRes.Data = struct {
		Key     string `json:"key"`
		Message string `json:"message"`
	}{
		Key:     key,
		Message: data,
	}
	return
}
