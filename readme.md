# Aplicação CRUD com Golang + Fiber (Web Framework) + MongoDB & Docker

Desenvolvimento de aplicação CRUD utilizando Golang + Fiber + MongoDB e Docker. <br>
Implementação de GitHub Actions e deploy no **[Microsoft Azure](https://azure.microsoft.com/pt-br/)**!

## Tecnologias Utilizadas 🚀
* **[Visual Studio Code](https://code.visualstudio.com/)**
* **[Golang - Visual Studio Code (Extensão)](https://code.visualstudio.com/docs/languages/go)**
* **[Golang](https://golang.org/)**
* **[Fiber](https://gofiber.io/)**
* **[MongoDB](https://github.com/mongodb/mongo-go-driver)**
* **[Robo3T](https://robomongo.org/)**
* **[Docker](https://www.docker.com/)**
* **[GitHub Actions](https://github.com/features/actions)**
* **[Microsoft Azure](https://azure.microsoft.com/pt-br/)**

## Como Executar? 🔥
1. Para executar o projeto local devemos ter o Golang instalado na máquina e executar o seguinte comando: 
```
> go run main.go
```
2. Após a execução do comando podemos ir ao Postman ou Browser de preferência e digitar: `http://localhost:3000/api`
## Visual Studio Code (Debug) 🐛
1. Para habilitar o debug no Visual Studio Code devemos criar o arquivo: `launch.json` na pasta `.vscode` e adicionar o seguinte conteúdo:
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "env": {
                "Environment": "Development",
                "ConnectionString": "mongodb://localhost:27017",
                "DbName": "persondb",
                "CollectionName": "persons"
            },
            "args": []
        }
    ]
}
```
## Docker 🐋
1. Para criação da imagem docker da nossa aplicação primeiro devemos criar o arquivo `Dockerfile`, após criação e configuração arquivo devemos executar o seguinte comando: 
```
> docker build -t jailtonjunior/api-person:v1 .
```
2. Para executar o container local devemos executar o seguinte comando:
```
> docker run --name api-person -p 3000:3000 -e Environment=Development -d jailtonjunior/api-person:v1
```
3. Para que possamos fazer o push (subir) nossa imagem para o Docker Hub, primeiramente precisamos ter uma conta e após criação da conta devemos fazer o login com: 
```
> docker login
```
4. Para fazer o push da imagem que criamos devemos executar o seguinte comando: 
```
> docker push jailtonjunior/api-person:v1
```
## Microsoft Azure ☁️
### Em desenvolvimento 🔨