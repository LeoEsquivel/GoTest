package database

import "sync"

type ZincSearch struct {
	User     string `default:"admin"`
	Password string `default:"Complexpass#123"`
	// mux      sync.Mutex
}

var zsInstance *ZincSearch

var once sync.Once

func GetInstance() *ZincSearch {
	once.Do(func() {
		zsInstance = &ZincSearch{
			User:     "admin",
			Password: "Complexpass#123",
		}
	})
	return zsInstance
}
