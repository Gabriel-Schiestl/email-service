package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Gabriel-Schiestl/email-service/internal/application/usecases"
	"github.com/Gabriel-Schiestl/email-service/internal/config"
	"github.com/Gabriel-Schiestl/email-service/internal/domain/models/rabbitmq"
	"github.com/Gabriel-Schiestl/email-service/internal/infra/repositories"
	"github.com/Gabriel-Schiestl/email-service/internal/infra/services"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var numWorkers = 5

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar env: %v", err)
	}

	rmq := rabbitmq.NewRabbitMQ("email-service")
	defer rmq.Close()

	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		log.Fatalf("Erro ao converter porta para int")
	}

	mailConfig := config.NewSenderConfig(os.Getenv("MAIL_HOST"), os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"), port)

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Erro ao converter porta para int")
	}

	dbConfig := config.NewDbConfig(os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), dbPort)
	
	db, err := gorm.Open(postgres.Open(dbConfig.ToString()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	
	fmt.Println("Conectado ao banco de dados")

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("Erro ao obter conexão com banco: %v", err)
	}

	defer sqlDb.Close()

	senderService := services.NewEmailService(mailConfig)
	repo := repositories.NewTemplateRepository(db)
	useCase := usecases.NewSendEmailUseCase(repo, senderService)

	msgs, err := rmq.Consume()
	if err != nil {
		log.Fatalf("Falha ao consumir mensagens: %v", err)
	}

	for i := 0; i < numWorkers; i++ {
		go func() {
			for msg := range msgs {
				if err := useCase.Execute(msg); err != nil {
					log.Printf("Erro ao processar mensagem %v: %v", msg, err)
				}
			}
		}()
	}

	for msg := range msgs {
		go useCase.Execute(msg)
	}
}