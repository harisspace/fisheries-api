package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

const (
	POSTGRES_USERNAME   = "POSTGRES_USERNAME"
	POSTGRES_PASSWORD   = "POSTGRES_PASSWORD"
	POSTGRES_HOST       = "POSTGRES_HOST"
	POSTGRES_DB_NAME    = "POSTGRES_DB_NAME"
	POSTGRES_PORT       = "POSTGRES_PORT"
	HTTP_PORT           = "HTTP_PORT"
	APPLICATION_NAME    = "APPLICATION_NAME"
	BASIC_AUTH_USERNAME = "BASIC_AUTH_USERNAME"
	BASIC_AUTH_PASSWORD = "BASIC_AUTH_PASSWORD"
)

type Env struct {
	PostrgresUsername string
	PostrgresPassword string
	PostgresHost      string
	PostgresDBName    string
	PostgresPort      uint32
	RootApp           string
	Port              uint32
	BasicAuthUsername string
	BasicAuthPassword string
	ApplicationName   string
}

var GlobalEnv Env

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable load env file: %v", err)
	}

	var ok bool

	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Can't get working directory: %v", err)
	}

	rootApp := strings.TrimSuffix(path, "config")

	os.Setenv("APP_PATH", rootApp)
	GlobalEnv.RootApp = rootApp

	if port, err := strconv.Atoi(os.Getenv(HTTP_PORT)); err != nil {
		panic(fmt.Sprintf("Missing %s environment", HTTP_PORT))
	} else {
		GlobalEnv.Port = uint32(port)
	}

	GlobalEnv.PostrgresUsername, ok = os.LookupEnv(POSTGRES_USERNAME)
	if !ok {
		panic(fmt.Sprintf("Missing %s environment", POSTGRES_USERNAME))
	}

	GlobalEnv.PostrgresPassword, ok = os.LookupEnv(POSTGRES_PASSWORD)
	if !ok {
		panic(fmt.Sprintf("Missing %s environment", POSTGRES_PASSWORD))
	}

	GlobalEnv.PostgresDBName, ok = os.LookupEnv(POSTGRES_DB_NAME)
	if !ok {
		panic(fmt.Sprintf("Missing %s environment", POSTGRES_DB_NAME))
	}

	if dbPort, err := strconv.Atoi(os.Getenv(POSTGRES_PORT)); err != nil {
		panic(fmt.Sprintf("Missing %s environtment", POSTGRES_PORT))
	} else {
		GlobalEnv.PostgresPort = uint32(dbPort)
	}

	GlobalEnv.PostgresHost, ok = os.LookupEnv(POSTGRES_HOST)
	if !ok {
		panic(fmt.Sprintf("Missing %s environment", POSTGRES_HOST))
	}

	GlobalEnv.ApplicationName, ok = os.LookupEnv(APPLICATION_NAME)
	if !ok {
		panic(fmt.Sprintf("Missing %s environment", APPLICATION_NAME))
	}

	GlobalEnv.BasicAuthUsername, ok = os.LookupEnv(BASIC_AUTH_USERNAME)
	if !ok {
		panic(fmt.Sprintf("Missing %s environment", BASIC_AUTH_USERNAME))
	}

	GlobalEnv.BasicAuthPassword, ok = os.LookupEnv(BASIC_AUTH_PASSWORD)
	if !ok {
		panic(fmt.Sprintf("Missing %s environment", BASIC_AUTH_PASSWORD))
	}
}
