// +build wireinject

package controllers

import (
	"github.com/google/wire"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/config"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/database"
	"github.com/keyeMyria/gin-wire-plate/internal/pkg/log"
	"github.com/keyeMyria/gin-wire-plate/internal/app/ratings/services"
	"github.com/keyeMyria/gin-wire-plate/internal/app/ratings/repositories"
)

var testProviderSet = wire.NewSet(
	log.ProviderSet,
	config.ProviderSet,
	database.ProviderSet,
	services.ProviderSet,
	//repositories.ProviderSet,
	ProviderSet,
)


func CreateRatingsController(cf string, sto repositories.RatingsRepository) (*RatingsController, error) {
	panic(wire.Build(testProviderSet))
}
