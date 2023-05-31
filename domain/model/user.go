// Code generated by sql2go. DO NOT EDIT.
package model


type User struct {
	Uid       int64 `gorm:"column:uid;primary_key" json:"uid"`              // 用户ID 系统生成
	Aid       string `gorm:"column:aid;NOT NULL" json:"aid"`                 // 账户ID 用户设置
	Password  string `gorm:"column:password;NOT NULL" json:"password"`            // 密码
	Did       string `gorm:"column:did;NOT NULL" json:"did"`                 // 注册设备唯一标识
	Status    int32   `gorm:"column:status;default:0;NOT NULL" json:"status"`    // 用户状态
	Nickname  string `gorm:"column:nickname;NOT NULL" json:"nickname"`            // 昵称
	Firstname string `gorm:"column:firstname;NOT NULL" json:"firstname"`           // firstname
	Lastname  string `gorm:"column:lastname;NOT NULL" json:"lastname"`            // lastname
	Gender    int8   `gorm:"column:gender;default:0;NOT NULL" json:"gender"`    // 性别
	Birth     int64  `gorm:"column:birth;default:0;NOT NULL" json:"birth"`     // 生日
	Email     string `gorm:"column:email;NOT NULL" json:"email"`               // Email
	Mobile    string `gorm:"column:mobile;NOT NULL" json:"mobile""`              // 手机号
	Platform  uint   `gorm:"column:platform;default:0;NOT NULL" json:"platform"`  // 注册平台/登录平台
	ServerId  int    `gorm:"column:server_id;default:0;NOT NULL" json:"server_id"` // 分配的ws服务器
	CityId    int    `gorm:"column:city_id;default:0;NOT NULL" json:"city_id"`   // 城市ID
	AvatarKey string `gorm:"column:avatar_key;NOT NULL" json:"avatar_key"`          // 小图 72*72

	Created   int64  `gorm:"column:created;autoCreateTime:milli;NOT NULL" json:"created"`
	Updated   int64  `gorm:"column:updated;autoCreateTime:milli;NOT NULL" json:"updated"`
	Deleted   int64  `gorm:"column:deleted;default:0;NOT NULL" json:"deleted"`
}

