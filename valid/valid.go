package valid

import (
	"github.com/go-playground/validator"
	"regexp"
)

type nValidator struct {
	*validator.Validate
}

var vd *nValidator

func init() {
	vd = &nValidator{
		validator.New(),
	}
	vd.setTagMobilePhoneWithPhoneNumber()
}

func GetValidate() *nValidator {
	return vd
}

// here is bind struct validate with custom tag
func (s *nValidator) CustomTagWithPhoneNumber(tag string) {
	s.Validate.RegisterValidation(tag, func(fl validator.FieldLevel) bool {
		return IsPhoneNumber(fl.Field().String())
	})
}
// set tag "MobilePhone" for CustomTagWithPhoneNumber
func (s *nValidator) setTagMobilePhoneWithPhoneNumber() {
	s.CustomTagWithPhoneNumber("MobilePhone")
}

// IsPhoneNumber
func IsPhoneNumber(value string) bool {
	//reg := `^1([3-9][0-9]|14[57]|5[^4])\d{8}$`
	reg := `^[1-9]\d{10}$`

	return regexp.MustCompile(reg).MatchString(value)
}

// IsFixedPhone
func IsFixedPhone(value string) bool {
	reg := `^(0[0-9]{2,3}-)?([2-9][0-9]{6,7})+(-[0-9]{1,4})?$`

	return regexp.MustCompile(reg).MatchString(value)
}

// bank account
func IsBankAccount(value string) bool {
	reg := `^([1-9]{1})(\d{15}|\d{16}|\d{17}|\d{18}|\d{19}|\d{20})$`

	return regexp.MustCompile(reg).MatchString(value)
}

// idcard
func IsIdcard(value string) bool {
	if len(value) == 15 {
		reg := `^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{2}[0-9Xx]$`
		return regexp.MustCompile(reg).MatchString(value)
	} else {
		reg := `^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`
		return regexp.MustCompile(reg).MatchString(value)
	}
	return false
}

// suport uri, url, urn_rfc2141
// support example: http://github.com/chrisho/sd-helper,
// 					//github.com/chrisho/sd-helper
//					/github.com/chrisho/sd-helper
// not support example: github.com/chrisho/sd-helper
func IsUrl(value string) bool {
	if err := vd.Var(value, "required,uri|url|urn_rfc2141"); err != nil {
		return false
	}
	return true
}

// chinese
func IsChinese(value string) bool {
	reg := `^[\p{Han}]+$`

	return regexp.MustCompile(reg).MatchString(value)
}

// 中文名字
func IsChineseName(value string) bool {
	reg := `^[\p{Han}（）()]+$`

	return regexp.MustCompile(reg).MatchString(value)
}

// has chinese
func HasChinese(value string) bool {
	reg, _ := regexp.Compile(`[\p{Han}]`)

	return reg.MatchString(value)
}

// chinese and letter and number
func IsChineseAndLetterAndNumber(value string) bool {
	reg := `^[a-zA-Z0-9\p{Han}]+$`

	return regexp.MustCompile(reg).MatchString(value)
}

// postcode
func IsPostcode(value string) bool {
	reg := `^[1-9][0-9]{5}$`

	return regexp.MustCompile(reg).MatchString(value)
}

// IsNumeric
func IsNumeric(value string) bool {
	if err := vd.Var(value, "required,numeric"); err != nil {
		return false
	}
	return true
}

// IsEmail
func IsEmail(value string) bool {
	if err := vd.Var(value, "required,email"); err != nil {
		return false
	}
	return true
}

// IsImei
func IsImei(value string) bool {
	reg := `^[0-9A-Za-z]+$`

	return regexp.MustCompile(reg).MatchString(value)
}