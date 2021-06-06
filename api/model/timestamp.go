package model

import "time"

// Timestamp は、 CreatedAt と UpdatedAt の組み合わせ
// gorm ではデフォルトで CreatedAt には作成日時が
// UpdatedAt には更新日時が記録されます。
type Timestamp struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
