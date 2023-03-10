package controller

import (
	"encoding/json"
	"fmt"
	"github.com/cloudRestaurant/model"
	"github.com/cloudRestaurant/param"
	"github.com/cloudRestaurant/service"
	"github.com/cloudRestaurant/tool"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"time"
)

type MemBerController struct {
}

func (mc *MemBerController) Router(engine *gin.Engine) {
	engine.GET("/api/sendcode", mc.sendSmsCode)
	engine.POST("/api/login_sms", mc.smsLogin)
	engine.GET("/api/captcha", mc.captcha)

	// postman测试
	engine.POST("/api/vertifycha", mc.vertifyCaptcha)
	//登录路由
	engine.POST("/api/login_pwd", mc.nameLogin)
	// 用户头像上传功能
	engine.POST("/api/upload/avator", mc.uploadAvator)

	//用户信息查询
	//engine.GET("/api/userinfo", mc.UserInfo)
}

// uploadAvator 该函数提供了上传用户图像的功能
func (mc MemBerController) uploadAvator(ctx *gin.Context) {

	// 1、解析上传的参数 file user_id
	userId := ctx.PostForm("user_id")
	if userId == "" {
		tool.Failed(ctx, "用户未登录")
		return
	}
	fmt.Println(userId)
	file, err := ctx.FormFile("avator")
	if err != nil {
		tool.Failed(ctx, "获取文件失败"+err.Error())
		return
	}
	//判断user_id是否已经登录，只有登录的才有权限
	sess := tool.GetSession(ctx, "user_"+userId)
	if sess == nil {
		tool.Failed(ctx, "参数不合法")
		return
	}
	var member model.Member
	err = json.Unmarshal(sess.([]byte), &member)
	if err != nil {
		tool.Failed(ctx, "unmarshall err")
		return
	}
	//2、file保存到本地
	fileName := "./upLoad/" + file.Filename + strconv.FormatInt(time.Now().Unix(), 10)
	err = ctx.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.Failed(ctx, "save err")
		return
	}
	//上传文件到fastDFS
	fileId := tool.UploadFile(fileName)
	memberService := service.MemberService{}
	if fileId != "" {
		os.Remove(fileId)
		// 3、更新数据库,返回结果
		//去点
		id, err := strconv.ParseInt(userId, 10, 64)
		if err != nil {
			tool.Failed(ctx, err.Error())
			return
		}
		path := memberService.UploadAvatar(id, fileId)
		if path != "" {
			tool.Success(ctx, tool.FileServerAddr()+path)
			return
		}
	}
	fileName = fileName[1:]
	id, _ := strconv.ParseInt(userId, 10, 64)
	r := memberService.UploadAvatar(id, fileName)
	if r != "" {
		tool.Success(ctx, "http://localhost:8080"+r)
		return
	}

	tool.Failed(ctx, "upload err")
}

// sendSmsCode 接受eg：http://localhost:8090/api/sendcode?phone=13572015398传递过来的参数
// 调用用户服务层的方法向用户发送代码
func (mc *MemBerController) sendSmsCode(ctx *gin.Context) {
	//发送验证码
	phone, exist := ctx.GetQuery("phone")
	if !exist {
		tool.Failed(ctx, "参数解析失败")
		return
	}
	//调用短信验证功能，发送验证码
	//实例化一个service层的对象，用来发消息
	ms := service.MemberService{}
	isSend := ms.SendCode(phone)
	if isSend {

		tool.Success(ctx, "发送成功")
		return
	}
	tool.Failed(ctx, "发送失败")
}

// smsLogin 验证用户输入的验证码是否和数据库中存储的是否相等
func (mc *MemBerController) smsLogin(ctx *gin.Context) {
	var smsLoginParm param.SmsLoginParam
	// 解析传入的参数
	err := tool.Decode(ctx.Request.Body, &smsLoginParm)
	if err != nil {
		tool.Failed(ctx, "参数解析失败")
	}
	//创建一个服务层的对象，用来调用登录功能，分层模型，控制层只处理逻辑，服务层去处理相关的请求
	us := service.MemberService{}
	member := us.SmsLogin(smsLoginParm)
	if member != nil {
		//登录成功,设置session
		sess, _ := json.Marshal(member)
		err := tool.SetSession(ctx, "user_"+strconv.FormatInt(member.Id, 10), sess)
		if err != nil {
			tool.Failed(ctx, "登录失败")
			return
		}
		ctx.SetCookie("cookie_user", strconv.FormatInt(member.Id, 10), 10*60, "/", "localhost", true, true)
		tool.Success(ctx, member)
		return
	}
	tool.Failed(ctx, "登录失败")
}

// captcha 生成验证码，并返回给客户端
func (mc *MemBerController) captcha(ctx *gin.Context) {
	//生成验证码并返回客户端
	tool.GenerateCaptchaHandler(ctx)
}

// vertifyCaptcha 验证验证码是否正确
func (mc *MemBerController) vertifyCaptcha(ctx *gin.Context) {
	//解析参数
	var captcha tool.CaptchaResult
	fmt.Println(ctx.Query("phone"))
	err := tool.Decode(ctx.Request.Body, &captcha)
	if err != nil {
		tool.Failed(ctx, "参数解析失败")
		return
	}
	r := tool.VertifyCaptcha(captcha.Id, captcha.VertifyValue)
	if r {
		fmt.Println("验证通过")
	} else {
		fmt.Println("验证失败")
	}
}

// nameLogin 用户名+密码登录系统
func (mc *MemBerController) nameLogin(ctx *gin.Context) {
	//解析参数
	var loginParam param.LoginParam
	if err := tool.Decode(ctx.Request.Body, &loginParam); err != nil {
		tool.Failed(ctx, "参数解析失败")
	}
	//验证验证码
	if r := tool.VertifyCaptcha(loginParam.Id, loginParam.Value); !r {
		tool.Failed(ctx, "验证码不正确")
		return
	}
	//	登录
	ms := service.MemberService{}
	member, isTrue := ms.Login(loginParam.Name, loginParam.Password)
	if member != nil && isTrue == true {
		//登录成功,设置
		sess, _ := json.Marshal(member)
		err := tool.SetSession(ctx, "user_"+strconv.FormatInt(member.Id, 10), sess)
		if err != nil {
			tool.Failed(ctx, "登录失败")
			return
		}
		tool.Success(ctx, &member)
		return
	} else {
		tool.Failed(ctx, "密码错误")
		return
	}

}

// UserInfo 获取用户的信息
func (mc *MemBerController) UserInfo(ctx *gin.Context) {
	//判断用户是否登录
	cookie, err := tool.CookieAuth(ctx)
	if err != nil {
		ctx.Abort()
		tool.Failed(ctx, "请先登录")
		return
	}
	memberService := service.MemberService{}
	member := memberService.GetUserInfo(cookie.Value)
	if member != nil {
		tool.Success(ctx, gin.H{
			"id":            member.Id,
			"user_name":     member.UserName,
			"mobile":        member.Mobile,
			"register_time": member.RegisterTime,
			"avatar":        member.Avatar,
			"balance":       member.Balance,
			"city":          member.City,
		})
		return
	}
	tool.Failed(ctx, "获取信息失败")
}
