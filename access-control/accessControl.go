package accessControl

import (
	"github.com/Swejal08/go-ggqlen/enums"
	"github.com/Swejal08/go-ggqlen/graph/services"
)

func Check(allowedRoles []enums.EventMembershipRole, userId int, eventId int) bool {

	memberShip := services.GetEventMembership(eventId, userId)

	if memberShip == nil {
		return false
	}

	var convertedRoles []string
	for _, role := range allowedRoles {
		convertedRoles = append(convertedRoles, enums.GetRoleDescription(int(role)))
	}

	hasAccess := hasAccess(convertedRoles, string(memberShip.Role))

	if !hasAccess {
		return false
	}

	return true

}

func hasAccess(roles []string, userRole string) bool {
	for _, role := range roles {
		if role == userRole {
			return true
		}
	}
	return false
}
