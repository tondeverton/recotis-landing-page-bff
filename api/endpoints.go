package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"recotis-landing-page-bff/internal"
)

func SignedUrlsHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		response.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req internal.SignedURLRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		log.Println("Error:", err)
		return
	}

	if !internal.IsAssetsValid(req.Assets) {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	urls, expiresIn, err := internal.GetUrlsByAssets(req.Assets)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		log.Println("Error:", err)
		return
	}

	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(internal.SignedURLResponse{URLs: urls, ExpiresIn: expiresIn})
}

func SendEmailHandler(response http.ResponseWriter, request *http.Request) {
	var req internal.SendEmailRequest

	if err := json.NewDecoder(request.Body).Decode(&req); err != nil {
		response.WriteHeader(http.StatusBadRequest)
		log.Println("Error:", err)
		return
	}

	if !internal.IsSendEmailRequestValid(req) {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	body := fmt.Sprintf("Name: %v\r\nEmail: %v\r\nMessage: %v", req.Name, req.Email, req.Message)

	if err := internal.SendEmail(req.Subject, body); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		log.Println("Error:", err.Error())
		return
	}

	response.WriteHeader(http.StatusOK)
}
