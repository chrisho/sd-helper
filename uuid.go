package sdhelper

import "github.com/rs/xid"

func GenerateUuid() string {
	return xid.New().String()
}