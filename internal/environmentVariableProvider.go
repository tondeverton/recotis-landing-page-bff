/*
This file is licensed under the Creative Commons Attribution-NonCommercial 4.0 International License.
You may obtain a copy of the license at https://creativecommons.org/licenses/by-nc/4.0/legalcode.txt
*/

package internal

import (
	"os"
	"strconv"
)

const (
	smtpUsernameEnvName        = "SMTP_USERNAME"
	smtpPasswordEnvName        = "SMTP_PASSWORD"
	smtpServerEnvName          = "SMTP_SERVER"
	smtpServerPortEnvName      = "SMTP_SERVER_PORT"
	emailTargetEnvName         = "EMAIL_TARGET"
	tokenPasswordEnvName       = "TOKEN_PASSWORD"
	serverPortEnvName          = "SERVER_PORT"
	bucketNameEnvName          = "BUCKET_NAME"
	bucketBasePathEnvName      = "BUCKET_BASE_PATH"
	signedUrlExpirationEnvName = "SIGNED_URL_EXPIRATION"
)

var environmentVariables map[string]*string

func GetServerPort() string {
	return getEnv(serverPortEnvName)
}

func GetTokenPassword() string {
	return getEnv(tokenPasswordEnvName)
}

func GetSmtpUsername() string {
	return getEnv(smtpUsernameEnvName)
}

func GetSmtpPassword() string {
	return getEnv(smtpPasswordEnvName)
}

func GetSmtpServer() string {
	return getEnv(smtpServerEnvName)
}

func GetSmtpServerPort() string {
	return getEnv(smtpServerPortEnvName)
}

func GetEmailTarget() string {
	return getEnv(emailTargetEnvName)
}

func GetBucketName() string {
	return getEnv(bucketNameEnvName)
}

func GetBucketBasePath() string {
	return getEnv(bucketBasePathEnvName)
}

func GetSignedUrlExpiration() int {
	return getEnvAsInt(signedUrlExpirationEnvName)
}

func LoadEnvironmentVariables() {
	environmentVariables = make(map[string]*string)

	environmentVariables[smtpUsernameEnvName] = nil
	environmentVariables[smtpPasswordEnvName] = nil
	environmentVariables[smtpServerEnvName] = nil
	environmentVariables[smtpServerPortEnvName] = nil
	environmentVariables[emailTargetEnvName] = nil
	environmentVariables[tokenPasswordEnvName] = nil

	defaultServerPort := "8080"
	environmentVariables[serverPortEnvName] = &defaultServerPort
	defaultBucketName := "www.recotis.com"
	environmentVariables[bucketNameEnvName] = &defaultBucketName
	defaultBucketBasePath := "protected-assets"
	environmentVariables[bucketBasePathEnvName] = &defaultBucketBasePath
	defaultSignedUrlExpiration := "180"
	environmentVariables[signedUrlExpirationEnvName] = &defaultSignedUrlExpiration

	for envName, value := range environmentVariables {
		if value != nil {
			continue
		}
		if _, exists := os.LookupEnv(envName); !exists {
			panic("Error: required variable (" + envName + ") doesn't exists")
		}
	}
}

func getEnv(envName string) string {
	if value, exists := os.LookupEnv(envName); exists {
		return value
	}

	if value := environmentVariables[envName]; value != nil {
		return *value
	}

	panic("Error: variable (" + envName + ") doesn't exists")
}

func getEnvAsInt(envName string) int {
	value := getEnv(envName)

	valueAsInt, err := strconv.Atoi(value)
	if err == nil {
		return valueAsInt
	}

	panic("Error: value from (" + envName + ") can't be converted to int: " + err.Error())
}
