package response

import (
	"encoding/json"
	"time"
)

type Response struct {
	Name       string        `json:"name"`
	Profession string        `json:"profession"`
	Driver     string        `json:"driver"`
	ID         string        `json:"id"`
	Create     time.Duration `json:"create"`
	Update     time.Duration `json:"update"`
}

type RawResponse json.RawMessage
