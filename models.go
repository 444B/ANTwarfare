package main

// WarState represents the current state of the war, including which factions are involved.
type WarState struct {
	WarId                string `json:"warId"`
	WarNumber            int    `json:"warNumber"`
	Winner               string `json:"winner"`
	ConquestStartTime    int64  `json:"conquestStartTime"`
	ConquestEndTime      int64  `json:"conquestEndTime"`
	ResistanceStartTime  int64  `json:"resistanceStartTime"`
	RequiredVictoryTowns int    `json:"requiredVictoryTowns"`
}

// Faction represents a player's faction in the game, such as Colonial or Warden.
type Faction string

const (
	FactionColonial Faction = "Colonial"
	FactionWarden   Faction = "Warden"
	FactionNone     Faction = "None" // For neutral or undefined faction
)

// Player represents a human actor, associated with a faction.
type Player struct {
	ID       string  // Unique identifier for the player
	Faction  Faction // The faction to which the player belongs
	Username string  // The player's in-game name
}

// MapItem represents non-human actors such as territories, resources, and bases.
type MapItem struct {
	ID            string  `json:"id"`            // Unique identifier for the item
	Type          string  `json:"type"`          // The type of item (e.g., resource node, base)
	Location      Point   `json:"location"`      // Geographical location on the map
	Faction       Faction `json:"faction"`       // Controlling faction, if applicable
	IsVictoryBase bool    `json:"isVictoryBase"` // Indicates if this is a victory base
	IsScorched    bool    `json:"isScorched"`    // Indicates if the item is scorched
}

// Point represents a location on the map.
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Interaction represents an event or relationship between actors.
type Interaction struct {
	Type         string   // The type of interaction (e.g., problematization, interessment, enrollment, mobilization)
	Participants []string // IDs of the actors involved
	Location     Point    // Where the interaction takes place, if applicable
	Timestamp    int64    // When the interaction occurred
}
	
// Actors represent all actors in the network, whether human or non-human.
type Actor struct {
	ID           string  `json:"id"`       // Unique ID for the Actant
	Type         string  `json:"iconType"` // The type of actor, if non-human
	Location     Point   `json:"location"` // Where the Actant is located, if applicable
	Faction      Faction `json:"faction"`  // Controlling faction, if applicable
}
