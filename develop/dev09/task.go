package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

var protocol string

func main() {
	flag.Parse()
	protocol = "http://"
	url := strings.TrimLeft(flag.Args()[0], "http://")
	site, file, _ := strings.Cut(url, "/")

	err := os.Mkdir(site, 0777)
	if err != nil {
		log.Println(err)
	}

	err = os.Chdir(site)
	if err != nil {
		log.Println(err)
	}
	var urls = make([][]byte, 1)
	urls[0] = []byte(url)

	fmt.Println(site, file)
	downloadFile("index.html", protocol+url)
	// for _, url := range urls {
	// 	if strings.Contains(string(url), site) {
	// 		doc, err := downloadFile(string(url), string(url))
	// 		doc, _ = os.Open(doc.Name())
	// 		data, err := io.ReadAll(doc)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		exp, _ := regexp.Compile(`<a href=".*">`)
	// 		urls = exp.FindAll(data, -1)
	// 		for i, _ := range urls {
	// 			urls[i] = urls[i][9 : len(urls[i])-2]
	// 			fmt.Println(string(urls[i]))
	// 		}
	// 		defer doc.Close()
	// 	}
	// }
	// doc, err := downloadFile("index.html", url)
	// doc, _ = os.Open(doc.Name())

	// fmt.Println(doc.Name())

	// defer doc.Close()
	// if err != nil {
	// 	log.Println(err)
	// }
	// data, err := io.ReadAll(doc)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// exp, _ := regexp.Compile(`<a href=".*">`)
	// urls = exp.FindAll(data, -1)

	// for i, _ := range urls {
	// 	urls[i] = urls[i][9 : len(urls[i])-2]
	// 	fmt.Println(string(urls[i]))
	// }

}

func downloadFile(filepath string, url string) (*os.File, error) {
	log.Printf("Downdloading %s as %s\n", url, filepath)
	var recursive = true
	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		log.Println("create", err)
		out, err = os.Open(filepath)
		if err != nil {
			log.Println("open", err)
			return out, err
		}

		//return out, err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Println("get", err)
		return out, err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		log.Printf("bad status: %s", resp.Status)
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println("copy", err)
		return out, err
	}

	log.Printf("Finished %s\n", url)

	if !recursive {
		return out, nil
	}
	out.Close()
	out, _ = os.Open(out.Name())
	data, err := io.ReadAll(out)
	if err != nil {
		log.Fatal(err)
	}
	exp, _ := regexp.Compile(`<a href=".*">`)
	urls := exp.FindAll(data, -1)

	//log.Printf("Downdloading %v\n", urls)
	for i, _ := range urls {
		urls[i] = urls[i][9 : len(urls[i])-2]
		parts := strings.Split(string(urls[i]), "/")
		filepath = parts[len(parts)-1] + ".html"

		downloadFile(filepath, string(urls[i]))
	}
	return out, nil

}
