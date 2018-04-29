package archiver

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func Test_zipper_Meta(t *testing.T) {
	tests := []struct {
		name string
		want archiverMetadata
	}{
		{
			name: "test zipper ext",
			want: archiverMetadata{
				Ext: "zip",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &zipper{}
			if got := z.Meta(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("zipper.Meta() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_zipper_Check(t *testing.T) {
	var ztest = "ztest.z"
	func() {
		if err := ioutil.WriteFile(ztest, zipHeader, 0644); err != nil {
			t.Errorf("couldnt create temp file : %v", ztest)
		}
	}()
	defer os.Remove(ztest)
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test zipfile name",
			args: args{
				file: "example.zip",
			},
			want: true,
		},
		{
			name: "test zipfile header code",
			args: args{
				file: ztest,
			},
			want: true,
		},
		{
			name: "false test",
			args: args{
				file: "archiver.go",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			z := &zipper{}
			if got := z.Check(tt.args.file); got != tt.want {
				t.Errorf("zipper.Check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unzipFile(t *testing.T) {
	ztest := "test.zip"
	defer os.RemoveAll("test")
	func() {
		z := &zipper{}
		if err := z.Compress("cmd", ztest); err != nil {
			t.Errorf("Couldnt create temp test.zip file")
		}
	}()
	defer os.Remove("test.zip")
	type args struct {
		file string
		dest string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "zip file ./cmd",
			args: args{
				file: ztest,
				dest: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := unzipFile(tt.args.file, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("unzipFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_zipFile(t *testing.T) {
	defer os.Remove("test.zip")
	type args struct {
		fname string
		dest  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "zip file ./cmd",
			args: args{
				fname: "cmd",
				dest:  "test.zip",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := zipFile(tt.args.fname, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("zipFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
