package service

import (
	"fmt"
	"testing"
)

func TestUserServiceImpl_GetUserBasicInfoById(t *testing.T) {
	userBasicInfo := userServiceImp.GetUserBasicInfoById(1)
	fmt.Println(userBasicInfo)
}

func TestUserServiceImpl_GetUserBasicInfoByName(t *testing.T) {
	userBasicInfo := userServiceImp.GetUserBasicInfoByName("qcj")
	fmt.Println(userBasicInfo)
}
