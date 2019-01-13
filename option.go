package main

import (
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/k0kubun/pp"
)

type config struct {
	baseurl string
	apikey  string

	apikeyfile string
	action     string
	parameter  string
	encode     string
	debug      bool
	full       bool
	bulk       bool
	dryrun     bool

	get  func(*backlog_request) (*http.Response, error)
	post func(*backlog_request) (*http.Response, error)
}

func (cfg *config) has_parameter() bool {
	return len(cfg.parameter) > 0
}

func get_config(args []string) *config {
	cfg := &config{}
	cfg.get = http_get
	cfg.post = http_post
	flag.Usage = helpmessage
	flag.StringVar(&cfg.baseurl, "U", "", "base url for this tool")
	flag.StringVar(&cfg.apikey, "K", "", "APIKEY")
	flag.StringVar(&cfg.apikeyfile, "F", "", "file for APIKEY")
	flag.StringVar(&cfg.action, "A", "", "action")
	flag.StringVar(&cfg.parameter, "C", "", "add parameter(one line comma separated format) or csv filename if --buil")
	flag.StringVar(&cfg.encode, "encode", "utf8", "file encoding")
	flag.BoolVar(&cfg.full, "full", false, "input full parameter")
	flag.BoolVar(&cfg.debug, "D", false, "enable debug flag")
	flag.BoolVar(&cfg.bulk, "bulk", false, "enable bulk add")
	flag.BoolVar(&cfg.dryrun, "dryrun", false, "enable dryrun mode (only post api)")
	flag.Parse()

	if cfg.apikeyfile != "" {
		key, err := read_apikey(cfg.apikeyfile)
		if err == nil {
			cfg.apikey = key
		} else {
			log.Printf("can not open apikeyfile: %s\n", cfg.apikeyfile)
		}
	}

	if cfg.dryrun {
		cfg.post = http_post_dummy
	}
	return cfg
}

func helpmessage() {
	flag.PrintDefaults()
}
func read_apikey(path string) (string, error) {
	fd, err := os.Open(path)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(fd)
	return strings.TrimSpace(string(data)), err
}

type http_post_dummy_body struct {
}

func (bo *http_post_dummy_body) Close() error {
	return nil
}

func (bo *http_post_dummy_body) Read(b []byte) (int, error) {
	return 0, io.EOF
}
func http_post_dummy(b *backlog_request) (*http.Response, error) {
	pp.Println("http_post_dummy")
	pp.Println(b.url)
	pp.Println(string(b.postdata))
	res := &http.Response{Body: &http_post_dummy_body{},
		Status: "200",
	}
	return res, nil
}
