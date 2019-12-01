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
type Suggestion struct {
    ID uuid.UUID `json:"id" db:"id"`
    Title string `json:"title" db:"title"`
    Body string `json:"body" db:"body"`
    DevComment nulls.String `json:"dev_comment" db:"dev_comment"`
    VersionWhenSuggested nulls.String `json:"version_when_suggested" db:"version_when_suggested"`
    UpvotedBy slices.Map `json:"upvoted_by" db:"upvoted_by"`
    DownvotedBy slices.Map `json:"downvoted_by" db:"downvoted_by"`
    Editable bool `json:"editable" db:"editable"`
    Keywords slices.Map `json:"keywords" db:"keywords"`
    Fulfilled bool `json:"fulfilled" db:"fulfilled"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (s Suggestion) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Suggestions is not required by pop and may be deleted
type Suggestions []Suggestion

// String is not required by pop and may be deleted
func (s Suggestions) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Suggestion) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: s.Title, Name: "Title"},
		&validators.StringIsPresent{Field: s.Body, Name: "Body"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Suggestion) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Suggestion) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
