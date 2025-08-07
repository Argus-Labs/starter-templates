package system

import (
	"chat/component"

	"github.com/argus-labs/go-ecs/pkg/ecs"
)

type ChatSearch = ecs.Exact[struct {
	UserTag ecs.Ref[component.UserTag]
	Chat    ecs.Ref[component.Chat]
}]
