package utils

import (
	"encoding/json"
	"log"
	"os"
)

func JsonDump(body *BodyCuaca, fileName string) {
	input, err := json.MarshalIndent(&body, "", "    ")
	if err != nil {
		log.Fatal(err.Error())
	}
	err2 := os.WriteFile(fileName, input, 0644)
	if err2 != nil {
		log.Fatalf("error when writing new data to [ %s ]\n", fileName)
	}
}

func JsonLoad(fileName string) *BodyCuaca {
	input, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error when reading new data from [ %s ]\n", fileName)
	}
	var body BodyCuaca

	err2 := json.Unmarshal(input, &body)
	if err2 != nil {
		log.Fatalf("error when loading file [ %s ]\n", fileName)
	}
	return &body
}
