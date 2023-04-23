package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/LingshengCode/awesomeProject/dao"
	"log"
	"sync"
)

type UserServiceImpl struct {
}

var (
	userServiceImp  *UserServiceImpl
	userServiceOnce sync.Once
)

func GetUserServiceInstance() *UserServiceImpl {
	userServiceOnce.Do(func() {
		userServiceImp = &UserServiceImpl{}
	})
	return userServiceImp
}

func (usi *UserServiceImpl) GetUserBasicInfoById(id int64) dao.UserBasicInfo {
	user, err := dao.GetUserBasicInfoById(id)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
		return user
	}
	log.Println("Query User Success")
	return user
}

func (usi *UserServiceImpl) GetUserBasicInfoByName(name string) dao.UserBasicInfo {
	user, err := dao.GetUserBasicInfoByName(name)
	if err != nil {
		log.Println("Err:", err.Error())
		log.Println("User Not Found")
		return user
	}
	log.Println("Query User Success")
	return user
}

func (usi *UserServiceImpl) InsertUser(user *dao.UserBasicInfo) bool {
	flag := dao.InsertUser(user)
	if flag == false {
		log.Println("Insert Fail!")
		return false
	}
	return true
}

// 给密码加密

func EnCoder(password string) string {
	h := hmac.New(sha256.New, []byte(password))
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}
