# Email Service com RabbitMQ

Este é um projeto em Go para envio de emails. O serviço se conecta ao **RabbitMQ** para consumir mensagens e processar o envio de emails com base nas informações recebidas.

## Funcionalidades

- **Envio de Emails:** Envia emails utilizando templates.
- **Integração com RabbitMQ:** Consome mensagens de uma fila RabbitMQ para processar e enviar emails de forma assíncrona.
- **Templates de Email:** Suporte para o envio de emails com templates personalizados.

## Pré-requisitos

Antes de executar este projeto, certifique-se de ter o seguinte:

1. **Go** - [Instale o Go](https://golang.org/dl/)
2. **RabbitMQ** - Um servidor RabbitMQ em execução. Você pode usar [Docker](https://www.docker.com/) para subir o RabbitMQ facilmente, ou configurar uma instância própria.
3. **Git** - Para clonar o repositório.
