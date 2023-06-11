package rabbitmq

import "fmt"

func failOnError(err error, message string) {
	if err != nil {
		panic(fmt.Sprintf("%s, err: %s", message, err))
	}
}
