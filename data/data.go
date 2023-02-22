package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Input struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Telephone string `json:"telephone"`
}

type Output struct {
	Results []Input `json:"results"`
}

func UploadTelephone(newInput Input) {
	output := GetTelephones()
	if len(output.Results) == 0 {
		newInput.Id = 1
	} else {
		newInput.Id = output.Results[len(output.Results)-1].Id + 1
	}
	output.Results = append(output.Results, newInput)
	writeTelephoneToFile(output)
}

func GetTelephones() Output {
	var output Output
	content, err := ioutil.ReadFile("data/telephones.json")
	if err != nil {
		return output
	}
	err = json.Unmarshal(content, &output)
	if err != nil {
		log.Fatal(err)
	}
	return output
}

func GetTelephoneById(id int) Input {
	output := GetTelephones()
	for _, result := range output.Results {
		if result.Id == id {
			return result
		}
	}
	return Input{}
}

func DeleteTelephone(id int) {
	output := GetTelephones()
	for i, result := range output.Results {
		if result.Id == id {
			output.Results = remove(output.Results, i)
		}
	}
	writeTelephoneToFile(output)
}

func remove(slice []Input, s int) []Input {
	return append(slice[:s], slice[s+1:]...)
}

func writeTelephoneToFile(output Output) {
	newOutput, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("data/telephones.json", newOutput, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
