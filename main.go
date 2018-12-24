package main

import (
	"github.com/joho/godotenv"
	"os"
	"github.com/astaxie/beego/orm"
	"net/url"
	"github.com/micro/go-micro"
	"fantasy-user-service/proto"
	"fantasy-user-service/handler"
	"github.com/micro/go-log"
	_ "github.com/go-sql-driver/mysql"
	k8s "github.com/micro/kubernetes/go/micro"
)

func init()  {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("no .env file")
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbDataBase := os.Getenv("DB_DATABASE")
	timeZone := "Asia/Shanghai"
	dsn:=dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbDataBase+"?charset=utf8"+"&loc="+url.QueryEscape(timeZone)
	orm.RegisterDataBase("default","mysql",dsn)
}
func main()  {

	service := k8s.NewService(
		micro.Name("fantasy.user"),
		micro.Version("latest"),
	)
	service.Init()
	user.RegisterUserServiceHandler(service.Server(),handler.NewHandler())
	if err := service.Run(); err != nil {
		log.Fatalf("failed to serve: %v",err)
	}
}
