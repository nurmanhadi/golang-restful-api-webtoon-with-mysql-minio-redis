package enum

type TYPE string
type STATUS string
type ROLE string

const (
	TYPE_MANGA  TYPE = "manga"
	TYPE_MANHUA TYPE = "manhua"
	TYPE_MANHWA TYPE = "manhwa"

	STATUS_COMPLETED STATUS = "completed"
	STATUS_HIATUS    STATUS = "hiatus"
	STATUS_ONGOING   STATUS = "ongoing"

	ROLE_ADMIN ROLE = "admin"
	ROLE_USER  ROLE = "user"
)
