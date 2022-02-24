package database

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var Session *session.Session

func LoadAWS() (err error) {
	Session, err = session.NewSession(
		&aws.Config{
			Region: aws.String(os.Getenv("REGION")),
		},
	)

	if err != nil {
		log.Fatal(fmt.Println(err))
		return err
	}

	return nil
}
