package model

import (
	"api/database"
	"time"
)

type RequestData struct {
	ZincSearch   database.ZincSearch
	Content_type string
	User_agent   string
}

type Query struct {
	Term       string    `json:"term"`
	Start_time time.Time `json:"start_time"`
	End_time   time.Time `json:"end_time"`
}

type QuerySearch struct {
	Search_type string        `json:"search_type"`
	Query_data  Query         `json:"query"`
	From        int           `json:"from"`
	Max_results int           `json:"max_results"`
	Source      []interface{} `json:"_source"`
}
