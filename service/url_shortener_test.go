package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateShortenURL(t *testing.T) {
	type args struct {
		url       string
		shortcode string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Success", args{"http://test1", "100001"}, false},
		{"Success", args{"http://test2", "100002"}, false},
		{"Success", args{"http://test3", "100003"}, false},
		{"FailedInvalidCode", args{"http://test", "1"}, true},
		{"FailedUsedCode", args{"http://test", "100001"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateShortenURL(tt.args.shortcode, tt.args.url)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestIsShortenURLExists(t *testing.T) {
	type args struct {
		shortcode string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"SuccessNotFound", args{"notfound1"}, false},
		{"SuccessNotFound", args{"notfound2"}, false},
		{"SuccessNotFound", args{"notfound3"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsShortenURLExists(tt.args.shortcode)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestRegisterURL(t *testing.T) {
	type args struct {
		url       string
		shortcode string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Success", args{"http://test1", "100001"}},
		{"Success", args{"http://test2", "100002"}},
		{"Success", args{"http://test3", "100003"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NotPanics(t, func() { RegisterURL(tt.args.url, tt.args.shortcode) })
		})
	}
}

func TestGetRegisteredURL(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"SuccessEmpty", args{"http://testnotfound1"}, ""},
		{"SuccessEmpty", args{"http://testnotfound2"}, ""},
		{"SuccessEmpty", args{"http://testnotfound3"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetRegisteredURL(tt.args.url)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsValidURL(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"SuccessHttp", args{"http://someurl"}, true},
		{"SuccessHttps", args{"https://someurl"}, true},
		{"FailedEmpty", args{""}, false},
		{"FailedProtocol", args{"ftp://"}, false},
		{"FailedProtocolOnly", args{"http://"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidURL(tt.args.str)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestIsValidShortcode(t *testing.T) {
	type args struct {
		shortcode string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"SuccessAllNumber", args{"123456"}, true},
		{"SuccessAllLowerCase", args{"abcdef"}, true},
		{"SuccessAllUpperCase", args{"ABCDEF"}, true},
		{"SuccessAllUnderscore", args{"______"}, true},
		{"SuccessAllCombination", args{"12cd_F"}, true},
		{"FailedLessThan6", args{"12345"}, false},
		{"FailedMoreThan6", args{"1234567"}, false},
		{"FailedUnwanted", args{"-!@#$&^"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidShortcode(tt.args.shortcode)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetNextShortcode(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Success", "1_____"},
		{"Success", "2_____"},
		{"Success", "3_____"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetNextShortcode()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGenerateShortcode(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Succcess", args{1}, "1_____"},
		{"Succcess", args{2}, "2_____"},
		{"Succcess", args{3}, "3_____"},
		{"Succcess", args{10}, "A_____"},
		{"Succcess", args{100}, "1c____"},
		{"Succcess", args{1000}, "G8____"},
		{"Succcess", args{100000}, "Q0u___"},
		{"Succcess", args{10000000}, "fxSK__"},
		{"Succcess", args{100000000}, "6laZE_"},
		{"Succcess", args{1000000000}, "15ftgG"},
		{"Succcess", args{56800235583}, "zzzzzz"},
		{"Succcess", args{56800235584}, "1000000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GenerateShortcode(tt.args.num)
			assert.Equal(t, tt.want, got)
		})
	}
}
