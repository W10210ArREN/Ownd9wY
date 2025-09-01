// 代码生成时间: 2025-09-01 15:30:25
package main

import (
    "crypto/sha1"
    "encoding/hex"
    "github.com/revel/revel"
    "strings"
)

// HashCalculator is the app module
var HashCalculator = revel.NewModule("
    
type HashCalculator struct {
    *revel.Controller
}

// CalculateSHA1 calculates SHA1 hash of a given string and returns the hash in hexadecimal format
func (c HashCalculator) CalculateSHA1(input string) revel.Result {
    if strings.TrimSpace(input) == "" {
        return c.RenderError(revel.NewError(revel.StatusInternalServerError, "Input string cannot be empty"))
    }
    
    // Calculate SHA1 hash
    h := sha1.New()
    h.Write([]byte(input))
    hash := hex.EncodeToString(h.Sum(nil))
    
    // Return the hash as a JSON response
    return c.RenderJson(map[string]string{"hash": hash})
}

func init() {
    // Initialize module routes
    revel.Router(
        (*HashCalculator)(nil),
        []revel.Route{
            {
                Name: "CalculateSHA1",
                Method: "GET",
                Path: "/hash",
                Action: "HashCalculator.CalculateSHA1",
            },
        },
    )
}
