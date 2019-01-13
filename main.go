package main

import (
	"fmt"
	"os"

	"github.com/k0kubun/pp"
	//"strconv"
	//"strings"
)

func main() {
	orignal_args := os.Args
	cfg := get_config(os.Args)
	os.Exit(run(cfg))
	os.Args = orignal_args
}
func h() {
	fmt.Println("\n-------------------------------------")
}

func run(cfg *config) int {
	if cfg.debug {
		pp.Println(cfg)
	}
	switch cfg.action {
	case "list_project":
		return run_list_projects(cfg)
	case "add_issues":
		return run_add_issues(cfg)
	case "list_issueTypes":
		return run_list_issuesType(cfg)
	case "list_categories":
		return run_list_categories(cfg)
	case "list_users":
		return run_list_users(cfg)
	case "add_issueTypes":
		return run_add_issueTypes(cfg)
	case "list_versions":
		return run_list_versions(cfg)
	default:
		helpmessage()
	}
	return 0
}
