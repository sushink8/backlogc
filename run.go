package main

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/k0kubun/pp"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func run_list_categories(cfg *config) int {
	h := func() {
		log.Println("parameter: projectid int")
	}
	if !cfg.has_parameter() {
		h()
		return 1
	}
	projectid, err := strconv.Atoi(cfg.parameter)
	if err != nil {
		h()
		return 1
	}
	output_categories(cfg, projectid, os.Stdout)
	return 0
}
func run_list_issuesType(cfg *config) int {
	h := func() {
		log.Println("parameter: projectid int")
	}
	if !cfg.has_parameter() {
		h()
		return 1
	}
	projectid, err := strconv.Atoi(cfg.parameter)
	if err != nil {
		h()
		return 1
	}
	output_issueTypes(cfg, projectid, os.Stdout)
	return 0
}

func run_list_users(cfg *config) int {
	output_users(cfg, os.Stdout)
	return 0
}

func run_list_projects(cfg *config) int {
	output_projects(cfg, os.Stdout)
	return 0
}

func run_list_versions(cfg *config) int {
	h := func() {
		log.Println("parameter: projectid int")
	}
	if !cfg.has_parameter() {
		h()
		return 1
	}
	projectid, err := strconv.Atoi(cfg.parameter)
	if err != nil {
		h()
		return 1
	}
	output_versions(cfg, projectid, os.Stdout)
	return 0
}

//////////////////
func run_add_issueTypes(cfg *config) int {
	h := func() {
		log.Println("parameter: projectid int,issueTypesName string,colorCode string")
		log.Println("colorCode: #e30000 #990000 #934981 #814fbc #2779ca #007e9a #7ea800 #ff9200 #ff3265 #666665")
	}
	if !cfg.has_parameter() {
		h()
		return 1
	}
	params := strings.Split(cfg.parameter, ",")
	if len(params) != 3 {
		h()
		return 1
	}
	project_id, err := strconv.Atoi(params[0])
	if err != nil {
		h()
		return 1
	}
	issueTypesName := params[1]
	colorCode := params[2]
	pd := &postparam{}
	pd.add("name", issueTypesName)
	pd.add("color", colorCode)
	jsondata, err := json.Marshal(pd)
	if err != nil {
		log.Println("json encoding error")
		return 1
	}
	res, err := add_issueTypes(cfg, jsondata, project_id)
	pp.Println(res, err)
	return 0
}
func run_add_issues(cfg *config) int {
	if !cfg.has_parameter() {
		log.Println(show_add_issues_csv_param_short())
		return 1
	}
	if cfg.bulk {
		return bulk_add_issues(cfg)
	} else {
		return run_add_one_issues(cfg, cfg.parameter)
	}
}

func bulk_add_issues(cfg *config) int {
	path := cfg.parameter

	fd, err := os.Open(path)
	if err != nil {
		log.Printf("open file error: %s\n", path)
		return 1
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	var res []int
	if scanner.Scan() {
		_ = scanner.Text() // headerは捨てる
	}
	for scanner.Scan() {
		line := scanner.Text()
		switch cfg.encode {
		case "cp932":
			line, err = cp932_to_utf8(line)
			if err != nil {
				log.Printf("encoding error")
				return 1
			}
		}
		l_res := run_add_one_issues(cfg, line)
		res = append(res, l_res)
	}
	return max(res)
}

func cp932_to_utf8(str string) (string, error) {
	cp932reader := strings.NewReader(str)
	transreader := transform.NewReader(cp932reader, japanese.ShiftJIS.NewDecoder())
	ret, err := ioutil.ReadAll(transreader)
	return string(ret), err
}

func max(nums []int) int {
	var num int
	for _, v := range nums {
		if v > num {
			num = v
		}
	}
	return num
}

func run_add_one_issues(cfg *config, param string) int {
	var headers string
	if cfg.full {
		headers = "projectId,summary,parentIssueId,description,startDate,dueDate,estimatedHours,actualHours,issueTypeId,categoryId[],versionId[],milestoneId[],priorityId,assigneeId,notifiedUserId[],attachmentId[]"
	} else {
		headers = "projectId,summary,description,priorityId,issueTypeId"
	}
	jsondata, err := csv_param_to_bytes(cfg, headers, param)
	if err != nil {
		log.Println(err)
		return 1
	}
	res, err := add_issues(cfg, jsondata)
	if err != nil {
		log.Println(err)
		return 1
	}
	pp.Println(res)
	return 0
}
