package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Путь к директории для загрузки файлов
var workingDir string

// Список загруженных файлов
var treeMap map[string]bool

// Протокол, по которому загружаются все файлы
var scheme string

// Максимальная глубина рекурсии. По умолчанию -1 для скачивания всего сайта.
// Скачиванием сайта назовем загрузку всех файлов, ссылки на которые содержатся в исходном файле(html-документе) и содержащие исходный домен.
var maxDepth int

func main() {
	flag.Parse()

	parsedURL, err := url.Parse(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	os.Mkdir(parsedURL.Host, 0777)
	os.Chdir(parsedURL.Host)
	workingDir, _ = os.Getwd()

	path := parsedURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]
	if fileName == "" {
		fileName = parsedURL.Path
	}
	treeMap = make(map[string]bool)
	scheme = parsedURL.Scheme
	maxDepth = -1
	downloadFile(parsedURL.Path, parsedURL.Host, parsedURL.Path, 0)
}

// Функция для рекурсивной загрузки файлов
func downloadFile(fileName string, host, path string, depth int) (*os.File, error) {
	log.Println(fileName, host, path, depth)
	var recursive = true
	if depth == maxDepth {
		recursive = false
	}
	log.Println(fileName, scheme, host, path)
	var download = scheme + "://" + host + path
	fileName = strings.Trim(fileName, "/")
	if fileName == "" {
		fileName = host
	}

	log.Printf("Downdloading %s as %s(%d)\n", download, fileName, len(fileName))

	// Get the data
	resp, err := http.Get(download)
	if err != nil {
		log.Println("get", err)
		return nil, err
	}
	defer resp.Body.Close()
	// Check server response
	if resp.StatusCode != http.StatusOK {
		log.Printf("bad status: %s", resp.Status)
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	log.Printf("Finished %s\n", download)

	// Create the file
	out := createFile(fileName)
	log.Printf("copying %s ...\n", out.Name())
	_, err = io.Copy(out, resp.Body)

	out.Close()

	out, err = os.Open(out.Name())

	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = os.Chdir(workingDir)
	if err != nil {
		log.Println(err)
	}

	document, err := goquery.NewDocumentFromReader(out)
	if err != nil {
		log.Println("Error loading HTTP response body. ", err)
		return nil, err
	}
	downloadStatic(host, document)
	if !recursive {
		return out, nil
	}
	log.Println("scanning links...")

	// Find all links and process them
	var links = make([]*url.URL, 0)
	document.Find("a").Each(func(index int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		if exists {
			parsedURL, err := url.Parse(href)
			if err == nil {
				if len(parsedURL.Host) == 0 {
					parsedURL.Host = host
				}
				if host == parsedURL.Host {
					links = append(links, parsedURL)
				}
			} else {
				log.Println(err)
			}
		}
	})
	log.Println(links)
	for _, plink := range links {
		newFileName := plink.Path
		if _, ok := treeMap[newFileName]; !ok {
			treeMap[newFileName] = true
			downloadFile(plink.Path, plink.Host, plink.Path, depth+1)
		}
	}
	return out, nil
}

// Создание дерева из директорий и открытие файла
func createFile(fileName string) (file *os.File) {
	path := strings.Split(fileName, "/")
	var err error
	var i int
	for i = 0; i < len(path)-1; i++ {
		os.Mkdir(path[i], 0777)
		os.Chdir(path[i])
	}
	if i == 0 {
		file, err = os.Create(fileName)
	} else {
		file, err = os.Create(path[i])
	}

	if err != nil {
		log.Printf("could not create file: %s", err)
		file, err = os.Create("index.html")
		if err != nil {
			log.Println("links duplicate")
			os.Exit(-1)
		}
	}
	log.Printf("created file: %s\n", file.Name())
	return
}

// Загрузка статических файлов
func downloadStatic(host string, document *goquery.Document) {
	document.Find("link").Each(func(i int, element *goquery.Selection) {
		href, exists := element.Attr("href")
		log.Println("static:    ", href, exists)
		if exists {
			parsedURL, err := url.Parse(href)
			if err == nil {
				if parsedURL.Host == "" {
					parsedURL.Host = host
				}
				downloadFile(parsedURL.Path, parsedURL.Host, parsedURL.Path, maxDepth)
			} else {
				log.Println(err)
			}
		}
	})
	document.Find("img").Each(func(i int, element *goquery.Selection) {
		href, exists := element.Attr("src")
		log.Println("static:    ", href, exists)
		if exists {
			parsedURL, err := url.Parse(href)
			if err == nil {
				if parsedURL.Host == "" {
					parsedURL.Host = host
				}
				downloadFile(parsedURL.Path, parsedURL.Host, parsedURL.Path, maxDepth)
			} else {
				log.Println(err)
			}
		}
	})

}
