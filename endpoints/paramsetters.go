package endpoint

import (
	"strconv"

	"github.com/ndesai96/nba-api/models/league"
	"github.com/ndesai96/nba-api/models/player"
)

func (b *Base) SetPlayerID(playerID int) {
	b.params.Set("PlayerID", strconv.Itoa(playerID))
}

func (b *Base) SetTeamID(teamID int) {
	b.params.Set("TeamID", strconv.Itoa(teamID))
}

func (b *Base) SetLeagueID(leagueID league.ID) {
	b.params.Set("LeagueID", string(leagueID))
}

func (b *Base) SetOnlyAllStar(onlyAllStar bool) {
	b.params.Set("AllStar", strconv.Itoa(boolToInt(onlyAllStar)))
}

func (b *Base) SetPlayerPosition(playerPosition player.Position) {
	b.params.Set("PlayerPosition", string(playerPosition))
}

func (b *Base) SetSeason(season league.Season) {
	b.params.Set("Season", string(season))
}

func (b *Base) SetHistorical(includeHistorical bool) {
	b.params.Set("Historical", strconv.Itoa(boolToInt(includeHistorical)))
}

func (b *Base) SetOnlyActive(onlyActive bool) {
	b.params.Set("Active", strconv.Itoa(boolToInt(onlyActive)))
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
