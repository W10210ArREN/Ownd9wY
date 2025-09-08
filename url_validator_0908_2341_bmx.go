// 代码生成时间: 2025-09-08 23:41:49
package controllers

import (
    "encoding/json"
    "net/http"
    "net/url"
    "revel"
)

// UrlValidator is the controller for checking URL validity.
type UrlValidator struct {
    revel.Controller
}

// Check validates the provided URL and returns a JSON response indicating its validity.
func (c UrlValidator) Check() revel.Result {
    // Get the URL from the request query parameter.
    urlStr := c.Params.Get("url")
    if urlStr == "" {
        // Return an error if the URL is empty.
        return c.RenderJSON(ErrorResponse{"error": "URL parameter is missing"})
    }

    // Parse the URL to check its validity.
    parsedUrl, err := url.ParseRequestURI(urlStr)
    if err != nil {
        // Return an error if the URL parsing fails.
        return c.RenderJSON(ErrorResponse{"error": "Invalid URL format"})
    }

    // Check if the scheme is valid.
    if parsedUrl.Scheme == "" || (parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https") {
        return c.RenderJSON(ErrorResponse{"error": "Unsupported URL scheme"})
    }

    // Return a success response indicating the URL is valid.
    return c.RenderJSON(SuccessResponse{"message": "The URL is valid"})
}

// ErrorResponse represents an error response.
type ErrorResponse struct {
    Error string `json:"error"`
}

// SuccessResponse represents a success response.
type SuccessResponse struct {
    Message string `json:"message"`
}
