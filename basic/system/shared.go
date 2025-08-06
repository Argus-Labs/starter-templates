package system

import (
	"world-engine/component"

	"github.com/argus-labs/monorepo/pkg/ecs"
)

type PlayerSearch = ecs.Exact[struct {
	Tag    ecs.Ref[component.PlayerTag]
	Health ecs.Ref[component.Health]
}]

type GraveSearch = ecs.Exact[struct {
	Grave ecs.Ref[component.Gravestone]
}]
