package apinto

import (
	"encoding/json"
	"time"
)

type getResponse json.RawMessage

// list接口待定
type listResponse []json.RawMessage

type baseResponse struct {
	Name       string        `json:"name"`
	Profession string        `json:"profession"`
	Driver     string        `json:"driver"`
	ID         string        `json:"id"`
	Create     time.Duration `json:"create"`
	Update     time.Duration `json:"update"`
}
type updateResponse baseResponse
type deleteResponse baseResponse
type createResponse baseResponse
