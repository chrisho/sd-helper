package sd_helper

import "github.com/rs/xid"

func GenerateUuid() string {
	return xid.New().String()
}