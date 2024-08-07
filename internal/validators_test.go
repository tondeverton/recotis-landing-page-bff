package internal

import "testing"

func TestIsAssetsValid_GivenEmptyArray_ShouldReturnsTrue(t *testing.T) {
	assets := []string{}
	if !IsAssetsValid(assets) {
		t.Errorf("TestIsAssetsValid_GivenEmptyArray_ShouldReturnsTrue: expect true for %v", assets)
	}
}

func TestIsAssetsValid_GivenArrayContainingOneFileFromValidExtension_ShouldReturnsTrue(t *testing.T) {
	properties := [][]string{
		{"file.jpg"},
		{"file.png"},
		{"file.mp4"},
	}

	for _, property := range properties {
		if !IsAssetsValid(property) {
			t.Errorf("TestIsAssetsValid_GivenArrayContainingOneFileFromValidExtension_ShouldReturnsTrue: expect true for %v", property)
		}
	}
}

func TestIsAssetsValid_GivenArrayContainingThreeFilesFromValidExtensions_ShouldReturnsTrue(t *testing.T) {
	properties := [][]string{
		{"file.jpg", "file.png", "file.mp4"},
		{"file.png", "file.jpg", "file.mp4"},
		{"file.mp4", "file.png", "file.jpg"},
	}

	for _, property := range properties {
		if !IsAssetsValid(property) {
			t.Errorf("TestIsAssetsValid_GivenArrayContainingThreeFilesFromValidExtensions_ShouldReturnsTrue: expect true for %v", property)
		}
	}
}

func TestIsAssetsValid_GivenArrayContainingOnePDFFile_ShouldReturnsFalse(t *testing.T) {
	assets := []string{"file.pdf"}
	if IsAssetsValid(assets) {
		t.Errorf("TestIsAssetsValid_GivenArrayContainingOnePDFFile_ShouldReturnsFalse: expect false for %v", assets)
	}
}

func TestIsAssetsValid_GivenArrayContainingTwoFilesFromValidExtensionsAndOneInvalid_ShouldReturnsFalse(t *testing.T) {
	properties := [][]string{
		{"file.pdf", "file.png", "file.mp4"},
		{"file.png", "file.pdf", "file.mp4"},
		{"file.mp4", "file.png", "file.pdf"},
	}

	for _, property := range properties {
		if IsAssetsValid(property) {
			t.Errorf("TestIsAssetsValid_GivenArrayContainingTwoFilesFromValidExtensionsAndOneInvalid_ShouldReturnsFalse: expect false for %v", property)
		}
	}
}

func getValidSendEmailRequest() SendEmailRequest {
	return SendEmailRequest{
		Name:    "Name",
		Subject: "Subject",
		Message: "Big message",
		Email:   "email@domain.com",
	}
}

func TestIsSendEmailRequestValid_GivenValidRequest_ShouldReturnsTrue(t *testing.T) {
	request := getValidSendEmailRequest()

	if !IsSendEmailRequestValid(request) {
		t.Errorf("TestIsSendEmailRequestValid_GivenValidRequest_ShouldReturnsTrue: expect true for %v", request)
	}
}

func TestIsSendEmailRequestValid_GivenRequestWithShortName_ShouldReturnsFalse(t *testing.T) {
	request := getValidSendEmailRequest()
	request.Name = "Na"

	if IsSendEmailRequestValid(request) {
		t.Errorf("TestIsSendEmailRequestValid_GivenRequestWithShortName_ShouldReturnsFalse: expect false for %v", request)
	}
}

func TestIsSendEmailRequestValid_GivenRequestWithNameContainingNumbers_ShouldReturnsFalse(t *testing.T) {
	request := getValidSendEmailRequest()
	request.Name = "Name 123 Surname"

	if IsSendEmailRequestValid(request) {
		t.Errorf("TestIsSendEmailRequestValid_GivenRequestWithNameContainingNumbers_ShouldReturnsFalse: expect false for %v", request)
	}
}

func TestIsSendEmailRequestValid_GivenRequestWithShortSubject_ShouldReturnsFalse(t *testing.T) {
	request := getValidSendEmailRequest()
	request.Subject = "Subj"

	if IsSendEmailRequestValid(request) {
		t.Errorf("TestIsSendEmailRequestValid_GivenRequestWithShortSubject_ShouldReturnsFalse: expect false for %v", request)
	}
}

func TestIsSendEmailRequestValid_GivenRequestWithShortMessage_ShouldReturnsFalse(t *testing.T) {
	request := getValidSendEmailRequest()
	request.Message = "Messag"

	if IsSendEmailRequestValid(request) {
		t.Errorf("TestIsSendEmailRequestValid_GivenRequestWithShortMessage_ShouldReturnsFalse: expect false for %v", request)
	}
}

func TestIsSendEmailRequestValid_GivenRequestWithInvalidEmail_ShouldReturnsFalse(t *testing.T) {
	request := getValidSendEmailRequest()
	request.Email = "email@mail"

	if IsSendEmailRequestValid(request) {
		t.Errorf("TestIsSendEmailRequestValid_GivenRequestWithInvalidEmail_ShouldReturnsFalse: expect false for %v", request)
	}
}

func TestIsSendEmailRequestValid_GivenRequestWithInvalidNameSubjectMessageEmail_ShouldReturnsFalse(t *testing.T) {
	request := SendEmailRequest{
		Name:    "Na",
		Subject: "Subj",
		Message: "Messag",
		Email:   "email@mail",
	}

	if IsSendEmailRequestValid(request) {
		t.Errorf("TestIsSendEmailRequestValid_GivenRequestWithInvalidNameSubjectMessageEmail_ShouldReturnsFalse: expect false for %v", request)
	}
}
