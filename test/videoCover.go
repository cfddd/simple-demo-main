package test

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"strings"
)

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	fmt.Println(1)
	buf := bytes.NewBuffer(nil)
	fmt.Println(1)
	err = ffmpeg.
		Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	fmt.Println(1)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	fmt.Println(1)
	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	fmt.Println(1)
	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}
	fmt.Println(1)
	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return
}

func main() {
	s, err := GetSnapshot("./public/bear.mp4", "./public/test", 1)
	if err != nil {
		return
	}
	fmt.Println(s)
}
