package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"github.com/gookit/color"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	URL0      string
	Coroutine int
	Time      int
	proxyFile string
)

// 代理检测结果存储，线程安全
type SafeSlice struct {
	sync.Mutex
	items []string
}

func (s *SafeSlice) Append(item string) {
	s.Lock()
	defer s.Unlock()
	s.items = append(s.items, item)
}

func NewSafeSlice() *SafeSlice {
	return &SafeSlice{}
}

func main() {
	flag.IntVar(&Coroutine, "c", 200, "设置验证代理的协程数量默认200")
	flag.IntVar(&Time, "t", 3, "设置验证代理超时时间默认3秒")
	flag.StringVar(&URL0, "v", "https://www.baidu.com/", "设置http proxy 验证用的url")
	flag.StringVar(&proxyFile, "f", "proxy.txt", "设置http proxy 文件默认为proxy.txt")
	flag.Parse()

	clientPool := &sync.Pool{
		New: func() interface{} {
			return &http.Client{
				Timeout: time.Duration(Time) * time.Second,
			}
		},
	}

	var liveres []string
	fileContents, err := os.ReadFile(proxyFile)
	if err != nil {
		panic(err)
	}
	proxies := strings.Split(strings.ReplaceAll(string(fileContents), "\r\n", "\n"), "\n")
	proxies = removeDuplicates(proxies)

	limiter := make(chan struct{}, Coroutine)
	var wg sync.WaitGroup

	for _, proxy := range proxies {
		proxy := proxy // 避免闭包捕获同一变量

		wg.Add(1)
		limiter <- struct{}{}
		go func() {
			defer wg.Done()
			if checkProxy(clientPool, proxy, URL0) {
				liveres = append(liveres, proxy)
			}
			<-limiter
		}()
	}

	wg.Wait()
	color.RGBStyleFromString("237,64,35").Printf("\n[+]一共获取存活代理:%d条", len(liveres))
	fmt.Println(liveres)

}

// checkProxy checks the proxy and returns true if the proxy is live.
func checkProxy(pool *sync.Pool, proxyIp string, targetURL string) bool {
	proxyURL := fmt.Sprintf("http://%s", proxyIp)
	proxy, err := url.Parse(proxyURL)
	if err != nil {
		return false
	}

	client := pool.Get().(*http.Client)
	client.Transport = &http.Transport{
		Proxy:           http.ProxyURL(proxy),
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	defer pool.Put(client)

	response, err := client.Get(targetURL)
	if err != nil {
		return false
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		fmt.Println("true proxy: " + proxy.String())
		return true
	}
	return false
}

// removeDuplicates removes duplicate entries from a slice of strings.
func removeDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
