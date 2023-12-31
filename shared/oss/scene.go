package oss

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"github.com/i-Things/things/shared/errors"
	"path"
	"strings"
)

type (
	SceneInfo struct {
		Business string
		Scene    string
		FilePath string
		FileName string
	}
)

// 产品管理
const (
	BusinessProductManage = "productManage" //产品管理
	SceneProductImg       = "productImg"    //产品图片
)

func GetSceneInfo(filePath string) (*SceneInfo, error) {
	paths := strings.Split(filePath, "/")
	if len(paths) < 3 {
		return nil, errors.Parameter.WithMsg("路径不对")
	}
	scene := &SceneInfo{
		Business: paths[0],
		Scene:    paths[1],
		FilePath: strings.Join(paths[2:], "/"),
		FileName: paths[len(paths)-1],
	}
	return scene, nil
}

func GetFilePath(scene *SceneInfo, rename bool) (string, error) {
	if rename == true {
		ext := path.Ext(scene.FilePath)
		if ext == "" {
			return "", errors.Parameter.WithMsg("未能获取文件后缀名")
		}
		uuid, er := uuid.GenerateUUID()
		if er != nil {
			err := errors.System.AddDetail(er)
			return "", err
		}
		scene.FilePath = uuid + ext
	} else {
		spcChar := []string{`,`, `?`, `*`, `|`, `{`, `}`, `\`, `$`, `、`, `·`, "`", `'`, `"`}
		if strings.ContainsAny(scene.FilePath, strings.Join(spcChar, "")) {
			return "", errors.Parameter.WithMsg("包含特殊字符")
		}
	}
	filePath := fmt.Sprintf("%s/%s/%s", scene.Business, scene.Scene, scene.FilePath)
	return filePath, nil
}

func CheckWithCopy(ctx context.Context, handle Handle, srcPath string, business, scene string) (string, error) {
	//如果第一次就提交了模型文件
	si, err := GetSceneInfo(srcPath)
	if err != nil {
		return "", err
	}
	if !(si.Business == business && si.Scene == scene) {
		return "", errors.Parameter.WithMsg(scene + "的路径不对")
	}
	nwePath, err := GetFilePath(si, false)
	if err != nil {
		return "", err
	}
	path, err := handle.CopyFromTempBucket(srcPath, nwePath)
	if err != nil {
		return "", errors.System.AddDetail(err)
	}
	return path, nil
}
