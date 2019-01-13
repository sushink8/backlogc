package main

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"strings"

	"github.com/k0kubun/pp"
)

func proc_each_project(res backlog_responses, cfg *config, fn func(*config, int) (backlog_responses, error)) error {
	ids, err := get_projectids(res)
	if err != nil {
		return err
	}
	for _, i := range ids {
		r, _ := fn(cfg, i)
		pp.Println(r)
	}
	return nil
}

func get_projectids(res backlog_responses) ([]int, error) {
	var ids []int
	for _, v := range res {
		ids = append(ids, int((v["id"].(float64))))
	}
	return ids, nil
}

func get_projectid_by_projectkey(cfg *config, projectkey string) (int, error) {
	res, err := get_projects_detail_by_key(cfg, projectkey)
	if err != nil {
		return -1, err
	}
	return int(res[0]["id"].(float64)), nil
}

func show_csv_header_add_issues() string {
	return "projectKey,summary,description,priorityId,issueTypeId"
}

func csv_param_to_bytes(cfg *config, header string, param string) ([]byte, error) {
	header_reader := csv.NewReader(strings.NewReader(header))
	header_reader.Comma = ','
	header_reader.LazyQuotes = true
	param_reader := csv.NewReader(strings.NewReader(param))
	param_reader.Comma = ','
	param_reader.LazyQuotes = true
	pd := &postparam{}
	for {
		header_record, err := header_reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return []byte{}, err
		}
		param_record, err := param_reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return []byte{}, err
		}
		for i, _ := range header_record {
			if len(header_record[i]) == 0 || len(param_record[i]) == 0 {
				continue
			}
			pd.add(header_record[i], strings.Replace(param_record[i], `\n`, "\n", -1))
		}
	}
	if cfg.debug {
		pp.Println(pd)
	}
	return json.Marshal(pd)
}

func show_csv_param(headers []string) string {
	// header の(以降を削除
	for i, v := range headers {
		stringindex := strings.Index(v, "(")
		if stringindex >= 0 {
			headers[i] = v[:stringindex]
		}
	}
	return strings.Join(headers, ",")
}

func show_add_issues_csv_param_short() string {
	headers := []string{
		"projectId",
		"summary",
		"description",
		"priorityId",
		"issueTypeId",
	}
	return show_csv_param(headers)
}
