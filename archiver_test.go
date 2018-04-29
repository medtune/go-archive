package archiver

import (
	"os"
	"reflect"
	"testing"
)

func Test_registerArchiver(t *testing.T) {
	type args struct {
		name     string
		archiver Archiver
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "first zip insert",
			args: args{
				name:     "zip",
				archiver: &zipper{},
			},
			wantErr: true,
		},
		{
			name: "different insert",
			args: args{
				name:     "diff",
				archiver: &zipper{},
			},
			wantErr: false,
		},
		{
			name: "second zip insert",
			args: args{
				name:     "zip",
				archiver: &zipper{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := registerArchiver(tt.args.name, tt.args.archiver); (err != nil) != tt.wantErr {
				t.Errorf("registerArchiver() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPick(t *testing.T) {
	type args struct {
		kind string
	}
	v, ok := supportPool["zip"]
	if !ok {
		t.Error("internal error : ")
	}
	tests := []struct {
		name    string
		args    args
		want    Archiver
		wantErr bool
	}{
		{
			name: "test zip type",
			args: args{
				kind: "zip",
			},
			want:    v,
			wantErr: false,
		},
		{
			name: "test another type",
			args: args{
				kind: "zippppx",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Pick(tt.args.kind)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pick() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pick() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_extractHeader(t *testing.T) {
	data := []byte("headerData")
	func() {
		file, err := os.Create("test_dir")
		if err != nil {
			t.Error("Cant create test dir")
		}
		defer file.Close()
		file.Write(data)
	}()
	defer os.Remove("test_dir")
	type args struct {
		file string
		size int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "test file",
			args: args{
				file: "test_dir",
				size: 10,
			},
			want:    data,
			wantErr: false,
		},
		{
			name: "test file",
			args: args{
				file: "test_di",
				size: 10,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := extractHeader(tt.args.file, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("extractHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
