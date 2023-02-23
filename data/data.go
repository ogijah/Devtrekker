package data

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Input struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstname" validate:"required"`
	LastName  string `json:"lastname"  validate:"required"`
	Telephone string `json:"telephone" validate:"required"`
}

type Output struct {
	Results []Input `json:"results"`
}

func UploadTelephone(newInput Input) error {
	if checkIfTelephoneExists(newInput.Telephone) {
		return errors.New("Phone number already exists!")
	}
	output, err := GetTelephones()
	if err != nil {
		return err
	}
	if len(output.Results) == 0 {
		newInput.Id = 1
	} else {
		newInput.Id = output.Results[len(output.Results)-1].Id + 1
	}
	output.Results = append(output.Results, newInput)
	return writeTelephoneToFile(output)
}

func GetTelephones() (Output, error) {
	var output Output
	content, err := ioutil.ReadFile("data/telephones.json")
	if err != nil {
		return output, nil
	}
	err = json.Unmarshal(content, &output)
	if err != nil {
		return output, err
	}
	return output, nil
}

func GetTelephoneById(id int) (Input, error) {
	output, err := GetTelephones()
	if err != nil {
		return Input{}, err
	}
	for _, result := range output.Results {
		if result.Id == id {
			return result, nil
		}
	}
	return Input{}, nil
}

func DeleteTelephone(id int) error {
	output, err := GetTelephones()
	if err != nil {
		return err
	}
	for i, result := range output.Results {
		if result.Id == id {
			output.Results = remove(output.Results, i)
		}
	}
	return writeTelephoneToFile(output)
}

func remove(slice []Input, s int) []Input {
	return append(slice[:s], slice[s+1:]...)
}

func writeTelephoneToFile(output Output) error {
	newOutput, err := json.Marshal(output)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("data/telephones.json", newOutput, 0644)
	if err != nil {
		return err
	}
	return nil
}

func checkIfTelephoneExists(telephone string) bool {
	output, err := GetTelephones()
	if err != nil {
		return false
	}
	for _, result := range output.Results {
		if result.Telephone == telephone {
			return true
		}
	}
	return false
}
