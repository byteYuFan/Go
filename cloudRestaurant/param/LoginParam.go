package param

// 用户名，密码和验证码登录
type LoginParam struct {
	Name     string `json:"name"`  //用户名
	Password string `json:"pwd"`   //密码
	Id       string `json:"id"`    // captchaId 验证码ID
	Value    string `json:"value"` //验证码
}
