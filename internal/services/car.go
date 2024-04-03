package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"parking-app-mobile-api-gateway/internal/utils"
)

type CarService struct {
	URL string
}

func NewCarService() *CarService {
	url, _ := utils.GetEnvVar("CAR_SERVICE_URL")

	return &CarService{
		URL: url,
	}
}

type Car struct {
	ID                string `json:"id"`
	Owner             string `json:"owner"`
	RegistrationPlate string `json:"registrationPlate"`
}

type CreateCarRequest struct {
	Owner             string `json:"owner"`
	RegistrationPlate string `json:"registrationPlate"`
}

func (c *CarService) GetCars() ([]Car, error) {

	req, err := http.NewRequest("GET", c.URL+"/cars/", nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var cars []Car

	err = json.NewDecoder(resp.Body).Decode(&cars)
	if err != nil {
		return make([]Car, 0), err
	}

	return cars, nil
}

func (c *CarService) CreateCar(createCarRequest CreateCarRequest) error {

	payload, err := json.Marshal(createCarRequest)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.URL+"/cars/", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *CarService) GetCarByID(id string) (*Car, error) {

	req, err := http.NewRequest("GET", c.URL+"/cars/"+id, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var car Car

	err = json.NewDecoder(resp.Body).Decode(&car)
	if err != nil {
		return nil, err
	}

	return &car, nil
}

func (c *CarService) DeleteCarByID(id string) error {

	req, err := http.NewRequest("DELETE", c.URL+"/cars/"+id, nil)
	if err != nil {
		return err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
