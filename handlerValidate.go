package main

// import (
// 	"encoding/json"
// 	"net/http"
// 	"strings"
// )

// type parameters struct {
// 	Body string `json:"body"`
// }

// func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
// 	type returnVals struct {
// 		CleanedBody string `json:"cleaned_body"`
// 	}
// 	decoder := json.NewDecoder(r.Body)
// 	params := parameters{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "error decoding parameters", err)
// 		return
// 	}

// 	if len(params.Body) > 140 {
// 		respondWithError(w, http.StatusBadRequest, "Chirp is too long", err)
// 		return
// 	}

// 	cleanedParams := profanityFilter(params)

// 	respondWithJSON(w, http.StatusOK, returnVals{
// 		CleanedBody: cleanedParams.Body,
// 	})
// }

// func profanityFilter(params parameters) parameters {
// 	theProfane := []string{
// 		"kerfuffle",
// 		"sharbert",
// 		"fornax",
// 	}

// 	words := strings.Fields(params.Body)
// 	for i, word := range words {
// 		lowercase := strings.ToLower(word)
// 		for _, profanity := range theProfane {
// 			if lowercase == profanity {
// 				words[i] = "****"
// 				continue
// 			}
// 		}
// 	}
// 	cleanedBody := strings.Join(words, " ")
// 	params.Body = cleanedBody
// 	return params
// }
