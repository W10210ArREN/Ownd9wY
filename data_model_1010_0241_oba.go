// 代码生成时间: 2025-10-10 02:41:27
package models

import (
    "encoding/json"
    "errors"
    "github.com/revel/revel"
)

// User represents a user in the system.
type User struct {
    ID       int    `db:"id" json:"id"`
    Username string `db:"username" json:"username"`
    Email    string `db:"email" json:"email"`
    Password string `db:"password" json:"-"` // Password is not included in JSON output for security reasons.
}

// Validate checks the validity of a User instance.
func (u *User) Validate(v *revel.Validation) {
    v.Check(u.Username, revel.ValidRequired(), revel.ValidMaxSize(100))
    v.Check(u.Email, revel.ValidRequired(), revel.ValidEmail())
    v.Check(u.Password, revel.ValidRequired(), revel.ValidMinSize(8))
}

// UserJSON is a helper struct for outputting User data in JSON, excluding sensitive fields.
type UserJSON struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}

// ToJSON converts a User to a UserJSON struct for JSON output.
func (u *User) ToJSON() UserJSON {
    return UserJSON{
        ID:       u.ID,
        Username: u.Username,
        Email:    u.Email,
    }
}

// Save persists the User to the database.
func (u *User) Save() error {
    // Here you would typically use a database ORM to save the user.
    // For the sake of this example, we're just returning a success.
    err := json.Marshal(u) // Simulate saving by marshaling to JSON.
    if err != nil {
        return errors.New("Failed to save user: " + err.Error())
    }
    return nil
}
