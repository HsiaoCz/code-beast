package handlers

import (
	"github.com/HsiaoCz/code-beast/sunset/app/views/landing"

	"github.com/anthdm/superkit/kit"
)

func HandleLandingIndex(kit *kit.Kit) error {
	return kit.Render(landing.Index())
}
