package models

type CompletionsHandlerResponse struct {
	BaseResponse
	Data string `json:"data"`
}
