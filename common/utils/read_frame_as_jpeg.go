package utils

import (
	"bytes"
	"douyin/common/global"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"go.uber.org/zap"
)

/*
ReadVideoAsImage
ffmpeg常用输入选项：

	-i filename：指定输入文件名
	-f fmt：强制设定文件格式，需使用能力集列表中的名称（缺省是根据扩展名选择的）
	-ss hh:mm:ss[.xxx]：设定输入文件的起始时间点，启动后将跳转到此时间点然后开始读取数据

ffmpeg常用输入选项：

	-f fmt：强制设定文件格式，需使用能力集列表中的名称（缺省是根据扩展名选择的）
	-acodec codec：指定声音的编码器，需使用能力集列表中的名称（编码器设定为”copy“表示不进行编解码）
	-vcodec codec：指定视频的编码器，需使用能力集列表中的名称（编解码器设定为”copy“表示不进行编解码）
	-r fps：设定视频编码器的帧率，整数，单位fps
	-pix_fmt format：设置视频编码器使用的图像格式（如RGB还是YUV）

frameNum参数解析：

	select 过滤器可以使用不同的条件来筛选帧或音频样本，以便对它们进行处理或输出。其中，gte 是 select 过滤器的一个条件选项之一。
	fmt.Sprintf("gte(n,%d)", frameNum) 使用了 fmt.Sprintf 函数来构建一个字符串，其中 %d 是一个占位符，表示帧的索引值。
	frameNum 是一个变量，表示要筛选的帧的索引

ffmpeg "github.com/u2takey/ffmpeg-go" 函数参数解析:

	Input参数：指定读取的视频路径
	Output参数：
		vframes：表示生成图片的帧数。"vframes": 1表示指定输出的图片帧数为 1，即仅生成一张图片
		其他参数见上述
	Filter参数：
		select：表示使用select过滤器
		gte(n,%d)：表示从 大于等于 %d 的帧数开始截取数据
*/
func ReadVideoAsImage(videoPath, coverPath string, frameNum int) (err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf).Run()
	if err != nil {
		global.Log.Error("生成视频缩略图失败", zap.Error(err))
		return err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		global.Log.Error("生成视频缩略图失败", zap.Error(err))
		return err
	}
	err = imaging.Save(img, coverPath)
	if err != nil {
		global.Log.Error("生成视频缩略图失败", zap.Error(err))
		return err
	}

	return nil
}
