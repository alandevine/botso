package footy

import "time"

// Generated using https://mholt.github.io/json-to-go/
type Standings struct {
	Errors []any  `json:"errors"`
	Get    string `json:"get"`
	Paging struct {
		Current int `json:"current"`
		Total   int `json:"total"`
	} `json:"paging"`
	Parameters struct {
		League string `json:"league"`
		Season string `json:"season"`
	} `json:"parameters"`
	Response []struct {
		League struct {
			Country   string `json:"country"`
			Flag      any    `json:"flag"`
			ID        int    `json:"id"`
			Logo      string `json:"logo"`
			Name      string `json:"name"`
			Season    int    `json:"season"`
			Standings [][]struct {
				All struct {
					Draw  int `json:"draw"`
					Goals struct {
						Against int `json:"against"`
						For     int `json:"for"`
					} `json:"goals"`
					Lose   int `json:"lose"`
					Played int `json:"played"`
					Win    int `json:"win"`
				} `json:"all"`
				Away struct {
					Draw  any `json:"draw"`
					Goals struct {
						Against any `json:"against"`
						For     any `json:"for"`
					} `json:"goals"`
					Lose   any `json:"lose"`
					Played any `json:"played"`
					Win    any `json:"win"`
				} `json:"away"`
				Description string `json:"description"`
				Form        string `json:"form"`
				GoalsDiff   int    `json:"goalsDiff"`
				Group       string `json:"group"`
				Home        struct {
					Draw  any `json:"draw"`
					Goals struct {
						Against any `json:"against"`
						For     any `json:"for"`
					} `json:"goals"`
					Lose   any `json:"lose"`
					Played any `json:"played"`
					Win    any `json:"win"`
				} `json:"home"`
				Points int    `json:"points"`
				Rank   int    `json:"rank"`
				Status string `json:"status"`
				Team   struct {
					ID   int    `json:"id"`
					Logo string `json:"logo"`
					Name string `json:"name"`
				} `json:"team"`
				Update time.Time `json:"update"`
			} `json:"standings"`
		} `json:"league"`
	} `json:"response"`
	Results int `json:"results"`
}
