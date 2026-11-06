package model

import "time"

// Decomposition 一条需求拆解记录。
// Result 存拆解结果的 JSON 字符串(前端 zod 校验后的结构化结果)。
type Decomposition struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Requirement string    `json:"requirement"`
	Result      string    `json:"result"`
	Model       string    `json:"model"`
	CreatedAt   time.Time `json:"created_at"`
}
