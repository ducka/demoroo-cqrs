package main

import (
	"github.com/ducka/demoroo/internal/core/config"
	"github.com/ducka/demoroo/internal/core/logging"
	"github.com/ducka/demoroo/internal/domain"
	"github.com/ducka/demoroo/internal/grpc"
	"go.uber.org/fx"
)

func main() {
	// NB: The intention here is to keep application registration clean and semantic. Dependencies
	// should be organised in a modular way, to promote easy reuse across different kinds of application
	// hosts, and to make future refactors of the service easier.
	fx.New(
		// Composition of crosscutting concerns
		config.RegisterConfig("grpc", grpc.ServerSettings{}),
		logging.RegisterLogger(),

		// Composition of the application host
		grpc.RegisterServer(),

		// Composition of the application logic itself
		domain.RegisterApplication(),
	).Run()
}
