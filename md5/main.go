package main

import (
	"crypto/md5"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
)

func main() {
	s1, err := getFileMD5("/Users/jinersong/Downloads/app_config.json")
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	fmt.Printf("s1:%v", s1)
}

func getFileMD5(path string) (string, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return "", errors.Wrapf(err, "get oss file path:%v err", path)
	}
	defer configFile.Close()

	data, err := io.ReadAll(configFile)
	if err != nil {
		return "", err
	}
	fmt.Printf("file_data:%v", string(data))
	m := md5.New()
	m.Write(data)
	//if _, err := io.Copy(m, configFile); err != nil {
	//	return "", err
	//}
	return fmt.Sprintf("%x", m.Sum(nil)), nil
}
