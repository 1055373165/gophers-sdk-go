package main

import (
	"context"
	"fmt"
	"io"
	"os"

	gopherssdk "github.com/1055373165/gophers-sdk-go"
)

func main() {
	config := gopherssdk.NewConfiguration()
	client := gopherssdk.NewAPIClient(config)

	// Check Health
	// When we call GophersApi.CheckHealth method, it return a string
	// equals to OK if the Gophers API is running and healthy

	res, err := client.GophersAPI.CheckHealth(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GophersApi.CheckHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res.StatusCode)
	}
	// response from `CheckHealth`: string
	fmt.Fprintf(os.Stdout, "Response from `GophersApi.CheckHealth`: %v\n", res.Status)

	// Get Gophers
	res, err = client.GophersAPI.GophersGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GophersApi.GophersGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", res.StatusCode)
	}

	// response from `GophersGet`: []Gopher
	buf, _ := io.ReadAll(res.Body)
	fmt.Fprintf(os.Stdout, "Response from `GophersApi.GophersGet`: %s\n", string(buf))
}
