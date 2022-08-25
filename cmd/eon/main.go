package main

import (
	"fmt"
	"log"

	"github.com/alwindoss/eon"
	"github.com/alwindoss/eon/internal/engine"
	"github.com/caarlos0/env/v6"
)

func main() {
	cfg := eon.Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	fmt.Printf("%+v\n", cfg)
	log.Fatal(engine.Run(&cfg))
}
