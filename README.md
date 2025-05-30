# Consultor de CEP

Um programa simples em Go que consulta informações de CEP utilizando a API do ViaCEP e salva os dados em um arquivo de texto.

## 📋 Funcionalidades

- Consulta informações de CEP através da API do ViaCEP
- Aceita múltiplos CEPs como argumentos
- Salva os dados do último CEP consultado em um arquivo `cidade.txt`
- Exibe as informações completas no terminal

## 🚀 Pré-requisitos

- Go 1.16 ou superior instalado
- Conexão com a internet para acessar a API do ViaCEP

## 📦 Instalação

1. Clone este repositório:

```bash
git clone https://github.com/seu-usuario/consultor-cep.git
```
cd consultor-cep

2. Certifique-se de que o Go está instalado:

```bash
go version
```

## 🔧 Como usar

### Execução básica

Para consultar um CEP específico:

```bash
go run main.go 12345678 //Use um CEP válido
```

### Consultando múltiplos CEPs

Você pode consultar vários CEPs de uma vez:

```bash
go run main.go 12345678 87654321 //Use CEPs válidos
```

### Compilando o programa

Para compilar o programa e gerar um executável:

```bash
go build -o consultor-cep main.go
```

Depois execute:

```bash
./consultor-cep 12345678 //Use um CEP válido
```

## 📝 Exemplo de uso

```bash
go run main.go 01310100
```
**Saída no terminal:**

```json
  {
  "cep": "01310-100",
  "logradouro": "Avenida Paulista",
  "complemento": "",
  "bairro": "Bela Vista",
  "localidade": "São Paulo",
  "uf": "SP",
  "ibge": "3550308",
  "gia": "1004",
  "regiao": "Sudeste",
  "estado": "São Paulo"
  }
```
**Arquivo criado (`cidade.txt`):**

```bash
CEP: 01310-100, Logradouro: Avenida Paulista, Bairro: Bela Vista, Localidade: São Paulo, UF: SP
```

## 📁 Estrutura do projeto

```bash
├── main.go # Arquivo principal do programa
├── README.md # Este arquivo
└── cidade.txt # Arquivo gerado com os dados do CEP (criado após execução)
```


## ⚙️ Como funciona

1. O programa recebe um ou mais CEPs como argumentos da linha de comando
2. Para cada CEP, faz uma requisição HTTP para a API do ViaCEP
3. Processa a resposta JSON e mapeia para a estrutura `ViaCEP`
4. Salva as informações do último CEP consultado no arquivo `cidade.txt`
5. Exibe a resposta completa da API no terminal

## 🛠️ Estrutura de dados

O programa utiliza a seguinte estrutura para mapear os dados do ViaCEP:

```go
type ViaCEP struct {
  Uf string json:"uf"
  Cep string json:"cep"
  Gia string json:"gia"
  Ibge string json:"ibge"
  Bairro string json:"bairro"
  Estado string json:"estado"
  Unidade string json:"unidade"
  Regiao string json:"regiao"
  Logradouro string json:"logradouro"
  Localidade string json:"localidade"
  Complemento string json:"complemento"
}
```


## ⚠️ Observações importantes

- O arquivo `cidade.txt` é sobrescrito a cada execução
- Se consultar múltiplos CEPs, apenas o último será salvo no arquivo
- CEPs inválidos podem retornar erro ou dados vazios
- É necessário conexão com a internet para funcionar

## 🐛 Tratamento de erros

O programa trata os seguintes tipos de erro:
- Falha na requisição HTTP
- Erro na leitura da resposta
- Erro no parsing do JSON
- Erro na criação/escrita do arquivo

## 🤝 Contribuindo

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -am 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo `LICENSE` para mais detalhes.

## 🔗 API utilizada

Este projeto utiliza a API gratuita do [ViaCEP](https://viacep.com.br/) para consulta de informações de CEP.