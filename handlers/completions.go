package handlers

import (
	"chat-openai/application"
	"chat-openai/models"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func CompletionsHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := application.GetClient()
	w.Header().Set("Content-Type", "application/json")
	var failResponse models.BaseResponse

	var messageRequest models.MessageSendRequest
	bodyError := json.NewDecoder(r.Body).Decode(&messageRequest)
	if bodyError != nil {

		failResponse.Success = false
		failResponse.Message = "invalid request, check body"

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(failResponse)
		return
	}

	response, err := application.GetCompletions(client, messageRequest.Content)
	if err != nil {

		failResponse.Success = false
		failResponse.Message = err.Error()

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(failResponse)
		return
	}

	var successResponse models.CompletionsHandlerResponse
	successResponse.Success = true
	successResponse.Message = "ok"
	successResponse.Data = response.Choices[0].Message.Content

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(successResponse)

}
