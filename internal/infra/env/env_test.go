package env

import (
	"os"
	"testing"
)

func TestGetInt(t *testing.T) {
	type args struct {
		key          string
		defaultValue int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "should get env value",
			args: args{key: "KEY_INT", defaultValue: 2},
			want: 1,
		},
		{
			name: "should get default value",
			args: args{key: "KEY_INT_MISSING", defaultValue: 2},
			want: 2,
		},
		{
			name: "should get default value when parse fail",
			args: args{key: "KEY_INT2", defaultValue: 2},
			want: 2,
		},
	}
	_ = os.Setenv("KEY_INT", "1")
	_ = os.Setenv("KEY_INT2", "3E0")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInt(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetString(t *testing.T) {
	type args struct {
		key          string
		defaultValue string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should get env value",
			args: args{key: "KEY_STRING", defaultValue: "B"},
			want: "A",
		},
		{
			name: "should get default value",
			args: args{key: "KEY_STRING_MISSING", defaultValue: "B"},
			want: "B",
		},
	}
	_ = os.Setenv("KEY_STRING", "A")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetString(tt.args.key, tt.args.defaultValue); got != tt.want {
				t.Errorf("GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}
