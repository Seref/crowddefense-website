package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Prequestionnaire is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Prequestionnaire struct {
	ID uuid.UUID `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Nationality string `json:"nationality" db:"nationality"`
    Age string `json:"age" db:"age"`
    Gender string `json:"gender" db:"gender"`
    Job string `json:"job" db:"job"`
    HoursAtPCPerWeek string `json:"hours_at_pc_per_week" db:"hours_at_pc_per_week"`
    HoursVideoGamesPerWeek string `json:"hours_video_games_per_week" db:"hours_video_games_per_week"`
    MostImportantVideoGameAspects string `json:"most_important_video_game_aspects" db:"most_important_video_game_aspects"`
    MostImportantVideoGameAspect string `json:"most_important_video_game_aspect" db:"most_important_video_game_aspect"`
    PlayedALotOfTowerDefense int `json:"played_alot_of_tower_defense" db:"played_alot_of_tower_defense"`
    LikesTowerDefense int `json:"likes_tower_defense" db:"likes_tower_defense"`
    PlaysMoreTowerDefenseThanOtherGames int `json:"plays_more_tower_defense_than_other_games" db:"plays_more_tower_defense_than_other_games"`
    HasDesignedVideoGame bool `json:"has_designed_video_game" db:"has_designed_video_game"`
    HasDesignedOtherApplication bool `json:"has_designed_other_application" db:"has_designed_other_application"`
    WantsToBeIncluded bool `json:"wants_to_be_included" db:"wants_to_be_included"`
    HowWantsToBeIncluded string `json:"how_wants_to_be_included" db:"how_wants_to_be_included"`
    EverWatchedTwitchPlays bool `json:"ever_watched_twitch_plays" db:"ever_watched_twitch_plays"`
    ParticipatedTwitchPlays bool `json:"participated_twitch_plays" db:"participated_twitch_plays"`
    HowLikesTwitchPlays int `json:"how_likes_twitch_plays" db:"how_likes_twitch_plays"`
    HeardOfPleaseBeNice bool `json:"heard_of_please_be_nice" db:"heard_of_please_be_nice"`
    PlayedPleaseBeNice bool `json:"played_please_be_nice" db:"played_please_be_nice"`
    LikedPleaseBeNice int `json:"liked_please_be_nice" db:"liked_please_be_nice"`
    IdeaImplementedPleaseBeNice bool `json:"idea_implemented_please_be_nice" db:"idea_implemented_please_be_nice"`
    HeardOfCrowdjump bool `json:"heard_of_crowdjump" db:"heard_of_crowdjump"`
    PlayedCrowdjump bool `json:"played_crowdjump" db:"played_crowdjump"`
    LikedCrowdjump int `json:"liked_crowdjump" db:"liked_crowdjump"`
    IdeaImplementedCrowdjump bool `json:"idea_implemented_crowdjump" db:"idea_implemented_crowdjump"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p Prequestionnaire) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Prequestionnaires is not required by pop and may be deleted
type Prequestionnaires []Prequestionnaire

// String is not required by pop and may be deleted
func (p Prequestionnaires) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Prequestionnaire) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.Age, Name: "Age"},
		&validators.StringIsPresent{Field: p.Gender, Name: "Gender"},
		&validators.StringIsPresent{Field: p.Job, Name: "Job"},
		&validators.StringIsPresent{Field: p.HoursAtPCPerWeek, Name: "HoursAtPCPerWeek"},
		&validators.StringIsPresent{Field: p.HoursVideoGamesPerWeek, Name: "HoursVideoGamesPerWeek"},
		&validators.StringIsPresent{Field: p.MostImportantVideoGameAspects, Name: "MostImportantVideoGameAspects"},
		&validators.StringIsPresent{Field: p.MostImportantVideoGameAspect, Name: "MostImportantVideoGameAspect"},
		&validators.IntIsPresent{Field: p.PlayedALotOfTowerDefense, Name: "PlayedALotOfTowerDefense"},
		&validators.IntIsPresent{Field: p.LikesTowerDefense, Name: "LikesTowerDefense"},
		&validators.IntIsPresent{Field: p.PlaysMoreTowerDefenseThanOtherGames, Name: "PlaysMoreTowerDefenseThanOtherGames"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Prequestionnaire) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Prequestionnaire) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
