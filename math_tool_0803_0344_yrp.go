// 代码生成时间: 2025-08-03 03:44:52
package controllers

import (
    "encoding/json"
    "github.com/revel/revel"
)

// MathTool provides a set of mathematical operations.
type MathTool struct {
    *revel.Controller
}

// Sum calculates the sum of two numbers.
func (c MathTool) Sum(a, b float64) revel.Result {
    sum := a + b
    
    result := map[string]float64{
        "result": sum,
    }
    return c.RenderJSON(result)
}

// Subtract calculates the difference of two numbers.
func (c MathTool) Subtract(a, b float64) revel.Result {
    diff := a - b
    
    result := map[string]float64{
        "result": diff,
    }
    return c.RenderJSON(result)
}

// Multiply calculates the product of two numbers.
func (c MathTool) Multiply(a, b float64) revel.Result {
    product := a * b
    
    result := map[string]float64{
        "result": product,
    }
    return c.RenderJSON(result)
}

// Divide calculates the quotient of two numbers.
// It returns an error if the divisor is zero.
func (c MathTool) Divide(a, b float64) revel.Result {
    if b == 0 {
        error := map[string]string{
            "error": "Divide by zero is not allowed",
        }
        return c.RenderJSON(error)
    }
    
    quotient := a / b
    
    result := map[string]float64{
        "result": quotient,
    }
    return c.RenderJSON(result)
}