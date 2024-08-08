/*
This file is licensed under the Creative Commons Attribution-NonCommercial 4.0 International License.
You may obtain a copy of the license at https://creativecommons.org/licenses/by-nc/4.0/legalcode.txt
*/

package middlewares

import "net/http"

func SecurityMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		response.Header().Set("Content-Security-Policy", "default-src 'self'")
		response.Header().Set("X-Content-Type-Options", "nosniff")
		response.Header().Set("X-Frame-Options", "DENY")
		response.Header().Set("X-XSS-Protection", "1; mode=block")
		response.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		response.Header().Set("Referrer-Policy", "no-referrer-when-downgrade")
		response.Header().Set("Feature-Policy", "geolocation 'self'; microphone 'self'")

		next.ServeHTTP(response, request)
	})
}
