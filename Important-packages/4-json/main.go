package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Account struct {
	Name string `json:"n"` //Usando essas tags a struct consegue saber que se o JSON tiver n , é o valor Name
	Fund int    `json:"f"`
}

func main() {

	account := Account{
		Name: "Augusto",
		Fund: 10,
	}

	// Dessa forma utilizando o marshal estamos pegando e declarando esse valor em uma variavel
	res, err := json.Marshal(account)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))

	/*
		Se não precisarmos da resposta em uma variavel podemos usar o encoder.
		o NewEncoder - definimos onde saira o resultado. No caso escolhemos o terminal
	*/
	err = json.NewEncoder(os.Stdout).Encode(account)
	if err != nil {
		panic(err)
	}

	/*
		Na forma abaixa , estamos fazendo ao contrario, pegando Json puro (Em go fica salvo em bytes).

		E convertendo para struct
		Obs: Se vc ver a variavel accountX foi declarada , mas é usado o & para apontar pro endereço da memoria, e salvar o dado

	*/

	// Forma normal
	// pureJson := []byte(`{"Name":"Augusto2","Fund":20}`)

	// Quano Json é diferente da struct
	pureJson := []byte(`{"n":"Augusto2","f":20}`)
	var accountX Account

	err = json.Unmarshal(pureJson, &accountX)
	if err != nil {
		panic(err)
	}

	fmt.Println(accountX)

}
