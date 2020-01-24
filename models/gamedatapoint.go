package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

type GameDataPoint struct {
	ID              uuid.UUID  `json:"id" db:"id"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" db:"updated_at"`
	Username        string     `json:"username" db:"username"`
	Version         string     `json:"version" db:"version"`
	Score           int        `json:"score" db:"score"`
	WaveSurvived    int        `json:"wave_survived" db:"wave_survived"`
	SecondsSurvived int        `json:"seconds_survived" db:"seconds_survived"`
	Win             bool       `json:"win" db:"win"`
	RestartedRound  bool       `json:"restarted_round" db:"restarted_round"`
	TutorialPressed int        `json:"tutorial_pressed" db:"tutorial_pressed"`
	MoneyEarned     int        `json:"money_earned" db:"money_earned"`
	MoneySpent      int        `json:"money_spent" db:"money_spent"`
	Additional      slices.Map `json:"additional" db:"additional"`
}

func (g *GameDataPoint) Create(tx *pop.Connection) (*validate.Errors, error) {
	return tx.ValidateAndCreate(g)
}

func (g GameDataPoint) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

type GameDataPoints []GameDataPoint

func (g GameDataPoints) String() string {
	jg, _ := json.Marshal(g)
	return string(jg)
}

func (g *GameDataPoint) Validate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(), err
}

func (g *GameDataPoint) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	var err error
	return validate.Validate(), err
}

func (g *GameDataPoint) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
