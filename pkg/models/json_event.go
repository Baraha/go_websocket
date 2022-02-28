package models

type Json_event struct {
	Category string      `json:"category"`
	Message  interface{} `json:"message"`
}
