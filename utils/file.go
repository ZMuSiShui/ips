package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// 确定文件是否存在
func FileExists(file string) bool {
	info, err := os.Stat(file)
	return err == nil && !info.IsDir()
}

// 确定目录是否存在
func DirExists(file string) bool {
	if file == "" {
		return false
	}
	info, err := os.Stat(file)
	return err == nil && info.IsDir()
}

// 创建嵌套文件
func CreatNestedFile(path string) (*os.File, error) {
	basePath := filepath.Dir(path)
	if !FileExists(basePath) {
		err := os.MkdirAll(basePath, 0700)
		if err != nil {
			log.Errorf("can't create foler，%s", err)
			return nil, err
		}
	}
	return os.Create(path)
}

// 将结构写入 JSON 文件
func WriteToJson(src string, object interface{}) error {
	data, err := json.MarshalIndent(object, "", "  ")
	if err != nil {
		log.Errorf("无法将 Object 转换为 []byte: %v", err.Error())
		return err
	}
	err = ioutil.WriteFile(src, data, 0777)
	if err != nil {
		log.Errorf("无法写入 JSON 文件: %v", err.Error())
		return err
	}
	return err
}

// 将结构写入 Yaml 文件
func WriteToYaml(src string, object interface{}) error {
	data, err := yaml.Marshal(object)
	if err != nil {
		log.Errorf("无法将 Object 转换为 []byte: %v", err.Error())
		return err
	}
	if err := ioutil.WriteFile(src, data, 0777); err != nil {
		log.Errorf("无法写入 Yaml 文件: %v", err.Error())
		return err
	}
	return nil
}

func ParsePath(path string) string {
	path = strings.TrimRight(path, "/")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return path
}
