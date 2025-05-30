package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Uf string `json:"uf"`
	Cep string `json:"cep"`
	Gia string `json:"gia"`
	Ibge string `json:"ibge"`
	Bairro string `json:"bairro"`
	Estado string `json:"estado"`
	Unidade string `json:"unidade"`
	Regiao string `json:"regiao"`
	Logradouro string `json:"logradouro"`
	Localidade string `json:"localidade"`
	Complemento string `json:"complemento"`
}

func handleError(feedback string, err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, feedback, err)
	}
}

func main() {
	for _, cep := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + cep + "/json/")
		handleError("Erro ao fazer requisição: %v\n", err)

		defer req.Body.Close()

		res, err := io.ReadAll(req.Body)
		handleError("Erro ao ler resposta: %v\n", err)

		var data ViaCEP

		err = json.Unmarshal(res, &data)
		handleError("Erro ao fazer parse da resposta: %v\n", err)

		file, err := os.Create(cep +".txt")
		handleError("Erro ao criar arquivo: %v\n", err)

		defer file.Close()

		_, err = file.WriteString(
			fmt.Sprintf("CEP: %s, Logradouro: %s, Bairro: %s, Localidade: %s, UF: %s",
				data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.Uf,
			),
		)
		handleError("Erro ao escrever no arquivo: %v\n", err)

		fmt.Println("Arquivo criado com sucesso!")

		println(string(res))
	}
}