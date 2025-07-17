// App binary starts HTTP-server.
package main

import (
	"github.com/sirupsen/logrus"

	"SubscriptionAggregator/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		logrus.Fatal(err)
	}
}
