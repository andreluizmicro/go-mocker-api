# GO MOCKER API

API desenvolvida em GO Lang que tem por objetivo, facilitar testes para desenvolvedores, QAs e etc. A API é responsável por criar mocks genéricos para testes de APIS, frontends, testes entre outros.

Trata-se de uma API que visa a simplicidade, seu funcionamento é extremamente simples, possui apenas dois endpoints:

### Funcionamento

para subir a aplicação basta rodar o arquivo executável com o comando abaixo:

> **Nota:** A porta `9006` indicada no comando é opcional, e pode ser informada qualquer porta, caso não informada, a aplicação subirá por padrão na porta `9000`.

    go run cmd/go-mocker-api/main.go 9002 validation


### Exemplo de requisições

- `POST /mock`: Cria um novo payload

Na requisição POST não existe um contrato de payload, pode ser passado qualquer tipo de payload desde que seja um json válido.

```json
{
    "name": "André Luiz",
    "email": "andreluizmicro@gmail.com"
}
```

- `GET /mock`: Retorna o último payload cadastrado

```json
{
    "name": "André Luiz",
    "email": "andreluizmicro@gmail.com"
}
```