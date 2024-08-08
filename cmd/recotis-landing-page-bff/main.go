/*
This file is licensed under the Creative Commons Attribution-NonCommercial 4.0 International License.
You may obtain a copy of the license at https://creativecommons.org/licenses/by-nc/4.0/legalcode.txt
*/

package main

import (
	"log"
	"net/http"

	"recotis-landing-page-bff/api"
	"recotis-landing-page-bff/api/middlewares"
	"recotis-landing-page-bff/internal"
)

func main() {
	internal.LoadEnvironmentVariables()

	mux := http.NewServeMux()
	mux.HandleFunc("/signed-urls", api.SignedUrlsHandler)
	mux.HandleFunc("/email", api.SendEmailHandler)

	handler := middlewares.CorsMiddleware(
		middlewares.TokenMiddleware(
			middlewares.SecurityMiddleware(mux),
		),
	)

	serverPort := internal.GetServerPort()
	log.Println("Starting server on :" + serverPort)
	if err := http.ListenAndServe(":"+serverPort, handler); err != nil {
		log.Fatalf("ERROR: Failed to start server: %v", err)
	}
}
