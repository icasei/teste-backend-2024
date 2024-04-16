# Teste iCasei: Backend Júnior/Pleno
Esta aplicação é um sistema de cadastro de produtos que mantém e sincroniza dados em duas bases de dados distintas usando mensageria (Kafka).

## Instruções
- Faça um fork deste projeto para sua conta pessoal do GitHub.
- Siga as especificações abaixo.
- Envie o link do repositório para o e-mail backend@icasei.com.br com o título Teste Backend Júnior/Pleno
- Utilizar Insomnia para teste da aplicação (Insomnia_teste_backend.json)

## Especificações Técnicas

### Ajustes na Aplicação Rails
1. Adicione um novo campo de estoque no SQLite.
2. O novo campo deve aceitar apenas números.
3. Ajuste a mensageria ('rails-to-go') com o novo campo.
4. Crie um consumidor da mensageria do projeto em Golang ('go-to-rails').
5. O consumidor deve criar e realizar a atualização no banco.

### Ajustes na Aplicação Golang
1. Adicione um novo campo de estoque no MongoDB.
2. O novo campo deve aceitar apenas números.
3. Crie a mensageria ('go-to-rails').
4. Crie um consumidor da mensageria do projeto em Rails ('rails-to-go').
5. O consumidor deve criar e realizar a atualização no banco.


*Nota: Existe uma "falha" que caso não seja tratada, as mensageria podem entrar em loop.

## Resultado Final Esperado
Ao realizar um POST/PATCH através da aplicação RAILS, o dado deve ser salvo no banco de dados (SQLite) e gerar uma mensageria com os dados do produto. A aplicação em Golang deve consumir esses dados enviados através da mensageria, criar ou alterar o registro em seu banco de dados (Mongo). O mesmo processo deve ocorrer na aplicação em Golang, ao realizar um POST/PATCH deve realizar o disparo de uma mensageria para aplicação RAILS que irá atualizar ou criar o registro.

## O que será avaliado?
- Funcionamento
- Organização do projeto
- Lógica do código
- Uso do Git
- Uso de componentização
- Testes unitários

## É obrigatório utilizar:
- Arquitetura de microserviços com APIs RESTFUL
- Banco de dados relacional e no-sql
- Ferramenta de mensageria
- Docker
