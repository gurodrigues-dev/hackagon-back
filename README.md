
# 🚀 Inicializando

1. Primeiro builde a imagem do Docker.

```sh
docker build -t hackagon-back .
```
2. Antes de subir a imagem na sua máquina, é necessário um arquivo `config.yaml`, com todas as configurações sensíveis solicitadas no arquivo `config.go`

- Informações do Servidor, Banco de Dados, Redis & AWS.  

3. Agora builde a imagem.

```sh
docker run -p 9672:9672 -v /path/config/config.yaml:/app/config/config.yaml hackagon-back:latest
```

## ⚙️ API Endpoints

Por padrão a API é executada na porta `9632` localmente.

A API usa o formato JSON no conteúdo do `body` seguindo os princípios REST.

### POST /user

Cria uma conta na API

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `nickname` | body | string  | Nickname do usuário. |
| `email`| body | string  | E-mail do usuário. |
| `senha`| body | string  | Senha do usuário. |      

**Resposta**

```json
{
    "message": "Usuário criado com sucesso"
}
```

---

### GET /user/\<uuid>

Pegue as informações do Usuário

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `nickname` | body | string  | Nickname do usuário. |
| `email`| body | string  | E-mail do usuário. |
| `senha`| body | string  | Senha do usuário. |      

**Resposta**

```json
{
    "message": "Usuário criado com sucesso",
    "user": {
        "email": "pass",
        "nome": "mwb",
        "user_id": 1
    }
}
```

### UPDATE /user

Atualize as informações do Usuário

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `nickname` | body | string  | Nickname do usuário. |
| `email`| body | string  | E-mail do usuário. |
| `senha`| body | string  | Senha do usuário. |      

**Resposta**

```json
{
    "message": "Usuário atualizado com sucesso"
}
```


---

### DELETE /user

Delete as informações do Usuário

**Parâmetros**

- Não são necessários parâmetros, apenas o JWT.

**Resposta**

```json
{
    "message": "Usuário deletado com sucesso"
}
```
---

### POST /question

Cria uma questão na API

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `title` | body | string  | Titulo da questão. |
| `description`| body | string  | Descrição da questão. |
| `date`| body | date  | Data de apresentação da questão. |      
| `level`| body | string  | Level da questão. |   
| `param1`| body | string  | Param1 da função. |  
| `param2`| body | string  | Param2 da função. |
| `response1`| body | string  | Resposta final da função. |
| `username`| body | string  | Username do Cognito. |
| `cognito`| body | string  | Password do Cognito. |

**Resposta**

```json
{
	"id": "018e1a74-29cd-79d4-ad07-cfea2c9f3f18",
	"title": "Soma de Casal5",
	"description": "José é casado com Maisa e querem ver juntos quanto de patrimônio possuem juntos, crie uma função que receba dois parametros e os some.",
	"date": "2024-03-07",
	"level": "easy",
	"inputs": {
		"test1": {
			"params": [
				"2",
				"2"
			],
			"response": "4"
		},
		"test2": {
			"params": [
				"3",
				"3"
			],
			"response": "6"
		},
		"test3": {
			"params": [
				"4",
				"4"
			],
			"response": "8"
		}
	},
	"username": "qzKMIlyh",
	"password": "BpNabUrZOs"
}
```

---

### GET /question

Pegue as informações da questão

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `title` | body | string  | Titulo da questão. |
| `description`| body | string  | Descrição da questão. |
| `date`| body | date  | Data de apresentação da questão. |      
| `level`| body | string  | Level da questão. |   
| `param1`| body | string  | Param1 da função. |  
| `param2`| body | string  | Param2 da função. |
| `response1`| body | string  | Resposta final da função. |
| `username`| body | string  | Username do Cognito. |
| `cognito`| body | string  | Password do Cognito. |  

**Resposta**

```json
{
	"id": "018e1a74-29cd-79d4-ad07-cfea2c9f3f18",
	"title": "Soma de Casal5",
	"description": "José é casado com Maisa e querem ver juntos quanto de patrimônio possuem juntos, crie uma função que receba dois parametros e os some.",
	"date": "2024-03-07",
	"level": "easy",
	"inputs": {
		"test1": {
			"params": [
				"2",
				"2"
			],
			"response": "4"
		},
		"test2": {
			"params": [
				"3",
				"3"
			],
			"response": "6"
		},
		"test3": {
			"params": [
				"4",
				"4"
			],
			"response": "8"
		}
	},
	"username": "qzKMIlyh",
	"password": "BpNabUrZOs"
}
```

### UPDATE /question

Atualize as informações da Questão

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `title` | body | string  | Titulo da questão. |
| `description`| body | string  | Descrição da questão. |
| `date`| body | date  | Data de apresentação da questão. |      
| `level`| body | string  | Level da questão. |   
| `param1`| body | string  | Param1 da função. |  
| `param2`| body | string  | Param2 da função. |
| `response1`| body | string  | Resposta final da função. |
| `username`| body | string  | Username do Cognito. |
| `cognito`| body | string  | Password do Cognito. |    

**Resposta**

```json
{
    "message": "Questão atualizada com sucesso"
}
```


---

### DELETE /question/\<uuid>

Delete as informações da Questão

**Parâmetros**

- Não são necessários parâmetros, apenas o PATH na URL.

**Resposta**

```json
{
    "message": "Questão deletado com sucesso"
}
```

---

### POST /user/login

Realiza o login na aplicação

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `nickname`| body | string  | Nickname do usuário. |
| `senha`| body | string  | Senha do usuário. |      

**Resposta**

```json
{
    "message": "Informações de login corretas",
    "user": {
        "email": "pass",
        "nome": "mwb",
        "user_id": 1
    },
    "jwt":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDk5Mzk2OTgsIm5pY2tuYW1lIjoiZmFicmljaW8ifQ.6nO4kR3mbmIfFfF5NmGtPJikS7Pjo-xKWChkkqH7jVQ"
}
```

---

### GET /ranking

Busca a classificação geral de todos os usuários.

**Parâmetros**

- Não são necessários parâmetros, apenas a requisição.  

**Resposta**

```json
[
	{
		"nickname": "rodrigues",
		"points": 130,
		"position": 1
	},
	{
		"nickname": "ozzett",
		"points": 15,
		"position": 2
	},
	{
		"nickname": "guiaugusto",
		"points": 10,
		"position": 3
	},
	{
		"nickname": "carodoso",
		"points": 0,
		"position": 4
	}
]
```

---

### POST /answer

Registra uma nova reposta do usuário.

**Parâmetros**

| Nome | Local | Tipo | Descrição
|-------------:|:--------:|:-------:| --- |
| `status`| body | string  | Estado da resposta. |
| `question_id`| body | string  | UUid da questão. |
| `points`| body | string  | Quantidade de pontos da questão. |

**Resposta**

```json
{
	"message": "answer created"
}
```
