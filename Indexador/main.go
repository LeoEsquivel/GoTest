package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type BulkJson struct {
	Index   string              `json:"index"`
	Records []map[string]string `json:"records"`
}

// TO DO: Add concurrency
func main() {
	startTime := time.Now()
	fmt.Println("Iniciando programa")
	ruta := "D:/Proyectos-Personales/Go/enron_mail_20110402"
	var waitGroup sync.WaitGroup

	jsonbulk2 := make([]map[string]string, 0)
	var err = filepath.Walk(ruta, func(path string, info fs.FileInfo, err error) error {
		var fileData = make(map[string]string)

		if !info.IsDir() {

			content, err := ioutil.ReadFile(path)

			check(err, "Hubo un problema al acceder a la ruta: "+path)

			lines := strings.Split(string(content), "\n")

			for _, line := range lines {
				line = strings.Replace(line, "\r", "", -1)
				keyval := strings.SplitN(line, ":", 2)
				if len(keyval) == 2 && len(fileData) < 15 {
					fileData[keyval[0]] = keyval[1]
				} else {
					fileData["message"] += keyval[0]
				}
			}
			check(err, "Eror al crear crear el JSON: ")

			jsonbulk2 = append(jsonbulk2, fileData)

		} else if len(jsonbulk2) > 0 {
			fmt.Printf("Visitando: %q\n", path)
			waitGroup.Add(1)
			go func() {
				mailsData := jsonbulk2
				jsonbulk2 = nil
				insertData(mailsData)
				waitGroup.Done()
			}()

			// insertData(jsonbulk2)
			// jsonbulk2 = nil
		}

		return nil
	})

	check(err, "Error recorriendo la ruta.")
	waitGroup.Wait()
	endTime := time.Now()
	fmt.Println("Fin de ejecucion:", endTime.Sub(startTime))

}

func insertData(fileDatas []map[string]string) {

	var data = &BulkJson{
		Index:   "enron_mail_concurerncia",
		Records: fileDatas,
	}

	json_data, err := json.Marshal(data)

	check(err, "Error al pasar la información a JSON")

	req, err := http.NewRequest("POST", "http://localhost:4080/api/_bulkv2", bytes.NewBuffer(json_data))

	check(err, "A ocurrido un error al intentar hacer la petición")

	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	check(err, "Error al hacer la peticion")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	check(err, "Error al obtener la respuesta")
	fmt.Println(string(body))
}

func check(e error, msg string) {
	if e != nil {
		fmt.Println(msg)
		log.Fatal(e)
	}
}
