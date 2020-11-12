package api

import "github.com/diamondburned/arikawa/v2/discord"

type PullVector struct {
	Decide *discord.Snowflake

	Before discord.Snowflake
	After  discord.Snowflake
	Around discord.Snowflake

	Limit uint
}

func (pv PullVector) CreateVector() {
	switch {
	case pv.Around > 0:
		pv.Decide = &pv.Around
	case pv.Before > 0:
		pv.Decide = &pv.Before
	default:
		pv.Decide = &pv.After
	}
}

func (pv PullVector) ReturnVector() (discord.Snowflake, uint) {
	return *pv.Decide, pv.Limit
}
