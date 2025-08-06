package main

import (
	"github.com/argus-labs/go-ecs/pkg/cardinal"
)

func main() {
	world := cardinal.NewWorld()

	// Register systems
	// cardinal.RegisterSystem(world, system.ExampleSystem)

	world.StartGame()
}
