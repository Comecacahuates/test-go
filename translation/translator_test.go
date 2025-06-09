package translation_test

import (
	"testing"

	"github.com/comecacahuates/test-go/translation"
)

func TestTranslate(t *testing.T) {
	// Arrange
	type TestCase struct {
		word     string
		lang     string
		expTrans string
	}
	testCases := []TestCase{
		{word: "hello", lang: "english", expTrans: "hello"},
		{word: "hello", lang: "german", expTrans: "hallo"},
		{word: "hello", lang: "finnish", expTrans: "hei"},
		{word: "hello", lang: "dutch", expTrans: ""},
		{word: "bye", lang: "dutch", expTrans: ""},
		{word: "hello", lang: "German", expTrans: "hallo"},
		{word: "Hello", lang: "german", expTrans: "hallo"},
		{word: "hello ", lang: "german", expTrans: "hallo"},
	}

	// Act
	for _, tc := range testCases {
		actTrans := translation.Translate(tc.word, tc.lang)

		// Assert
		if actTrans != tc.expTrans {
			t.Errorf("expected '%s' to be '%s' from '%s' but got '%s'\n", tc.word, tc.lang, tc.expTrans, actTrans)
		}
	}
}
