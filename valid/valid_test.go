package valid

import (
	"testing"
)

func TestIsPhoneNumber(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPhoneNumber(tt.args.value); got != tt.want {
				t.Errorf("IsPhoneNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFixedPhone(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsFixedPhone(tt.args.value); got != tt.want {
				t.Errorf("IsFixedPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsBankAccount(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBankAccount(tt.args.value); got != tt.want {
				t.Errorf("IsBankAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsIdcard(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsIdcard(tt.args.value); got != tt.want {
				t.Errorf("IsIdcard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUrl(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Normal", args{"http://github.com/chrisho/sd-helper"}, true},
		{"NoSchema", args{"github.com/chrisho/sd-helper"}, false},
		{"NoSchema", args{"/github.com/chrisho/sd-helper"}, true},
		{"D2Slash", args{"//github.com/chrisho/sd-helper"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUrl(tt.args.value); got != tt.want {
				t.Errorf("IsUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsChinese(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChinese(tt.args.value); got != tt.want {
				t.Errorf("IsChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsChineseName(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChineseName(tt.args.value); got != tt.want {
				t.Errorf("IsChineseName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasChinese(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasChinese(tt.args.value); got != tt.want {
				t.Errorf("HasChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsChineseAndLetterAndNumber(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChineseAndLetterAndNumber(tt.args.value); got != tt.want {
				t.Errorf("IsChineseAndLetterAndNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPostcode(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPostcode(tt.args.value); got != tt.want {
				t.Errorf("IsPostcode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumeric(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{"123456"}, true},
		{"minus", args{"-123456"}, true},
		{"float", args{"123456.1324"}, true},
		{"zero", args{"0.00"}, true},
		{"a-z", args{"abcdfef"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumeric(tt.args.value); got != tt.want {
				t.Errorf("IsNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmail(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{"chrisho@github.com"}, true},
		{"noname", args{"@github.com"}, false},
		{"nosuffix", args{"chrisho@github"}, false},
		{"nodomain", args{"chrisho@.com"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmail(tt.args.value); got != tt.want {
				t.Errorf("IsEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsImei(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsImei(tt.args.value); got != tt.want {
				t.Errorf("IsImei() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CustomTagWithIsPhoneNumber(t *testing.T) {

	type args struct {
		Phone string `validate:"MobilePhone"`
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"normal", args{"18688381114"}, true},
		{"string", args{"abc"}, false},
		{"less11", args{"1868888"}, false},
		{"sortNums", args{"12345678"}, false},
	}

	validate := GetValidate()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Struct(tt.args)
			if  (err != nil && tt.want != false) || (err == nil && tt.want != true) {
				t.Errorf("mobilePhone = %v want %v", false, tt.want)
			}
		})
	}
}
