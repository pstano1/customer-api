package pkg

type Customer struct {
	Id   string `gorm:"primarykey"`
	Name string `gorm:"not null"`
	Tag  string `gorm:"unique;not null"`
}
