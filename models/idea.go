package models

import (
	"encoding/json"
	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Idea is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Idea struct {
    ID uuid.UUID `json:"id" db:"id"`
    Title string `json:"title" db:"title"`
    Description nulls.String `json:"description" db:"description"`
    DevComment string `json:"dev_comment" db:"dev_comment"`
    VersionWhenSuggested string `json:"version_when_suggested" db:"version_when_suggested"`
    UpvotedBy slices.Map `json:"upvoted_by" db:"upvoted_by"`
    DownvotedBy slices.Map `json:"downvoted_by" db:"downvoted_by"`
    Editable bool `json:"editable" db:"editable"`
    Keywords string `json:"keywords" db:"keywords"`
    Fullfilled bool `json:"fullfilled" db:"fullfilled"`
    SuggestedBy string `json:"suggested_by" db:"suggested_by"`
    Picked bool `json:"picked" db:"picked"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (i Idea) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Ideas is not required by pop and may be deleted
type Ideas []Idea

// String is not required by pop and may be deleted
func (i Ideas) String() string {
	ji, _ := json.Marshal(i)
	return string(ji)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (i *Idea) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: i.Title, Name: "Title"},
		&validators.StringIsPresent{Field: i.VersionWhenSuggested, Name: "VersionWhenSuggested"},
		&validators.StringIsPresent{Field: i.SuggestedBy, Name: "SuggestedBy"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (i *Idea) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (i *Idea) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
