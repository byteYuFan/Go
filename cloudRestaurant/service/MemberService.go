package service

import (
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/cloudRestaurant/dao"
	"github.com/cloudRestaurant/model"
	"github.com/cloudRestaurant/param"
	"github.com/cloudRestaurant/tool"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type MemberService struct {
}

func (ms *MemberService) GetUserInfo(userId string) *model.Member {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil
	}
	memberDao := dao.MemberDao{Engine: tool.DbEngine}
	return memberDao.QueryByID(int64(id))
}

// SendCode SendCode将6位数的验证码发给用户
// 成功返回 true 失败返回 false
func (ms *MemberService) SendCode(phone string) bool {
	//要发送，首先产生一个验证码
	code := fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
	//调用阿里云sdk
	config := tool.GetConfig().Sms
	client, err := dysmsapi.NewClientWithAccessKey(config.RegionId, config.AppKey, config.AppSecret)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	//接受返回结果
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.SignName
	request.TemplateCode = config.TemplateCode
	request.PhoneNumbers = phone
	par, err := json.Marshal(gin.H{
		"code": code,
	})
	request.TemplateParam = string(par)
	response, err := client.SendSms(request)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	if response.Code == "OK" {
		// 说明短信发送成功
		//将内容保存到数据库中去 以便之后对验证码进行校验
		smsCode := model.SmsCode{
			Phone:      phone,
			Code:       code,
			BizId:      string(response.BizId),
			CreateTime: time.Now().Unix(),
		}
		memberDao := dao.MemberDao{
			Engine: tool.DbEngine,
		}
		result := memberDao.InsertCode(smsCode)
		return result > 0
	}
	return false
}

// SmsLogin 将前端传来的loginParam的参数传递到后端，
// 进行校验
func (ms *MemberService) SmsLogin(loginParam param.SmsLoginParam) *model.Member {
	//创建数据库对象,用于操作数据库来查询验证码
	memberDao := dao.MemberDao{Engine: tool.DbEngine}
	smsCode := memberDao.ValidateSms(loginParam.Phone, loginParam.Code)
	if smsCode == nil {
		return nil
	}
	member := memberDao.QueryByPhone(loginParam.Phone)
	if member != nil {
		//查询到了
		return member
	}
	//创建新的会员并加入到member表中去
	user := model.Member{
		UserName:     loginParam.Phone,
		Mobile:       loginParam.Phone,
		RegisterTime: time.Now().Unix(),
	}
	res := memberDao.InsertMember(user)
	if res != 1 {
		log.Fatal("插入用户失败")
		return nil
	}
	return &user
}

// 登录功能
func (ms *MemberService) Login(name, password string) (*model.Member, bool) {
	//使用用户名查询用户信息，如果存在的话，直接返回
	memberDao := dao.MemberDao{
		Engine: tool.DbEngine,
	}
	member, isTrue := memberDao.Query(name, password)
	if member != nil && isTrue == true {
		return member, true
	}
	if member == nil && isTrue == true {
		//说明密码错误
		return member, false
	}

	if member != nil && isTrue == false {
		//说明数据库已经有这个用户名了，只是没有设置密码，调用修改密码函数
		if ok := memberDao.SetPassword(member.Id, password); ok == false {
			member.Password = password
			return member, ok
		} else {
			return member, ok
		}
	}
	//不存在，作为新的用户插入数据
	user := model.Member{
		UserName:     name,
		Password:     password,
		RegisterTime: time.Now().Unix(),
	}
	memberDao.InsertMember(user)
	return &user, true
}

// UploadAvatar
func (ms *MemberService) UploadAvatar(id int64, fileName string) string {
	memberDao := dao.MemberDao{
		Engine: tool.DbEngine,
	}
	r := memberDao.UpdateMember(id, fileName)
	if !r {
		return ""
	}
	return fileName
}
