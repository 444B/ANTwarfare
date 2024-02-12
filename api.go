// api.go

package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// fetchWarState retrieves the current state of the war from the API.
func fetchWarState() (*WarState, error) {
    url := "https://war-service-live.foxholeservices.com/api/worldconquest/war"
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error fetching war state: %w", err)
    }
    defer resp.Body.Close()

    var state WarState
    if err := json.NewDecoder(resp.Body).Decode(&state); err != nil {
        return nil, fmt.Errorf("error decoding JSON: %w", err)
    }

    return &state, nil
}

// Additional API fetching functions can be added below.
