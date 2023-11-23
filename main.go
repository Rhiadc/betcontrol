package main

import (
	"fmt"

	"github.com/rhiadc/betcontrol/api"
	"github.com/rhiadc/betcontrol/config"
	"github.com/rhiadc/betcontrol/domain"
)

func main() {
	envs := config.LoadEnvVars()
	fmt.Println(envs)
	betfairService := domain.NewBetfairService(envs)
	_ = api.NewServer(*betfairService, envs)
}
