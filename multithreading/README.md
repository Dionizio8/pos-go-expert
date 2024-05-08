# Desafio Multithreading

## Descrição
 Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://brasilapi.com.br/api/cep/v1/01153000 + cep

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Como executar

1. entrar na pasta **multithreading/cmd**
```bash
cd multithreading/cmd
```

2. executar o arquivo **main.go**
```bash
go run main.go
```

3. responsta exemplo
```json
{
    "address": {
        "cep": "04094-050",
        "state": "SP",
        "city": "São Paulo",
        "neighborhood": "Parque Ibirapuera",
        "street": "Avenida Pedro Álvares Cabral"
    },
    "api": "via_cep",
    "created_at": "2024-05-07T22:41:53.652256-03:00"
}

## Error GOPATH

Como o projeto tem varios exemplos e dependendo de como esta configurado o GO ou editor em sua maquina recomendo abrir o projeto diretamente na pasta **multithreading**. 

1. entrar na pasta **multithreading**
```bash
cd multithreading
```

2. abrir o projeto diretamente nesta pasta (exemplo vscode)
```bash
code .
```

3. entrar na pasta **/cmd**
```bash
cd cmd
```

4. executar o arquivo **main.go** com o CEP (apenas numeros)
```bash
go run main.go 04094050
```

5. responsta exemplo
```json
{
    "address": {
        "cep": "04094-050",
        "state": "SP",
        "city": "São Paulo",
        "neighborhood": "Parque Ibirapuera",
        "street": "Avenida Pedro Álvares Cabral"
    },
    "api": "via_cep",
    "created_at": "2024-05-07T22:41:53.652256-03:00"
}
```
