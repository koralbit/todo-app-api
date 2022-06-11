package entities

type TodoList struct {
	Id          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:20"`
	Description string `gorm:"size:50"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli"`
}
