package account

type STATUS string
type TYPE string

const (
	CSGO_KEY  = TYPE("CSGO_KEY")
	CSGO_CASE = TYPE("CSGO_CASE")
)

const (
	ACTIVE    = STATUS("ACTIVE")
	ACCEPTED  = STATUS("ACCEPTED")
	COUNTERED = STATUS("COUNTERED")
	EXPIRED   = STATUS("EXPIRED")
	CANCELLED = STATUS("CANCELLED")
	DECLINED  = STATUS("DECLINED")
	COMPLETED = STATUS("COMPLETED")
	UNPAID    = STATUS("UNPAID")
)
