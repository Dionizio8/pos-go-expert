package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/Dionizio8/pos-go-expert/multithreading/configs"
	"github.com/Dionizio8/pos-go-expert/multithreading/internal/dto"
	"github.com/Dionizio8/pos-go-expert/multithreading/internal/entity"
)

const (
	BRASILAPI = "brasil_api"
	VIACEP    = "via_cep"
)

type CMDResponseAddress struct {
	Address   entity.Address `json:"address"`
	API       string         `json:"api"`
	CreatedAt time.Time      `json:"created_at"`
}

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic("unable to initialize settings")
	}

	responseBrasilApi := make(chan *CMDResponseAddress)
	responseViaCep := make(chan *CMDResponseAddress)
	errorHandler := make(chan error, 2)

	cep := os.Args[1]

	go getAddress(BRASILAPI, config.APIURLBrasilApi, cep, responseBrasilApi, errorHandler)
	go getAddress(VIACEP, config.APIURLViacep, cep, responseViaCep, errorHandler)

	select {
	case msg1 := <-responseBrasilApi:
		resp, _ := json.Marshal(msg1)
		println(string(resp))
	case msg2 := <-responseViaCep:
		resp, _ := json.Marshal(msg2)
		println(string(resp))
	case <-time.After(time.Second * 1):
		println("response timeout exceeded")
	case errors := <-errorHandler:
		println("errors: " + errors.Error())
	}

}

func getAddress(name, url, cep string, api chan<- *CMDResponseAddress, apierr chan<- error) {
	url = fmt.Sprintf(url, cep)
	req, err := http.NewRequestWithContext(context.Background(), "GET", url, nil)
	if err != nil {
		apierr <- err
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		apierr <- err
		return
	}

	if resp.StatusCode != http.StatusOK {
		apierr <- errors.New("not found")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		apierr <- err
		return
	}

	var address entity.Address
	switch name {
	case BRASILAPI:
		respAPI := dto.BrasilAPIOutput{}
		err = json.Unmarshal(body, &respAPI)
		if err != nil {
			apierr <- err
			return
		}
		address = respAPI.ToEntity()
	case VIACEP:
		respAPI := dto.ViacepAPIOutput{}
		err = json.Unmarshal(body, &respAPI)
		if err != nil {
			apierr <- err
			return
		}

		if respAPI.Cep == nil {
			apierr <- errors.New("not found")
			return
		}

		address = respAPI.ToEntity()
	}

	cmdResp := CMDResponseAddress{
		Address:   address,
		API:       name,
		CreatedAt: time.Now(),
	}

	api <- &cmdResp
}
