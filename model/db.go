package model

import (
	"fmt"
	"gin-vue-blog/utils"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 全局变量，db和error在其他文件里也需要使用
var (
	db  *gorm.DB
	err error
)

func InitDb() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		utils.DbHost,
		utils.DbUser,
		utils.DbPassword,
		utils.DbName,
		utils.DbPort,
	)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		//禁用默认表名的复数形式
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		fmt.Println("数据库连接失败，请检查连接参数", err)

	}

	//

	//数据库自动迁移
	db.AutoMigrate(&User{}, &Category{}, &Article{})

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("数据库连接设置出错，请检查连接参数", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	//不能超过 gin 框架的连接超时时间
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	//sqlDB.Close()
}
