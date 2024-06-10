# Microservices Project

Este projeto implementa um sistema de mensageria utilizando Docker para conteinerização, Kafka como sistema de mensageria, um produtor em Python e um consumidor em Go.

## Requisitos

Certifique-se de ter o Docker e o Docker Compose instalados em sua máquina.

## Instruções para Executar a Aplicação

1. Clone o repositório:

   git clone https://github.com/jrodrigues707/microservices-project.git

2. Navegue até o diretório do projeto:

   cd microservices-project

3. Execute o comando a seguir para iniciar todos os serviços:

   docker-compose up --build

4. O produtor em Python começará a enviar mensagens para o Kafka.

5. O consumidor em Go começará a receber e exibir as mensagens no console.

## Estrutura do Projeto

- producer.py: Script em Python que gera e envia mensagens fictícias para o Kafka.
- consumer.go: Script em Go que consome e exibe as mensagens do Kafka.
- Dockerfile: Arquivos Docker para conteinerizar as aplicações de produtor e consumidor.
- docker-compose.yml: Arquivo para orquestrar os serviços com Docker Compose.
