package enums

type EventMembershipRole int

const (
	Owner EventMembershipRole = iota
	Admin
	Contributor
	Attendee
)

func GetRoleDescription(role int) string {
	switch role {
	case int(Owner):
		return "owner"
	case int(Admin):
		return "admin"
	case int(Contributor):
		return "contributor"
	case int(Attendee):
		return "attendee"
	default:
		return "Invalid role"
	}
}
