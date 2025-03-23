# Email Service with RabbitMQ

This is a project developed in Golang in order to send emails. The service connects to **RabbitMQ** to consume messages and send emails based on the information received.

## Features

- **Sending Emails:** Send emails using templates.
- **Integration with RabbitMQ:** Consume incoming messages from a RabbitMQ queue to process and send emails asynchronously.
- **Email Templates:** Based on a database with custom templates to the emails.

## Prerequisites

Before running this project, ensure you have the following:

1. **Go** - [Install Go](https://golang.org/dl/)
2. **RabbitMQ** - A RabbitMQ server executing. You can use [Docker](https://www.docker.com/) to instantiate locally.
3. **Git** - In order to clone this repository.

## Step by step

1. Clone this repository:

```sh
     git clone https://github.com/Gabriel-Schiestl/email-service.git
```

2. Add a .env file in the root directory with these credentials:

```
   AMQP_URL=amqp://guest:guest@localhost:5672/
   MAIL_HOST=smtp.your-domain.com
   MAIL_PORT=your-port
   MAIL_USERNAME=your-email@domain.com
   MAIL_PASSWORD=your-password
   DB_HOST=localhost
   DB_PORT=5432
   DB_NAME=templates
   DB_USER=postgres
   DB_PASSWORD=postgres
```

3. Install the dependencies with `go mod tidy`
4. Run the service with `go run main.go`
5. Publish a message following this pattern(at params section you need to send all the template variables):

```
   {
   "to": "email.recipient@domain.com",
   "templateId": 2,
   "subject": "TestSubject",
   "params": {
   "token": "test"
   }
   }
```

This way, every time you publish a message on RabbitMQ, this one will be read and and email will be sent to the recipient.
