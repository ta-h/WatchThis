package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ProblemDetail struct {
	Type     string `json:"type"`               // A URI reference that identifies the problem type.
	Title    string `json:"title"`              // A short, human-readable summary of the problem type.
	Status   int    `json:"status"`             // The HTTP status code associated with the problem.
	Detail   string `json:"detail"`             // A human-readable explanation specific to this occurrence of the problem.
	Instance string `json:"instance,omitempty"` // A URI reference that identifies the specific occurrence of the problem (optional).
}

func (problemDetail *ProblemDetail) writeToResponse(res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/problem+json")
	res.WriteHeader(problemDetail.Status)
	json.NewEncoder(res).Encode(problemDetail)
}

func respondWithParsingError(res http.ResponseWriter, err error) {
	problem := ProblemDetail{
		Type:   "urn:watchlist:problem:invalid-json",
		Title:  "Invalid JSON format",
		Status: http.StatusBadRequest,
		Detail: fmt.Sprintf("Error parsing JSON: %v", err),
	}
	problem.writeToResponse(res)
}

func respondWithValidationError(res http.ResponseWriter, details string) {
	problem := ProblemDetail{
		Type:   "urn:watchlist:problem:invalid-payload",
		Title:  "Invalid data",
		Status: http.StatusBadRequest,
		Detail: details,
	}
	problem.writeToResponse(res)
}

func respondWithNotFound(res http.ResponseWriter) {
	problem := ProblemDetail{
		Type:   "urn:watchlist:problem:404-not-found",
		Title:  "404 not found",
		Status: http.StatusNotFound,
		Detail: "The requested resource is invalid or does not exist",
	}
	problem.writeToResponse(res)
}

func writeJSONResponse(res http.ResponseWriter, watchList *WatchList, statusCode int) {
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)

	if err := json.NewEncoder(res).Encode(&watchList); err != nil {
		http.Error(res, "Failed to encode response", http.StatusInternalServerError)
	}
}
