package migrations

import (
	userPD "github.com/gomsa/socialite/proto/user"
	db "github.com/gomsa/socialite/providers/database"
)

func init() {
	user()
}

// user 用户数据迁移
func user() {
	user := &userPD.User{}
	if !db.DB.HasTable(&user) {
		db.DB.Exec(`
			CREATE TABLE users (
			id varchar(36) NOT NULL,
			origin varchar(64) DEFAULT NULL,
			openid varchar(64) DEFAULT NULL,
			session varchar(64) DEFAULT NULL,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY origin_openid (origin,openid)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}
