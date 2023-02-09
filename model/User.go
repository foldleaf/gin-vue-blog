package model

import (
	"encoding/base64"
	"gin-vue-blog/utils/errmsg"
	"log"

	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"type:varchar(20);not null " json:"username"`
	Password string `gorm:"type:varchar(20);not null " json:"password"`
	Role     int    `gorm:"type:int; " json:"role"`
}

// 查询用户是否存在
func CheckUser(name string) (code int) {
	var users User
	//查询使用该用户名的 id
	db.Select("id").Where("username =?", name).First(&users)
	//存在使用该用户名的 id
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCESS //200
}

// 新增用户   返回状态码
func CreateUser(data *User) int {
	// scrypt 加密
	data.Password = ScryptPassword(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR //500
	}
	return errmsg.SUCCESS //200

}

// 查询用户列表 分页
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	err = db.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}

// 密码加密
func ScryptPassword(password string) string {
	const KeyLen = 10 //长度
	salt := make([]byte, 8)

	salt = []byte{12, 22, 111, 46, 82, 3, 7, 21}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	finalPw := base64.StdEncoding.EncodeToString(HashPw)
	return finalPw
}

// 删除用户

func DeleteUser(id int) int {
	err = db.Where("id=?", id).Delete(&User{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 编辑用户(密码除外)
func EditUser(id int, data *User) int {
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&User{}).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
