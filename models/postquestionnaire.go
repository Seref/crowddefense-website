package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Postquestionnaire is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Postquestionnaire struct {
	ID uuid.UUID `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
    PlayingWasFun int `json:"playing_was_fun" db:"playing_was_fun"`
    LikedPlaying int `json:"liked_playing" db:"liked_playing"`
    EnjoyedPlayingVeryMuch int `json:"enjoyed_playing_very_much" db:"enjoyed_playing_very_much"`
    GameExperienceWasPleasurable int `json:"game_experience_was_pleasurable" db:"game_experience_was_pleasurable"`
    PlayingIsEntertaining int `json:"playing_is_entertaining" db:"playing_is_entertaining"`
    WouldPlayForOwnSake int `json:"would_play_for_own_sake" db:"would_play_for_own_sake"`
    PlayingMadeMeForgetWhereIAm int `json:"playing_made_me_forget_where_iam" db:"playing_made_me_forget_where_iam"`
    PlayingMadeMeForgetSurroundings int `json:"playing_made_me_forget_surroundings" db:"playing_made_me_forget_surroundings"`
    PlayingFeltLikeJourney int `json:"playing_felt_like_journey" db:"playing_felt_like_journey"`
    PlayingGotMeAwayFromItAll int `json:"playing_got_me_away_from_it_all" db:"playing_got_me_away_from_it_all"`
    WhilePlayingCompletelyOblivious int `json:"while_playing_completely_oblivious" db:"while_playing_completely_oblivious"`
    WhilePlayingLostTrackOfTime int `json:"while_playing_lost_track_of_time" db:"while_playing_lost_track_of_time"`
    PlayingSparkedImagination int `json:"playing_sparked_imagination" db:"playing_sparked_imagination"`
    WhilePlayingFeltCreative int `json:"while_playing_felt_creative" db:"while_playing_felt_creative"`
    WhilePlayingFeltExplorative int `json:"while_playing_felt_explorative" db:"while_playing_felt_explorative"`
    WhilePlayingFeltAdventurous int `json:"while_playing_felt_adventurous" db:"while_playing_felt_adventurous"`
    WhilePlayingFeltActivated int `json:"while_playing_felt_activated" db:"while_playing_felt_activated"`
    WhilePlayingFeltJittery int `json:"while_playing_felt_jittery" db:"while_playing_felt_jittery"`
    WhilePlayingFeltFrenzied int `json:"while_playing_felt_frenzied" db:"while_playing_felt_frenzied"`
    WhilePlayingFeltExcited int `json:"while_playing_felt_excited" db:"while_playing_felt_excited"`
    WhilePlayingFeltUpset int `json:"while_playing_felt_upset" db:"while_playing_felt_upset"`
    WhilePlayingFeltHostile int `json:"while_playing_felt_hostile" db:"while_playing_felt_hostile"`
    WhilePlayingFeltFrustrated int `json:"while_playing_felt_frustrated" db:"while_playing_felt_frustrated"`
    WhilePlayingFeltInCharge int `json:"while_playing_felt_in_charge" db:"while_playing_felt_in_charge"`
    WhilePlayingFeltInfluential int `json:"while_playing_felt_influential" db:"while_playing_felt_influential"`
    WhilePlayingFeltAutonomous int `json:"while_playing_felt_autonomous" db:"while_playing_felt_autonomous"`
    WhilePlayingFeltConfident int `json:"while_playing_felt_confident" db:"while_playing_felt_confident"`
    EnjoyedPlayingCrowdDefenseVeryMuch int `json:"enjoyed_playing_crowd_defense_very_much" db:"enjoyed_playing_crowd_defense_very_much"`
    PlayingWasFunToDo int `json:"playing_was_fun_to_do" db:"playing_was_fun_to_do"`
    WasBoringActivity int `json:"was_boring_activity" db:"was_boring_activity"`
    DidNotHoldMyAttention int `json:"did_not_hold_my_attention" db:"did_not_hold_my_attention"`
    WouldDescribeAsInteresting int `json:"would_describe_as_interesting" db:"would_describe_as_interesting"`
    ThoughtWasEnjoyable int `json:"thought_was_enjoyable" db:"thought_was_enjoyable"`
    ThoughtWasEnjoyableWhilePlaying int `json:"thought_was_enjoyable_while_playing" db:"thought_was_enjoyable_while_playing"`
    DidWell int `json:"did_well" db:"did_well"`
    DidWellInComparisonToOthers int `json:"did_well_in_comparison_to_others" db:"did_well_in_comparison_to_others"`
    FeltCompetentAfterAWhile int `json:"felt_competent_after_awhile" db:"felt_competent_after_awhile"`
    SatisfiedWithPerformance int `json:"satisfied_with_performance" db:"satisfied_with_performance"`
    WasPrettySkilled int `json:"was_pretty_skilled" db:"was_pretty_skilled"`
    CouldntDoVeryWell int `json:"couldnt_do_very_well" db:"couldnt_do_very_well"`
    DidntFeelNervousPlaying int `json:"didnt_feel_nervous_playing" db:"didnt_feel_nervous_playing"`
    FeltTenseWhilePlaying int `json:"felt_tense_while_playing" db:"felt_tense_while_playing"`
    FeltRelaxedWhilePlaying int `json:"felt_relaxed_while_playing" db:"felt_relaxed_while_playing"`
    WasAnxiousWhilePlaying int `json:"was_anxious_while_playing" db:"was_anxious_while_playing"`
    FeltPressuredWhilePlaying int `json:"felt_pressured_while_playing" db:"felt_pressured_while_playing"`
    BelievedHadChoiceAboutPlaying int `json:"believed_had_choice_about_playing" db:"believed_had_choice_about_playing"`
    FeltLikeNotMyChoiceToPlay int `json:"felt_like_not_my_choice_to_play" db:"felt_like_not_my_choice_to_play"`
    DidntHaveChoiceAboutPlaying int `json:"didnt_have_choice_about_playing" db:"didnt_have_choice_about_playing"`
    FeltLikeHadToPlay int `json:"felt_like_had_to_play" db:"felt_like_had_to_play"`
    PlayedBecauseNoChoice int `json:"played_because_no_choice" db:"played_because_no_choice"`
    PlayedBecauseWantedTo int `json:"played_because_wanted_to" db:"played_because_wanted_to"`
    PlayedBecauseHadTo int `json:"played_because_had_to" db:"played_because_had_to"`
    FeltDistantToOthers int `json:"felt_distant_to_others" db:"felt_distant_to_others"`
    DoubtFriendshipWithOthers int `json:"doubt_friendship_with_others" db:"doubt_friendship_with_others"`
    FeltLikeCouldTrustOthers int `json:"felt_like_could_trust_others" db:"felt_like_could_trust_others"`
    WantMorePlayerInteraction int `json:"want_more_player_interaction" db:"want_more_player_interaction"`
    PreferNoInteraction int `json:"prefer_no_interaction" db:"prefer_no_interaction"`
    FeltLikeCouldntTrustOthers int `json:"felt_like_couldnt_trust_others" db:"felt_like_couldnt_trust_others"`
    CouldSeeFriendship int `json:"could_see_friendship" db:"could_see_friendship"`
    FeelsCloseToOtherPlayers int `json:"feels_close_to_other_players" db:"feels_close_to_other_players"`
    WouldLikeToUseFrequently int `json:"would_like_to_use_frequently" db:"would_like_to_use_frequently"`
    FoundUnnecessaryComplex int `json:"found_unnecessary_complex" db:"found_unnecessary_complex"`
    FoundEasyToUse int `json:"found_easy_to_use" db:"found_easy_to_use"`
    WouldNeedAssistanceFromTechnicalPerson int `json:"would_need_assistance_from_technical_person" db:"would_need_assistance_from_technical_person"`
    FunctionsWereWellIntegrated int `json:"functions_were_well_integrated" db:"functions_were_well_integrated"`
    TooMuchInconsistency int `json:"too_much_inconsistency" db:"too_much_inconsistency"`
    WouldImagineFlatLearningCurve int `json:"would_imagine_flat_learning_curve" db:"would_imagine_flat_learning_curve"`
    SystemWasCumbersome int `json:"system_was_cumbersome" db:"system_was_cumbersome"`
    FeltConfidentUsing int `json:"felt_confident_using" db:"felt_confident_using"`
    SteepLearningCurve int `json:"steep_learning_curve" db:"steep_learning_curve"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p Postquestionnaire) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Postquestionnaires is not required by pop and may be deleted
type Postquestionnaires []Postquestionnaire

// String is not required by pop and may be deleted
func (p Postquestionnaires) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Postquestionnaire) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: p.PlayingWasFun, Name: "PlayingWasFun"},
		&validators.IntIsPresent{Field: p.LikedPlaying, Name: "LikedPlaying"},
		&validators.IntIsPresent{Field: p.EnjoyedPlayingVeryMuch, Name: "EnjoyedPlayingVeryMuch"},
		&validators.IntIsPresent{Field: p.GameExperienceWasPleasurable, Name: "GameExperienceWasPleasurable"},
		&validators.IntIsPresent{Field: p.PlayingIsEntertaining, Name: "PlayingIsEntertaining"},
		&validators.IntIsPresent{Field: p.WouldPlayForOwnSake, Name: "WouldPlayForOwnSake"},
		&validators.IntIsPresent{Field: p.PlayingMadeMeForgetWhereIAm, Name: "PlayingMadeMeForgetWhereIAm"},
		&validators.IntIsPresent{Field: p.PlayingMadeMeForgetSurroundings, Name: "PlayingMadeMeForgetSurroundings"},
		&validators.IntIsPresent{Field: p.PlayingFeltLikeJourney, Name: "PlayingFeltLikeJourney"},
		&validators.IntIsPresent{Field: p.PlayingGotMeAwayFromItAll, Name: "PlayingGotMeAwayFromItAll"},
		&validators.IntIsPresent{Field: p.WhilePlayingCompletelyOblivious, Name: "WhilePlayingCompletelyOblivious"},
		&validators.IntIsPresent{Field: p.WhilePlayingLostTrackOfTime, Name: "WhilePlayingLostTrackOfTime"},
		&validators.IntIsPresent{Field: p.PlayingSparkedImagination, Name: "PlayingSparkedImagination"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltCreative, Name: "WhilePlayingFeltCreative"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltExplorative, Name: "WhilePlayingFeltExplorative"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltAdventurous, Name: "WhilePlayingFeltAdventurous"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltActivated, Name: "WhilePlayingFeltActivated"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltJittery, Name: "WhilePlayingFeltJittery"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltFrenzied, Name: "WhilePlayingFeltFrenzied"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltExcited, Name: "WhilePlayingFeltExcited"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltUpset, Name: "WhilePlayingFeltUpset"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltHostile, Name: "WhilePlayingFeltHostile"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltFrustrated, Name: "WhilePlayingFeltFrustrated"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltInCharge, Name: "WhilePlayingFeltInCharge"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltInfluential, Name: "WhilePlayingFeltInfluential"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltAutonomous, Name: "WhilePlayingFeltAutonomous"},
		&validators.IntIsPresent{Field: p.WhilePlayingFeltConfident, Name: "WhilePlayingFeltConfident"},
		&validators.IntIsPresent{Field: p.EnjoyedPlayingCrowdDefenseVeryMuch, Name: "EnjoyedPlayingCrowdDefenseVeryMuch"},
		&validators.IntIsPresent{Field: p.PlayingWasFunToDo, Name: "PlayingWasFunToDo"},
		&validators.IntIsPresent{Field: p.WasBoringActivity, Name: "WasBoringActivity"},
		&validators.IntIsPresent{Field: p.DidNotHoldMyAttention, Name: "DidNotHoldMyAttention"},
		&validators.IntIsPresent{Field: p.WouldDescribeAsInteresting, Name: "WouldDescribeAsInteresting"},
		&validators.IntIsPresent{Field: p.ThoughtWasEnjoyable, Name: "ThoughtWasEnjoyable"},
		&validators.IntIsPresent{Field: p.ThoughtWasEnjoyableWhilePlaying, Name: "ThoughtWasEnjoyableWhilePlaying"},
		&validators.IntIsPresent{Field: p.DidWell, Name: "DidWell"},
		&validators.IntIsPresent{Field: p.DidWellInComparisonToOthers, Name: "DidWellInComparisonToOthers"},
		&validators.IntIsPresent{Field: p.FeltCompetentAfterAWhile, Name: "FeltCompetentAfterAWhile"},
		&validators.IntIsPresent{Field: p.SatisfiedWithPerformance, Name: "SatisfiedWithPerformance"},
		&validators.IntIsPresent{Field: p.WasPrettySkilled, Name: "WasPrettySkilled"},
		&validators.IntIsPresent{Field: p.CouldntDoVeryWell, Name: "CouldntDoVeryWell"},
		&validators.IntIsPresent{Field: p.DidntFeelNervousPlaying, Name: "DidntFeelNervousPlaying"},
		&validators.IntIsPresent{Field: p.FeltTenseWhilePlaying, Name: "FeltTenseWhilePlaying"},
		&validators.IntIsPresent{Field: p.FeltRelaxedWhilePlaying, Name: "FeltRelaxedWhilePlaying"},
		&validators.IntIsPresent{Field: p.WasAnxiousWhilePlaying, Name: "WasAnxiousWhilePlaying"},
		&validators.IntIsPresent{Field: p.FeltPressuredWhilePlaying, Name: "FeltPressuredWhilePlaying"},
		&validators.IntIsPresent{Field: p.BelievedHadChoiceAboutPlaying, Name: "BelievedHadChoiceAboutPlaying"},
		&validators.IntIsPresent{Field: p.FeltLikeNotMyChoiceToPlay, Name: "FeltLikeNotMyChoiceToPlay"},
		&validators.IntIsPresent{Field: p.DidntHaveChoiceAboutPlaying, Name: "DidntHaveChoiceAboutPlaying"},
		&validators.IntIsPresent{Field: p.FeltLikeHadToPlay, Name: "FeltLikeHadToPlay"},
		&validators.IntIsPresent{Field: p.PlayedBecauseNoChoice, Name: "PlayedBecauseNoChoice"},
		&validators.IntIsPresent{Field: p.PlayedBecauseWantedTo, Name: "PlayedBecauseWantedTo"},
		&validators.IntIsPresent{Field: p.PlayedBecauseHadTo, Name: "PlayedBecauseHadTo"},
		&validators.IntIsPresent{Field: p.FeltDistantToOthers, Name: "FeltDistantToOthers"},
		&validators.IntIsPresent{Field: p.DoubtFriendshipWithOthers, Name: "DoubtFriendshipWithOthers"},
		&validators.IntIsPresent{Field: p.FeltLikeCouldTrustOthers, Name: "FeltLikeCouldTrustOthers"},
		&validators.IntIsPresent{Field: p.WantMorePlayerInteraction, Name: "WantMorePlayerInteraction"},
		&validators.IntIsPresent{Field: p.PreferNoInteraction, Name: "PreferNoInteraction"},
		&validators.IntIsPresent{Field: p.FeltLikeCouldntTrustOthers, Name: "FeltLikeCouldntTrustOthers"},
		&validators.IntIsPresent{Field: p.CouldSeeFriendship, Name: "CouldSeeFriendship"},
		&validators.IntIsPresent{Field: p.FeelsCloseToOtherPlayers, Name: "FeelsCloseToOtherPlayers"},
		&validators.IntIsPresent{Field: p.WouldLikeToUseFrequently, Name: "WouldLikeToUseFrequently"},
		&validators.IntIsPresent{Field: p.FoundUnnecessaryComplex, Name: "FoundUnnecessaryComplex"},
		&validators.IntIsPresent{Field: p.FoundEasyToUse, Name: "FoundEasyToUse"},
		&validators.IntIsPresent{Field: p.WouldNeedAssistanceFromTechnicalPerson, Name: "WouldNeedAssistanceFromTechnicalPerson"},
		&validators.IntIsPresent{Field: p.FunctionsWereWellIntegrated, Name: "FunctionsWereWellIntegrated"},
		&validators.IntIsPresent{Field: p.TooMuchInconsistency, Name: "TooMuchInconsistency"},
		&validators.IntIsPresent{Field: p.WouldImagineFlatLearningCurve, Name: "WouldImagineFlatLearningCurve"},
		&validators.IntIsPresent{Field: p.SystemWasCumbersome, Name: "SystemWasCumbersome"},
		&validators.IntIsPresent{Field: p.FeltConfidentUsing, Name: "FeltConfidentUsing"},
		&validators.IntIsPresent{Field: p.SteepLearningCurve, Name: "SteepLearningCurve"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Postquestionnaire) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Postquestionnaire) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
