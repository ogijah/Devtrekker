package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Input struct {
	Id        int64  `json:"id"`
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
	writeTelephoneToFile(newInput)
}

func GetTelephones() Output {
	var output Output
	var results []Input
	content, err := ioutil.ReadFile("telephones.json")
	if err != nil {
		return output
	}
	err = json.Unmarshal(content, &results)
	if err != nil {
		log.Fatal(err)
	}
	output.Results = results
	return output
}

func GetTelephoneById(id int64) Input {
	output := GetTelephones()
	for _, result := range output.Results {
		if result.Id == id {
			return result
		}
	}
	return Input{}
}

func DeleteTelephone(id int64) {
	output := GetTelephones()
	for i, result := range output.Results {
		if result.Id == id {
			output.Results = remove(output.Results, i)
		}
	}

}

func remove(slice []Input, s int) []Input {
	return append(slice[:s], slice[s+1:]...)
}

func writeTelephoneToFile(input Input) {
	output := GetTelephones()
	output.Results = append(output.Results, input)
	newOutput, err := json.Marshal(output)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("telephones.json", newOutput, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
