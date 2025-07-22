package lib

import (
	"log"
	"os"
)

type Env struct {
	DB_SSL_MODE    string
	DB_CLIENT_CERT string

	PG_HOST     string
	PG_USER     string
	PG_PASSWORD string
	PG_DB       string
	PG_PORT     string

	APP_PORT    string
	SOCKET_PORT string

	JWT_ACCESS_KEY  string
	JWT_REFRESH_KEY string

	GH_APP_SLUG          string
	GH_APP_ID            string
	GH_APP_CLIENT_ID     string
	GH_APP_CLIENT_SECRET string
	GH_PRIVATE_KEY       string

	DEFAULT_ADMIN_EMAIL string
	DEFAULT_ADMIN_PASS  string

	AUTH0_DOMAIN    string
	AUTH0_CLIENT_ID string

	RABBITMQ_URL string

	K8S_CLUSTER_CONFIG string

	INGRESS_ROOT_DOMAIN string

	REDIS_URL string

	ABLY_API_KEY string

	RODELAR_URL        string
	RODELAR_API_KEY    string
	RODELAR_API_KEY_ID string

	FLUTTERWAVE_API_URL       string
	FLUTTERWAVE_SECRET_KEY    string
	FLUTTERWAVE_ENCRYTION_KEY string
}

func NewEnv() Env {
	var secrets, err = SecretsFactory(UseDefault(os.Getenv("SECRET_FROM"), "local"))
	if err != nil {
		log.Println(err)
	}

	return Env{
		DB_SSL_MODE:    UseDefault(secrets.Get("DB_SSL_MODE"), "require"),
		DB_CLIENT_CERT: secrets.Get("DB_CLIENT_CERT"),

		PG_HOST:     secrets.GetOrPanic("PG_HOST"),
		PG_USER:     secrets.GetOrPanic("PG_USER"),
		PG_PASSWORD: secrets.GetOrPanic("PG_PASSWORD"),
		PG_DB:       secrets.GetOrPanic("PG_DB"),
		PG_PORT:     secrets.GetOrPanic("PG_PORT"),

		APP_PORT:    UseDefault(secrets.Get("APP_PORT"), "3000"),
		SOCKET_PORT: UseDefault(secrets.Get("SOCKET_PORT"), "9090"),

		JWT_ACCESS_KEY:  secrets.GetOrPanic("JWT_ACCESS_KEY"),
		JWT_REFRESH_KEY: secrets.GetOrPanic("JWT_REFRESH_KEY"),

		GH_APP_SLUG:          secrets.GetOrPanic("GH_APP_SLUG"),
		GH_APP_ID:            secrets.GetOrPanic("GH_APP_ID"),
		GH_APP_CLIENT_ID:     secrets.GetOrPanic("GH_APP_CLIENT_ID"),
		GH_APP_CLIENT_SECRET: secrets.GetOrPanic("GH_APP_CLIENT_SECRET"),
		GH_PRIVATE_KEY:       secrets.GetOrPanic("GH_PRIVATE_KEY"),

		DEFAULT_ADMIN_EMAIL: secrets.GetOrPanic("DEFAULT_ADMIN_EMAIL"),
		DEFAULT_ADMIN_PASS:  secrets.GetOrPanic("DEFAULT_ADMIN_PASS"),

		AUTH0_DOMAIN:    secrets.GetOrPanic("AUTH0_DOMAIN"),
		AUTH0_CLIENT_ID: secrets.GetOrPanic("AUTH0_CLIENT_ID"),

		RABBITMQ_URL: secrets.GetOrPanic("RABBITMQ_URL"),
		REDIS_URL:    secrets.GetOrPanic("REDIS_URL"),

		K8S_CLUSTER_CONFIG:  secrets.GetOrPanic("K8S_CLUSTER_CONFIG"),
		INGRESS_ROOT_DOMAIN: secrets.GetOrPanic("INGRESS_ROOT_DOMAIN"),

		ABLY_API_KEY: secrets.GetOrPanic("ABLY_API_KEY"),

		RODELAR_URL:        secrets.GetOrPanic("RODELAR_URL"),
		RODELAR_API_KEY:    secrets.Get("RODELAR_API_KEY"),
		RODELAR_API_KEY_ID: secrets.Get("RODELAR_API_KEY_ID"),

		FLUTTERWAVE_API_URL:       secrets.GetOrPanic("FLUTTERWAVE_API_URL"),
		FLUTTERWAVE_SECRET_KEY:    secrets.GetOrPanic("FLUTTERWAVE_SECRET_KEY"),
		FLUTTERWAVE_ENCRYTION_KEY: secrets.GetOrPanic("FLUTTERWAVE_ENCRYTION_KEY"),
	}
}
