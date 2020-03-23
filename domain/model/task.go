package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// 一番内側のDomain層に Interface な層の gorm をimport している
// gorm.Model は結局 ID, CreatedAt... とかなものなので、剥がしたくなったときに同様なものをここで定義すれば良さそう
// `gorm:"..."` な表記も tag でしかないため、内側のロジックに影響を与えない

type Task struct {
	gorm.Model
	Title  string     `gorm:"NOT NULL"`
	Done   bool       `gorm:"NOT NULL"`
	DoneAt *time.Time `gorm:"DEFAULT NULL"`
}
