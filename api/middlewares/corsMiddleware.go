/*
This file is licensed under the Creative Commons Attribution-NonCommercial 4.0 International License.
You may obtain a copy of the license at https://creativecommons.org/licenses/by-nc/4.0/legalcode.txt
*/

package middlewares

import "net/http"

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Access-Control-Allow-Origin", "*")
		response.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		response.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if request.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(response, request)
	})
}
