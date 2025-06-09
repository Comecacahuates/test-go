package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/comecacahuates/test-go/handlers/rest"
)

func TestTranslateAPI(t *testing.T) {
	// Arrange
	type TestCase struct {
		endpoint            string
		statusCode          int
		expectedLanguage    string
		expectedTranslation string
	}
	testCases := []TestCase{
		{
			endpoint:            "/hello",
			statusCode:          http.StatusOK,
			expectedLanguage:    "english",
			expectedTranslation: "hello",
		},
		{
			endpoint:            "/hello?language=german",
			statusCode:          http.StatusOK,
			expectedLanguage:    "german",
			expectedTranslation: "hallo",
		},
		{
			endpoint:            "/hello?language=dutch",
			statusCode:          http.StatusNotFound,
			expectedLanguage:    "",
			expectedTranslation: "",
		},
	}

	handler := http.HandlerFunc(rest.TranslateHandler)

	for _, testCase := range testCases {
		responseRecorder := httptest.NewRecorder()
		request, _ := http.NewRequest("GET", testCase.endpoint, nil)

		// Act
		handler.ServeHTTP(responseRecorder, request)

		// Assert
		if responseRecorder.Code != testCase.statusCode {
			t.Errorf("expected status 200 but got %d", responseRecorder.Code)
		}

		var response rest.Response
		json.Unmarshal(responseRecorder.Body.Bytes(), &response)

		if response.Language != testCase.expectedLanguage {
			t.Errorf("expected language '%s' but got '%s'", testCase.expectedLanguage, response.Language)
		}

		if response.Translation != testCase.expectedTranslation {
			t.Errorf("expected translation '%s' but got '%s'", testCase.expectedTranslation, response.Translation)
		}
	}

}
