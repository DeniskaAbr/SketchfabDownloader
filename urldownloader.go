package main

import (
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"time"

	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)


// загрузчик файлов по сцене
func downloadservice(urls []string) error {

	var downloaderWG sync.WaitGroup // downloader goroutines wait group
	var g errgroup.Group

	fmt.Println("Downloader: Waiting for downloading workers to finish")

	for _, url := range urls {
		// задержка запуска скачивалок
		time.Sleep(DownloadTime)

		fmt.Println("Downloader: Starting downloading worker")
		downloaderWG.Add(1)


		// TODO: если происходи тошибка при скачивании метка все равно ставится что скачало
		//go downloader(&downloaderWG, url)


		g.Go(func() error {
			// Fetch the URL.
			err := downloader(&downloaderWG, url)
			return err
		})
	}

	downloaderWG.Wait()

	fmt.Println("Downloader: Downloading completed")
	return nil
}

func downloader(dwg *sync.WaitGroup, url string) error {
	defer dwg.Done()

	if dwg == nil || len(url) < 0 {
		return errors.New("no variables for download")
	}

	// по полученной строке формируем путь сохранения файла
	fmt.Printf(
		"Download worker started: %s \n", url,
	)

	// отрезаем из URL часть с доменом
	// формируем путь до файла в файловом хранилище
	ptf := FILESTORAGE_PATH + strings.Replace(url, "https://media.sketchfab.com", "", -1) // https://media.sketchfab.com

	// fmt.Println(ptf)

	fullpath := filepath.FromSlash(strings.Replace(ptf, string(filepath.Separator), "/", -1))

	// fmt.Println(fullpath)

	// запрашиваем файл с сервера
	err := DownloadFile(fullpath, url)

	// ловим ошибки или сохраняем файл
	// Check(err)
	if err!=nil {
		return err
	}

	fmt.Printf("Download worker finished: %s \n", url)
	return nil
}

func DownloadFile(fullpath string, url string) error {

	// формируем параметры запроса
	req, err := http.NewRequest(http.MethodGet, url, nil)
	// проверяем на ошибку
	// Check(err)
	if err != nil {
		return err
	}

	// устанавливаем хидер для запроса
	req.Header.Set("User-Agent", DEFAULT_USER_AGENT)
	req.Header.Set("Host", "")
	cookie := http.Cookie{Name: "sb_allows_age_restricted", Value: "true"}
	req.AddCookie(&cookie)

	// запрашиваем JSON с сайта
	resp, err := httpClient.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	/*
		if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
			// path/to/whatever does not exist
		}

		if _, err := os.Stat("/path/to/whatever"); !os.IsNotExist(err) {
			// path/to/whatever exists
		}
	*/
	//d1 := []byte("")
	//err = os.Create(
	//	pathtofile,
	//	d1,
	//	0644,
	//)
	//Check(err)

	dir, file := filepath.Split(fullpath)
	// fmt.Println("Dir:", dir)   //Dir: /some/path/to/remove/
	// fmt.Println("File:", file) //File: file.name

	// если путь не найден то создаем директорию
	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}

	//TODO: иногда происходит ошибка доступа к папке или каталогу из за блокировки
	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		// path/to/whatever does not exist
		// Create the file

		if _, err := os.Stat(fullpath + ".part"); os.IsNotExist(err) {
			out, err := os.Create(fullpath + ".part")
			// Check(err)

			if err != nil {
				out.Close()
				return err
			}

			// Write the body to file
			_, err = io.Copy(out, resp.Body)

			if err != nil {
				out.Close()
				return err
			}

			// Check(err)
			out.Close()
		}
	}

	if file == "file.osgjs.gz" {
		modelfiles := []string{"model_file.bin.gz", "model_file_wireframe.bin.gz"}
		for _, modelfile := range modelfiles {

			// TODO: replace

			modelfileurl := strings.Replace(url, "file.osgjs.gz", modelfile, -1)

			if _, err := os.Stat(dir + modelfile); os.IsNotExist(err) {

				// формируем параметры запроса
				req, err := http.NewRequest(http.MethodGet, modelfileurl, nil)
				// проверяем на ошибку
				// Check(err)

				if err != nil {return err}

				// устанавливаем хидер для запроса
				req.Header.Set("User-Agent", DEFAULT_USER_AGENT)
				req.Header.Set("Host", "")
				cookie := http.Cookie{Name: "sb_allows_age_restricted", Value: "true"}
				req.AddCookie(&cookie)

				// Get the data
				resp, err := httpClient.Do(req)

				if err != nil {
					return err
				}
				defer resp.Body.Close()

				out, err := os.Create(dir + modelfile + ".part")
				if err != nil {
					return err
				}
				// Check(err)

				// Write the body to file
				_, err = io.Copy(out, resp.Body)

				// Check(err)
				if err != nil {
					return err
					out.Close()
				}

				out.Close()

				if _, err := os.Stat(dir + modelfile + ".part"); !os.IsNotExist(err) {
					err = os.Rename(dir+modelfile+".part", dir+modelfile)
				}
			}
		}
	}

	if _, err := os.Stat(fullpath + ".part"); !os.IsNotExist(err) {
		err = os.Rename(fullpath+".part", fullpath)
	}

	// Check(err)
	return err
}
