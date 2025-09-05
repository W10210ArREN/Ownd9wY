// 代码生成时间: 2025-09-05 11:53:25
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "revel"
)

// WebScraper struct defines the properties required for web scraping
type WebScraper struct {
    // The revel controller structure
    *revel.Controller
}

// GetWebContent is the action that handles the HTTP request to scrape web content
func (c WebScraper) GetWebContent(url string) revel.Result {
    resp, err := http.Get(url)
    if err != nil {
        return c.RenderError(err)
    }
    defer resp.Body.Close()

    // Ensure the request was successful
    if resp.StatusCode != http.StatusOK {
        return c.RenderError(fmt.Errorf("Failed to retrieve content: Status code %d", resp.StatusCode))
    }

    // Read the body of the response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return c.RenderError(err)
    }

    // Render the web content as a string
    return c.RenderText(string(body))
}

func init() {
    // Register the controller actions
    revel.RegisterController((*WebScraper)(nil), "WebScraper")
}

func main() {
    // Initialize the Revel framework
    revel.Run()
}
