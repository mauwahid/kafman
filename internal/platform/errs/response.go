package errs

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mauwahid/kafman/internal/presenter/dto"
)

var RespBadRequest = dto.PubResponse{
	Status:  strconv.Itoa(http.StatusBadRequest),
	Message: "invalid parameter",
}

func Error(data []byte, err error) (pubRes dto.PubResponse) {
	pubRes.Status = "99"
	pubRes.Message = err.Error()
	pubRes.Data = getUnmarshal(data)
	return
}

func Success(key string, data []byte) (pubRes dto.PubResponse) {
	pubRes.Status = "00"
	pubRes.Message = "success"
	pubRes.Data = struct {
		Key     string      `json:"key"`
		Message interface{} `json:"message"`
	}{
		Key:     key,
		Message: getUnmarshal(data),
	}
	return
}

func getUnmarshal(data []byte) (dataResp interface{}) {
	rawIn := json.RawMessage(data)
	bytes, err := rawIn.MarshalJSON()
	if err != nil {
		dataResp = string(data)
	}

	err = json.Unmarshal(bytes, &dataResp)
	if err != nil {
		dataResp = string(data)
	}
	return dataResp
}
