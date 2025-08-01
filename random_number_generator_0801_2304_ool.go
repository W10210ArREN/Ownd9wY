// 代码生成时间: 2025-08-01 23:04:39
package main

import (
    "crypto/rand"
    "encoding/binary"
    "math/big"
    "net/http"
    "revel"
)

// RandomNumberGenerator is a Revel controller for generating random numbers.
type RandomNumberGenerator struct {
    *revel.Controller
}

// Generate handles the HTTP request to generate a random number between min and max.
func (c RandomNumberGenerator) Generate(min, max int) revel.Result {
    // Validate input parameters
    if min >= max {
        return c.RenderError(http.StatusBadRequest, "Invalid range: min must be less than max")
    }

    // Generate a random number between min and max
    randomNumber, err := generateRandomNumber(min, max)
    if err != nil {
        return c.RenderError(http.StatusInternalServerError, "Failed to generate random number")
    }

    // Render the result as JSON
    return c.RenderJSON(RandomNumberResult{Number: randomNumber})
}

// generateRandomNumber generates a random number between min and max using crypto/rand for cryptographic randomness.
func generateRandomNumber(min, max int) (int, error) {
    maxInt := big.NewInt(int64(max))
    minInt := big.NewInt(int64(min))
    rangeInt := new(big.Int).Sub(maxInt, minInt)
    rangeInt = new(big.Int).Add(rangeInt, big.NewInt(1)) // inclusive range

    // Generate a random number in the range [0, rangeInt)
    randNum, err := rand.Int(rand.Reader, rangeInt)
    if err != nil {
        return 0, err
    }

    // Convert to a number in the range [min, max]
    return int(randNum.Int64()) + min, nil
}

// RandomNumberResult is a struct for rendering the random number as JSON.
type RandomNumberResult struct {
    Number int `json:"number"`
}

func init() {
    // Register the RandomNumberGenerator controller and its methods.
    revel.Router(
        (*RandomNumberGenerator)(nil),
        []string{"random/generate/{min}/{max}"},
        "Generate",
        [2]int{0, 1}, // min and max are route parameters
    )
}
