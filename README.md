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

## Passo a Passo

1. Clone o repositório.
2. Adicione um arquivo .env na raiz do projeto com as credenciais:
     MAIL_URL=smtp.seu-dominio.com
     MAIL_PORT=587
     MAIL_USERNAME=seu-email@dominio.com
     MAIL_PASSWORD=sua-senha
     DB_HOST=localhost
     DB_PORT=5432
     DB_NAME=templates
     DB_USER=postgres
     DB_PASSWORD=postgres
4. Instale as dependências com `go mod tidy`
5. Execute com `go run main.go`
6. Publique uma mensagem seguindo este padrão:
     {
        "to": "email.destino@dominio.com",
        "templateId": 2,
        "subject": "Assunto"
     }

Assim, sempre que publicar uma mensagem no RabbitMQ, esta será lida e será enviado um e-mail automático para o destinatário informado!
