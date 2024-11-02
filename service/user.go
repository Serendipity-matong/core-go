package service

import (
	"MiMengCore/config"
	"MiMengCore/global"
	"MiMengCore/model"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"log"
	"regexp"
	"time"
)

const (
	UserIDFormatError  = "用户ID格式错误"      // 用户ID格式错误
	UserIDIsTakenError = "user_id_taken" // 用户ID已被占用
)

// CreateUser 创建用户
func CreateUser(user *model.User) (err error) {
	if err = global.DB.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// CheckUserLogin 检查用户名和密码是否匹配
func CheckUserLogin(userID, password string) (*model.User, bool, string) {
	// 查找用户
	var user model.User
	result := global.DB.Where("user_id = ? AND password = ?", userID, password).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, false, "用户不存在或密码错误"
		}
		return nil, false, "数据库查询出错"
	}
	if result.RowsAffected == 0 {
		return nil, false, "用户不存在或密码错误"
	}
	return &user, true, ""
}

// GenerateToken 生成JWT令牌
func GenerateToken(user *model.User) (string, error) {
	// 设置令牌的过期时间
	expireTime := time.Now().Add(time.Duration(config.ExpireTime) * time.Second)

	// 创建一个令牌实例
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.UserID,
		"username": user.UserName,
		"exp":      expireTime.Unix(),
	})

	// 使用密钥对令牌进行签名
	tokenString, err := token.SignedString([]byte(config.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateUser 校验用户信息是否有效
func ValidateUser(userID string, userName string, password string, qq int) (bool, string) {
	// 校验用户ID
	valid, msg := CheckUserID(userID)
	if !valid {
		return false, msg
	}

	// 校验用户名
	if !CheckUserName(userName) {
		return false, "用户名格式不正确"
	}

	// 校验密码
	if !CheckPassword(password) {
		return false, "密码格式不正确"
	}

	// 校验QQ
	if !CheckQQ(qq) {
		return false, "QQ格式不正确"
	}

	return true, ""
}

// validateUserID 检查用户ID是否只包含数字、字母和下划线
func validateUserID(userID string) bool {
	//ID在6~20位之间
	fmt.Println(userID)
	if len(userID) < 6 || len(userID) > 20 {
		return false
	}
	//ID只能包含数字字母下划线
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return re.MatchString(userID)
}

// validatePassword 检查密码是否只包含数字、字母和下划线
func validatePassword(password string) bool {
	//密码经加密后为64位
	if len(password) != 64 {
		return false
	}
	//密码只能包含数字字母
	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return re.MatchString(password)
}

// validateUserName 检查用户名是否合法
func validateUserName(name string) bool {
	//用户名在1~12位之间
	if len(name) < 1 || len(name) > 12 {
		return false
	}
	return true
}

// validateQQ 检查QQ号是否合法
func validateQQ(qq int) bool {
	if qq < 10000 || qq > 10000000000 {
		return false
	}
	return true
}

// isExistUserID 检查用户ID是否已存在
func isExistUserID(userID string) bool {
	// 定义变量来接收查询结果
	var count int64

	// 执行查询，计算匹配的用户ID数量
	result := global.DB.Model(&model.User{}).Where("user_id = ?", userID).Count(&count)
	if result.Error != nil {
		log.Printf("查询数据库时发生错误: %v", result.Error)
		return false
	}

	// 如果计数大于0，则表示用户ID存在
	return count > 0
}

// CheckUserID 校验用户账号合法
func CheckUserID(userID string) (bool, string) {
	// 校验账号合法性
	if !validateUserID(userID) {
		return false, UserIDFormatError
	}
	// 确保用户ID唯一
	if isExistUserID(userID) {
		return false, UserIDIsTakenError
	}
	return true, ""
}

// CheckPassword 校验密码合法
func CheckPassword(password string) bool {
	return validatePassword(password)
}

// CheckUserName 校验用户名合法
func CheckUserName(name string) bool {
	return validateUserName(name)
}

// CheckQQ 校验QQ号合法
func CheckQQ(qq int) bool {
	return validateQQ(qq)
}
