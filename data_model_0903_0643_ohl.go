// 代码生成时间: 2025-09-03 06:43:31
package models

import (
    "encoding/json"
    "errors"
    "revel"
)

// Define the data structure for our application.
// In this example, we'll create a simple User model.
type User struct {
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Email  string `json:"email"`
    Active bool   `json:"active"`
}

// Validate checks the validity of the user data.
func (u *User) Validate(v *revel.ValidationResults) {
    // Validate the name field.
    v.Check(u.Name, revel.Required, revel.MinSize(2))
    // Validate the email field.
    v.Check(u.Email, revel.Required, revel.Email)
    // Validate the active field.
    v.Check(u.Active, revel.Required)
}

// String returns a string representation of the user.
func (u *User) String() string {
    return json.MarshalIndent(u, "", "    ")
}

// NewUser creates a new user instance with default values.
func NewUser() *User {
    return &User{
        Active: true, // default value for Active
    }
}

// Error messages for different validation errors.
var (
    ErrNameRequired    = errors.New("Name is required")
    ErrEmailRequired   = errors.New("Email is required")
    ErrEmailInvalid    = errors.New("Email is invalid")
    ErrActiveRequired  = errors.New("Active status is required")
)

// ValidateName checks if the name is not empty.
func ValidateName(name string) error {
    if name == "" {
        return ErrNameRequired
    }
    return nil
}

// ValidateEmail checks if the email is in a proper format.
func ValidateEmail(email string) error {
    // Simple email validation for example purposes.
    if len(email) < 3 || len(email) > 255 || !strings.Contains(email, "@") {
        return ErrEmailInvalid
    }
    return nil
}

// ValidateActive checks if the active status is true.
func ValidateActive(active bool) error {
    if !active {
        return ErrActiveRequired
    }
    return nil
}
