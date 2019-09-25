package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
	"time"
)
type User struct {
    ID 				uuid.UUID 	`json:"id" db:"id"`
    PhoneNumber 	string 		`json:"phone_number" db:"phone_number"`
    HasPermission 	bool		`json:"has_permission" db:"has_permission"`
    CreatedAt 		time.Time 	`json:"created_at" db:"created_at"`
    UpdatedAt 		time.Time 	`json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error

	return validate.Validate(
		&validators.StringIsPresent{Field: u.PhoneNumber, Name: "Phone number"},
		&validators.StringLengthInRange{Name: "Password", Field: u.PhoneNumber, Message: "Phone number doesn't appear " +
			"to be valid", Min: 9, Max: 9},
	), err
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
