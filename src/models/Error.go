package models

type Error struct {
	ResponseCode        string `json:"responsecode"`
	ResponseDescription string `json:"responsedescription"`
}
