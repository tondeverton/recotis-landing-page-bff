/*
This file is licensed under the Creative Commons Attribution-NonCommercial 4.0 International License.
You may obtain a copy of the license at https://creativecommons.org/licenses/by-nc/4.0/legalcode.txt
*/

package internal

import (
	"regexp"
)

var regexEmailMatcher = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

var regexAssetsMatcher = regexp.MustCompile("^[a-z]+(_)*[a-z]+.(jpg|png|mp4)$")

func IsAssetsValid(assets []string) bool {
	for i := 0; i < len(assets); i++ {
		asset := assets[i]
		if !regexAssetsMatcher.MatchString(asset) {
			return false
		}
	}

	return true
}

var regexSendEmailContactNameMatcher = regexp.MustCompile(`^[a-zA-ZÀ-ÖØ-öø-ÿ' .\-]{3,}$`)
var regexSendEmailSubjectMatcher = regexp.MustCompile("^.{5,}$")
var regexSendEmailMessageMatcher = regexp.MustCompile("^.{7,}$")

func IsSendEmailRequestValid(request SendEmailRequest) bool {
	if !regexSendEmailContactNameMatcher.MatchString(request.Name) ||
		!regexSendEmailSubjectMatcher.MatchString(request.Subject) ||
		!regexSendEmailMessageMatcher.MatchString(request.Message) ||
		!regexEmailMatcher.MatchString(request.Email) {
		return false
	}

	return true
}
