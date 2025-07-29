package entity

type Service struct {
	// service uuid
	ID string `gorm:"id;primaryKey;type:uuid"`
	// service name
	Name string `gorm:"name;not null;uniqueIndex"`
}
