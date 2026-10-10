package models

import (
	"time"

	"gorm.io/gorm"
)

// ------------------------------------------------------------
// HOW TO READ THE TAGS (the text inside `backticks`)
//
// json:"..." -> read by Go's JSON encoder.
//               Controls the field name in API responses / requests.
//               json:"-" = NEVER send/receive this field as JSON.
//
// gorm:"..." -> read by GORM (the ORM).
//               Controls how this field becomes a DATABASE COLUMN.
//               Several gorm options are separated by ";"
//               e.g. gorm:"uniqueIndex;not null"
//
// json 标签 = API 返回时的字段名；gorm 标签 = 数据库字段的规则。
// ------------------------------------------------------------

// User represents a customer or admin account in the shop.
// GORM maps this struct to a table called "users" (plural, lowercase).
type User struct {
	// gorm:"primaryKey" -> this column is the table's PRIMARY KEY.
	// Unique ID for each row; with uint, PostgreSQL auto-increments it (1, 2, 3...).
	// 主键，自动递增。
	ID uint `json:"id" gorm:"primaryKey"`

	// gorm:"uniqueIndex" -> creates a UNIQUE index: no two users can have the same email.
	//                       (Also makes searching by email fast.)
	// gorm:"not null"    -> this column can never be empty (NULL) in the database.
	// 唯一索引 + 不能为空。
	Email string `json:"email" gorm:"uniqueIndex;not null"`

	// json:"-"        -> NEVER included in JSON, so the password is never sent to the frontend.
	// gorm:"not null" -> must have a value. (Store a HASHED password here, never plain text.)
	// 密码不会出现在 API 返回里。
	Password string `json:"-" gorm:"not null"`

	// json:"first_name" -> in JSON this is "first_name" instead of "FirstName".
	// gorm:"not null"   -> required column.
	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`

	// No gorm tag -> normal column with default rules (can be empty).
	Phone string `json:"phone"`

	// gorm:"default:true" -> if not set when creating a user, the DATABASE stores true.
	// 默认值 true（账号默认启用）。
	IsActive bool `json:"is_active" gorm:"default:true"`

	// gorm:"default:customer" -> new users get the role "customer" unless set otherwise.
	// UserRole is a custom type (defined later), e.g. "customer" or "admin".
	// 默认角色 customer。
	Role UserRole `json:"role" gorm:"default:customer"`

	// GORM fills these AUTOMATICALLY (special field names, no tag needed):
	// CreatedAt -> set when the row is first created.
	// UpdatedAt -> updated every time the row is saved.
	// 创建时间/更新时间，GORM 自动填写。
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// gorm.DeletedAt -> enables SOFT DELETE.
	//   db.Delete(&user) does NOT remove the row; it sets this time instead,
	//   and GORM hides the row from normal queries.
	// gorm:"index"   -> creates a normal index on this column, so
	//   "WHERE deleted_at IS NULL" (which GORM adds to every query) stays fast.
	// json:"-"       -> not shown in API responses.
	// 软删除：数据还在，只是标记为已删除。
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// ------------------------------------------------------------
	// Relationships (links to OTHER tables)
	// These are NOT columns in the users table. GORM uses them to
	// load related rows (e.g. with db.Preload("Orders")).
	// json:"-" -> kept out of JSON to avoid huge nested responses.
	// 关联关系：不是 users 表的字段，而是连接到其他表。
	// ------------------------------------------------------------

	// One user has MANY refresh tokens (one-to-many).
	// The refresh_tokens table will have a "user_id" column pointing back here.
	RefreshTokens []RefreshToken `json:"-"`

	// One user has MANY orders (one-to-many).
	Orders []Order `json:"-"`

	// One user has ONE cart (one-to-one).
	Cart Cart `json:"-"`
}