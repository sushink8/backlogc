package main

import (
	"io"
	"strconv"
	"strings"
)

func output_issues(cfg *config, query_string string, w io.Writer) error {
	res, err := get_issues(cfg, query_string)
	if err != nil {
		return err
	}
	var ret []string
	ret = append(ret, strings.Join([]string{"id", "issueKey", "issueType", "summary", "status"}, ","))
	for _, v := range res {
		id := strconv.Itoa(int(v["id"].(float64)))
		issueKey := v["issueKey"].(string)
		issueType := v["issueType"].(map[string]interface{})["name"].(string)
		summary := v["summary"].(string)
		status := v["status"].(map[string]interface{})["name"].(string)
		ret = append(ret, strings.Join([]string{id, issueKey, issueType, summary, status}, ","))
	}
	w.Write([]byte(strings.Join(ret, "\n")))
	return nil
}

func output_projects(cfg *config, w io.Writer) error {
	res, err := get_projects(cfg)
	if err != nil {
		return err
	}
	var ret []string
	ret = append(ret, strings.Join([]string{"id", "projectKey", "name"}, ","))
	for _, v := range res {
		id := strconv.Itoa(int(v["id"].(float64)))
		projectKey := v["projectKey"].(string)
		name := v["name"].(string)
		ret = append(ret, strings.Join([]string{id, projectKey, name}, ","))
	}
	w.Write([]byte(strings.Join(ret, "\n")))
	return nil
}

func output_users(cfg *config, w io.Writer) error {
	res, err := get_users(cfg)
	if err != nil {
		return err
	}
	var ret []string
	ret = append(ret, strings.Join([]string{"id", "userId", "name", "mailAddress"}, ","))
	for _, v := range res {
		id := strconv.Itoa(int(v["id"].(float64)))
		userId := v["userId"].(string)
		name := v["name"].(string)
		mailAddress := v["mailAddress"].(string)
		ret = append(ret, strings.Join([]string{id, userId, name, mailAddress}, ","))
	}
	w.Write([]byte(strings.Join(ret, "\n")))
	return nil
}

func output_versions(cfg *config, projectid int, w io.Writer) error {
	res, err := get_versions(cfg, projectid)
	if err != nil {
		return err
	}
	var ret []string
	ret = append(ret, strings.Join([]string{}, ","))
	for _, v := range res {
		id := strconv.Itoa(int(v["id"].(float64)))
		projectid := strconv.Itoa(int(v["projectId"].(float64)))
		name := v["name"].(string)
		ret = append(ret, strings.Join([]string{projectid, id, name}, ","))
	}
	w.Write([]byte(strings.Join(ret, "\n")))
	return nil
}

func output_categories(cfg *config, projectid int, w io.Writer) error {
	res, err := get_categories(cfg, projectid)
	if err != nil {
		return err
	}
	var ret []string
	ret = append(ret, strings.Join([]string{}, ","))
	for _, v := range res {
		id := strconv.Itoa(int(v["id"].(float64)))
		name := v["name"].(string)
		ret = append(ret, strings.Join([]string{id, name}, ","))
	}
	w.Write([]byte(strings.Join(ret, "\n")))
	return nil
	/*
	  {
	    "id":           21.000000,
	    "name":         "【運用】お知らせ管理",
	    "displayOrder": 2147483646.000000,
	  },
	*/
}
func output_issueTypes(cfg *config, projectid int, w io.Writer) error {
	res, err := get_issueTypes(cfg, projectid)
	if err != nil {
		return err
	}
	var ret []string
	ret = append(ret, strings.Join([]string{"id", "name"}, ","))
	for _, v := range res {
		id := strconv.Itoa(int(v["id"].(float64)))
		name := v["name"].(string)
		ret = append(ret, strings.Join([]string{id, name}, ","))
	}
	w.Write([]byte(strings.Join(ret, "\n")))
	return nil
	/*
	  {
	    "displayOrder": 0.000000,
	    "id":           81.000000,
	    "projectId":    17.000000,
	    "name":         "課題",
	    "color":        "#e30000",
	  },
	*/
}
