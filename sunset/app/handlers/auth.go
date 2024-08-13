package handlers

import (
	"github.com/HsiaoCz/code-beast/sunset/app/types"

	"github.com/anthdm/superkit/kit"
)

func HandleAuthentication(kit *kit.Kit) (kit.Auth, error) {
	return types.AuthUser{}, nil
}
