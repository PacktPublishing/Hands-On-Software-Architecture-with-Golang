package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// CountVowels counts  vowels in  strings.
type VowelsService interface {
	Count(context.Context, string) int
}

// VowelsService is a concrete implementation of VowelsService
type VowelsServiceImpl struct{}

var vowels = map[rune]bool{
	'a': true,
	'e': true,
	'i': true,
	'o': true,
	'u': true,
}

func (VowelsServiceImpl) Count(_ context.Context, s string) int {
	count := 0
	for _, c := range s {
		if _, ok := vowels[c]; ok {
			count++
		}
	}

	return count
}

// ErrEmpty is returned when an input string is empty.
var ErrEmpty = errors.New("empty string")

// For each method, we define request and response structs
type countVowelsRequest struct {
	Input string `json:"input"`
}

type countVowelsResponse struct {
	Result int `json:"result"`
}

// An endpoint represents a single RPC  in the service interface
func makeEndpoint(svc VowelsService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countVowelsRequest)
		result := svc.Count(ctx, req.Input)
		return countVowelsResponse{result}, nil
	}
}

// Transports expose the service to the network.
func main() {
	svc := VowelsServiceImpl{}

	countHandler := httptransport.NewServer(
		makeEndpoint(svc),
		decodecountVowelsRequest,
		encodeResponse,
	)

	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func decodecountVowelsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countVowelsRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
