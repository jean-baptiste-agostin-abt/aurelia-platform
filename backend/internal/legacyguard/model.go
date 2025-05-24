package legacyguard

import "github.com/jinzhu/gorm"

type LegacyGuard struct {
	gorm.Model
	CapsuleID uint
	Triggered bool
}
