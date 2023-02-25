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

func UploadTelephone(newInput Input, path string) (Output, error) {
	if checkIfTelephoneExists(newInput.Telephone, path) {
		return Output{}, errors.New("Phone number already exists!")
	}
	output, err := GetTelephones(path)
	if err != nil {
		return Output{}, err
	}
	if len(output.Results) == 0 {
		newInput.Id = 1
	} else {
		newInput.Id = output.Results[len(output.Results)-1].Id + 1
	}
	output.Results = append(output.Results, newInput)
	return writeTelephoneToFile(output, path)
}

func GetTelephones(path string) (Output, error) {
	var output Output
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return output, nil
	}
	err = json.Unmarshal(content, &output)
	if err != nil {
		return output, err
	}
	return output, nil
}

func GetTelephoneById(id int, path string) (Input, error) {
	output, err := GetTelephones(path)
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

func DeleteTelephone(id int, path string) (Output, error) {
	output, err := GetTelephones(path)
	if err != nil {
		return Output{}, err
	}
	for i, result := range output.Results {
		if result.Id == id {
			output.Results = remove(output.Results, i)
		}
	}
	return writeTelephoneToFile(output, path)
}

func remove(slice []Input, s int) []Input {
	return append(slice[:s], slice[s+1:]...)
}

func writeTelephoneToFile(output Output, path string) (Output, error) {
	newOutput, err := json.Marshal(output)
	if err != nil {
		return Output{}, err
	}
	err = ioutil.WriteFile(path, newOutput, 0644)
	if err != nil {
		return Output{}, err
	}
	return output, nil
}

func checkIfTelephoneExists(telephone string, path string) bool {
	output, err := GetTelephones(path)
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
