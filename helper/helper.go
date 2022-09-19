package helper

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	uuid "github.com/google/uuid"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,

		Data: data,
	}

	return jsonResponse
}

func FormatFile(fn string) string {

	ext := path.Ext(fn)
	u := uuid.New()
	// fmt.Println("cek ext", ext)
	newFileName := u.String() + ext

	return newFileName
}

func UploadFile(Folder string, file *multipart.FileHeader) (string, error) {
	// val, err := file
	// file := file
	newpath := filepath.Join(".", "public/", Folder)
	err := os.MkdirAll(newpath, os.ModePerm)
	if err != nil {
		// panic("Create Folder Failed!")
		return "", err
	}
	src, err := file.Open()
	if err != nil {
		return "", errors.New("Please Upload a valid image")
	}
	defer src.Close()
	path := fmt.Sprintf("public/%s/%s", Folder, FormatFile(file.Filename))
	nameFile := FormatFile(file.Filename)
	size := file.Size
	// fmt.Println("the size: ", size)

	if size > int64(512000) {
		// fmt.Println("Sorry, Please u/*  */pload an Image of 500KB or less")
		// errList = "Sorry, Please upload an Image of 500KB or less"/*  */
		return "", errors.New("Sorry, Please upload an Image of 500KB or less")

	}
	fileType := strings.Split(file.Header.Get("Content-Type"), "/")[0]

	if !strings.HasPrefix(fileType, "image") {
		// fmt.Println("Please Upload a valid image")
		return "", errors.New("Please Upload a valid image")
	}

	// if err != nil {
	// 	return "", err
	// }
	// defer src.Close()

	out, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}

	return nameFile, nil

}
func RemoveFile(Folder string, fileName string) error {
	path := fmt.Sprintf("public/%s/%s", Folder, fileName)
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func DateInputFormat(tanggal string) string {
	// fmt.Println(tanggal)
	// layout := "2006-01-02T15:04:05.000Z"
	tanggal_saat_ini, _ := time.Parse(time.RFC3339Nano, tanggal)
	hasil := tanggal_saat_ini.Format("02/01/2006")
	// fmt.Println(hasil)
	return hasil
}
