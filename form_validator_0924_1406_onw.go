// 代码生成时间: 2025-09-24 14:06:19
package controllers

import (
    "encoding/json"
    "errors"
    "fmt"
    "revel"
    "revel/captcha"
)

// FormValidator is a simple form validator for the Revel framework.
type FormValidator struct {
    Captcha      *captcha.Captcha // Captcha validation
    Username     string
    Password     string
    Email        string
}

// Validate checks the form fields to ensure they meet the required criteria.
func (v *FormValidator) Validate() ([]string, error) {
    errorsFound := []string{}

    // Check if the username is provided and is not empty.
    if len(v.Username) < 3 || len(v.Username) > 20 {
        errorsFound = append(errorsFound, "Username must be between 3 and 20 characters.")
    }

    // Check if the password is provided and meets the complexity requirements.
    if len(v.Password) < 6 || len(v.Password) > 50 {
        errorsFound = append(errorsFound, "Password must be between 6 and 50 characters.")
    }

    // Check if the email is provided and is valid.
    if len(v.Email) == 0 || !revel.Valid(v.Email) {
        errorsFound = append(errorsFound, "Invalid email address.")
    }

    // Check if the captcha is provided and is correct.
    if v.Captcha == nil || !v.Captcha.AnswerCorrect() {
        errorsFound = append(errorsFound, "Captcha validation failed.")
    }

    if len(errorsFound) > 0 {
        return errorsFound, errors.New("Validation errors found.")
    }

    return nil, nil
}

// NewFormValidator creates a new instance of FormValidator.
func NewFormValidator(username, password, email string, captcha *captcha.Captcha) *FormValidator {
    return &FormValidator{
        Username: username,
        Password: password,
        Email:    email,
        Captcha:  captcha,
    }
}