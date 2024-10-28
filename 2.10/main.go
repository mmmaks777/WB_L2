package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

var visitedPages = make(map[string]bool)
var visitedResources = make(map[string]bool)

func main() {
	startURL := flag.String("url", "", "Download site URL")
	outputDir := flag.String("output", ".", "Directory for saving files")
	maxDepth := flag.Int("depth", 1, "Recursion depth")
	flag.Parse()

	if *startURL == "" {
		fmt.Println("Specify the URL using the -url flag")
		return
	}

	os.MkdirAll(*outputDir, os.ModePerm)

	err := downloadPage(*startURL, *outputDir, *startURL, *maxDepth)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func downloadPage(pageURL, dir, baseURL string, depth int) error {
	if depth <= 0 {
		return nil
	}

	if visitedPages[pageURL] {
		return nil
	}
	visitedPages[pageURL] = true

	fmt.Println("Downloading page:", pageURL)

	resp, err := http.Get(pageURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download page: %s, status: %d", pageURL, resp.StatusCode)
	}

	u, err := url.Parse(pageURL)
	if err != nil {
		return err
	}

	var filename string
	var saveDir string

	if path.Ext(u.Path) == "" || strings.HasSuffix(u.Path, "/") {
		filename = "index.html"
		saveDir = filepath.Join(dir, u.Host, u.Path)
	} else {
		filename = path.Base(u.Path)
		saveDir = filepath.Join(dir, u.Host, path.Dir(u.Path))
	}

	os.MkdirAll(saveDir, os.ModePerm)

	filePath := filepath.Join(saveDir, filename)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	bodyString := string(bodyBytes)

	doc, err := html.Parse(strings.NewReader(bodyString))
	if err != nil {
		_, err = file.Write(bodyBytes)
		return err
	}

	var pageLinks []string
	var resourceLinks []string

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "href" || a.Key == "src" {
					if strings.HasPrefix(a.Val, "#") {
						continue
					}

					resourceURL, err := url.Parse(a.Val)
					if err != nil {
						continue
					}
					absoluteURL := resp.Request.URL.ResolveReference(resourceURL)

					localPath := getLocalPath(absoluteURL, dir)
					a.Val = filepath.ToSlash(localPath[len(dir)+1:])

					if isHTMLPage(absoluteURL) {
						pageLinks = append(pageLinks, absoluteURL.String())
					} else {
						resourceLinks = append(resourceLinks, absoluteURL.String())
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	err = html.Render(file, doc)
	if err != nil {
		return err
	}

	for _, resURL := range resourceLinks {
		if !visitedResources[resURL] {
			err := downloadResource(resURL, dir)
			if err != nil {
				fmt.Println("Error while downloading resource:", resURL, err)
			}
		}
	}

	for _, link := range pageLinks {
		if strings.HasPrefix(link, baseURL) {
			err := downloadPage(link, dir, baseURL, depth-1)
			if err != nil {
				fmt.Println("Error while downloading page:", link, err)
			}
		}
	}

	return nil
}

func downloadResource(resURL, dir string) error {
	if visitedResources[resURL] {
		return nil
	}
	visitedResources[resURL] = true

	fmt.Println("Downloading resourse:", resURL)

	resp, err := http.Get(resURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download resource: %s, status: %d", resURL, resp.StatusCode)
	}

	u, err := url.Parse(resURL)
	if err != nil {
		return err
	}

	savePath := getLocalPath(u, dir)

	os.MkdirAll(filepath.Dir(savePath), os.ModePerm)

	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func isHTMLPage(u *url.URL) bool {
	ext := strings.ToLower(path.Ext(u.Path))
	if ext == ".html" || ext == ".htm" || ext == "" {
		return true
	}
	return false
}

func getLocalPath(u *url.URL, dir string) string {
	var filename string
	var saveDir string

	if path.Ext(u.Path) == "" || strings.HasSuffix(u.Path, "/") {
		filename = "index.html"
		saveDir = filepath.Join(dir, u.Host, u.Path)
	} else {
		filename = path.Base(u.Path)
		saveDir = filepath.Join(dir, u.Host, path.Dir(u.Path))
	}

	return filepath.Join(saveDir, filename)
}
