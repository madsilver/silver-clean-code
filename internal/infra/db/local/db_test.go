package local

import (
	"reflect"
	"testing"
)

func TestLocalDB_FindByID(t *testing.T) {
	tableName := "test"
	targetValue := "test-value"
	type args struct {
		tableName string
		id        any
	}
	tests := []struct {
		name    string
		args    args
		want    any
		wantErr bool
	}{
		{
			name:    "should return a value",
			args:    args{tableName: tableName, id: 123},
			want:    targetValue,
			wantErr: false,
		},
		{
			name:    "should return an error",
			args:    args{tableName: tableName, id: 111},
			want:    nil,
			wantErr: true,
		},
	}
	l := NewLocalDB()
	l.storage = map[string]map[any]any{
		tableName: {
			123: targetValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := l.FindByID(tt.args.tableName, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalDB_Create(t *testing.T) {
	tableName := "test"
	targetValue := "test-value"
	type args struct {
		tableName string
		id        any
		value     any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "should add a value to table",
			args:    args{tableName: tableName, id: 1, value: targetValue},
			wantErr: false,
		},
	}
	l := NewLocalDB()
	l.storage = map[string]map[any]any{
		tableName: {},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := l.Add(tt.args.tableName, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalDB_CreateTable(t *testing.T) {
	tableName := "test"
	type args struct {
		tableName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "should create a table",
			args:    args{tableName: tableName},
			wantErr: false,
		},
		{
			name:    "should return an error when table exists",
			args:    args{tableName: tableName},
			wantErr: true,
		},
	}
	l := NewLocalDB()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := l.CreateTable(tt.args.tableName); (err != nil) != tt.wantErr {
				t.Errorf("CreateTable() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
