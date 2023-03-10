package dao

import (
	"github.com/cloudRestaurant/model"
	"github.com/cloudRestaurant/tool"
)

type MemberDao struct {
	Engine *tool.GormEngine
}

func (md *MemberDao) QueryByID(userId int64) *model.Member {
	var member model.Member
	res := md.Engine.DB.Where("id=?", userId).First(&member)
	if res.RowsAffected == 0 {
		return nil
	}
	return &member
}

// InsertCode InsertCode 将用户的手机号和验证码 存储到 数据库中去 方便之后和用户传过来的验证码进行对比，校验
// 返回影响的行数
func (md *MemberDao) InsertCode(sms model.SmsCode) int64 {
	r := md.Engine.DB.Create(&sms)
	return r.RowsAffected
}

// ValidateSms ValidateSms 按照string和code在数据库查找对应的信息 查到返回 信息结构体指针
// 失败返回 nil
func (md *MemberDao) ValidateSms(phone string, code string) *model.SmsCode {
	var sms model.SmsCode
	res := md.Engine.DB.Where("phone=? AND code=?", phone, code).First(&sms)
	if res.RowsAffected != 1 {
		return nil
	}
	return &sms
}

// QueryByPhone 按照用户的手机号查询数据，成功返回新的学生对象，失败返回空 nil
func (md *MemberDao) QueryByPhone(phone string) *model.Member {
	var member model.Member
	res := md.Engine.DB.Where("mobile=?", phone).First(&member)
	if res.RowsAffected == 0 {
		return nil
	}
	return &member
}

// InsertMember 插入成员功能，将成员member 插入到数据库中去，返回受到影响的行数
// 0 代表0 行受到影响 1 代表 1 行受到影响，即插入成功
// 分层模型的搭建，各个函数各司其职，插入时不必考虑是否冲突的问题
func (md *MemberDao) InsertMember(member model.Member) int64 {
	r := md.Engine.DB.Create(&member)
	return r.RowsAffected
}

// Query 根据用户名和密码到数据库中查询第一条匹配的数据，如果没有查到，返回一个空值
// 如果查询到了，则返回这个用户的结构体指针
func (md *MemberDao) Query(name, password string) (*model.Member, bool) {
	var member model.Member
	tx := md.Engine.DB.Where("user_name=?", name).First(&member)
	if tx.RowsAffected == 0 {
		return nil, false
	}
	if member.Password == "" {
		return &member, false
	}
	if member.Password != password {
		return nil, true
	}
	return &member, true
}

func (md *MemberDao) SetPassword(Id int64, password string) bool {
	var member model.Member
	r := md.Engine.DB.Model(&member).Where("id=?", Id).Update("password", password)
	return r.RowsAffected == 1
}

func (md *MemberDao) UpdateMember(Id int64, fileName string) bool {
	var member model.Member
	r := md.Engine.DB.Model(&member).Where("id=?", Id).Update("avatar", fileName)
	return r.RowsAffected == 1
}
