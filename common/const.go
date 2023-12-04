package common

import "time"

const MaxLengthIdCanGenerate = 12

const RoleAdminId = "admin"

const DefaultPass = "app123"

const CurrentUserStr = "current_user"

var (
	VietNamLocation = time.FixedZone("UTC+7", 7*60*60)
)
