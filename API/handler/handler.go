package handler

import (
	"api/database"
	m "api/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
)

var (
	request_data = m.RequestData{
		ZincSearch:   *database.GetInstance(),
		Content_type: "application/json",
		User_agent:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36",
	}
	url_BASE string = "http://localhost:4080/api/"
	index    string = "enron_mail"
)

func GetLastMails() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		getSkip := chi.URLParam(r, "skip")
		getLimit := chi.URLParam(r, "limit")

		skip, _ := strconv.Atoi(getSkip)
		limit, _ := strconv.Atoi(getLimit)

		if limit == 0 {
			limit = 15
		}

		var query = m.QuerySearch{
			Search_type: "matchall",
			Query_data: m.Query{
				Term:       "",
				Start_time: time.Now().AddDate(0, -1, 0),
				End_time:   time.Now(),
			},
			From:        skip,
			Max_results: limit,
			Source:      []interface{}{},
		}

		query_json, err := json.Marshal(query)

		if err != nil {
			log.Fatal(err)
		}

		var url = url_BASE + index + "/_search"
		req, err := http.NewRequest(
			"POST",
			url,
			strings.NewReader(string(query_json)))

		if err != nil {
			log.Fatal(err)
		}

		req.SetBasicAuth(request_data.ZincSearch.User, request_data.ZincSearch.Password)
		req.Header.Set("Content-Type", request_data.Content_type)
		req.Header.Set("User-Agent", request_data.User_agent)

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		if err != nil {
			log.Fatal(err)
		}
		io.Copy(w, resp.Body)

	}

}

func Search() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		getSkip := chi.URLParam(r, "skip")
		getLimit := chi.URLParam(r, "limit")
		getTerm := chi.URLParam(r, "term")
		skip, _ := strconv.Atoi(getSkip)
		limit, _ := strconv.Atoi(getLimit)

		if limit == 0 {
			limit = 15
		}

		var query = m.QuerySearch{
			Search_type: "term",
			Query_data: m.Query{
				Term:       getTerm,
				Start_time: time.Now().AddDate(-1, 0, 0),
				End_time:   time.Now(),
			},
			From:        skip,
			Max_results: limit,
			Source:      []interface{}{},
		}

		query_json, err := json.Marshal(query)

		if err != nil {
			log.Fatal(err)
		}

		var url = url_BASE + index + "/_search"
		req, err := http.NewRequest(
			"POST",
			url,
			strings.NewReader(string(query_json)))

		if err != nil {
			log.Fatal(err)
		}

		req.SetBasicAuth(request_data.ZincSearch.User, request_data.ZincSearch.Password)
		req.Header.Set("Content-Type", request_data.Content_type)
		req.Header.Set("User-Agent", request_data.User_agent)

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		if err != nil {
			log.Fatal(err)
		}
		io.Copy(w, resp.Body)

	}
}
