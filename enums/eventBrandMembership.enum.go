package enums

type EventMembershipRole int

const (
	Admin EventMembershipRole = iota
	Contributor
	Attendee
)

func GetRoleDescription(role int) string {
	switch role {

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
