package system

import (
	"world-engine/component"

	"github.com/argus-labs/monorepo/pkg/ecs"
)

type ChatSearch = ecs.Exact[struct {
	UserTag ecs.Ref[component.UserTag]
	Chat    ecs.Ref[component.Chat]
}]
