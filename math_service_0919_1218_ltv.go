// 代码生成时间: 2025-09-19 12:18:04
package controllers

import (
    "math"
    "github.com/revel/revel"
)

// MathController is a Revel controller responsible for handling math-related calculations.
type MathController struct {
    *revel.Controller
}

// Add handles the addition of two numbers.
func (c MathController) Add(a, b float64) revel.Result {
    if a < 0 || b < 0 {
        return c.RenderError(revel.NewError(400, "Non-negative numbers are expected"))
    }
    result := a + b
    return c.RenderJSON(revel.Result{
        "filename": "math_service.go",
        "code": "{"result": " + strconv.FormatFloat(result, 'f', 2, 64) + "}",
    })
}

// Subtract handles the subtraction of two numbers.
func (c MathController) Subtract(a, b float64) revel.Result {
    if a < 0 || b < 0 {
        return c.RenderError(revel.NewError(400, "Non-negative numbers are expected"))
    }
    result := a - b
    return c.RenderJSON(revel.Result{
        "filename": "math_service.go",
        "code": "{"result": " + strconv.FormatFloat(result, 'f', 2, 64) + "}",
    })
}

// Multiply handles the multiplication of two numbers.
func (c MathController) Multiply(a, b float64) revel.Result {
    if a < 0 || b < 0 {
        return c.RenderError(revel.NewError(400, "Non-negative numbers are expected"))
    }
    result := a * b
    return c.RenderJSON(revel.Result{
        "filename": "math_service.go",
        "code": "{"result": " + strconv.FormatFloat(result, 'f', 2, 64) + "}",
    })
}

// Divide handles the division of two numbers.
func (c MathController) Divide(a, b float64) revel.Result {
    if a < 0 || b <= 0 {
        return c.RenderError(revel.NewError(400, "Non-negative numbers are expected and denominator must be greater than zero"))
    }
    result := a / b
    return c.RenderJSON(revel.Result{
        "filename": "math_service.go",
        "code": "{"result": " + strconv.FormatFloat(result, 'f', 2, 64) + "}",
    })
}

// Sqrt handles the square root calculation of a number.
func (c MathController) Sqrt(a float64) revel.Result {
    if a < 0 {
        return c.RenderError(revel.NewError(400, "Only non-negative numbers are expected"))
    }
    result := math.Sqrt(a)
    return c.RenderJSON(revel.Result{
        "filename": "math_service.go",
        "code": "{"result": " + strconv.FormatFloat(result, 'f', 2, 64) + "}",
    })
}
