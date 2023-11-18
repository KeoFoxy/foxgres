package handler

type AccessLevel struct {
	AccessLevel int `json:"access_level,omitempty"`
} // @name AccessLevel

type ResponseErr struct {
	Message string `json:"message,omitempty"`
} // @name ResponseErr
