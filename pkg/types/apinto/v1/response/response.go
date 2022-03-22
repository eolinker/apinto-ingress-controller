package response

import "encoding/json"

type Response struct {
	Name       string `json:"name"`
	Profession string `json:"profession"`
	Driver     string `json:"driver"`
	ID         string `json:"id"`
	Create     string `json:"create"`
	Update     string `json:"update"`
}

type RawResponse json.RawMessage
