package footy

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"adevine/botso/cfg"
)

var (
	FootyToken   string
	Participents cfg.Config
)

func DoRankings() (string, map[string]int) {
	url := "https://api-football-v1.p.rapidapi.com/v3/standings?league=4&season=2024"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-key", FootyToken)
	req.Header.Add("x-rapidapi-host", "api-football-v1.p.rapidapi.com")

	response, _ := http.DefaultClient.Do(req)

	var standings Standings

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", nil
	}

	json.Unmarshal(body, &standings)

	teamPoints := GetTeamPoints(standings)
	participentPoints := make(map[string]int)

	for _, participent := range Participents.Participents {
		fmt.Println(participent.Name)
		pointsSum := 0
		for _, team := range participent.Teams {
			pointsSum += teamPoints[team]
		}
		participentPoints[participent.Name] = pointsSum
	}

	return GetLeader(standings), participentPoints
}

func GetTeamPoints(standingsResponse Standings) map[string]int {
	totalScores := make(map[string]int)

	for _, standings := range standingsResponse.Response[0].League.Standings {
		for _, standing := range standings {
			totalScores[standing.Team.Name] = standing.Points
		}
	}

	return totalScores
}

func GetLeader(standingsResponse Standings) string {
	for _, standings := range standingsResponse.Response[0].League.Standings {
		for _, standing := range standings {
			if standing.Rank == 1 {
				return standing.Team.Name
			}
		}
	}

	return ""
}
