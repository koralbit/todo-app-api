package entities

type TodoItem struct {
	Id          uint   `gorm:"primaryKey"`
	Description string `gorm:"size:50"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli"`
	TodoListId  uint
	TodoList    TodoList
}
