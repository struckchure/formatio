package lib

import (
	"log"

	"github.com/overal-x/rodelar-go-sdk"
)

func NewRodelarClient(env Env) rodelar.IRodelarClient {
	client, err := rodelar.NewRodelarClient(
		rodelar.RodelarClientConfig{
			Url: env.RODELAR_URL,
		},
	)
	if err != nil {
		log.Println(err)
	}

	return client
}
