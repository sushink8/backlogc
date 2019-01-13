package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/k0kubun/pp"
)

type backlog_request struct {
	url      string
	postdata []byte
}

func (cfg *config) new_backlog_request(apipath string, query_string string, postdata []byte) *backlog_request {
	url := cfg.baseurl + apipath + "?apiKey=" + cfg.apikey
	if query_string != "" {
		url += url + "&" + query_string
	}
	return &backlog_request{url, postdata}
}

type backlog_responses []map[string]interface{}

func http_get(req *backlog_request) (*http.Response, error) {
	return http.Get(req.url)
}

func http_post(req *backlog_request) (*http.Response, error) {
	return http.Post(req.url, "application/json", bytes.NewBuffer(req.postdata))
}

func get_backlog_response(access func(*backlog_request) (*http.Response, error), req *backlog_request, cfg *config) (backlog_responses, error) {
	var (
		httpresponse    *http.Response
		err             error
		backlog_content []byte
		ret             backlog_responses
	)
	ret = backlog_responses{make(map[string]interface{})}
	if cfg.debug {
		pp.Println(req)
	}
	if httpresponse, err = access(req); err != nil {
		log.Println("http.Get error")
		return backlog_responses{}, err
	}
	defer httpresponse.Body.Close()
	if backlog_content, err = ioutil.ReadAll(httpresponse.Body); err != nil {
		log.Println("content read error")
		return backlog_responses{}, err
	}
	if httpresponse.StatusCode >= 400 {
		errmess := fmt.Sprintf("status code %d\n%s\n", httpresponse.StatusCode, backlog_content)
		//log.Printf(errmess)
		return backlog_responses{}, errors.New(errmess)
	}
	//pp.Println(string(backlog_content))
	if err := json.Unmarshal(backlog_content, &ret); err == nil {
		return ret, nil
	}

	if err := json.Unmarshal(backlog_content, &(ret[0])); err != nil {
		log.Println("json unmarshal error")
		return backlog_responses{}, err
	}
	return ret, nil
}
