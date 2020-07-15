package utils

import "github.com/rs/xid"

// GenerateUUID ...
func GenerateUUID() string {
	uuid := xid.New()
	return uuid.String()
}
