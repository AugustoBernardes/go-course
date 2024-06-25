package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", SearchCEPHandler)
	http.ListenAndServe(":8880", nil)
}

func SearchCEPHandler(responseWriter http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		responseWriter.WriteHeader(http.StatusBadRequest)
		return
	}
	cepParam := request.URL.Query().Get("cep")

	if cepParam != "" {
		cepData, err := GetCepData(cepParam)

		if err != nil {
			responseWriter.WriteHeader(http.StatusBadRequest)
		}
		fmt.Println(cepData)
		jsonData, err := json.Marshal(&cepData)

		responseWriter.Header().Set("Content-Type", "application/json")
		responseWriter.Write(jsonData)
		return
	}

	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(http.StatusAccepted)
	responseWriter.Write([]byte("Hello World!"))
}

func GetCepData(cep string) (*CepResponse, error) {

	response, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var cepData CepResponse
	err = json.Unmarshal(body, &cepData)

	if err != nil {
		return nil, err
	}

	return &cepData, nil
}
