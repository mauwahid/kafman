package dto

type PubRequest struct {
	Topic   string
	Message []byte
}

type PubResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (p PubRequest) Validate() bool {
	if p.Topic == "" || p.Message == nil {
		return false
	}
	return true
}
