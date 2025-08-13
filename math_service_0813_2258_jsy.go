// 代码生成时间: 2025-08-13 22:58:27
 * It follows GoLang best practices and is structured for easy understanding and maintenance.
 */

package main

import (
    "math"
    "revel"
)

// MathService is the struct type for the math service.
type MathService struct {
    *revel.Controller
}

// Add performs addition of two numbers.
func (c MathService) Add(a, b float64) (float64, error) {
    result := a + b
    return result, nil
}

// Subtract performs subtraction of two numbers.
func (c MathService) Subtract(a, b float64) (float64, error) {
    result := a - b
    return result, nil
}

// Multiply performs multiplication of two numbers.
func (c MathService) Multiply(a, b float64) (float64, error) {
    result := a * b
    return result, nil
}

// Divide performs division of two numbers.
func (c MathService) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, revel.NewErrorFromError(fmt.Errorf("division by zero"))
    }
    result := a / b
    return result, nil
}

// sqrt calculates the square root of a number.
func (c MathService) Sqrt(value float64) (float64, error) {
    if value < 0 {
        return 0, revel.NewErrorFromError(fmt.Errorf("cannot calculate square root of negative number"))
    }
    return math.Sqrt(value), nil
}
