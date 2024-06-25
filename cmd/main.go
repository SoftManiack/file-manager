package main

import (
	. "file-manager"
	"file-manager/pkg/handler"
	"file-manager/pkg/repository"
	"file-manager/pkg/service"
	"flag"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	var mode string

	flag.StringVar(&mode, "mode", "dev", "")
	flag.Parse()

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Println(err.Error())
	}

	if mode == "dev" {
		if err := godotenv.Load(".env.dev"); err != nil {
			logrus.Fatalf("error loading env variables %s", err.Error())
		}
	} else {
		if err := godotenv.Load(".env"); err != nil {
			logrus.Fatalf("error loading env variables %s", err.Error())
		}
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

	srv := new(Server)

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
// docker exec -it 013fb1dd0336 bash
// psql -h localhost -p 5432 -U postgres
// C:\Users\Артем\scoop\apps\migrate\4.15.1\migrate create -ext sql -dir ./schema -seq init
// docker run -it -p 80:80 \ -v /$PWD/dist/angular-nginx-docker://usr/share/nginx/html:ro \nginx:alpine
// \d+ table_name
// docker run --name nginx -p 80:80 -v /etc/nginx/sites-enabled:/etc/nginx/conf.d -v /etc/nginx/certs-enabled:/etc/nginx/certs -v /var/log/nginx:/var/log/nginx nginx
// tasklist /fi "imagename eq nginx.exe"
// nginx -s reload

//go test ./….
//go test ./... -cover
//go test ./... -coverprofile=coverage.txt
//go tool cover -html coverage.txt -o index.html

// goose -dir db/migrations postgres "postgres://postgres:123@localhost:5432/fm" down
// go mod download