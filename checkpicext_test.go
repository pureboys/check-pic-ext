package checkpicext

import (
	"os"
	"testing"
)

var filePath = "testdata"

func TestGetImageFormat(t *testing.T) {

	image1, _ := os.ReadFile(filePath + "/image1.heic")
	image2, _ := os.ReadFile(filePath + "/image2.png")

	type args struct {
		image []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"image1",
			struct{ image []byte }{image: image1},
			"heic",
		},
		{
			"image2",
			struct{ image []byte }{image: image2},
			"png",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetImageFormat(tt.args.image); got != tt.want {
				t.Errorf("GetImageFormat() = %v, want %v", got, tt.want)
			}
		})
	}

}
