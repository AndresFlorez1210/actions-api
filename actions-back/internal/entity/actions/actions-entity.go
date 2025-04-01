package entity

import "github.com/uptrace/bun"

type Action struct {
	bun.BaseModel `bun:"actions"`

	Ticker     string `json:"ticker"`
	TargetFrom string `json:"target_from"`
	TargetTo   string `json:"target_to"`
	Company    string `json:"company"`
	Action     string `json:"action"`
	Brokerage  string `json:"brokerage"`
	RatingFrom string `json:"rating_from"`
	RatingTo   string `json:"rating_to"`
	Time       string `json:"time"`
}

type FilterAction struct {
	Key   string `json:key`
	Value string `json:value`
}
