package enums

type EventMembershipRole int

const (
	Owner EventMembershipRole = iota
	Admin
	Contributor
	Attendee
)
