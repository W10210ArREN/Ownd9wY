// 代码生成时间: 2025-09-05 07:01:00
package main

import (
    "fmt"
    "log"
    "net/url"
    "revel"
)

// UrlValidator is a simple Revel controller that validates URL links.
type UrlValidator struct {
    revel.Controller
}

// Validate checks if the provided URL is valid or not.
func (c UrlValidator) Validate(url string) revel.Result {
    // Parse the URL to check its validity.
    parsedUrl, err := url.ParseRequestURI(url)
    if err != nil {
        // If there's an error, return a JSON response with an error message.
        return c.RenderJSON(map[string]string{
            "error": "Invalid URL format",
        })
    }

    // Check if the scheme is HTTP or HTTPS.
    if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
        return c.RenderJSON(map[string]string{
            "error": "URL scheme must be HTTP or HTTPS",
        })
    }

    // Return a JSON response indicating the URL is valid.
    return c.RenderJSON(map[string]bool{
        "valid": true,
    })
}

func init() {
    // Add the validate method to the UrlValidator controller with a route.
    revel.Router(
        (*UrlValidator)(nil),
        "validate",
        "POST /validate",
        func(c *UrlValidator, url string) revel.Result { return c.Validate(url) },
    )
}

func main() {
    // Initialize the Revel framework and start the server.
    revel.Run()
}