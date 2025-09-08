// 代码生成时间: 2025-09-08 16:46:26
 * This tool allows users to calculate hashes of strings using different algorithms like SHA256.
 *
 */

package main

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "revel"
)

// HashCalculatorController handles the request to calculate hash values.
type HashCalculatorController struct {
    *revel.Controller
}

// Index renders the form for inputting the string to hash.
func (c HashCalculatorController) Index() revel.Result {
    return c.Render()
}

// CalculateHash handles the POST request to calculate the hash of the input string.
func (c HashCalculatorController) CalculateHash(input string) revel.Result {
    hash := sha256.Sum256([]byte(input))
    return c.RenderJSON(HashingResult{Hash: hex.EncodeToString(hash[:]), Original: input}) // Render the result as JSON.
}

// HashingResult is the struct to hold the result of the hash calculation.
type HashingResult struct {
    Hash      string `json:"hash"` // The calculated hash value.
    Original  string `json:"original"` // The original string that was hashed.
}

// main function to initialize and run the Revel application.
func main() {
    revel.Filters = []revel.Filter{
        revel.Router,
        revel.Params,
        revel.SessionFilter,
        revel.FlashFilter,
        revel.ValidationFilter,
        revel.I18nFilter,
        // Add more filters here if needed.
    }
    revel.RunProd()
}
