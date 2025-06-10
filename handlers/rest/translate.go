package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/comecacahuates/test-go/translation"
)

type Response struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

const defaultLanguage = "english"

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = defaultLanguage
	}
	word := strings.ReplaceAll(r.URL.Path, "/", "")
	translation := translation.Translate(word, language)

	if translation == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	response := Response{
		Language:    language,
		Translation: translation,
	}
	if err := encoder.Encode(response); err != nil {
		panic("unable to encode response")
	}
}
