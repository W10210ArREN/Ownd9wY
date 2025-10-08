// 代码生成时间: 2025-10-08 17:41:47
package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "os"
    "revel"
)

// HashCalculator is a simple Revel controller for hash calculations.
type HashCalculator struct {
    revel.Controller
}

// Index renders the form for entering a string to hash.
func (c *HashCalculator) Index() revel.Result {
    return c.Render()
}

// Calculate hashes the input string using SHA-256 and returns the result.
func (c *HashCalculator) Calculate(input string) revel.Result {
    // Generate the hash of the provided input string.
    hash := sha256.Sum256([]byte(input))

    // Convert the hash to a hex string.
    hexHash := hex.EncodeToString(hash[:])

    // Return the result as a JSON response.
    return c.RenderJSON(map[string]string{
        "input": input,
        "hash": hexHash,
    })
}

// CheckHash compares the provided hash with the hash of the input string.
func (c *HashCalculator) CheckHash(input, hash string) revel.Result {
    // Check if the input and the provided hash are valid.
    if input == "" || hash == "" {
        return c.RenderJSON(map[string]string{
            "error": "Input and hash cannot be empty.",
        })
    }

    // Generate the hash of the provided input string.
    calculatedHash := sha256.Sum256([]byte(input))
    hexCalculatedHash := hex.EncodeToString(calculatedHash[:])

    // Compare the calculated hash with the provided hash.
    if hexCalculatedHash == hash {
        return c.RenderJSON(map[string]bool{
            "match": true,
        })
    } else {
        return c.RenderJSON(map[string]bool{
            "match": false,
        })
    }
}

func init() {
    // Add route for the Index action.
    revel.Router["/"].Name = "hashcalculator.index"
    // Add route for the Calculate action.
    revel.Router["/hashcalculator/calculate/:input"].Name = "hashcalculator.calculate"
    // Add route for the CheckHash action.
    revel.Router["/hashcalculator/checkhash/:input/:hash"].Name = "hashcalculator.checkhash"
}
