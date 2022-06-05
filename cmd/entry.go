package cmd

import (
	"context"
	"github.com/karankumarshreds/go-hex-tests/env"
)

func start() {
	// App context
	ctx := context.Background()

	// Env file
	_env := env.GetEnv(".env.development")

	// Run the database on the basis of the environment config

}
