package accessControl

import (
	"fmt"

	"github.com/Swejal08/go-ggqlen/enums"
	"github.com/Swejal08/go-ggqlen/graph/services"
)

func Check(allowedRoles []enums.EventMembershipRole, userId string, eventId string) error {
	fmt.Println(userId, eventId)

	memberShip, err := services.GetEventMembership(eventId, userId)

	if memberShip == nil || err != nil {
		return fmt.Errorf("This resource is forbidden")
	}

	var convertedRoles []string
	for _, role := range allowedRoles {
		convertedRoles = append(convertedRoles, enums.GetRoleDescription(int(role)))
	}

	hasAccess := hasAccess(convertedRoles, string(memberShip.Role))

	if !hasAccess {
		return fmt.Errorf("This resource is forbidden")

	}

	return nil

}

func hasAccess(roles []string, userRole string) bool {
	for _, role := range roles {
		if role == userRole {
			return true
		}
	}
	return false
}
