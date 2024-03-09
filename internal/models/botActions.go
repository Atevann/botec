package models

// BotAction Сущность экшена
type BotAction struct {
	Id          uint
	BotId       uint
	ParentId    uint
	Name        string
	Condition   string
	Action_name string
	Action_data string
}
