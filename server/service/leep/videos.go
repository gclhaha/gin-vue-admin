package leep

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/leep"
	leepReq "github.com/flipped-aurora/gin-vue-admin/server/model/leep/request"
	"github.com/google/uuid"
)

type VideosService struct{}

// CreateVideos 创建视频管理记录
// Author [yourname](https://github.com/yourname)
func (videosService *VideosService) CreateVideos(videos *leep.Videos) (err error) {
	videos.Id = strings.Replace(uuid.New().String(), "-", "", -1)
	err = global.MustGetGlobalDBByDBName("leep").Create(videos).Error
	return err
}

// DeleteVideos 删除视频管理记录
// Author [yourname](https://github.com/yourname)
func (videosService *VideosService) DeleteVideos(id string) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Delete(&leep.Videos{}, "id = ?", id).Error
	return err
}

// DeleteVideosByIds 批量删除视频管理记录
// Author [yourname](https://github.com/yourname)
func (videosService *VideosService) DeleteVideosByIds(ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Delete(&[]leep.Videos{}, "id in ?", ids).Error
	return err
}

// UpdateVideos 更新视频管理记录
// Author [yourname](https://github.com/yourname)
func (videosService *VideosService) UpdateVideos(videos leep.Videos) (err error) {
	err = global.MustGetGlobalDBByDBName("leep").Model(&leep.Videos{}).Where("id = ?", videos.Id).Updates(&videos).Error
	return err
}

// GetVideos 根据id获取视频管理记录
// Author [yourname](https://github.com/yourname)
func (videosService *VideosService) GetVideos(id string) (videos leep.Videos, err error) {
	err = global.MustGetGlobalDBByDBName("leep").Where("id = ?", id).First(&videos).Error
	return
}

// GetVideosInfoList 分页获取视频管理记录
// Author [yourname](https://github.com/yourname)
func (videosService *VideosService) GetVideosInfoList(info leepReq.VideosSearch) (list []leep.Videos, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("leep").Model(&leep.Videos{})
	var videoss []leep.Videos
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&videoss).Error
	return videoss, total, err
}
func (videosService *VideosService) GetVideosPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// UploadVideos 上传视频文件
func (videosService *VideosService) UploadVideos(file *multipart.FileHeader) (videoUrl string, err error) {
	// 获取文件扩展名
	fileExt := filepath.Ext(file.Filename)
	// 使用时间戳创建新的文件名，避免文件名冲突
	newFileName := fmt.Sprintf("%s-%d%s", file.Filename[0:len(file.Filename)-len(fileExt)], time.Now().Unix(), fileExt)

	filePath := fmt.Sprintf("temp/%s", newFileName)

	// 创建目录
	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("创建目录失败: %w", err)
	}

	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("打开文件失败: %w", err)
	}
	defer src.Close()

	// 创建目标文件
	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("创建文件失败: %w", err)
	}
	defer out.Close()

	// 复制文件
	_, err = io.Copy(out, src)
	if err != nil {
		return "", fmt.Errorf("复制文件失败: %w", err)
	}

	return "/" + filePath, nil // 返回相对路径，并添加前导斜杠
}

// LoadVideo 加载视频
// Author [yourname](https://github.com/yourname)
func (videosService *VideosService)LoadVideo() (err error) {
	// 请在这里实现自己的业务逻辑
	db := global.MustGetGlobalDBByDBName("leep").Model(&leep.Videos{})
    return db.Error
}

