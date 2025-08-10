package enum

type TYPE string
type STATUS string
type ROLE string

const (
	MANGA  TYPE = "manga"
	MANHUA TYPE = "manhua"
	MANHWA TYPE = "manhwa"

	COMPLETED STATUS = "completed"
	HIATUS    STATUS = "hiatus"
	ONGOING   STATUS = "ongoing"

	ADMIN ROLE = "admin"
	USER  ROLE = "user"
)
