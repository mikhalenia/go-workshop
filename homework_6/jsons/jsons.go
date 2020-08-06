package jsons

import (
	"encoding/json"
	"log"
	"os"
)

type Car struct { // Struct from homework_3
	Number        string `json:"number"`
	Brand         string `json:"brand"`
	Color         string `json:"color"`
	Year          int `json:"year"`
	Cost          float32 `json:"cost"`
	IsAvailable bool `json:"available"`
}

func toJson(data Car) string {
	result, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return string(result)
}
func fromJson(lineToUnmarshal string) (car Car){
	err := json.Unmarshal([]byte(lineToUnmarshal), &car)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return car
}
