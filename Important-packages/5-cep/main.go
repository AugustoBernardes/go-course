package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	for _, cep := range os.Args[1:] {

		req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro na requisição %v\n", err)
			continue
		}
		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro na conversão do arquivo %v\n", err)
			continue
		}

		var cepData CepResponse
		err = json.Unmarshal(res, &cepData)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Falha ao converter CEP para struct %v\n", err)
			continue
		}

		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Falha ao converter criararquivo %v\n", err)
			continue
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s , Bairro: %s, Uf: %s, Cidade: %s", cepData.Cep, cepData.Bairro, cepData.Uf, cepData.Localidade))
		fmt.Printf("Arquivo criado! \n")
		fmt.Println("Cidade: ", cepData.Localidade)

	}
}
