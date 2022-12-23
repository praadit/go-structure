package services

// import (
// 	"errors"
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url"
// 	"os"
// 	"strings"

// 	"github.com/pdfcpu/pdfcpu/pkg/api"
// )

// type PdfCpuService struct{}

// func (service *PdfCpuService) MergeFiles(group string, identifier string, filename string, files ...string) (filepath string, err error) {
// 	if group == "" {
// 		err = errors.New("Group name should be filled!")
// 		return
// 	}
// 	if identifier == "" {
// 		err = errors.New("Identifier should be filled by unique value!")
// 		return
// 	}
// 	if filename == "" {
// 		err = errors.New("Filename should be filled")
// 		return
// 	}
// 	if len(files) < 1 {
// 		err = errors.New("No files found")
// 	}

// 	filename = strings.ToLower(filename)
// 	err = os.MkdirAll(fmt.Sprintf("./tmp/%s/%s", group, identifier), os.ModePerm)
// 	if err != nil {
// 		return
// 	}

// 	filepath = fmt.Sprintf("./tmp/%s/%s/%s", group, identifier, filename)
// 	err = api.MergeCreateFile(files, filepath, nil)
// 	defer service.deleteTmpFile(files...)

// 	return
// }

// func (service *PdfCpuService) DeleteUploadedFile(paths ...string) {
// 	for _, path := range paths {
// 		e := os.Remove(path)
// 		if e != nil {
// 			log.Println(e)
// 		}
// 	}
// }

// func (service *PdfCpuService) deleteTmpFile(paths ...string) {
// 	for _, path := range paths {
// 		e := os.Remove(path)
// 		if e != nil {
// 			log.Println(e)
// 		}
// 	}
// }

// func (service *PdfCpuService) OnlineImageToPdf(group string, identifier string, filename string, url string) (filepath string, err error) {
// 	if group == "" {
// 		err = errors.New("Group name should be filled!")
// 		return
// 	}
// 	if identifier == "" {
// 		err = errors.New("Identifier should be filled by unique value!")
// 		return
// 	}
// 	if filename == "" {
// 		err = errors.New("Filename should be filled")
// 		return
// 	}

// 	filename = strings.ToLower(filename)

// 	imgFilePath, err := service.DownloadFile(group, identifier, filename, url)
// 	if err != nil {
// 		return
// 	}
// 	defer service.deleteTmpFile(imgFilePath)

// 	dirPath := fmt.Sprintf("./tmp/%s/%s", group, identifier)
// 	filepath, err = ConvertImageToPdf(imgFilePath, dirPath, filename)
// 	if err != nil {
// 		return
// 	}

// 	return
// }

// func (service *PdfCpuService) DownloadFile(group string, identifier string, filename string, url string) (string, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != 200 {
// 		return "", errors.New("Status code not OK")
// 	}

// 	dirPath := fmt.Sprintf("./tmp/%s/%s", group, identifier)
// 	err = os.MkdirAll(dirPath, os.ModePerm)
// 	if err != nil {
// 		return "", err
// 	}

// 	filePath := fmt.Sprintf("./tmp/%s/%s/%s", group, identifier, filename)
// 	out, err := os.Create(filePath)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer out.Close()

// 	_, err = io.Copy(out, resp.Body)
// 	return filePath, err
// }

// func (service *PdfCpuService) GetFileExtensionFromUrl(rawUrl string) (string, error) {
// 	u, err := url.Parse(rawUrl)
// 	if err != nil {
// 		return "", err
// 	}
// 	pos := strings.LastIndex(u.Path, ".")
// 	if pos == -1 {
// 		return "", errors.New("couldn't find a period to indicate a file extension")
// 	}
// 	return u.Path[pos+1 : len(u.Path)], nil
// }
