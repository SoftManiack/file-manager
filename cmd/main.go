package main

import (
	filemanager "file-manager"
	"file-manager/pkg/handler"
	"file-manager/pkg/repository"
	"file-manager/pkg/service"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Println(err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(filemanager.Server)

	logrus.Print("App Start")

	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

// go mod init file-manager
// docker run --name psql-container -p 5432:5432 -e POSTGRES_PASSWORD=123 -e PGDATA=/var/lib/postgresql/data/pgdata -d postgres
// docker exec -it 25edc7ab257a bash
// psql -h localhost -p 5432 -U postgres -W
// C:\Users\Артем\scoop\apps\migrate\4.15.1\migrate create -ext sql -dir ./schema -seq init
