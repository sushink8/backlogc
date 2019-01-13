package main

import (
	"encoding/json"
	"fmt"
)

func get_priorityId(cfg *config) (backlog_responses, error) {
	req := cfg.new_backlog_request("/api/v2/priorities", "", []byte{})
	return get_backlog_response(cfg.get, req, cfg)
}

func get_projects(cfg *config) (backlog_responses, error) {
	req := cfg.new_backlog_request("/api/v2/projects", "", []byte{})
	return get_backlog_response(cfg.get, req, cfg)
}

func get_projects_detail_by_id(cfg *config, id int) (backlog_responses, error) {
	req := cfg.new_backlog_request(fmt.Sprintf("/api/v2/projects/%d", id), "", []byte{})
	return get_backlog_response(cfg.get, req, cfg)
}

func get_projects_detail_by_key(cfg *config, key string) (backlog_responses, error) {
	req := cfg.new_backlog_request(fmt.Sprintf("/api/v2/projects/%s", key), "", []byte{})
	return get_backlog_response(cfg.get, req, cfg)
}

func get_users(cfg *config) (backlog_responses, error) {
	req := cfg.new_backlog_request("/api/v2/users", "", []byte{})
	return get_backlog_response(cfg.get, req, cfg)
}

func get_issueTypes(cfg *config, id int) (backlog_responses, error) {
	req := cfg.new_backlog_request(fmt.Sprintf("/api/v2/projects/%d/issueTypes", id), "", []byte{})
	return get_backlog_response(cfg.get, req, cfg)
}

func get_categories(cfg *config, id int) (backlog_responses, error) {
	req := cfg.new_backlog_request(fmt.Sprintf("/api/v2/projects/%d/categories", id), "", []byte{})
	return get_backlog_response(cfg.get, req, cfg)
}

func get_issues(cfg *config, query_string string) (backlog_responses, error) {
	req := cfg.new_backlog_request("/api/v2/issues", query_string, []byte{})
	return get_backlog_response(cfg.get, req, cfg)
}

func get_issues_detail(cfg *config, query_string string, issuesid int) (backlog_responses, error) {
	req := cfg.new_backlog_request(fmt.Sprintf("/api/v2/issues/%d", issuesid), "", []byte{})
	return get_backlog_response(cfg.get, req, cfg)
}

func get_versions(cfg *config, id int) (backlog_responses, error) {
	req := cfg.new_backlog_request(fmt.Sprintf("/api/v2/projects/%d/versions", id), "", []byte{})
	return get_backlog_response(cfg.get, req, cfg)
}

type backlog_issues_add struct {
	ProjectId int    `json:"projectId"`
	Summary   string `json:"summary"`
	// ParentIssueId  int    `json:"parentIssueId"`
	Description    string `json:"description"`
	StartDate      string `json:"startDate"`
	DueDate        string `json:"dueDate"`
	EstimatedHours int    `json:"estimatedHours"`
	ActualHours    int    `json:"actualHours"`
	IssueTypeId    int    `json:"issueTypeId"`
	CategoryId     []int  `json:"categoryId"`
	VersionId      []int  `json:"versionId"`
	MilestoneId    []int  `json:"milestoneId"`
	PriorityId     int    `json:"priorityId"`
	AssigneeId     int    `json:"assigneeId"`
	NotifiedUserId []int  `json:"notifiedUserId"`
	AttachmentId   []int  `json:"attachmentId"`
}

type postparam map[string]interface{}

func (pd *postparam) add(key string, value interface{}) *postparam {
	map[string]interface{}(*pd)[key] = value
	return pd
}

func (pd *postparam) to_json() ([]byte, error) {
	return json.Marshal(*pd)
}

func add_issues(cfg *config, jsondata []byte) (backlog_responses, error) {
	req := cfg.new_backlog_request("/api/v2/issues", "", jsondata)
	return get_backlog_response(cfg.post, req, cfg)
}

func add_issueTypes(cfg *config, jsondata []byte, project_id int) (backlog_responses, error) {
	req := cfg.new_backlog_request(fmt.Sprintf("/api/v2/projects/%d/issueTypes", project_id), "", jsondata)
	return get_backlog_response(cfg.post, req, cfg)
}
