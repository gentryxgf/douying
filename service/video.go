package service

import (
	"douyin/common/global"
	"douyin/common/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/fs"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"time"
)

type VideoService struct{}

var (
	WhiteVideoList = []string{
		"flv",
		"avi",
		"mov",
		"mp4",
		"wmv",
	}
)

func (VideoService) UploadVideo(video *multipart.FileHeader, title string, userID int64, c *gin.Context) (err error) {

	// 判断保存视频的文件路径是否存在，如果不存在就递归创建
	baseVideoPath := global.Config.UploadConf.VideoLocalPath
	if _, err := os.ReadDir(baseVideoPath); err != nil {
		// 不存在文件路径递归创建
		err = os.MkdirAll(baseVideoPath, os.ModePerm)
		if err != nil {
			global.Log.Error("VideoService.UploadVideo 创建文件路径失败", zap.Error(err))
			return err
		}
	}
	baseCoverPath := global.Config.UploadConf.CoverLocalPath
	if _, err := os.ReadDir(baseCoverPath); err != nil {
		// 不存在文件路径递归创建
		err = os.MkdirAll(baseCoverPath, fs.ModePerm)
		if err != nil {
			global.Log.Error("VideoService.UploadVideo 创建文件路径失败", zap.Error(err))
			return err
		}
	}

	// 判断当前文件是否为视频文件
	// 获取视频名称
	name := video.Filename
	// 获取文件后缀
	nameList := strings.Split(name, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if ok := utils.IsInList(WhiteVideoList, suffix); !ok {
		global.Log.Error("VideoService.UploadVideo 上传的视频文件不合规", zap.Error(err))
		return errors.New("视频文件不合规")
	}

	// 视频本地存储，防止视频重名，所以要给上传的视频重新命名
	// 命名规则为：userID_timestamp
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	userIDStr := strconv.FormatInt(userID, 10)
	videoPath := utils.JoinStringToPath(baseVideoPath, userIDStr, timeStamp, suffix)
	coverPath := utils.JoinStringToPath(baseCoverPath, userIDStr, timeStamp, "jpg")

	// 保存视频到本地
	err = c.SaveUploadedFile(video, videoPath)
	if err != nil {
		global.Log.Error("VideoService.UploadVideo USE gin.Context.SaveUploadedFile ERROR", zap.Error(err))
		return err
	}
	// 保存封面到本地
	err = utils.ReadVideoAsImage(videoPath, coverPath, 1)
	if err != nil {
		global.Log.Error("VideoService.UploadVideo USE utils.ReadVideoAsImage ERROR", zap.Error(err))
		return err
	}

	// 保存视频数据到数据库
	err = VideoDao.CreateVideo(userID,
		utils.JoinString(global.Config.UploadConf.UrlPath, videoPath),
		utils.JoinString(global.Config.UploadConf.UrlPath, coverPath),
		title,
	)
	if err != nil {
		global.Log.Error("VideoService.UploadVideo USE VideoDao.CreateVideo ERROR", zap.Error(err))
		return err
	}

	return nil
}
