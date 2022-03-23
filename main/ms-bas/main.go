package main

import (
	"fmt"
	"github.com/cube-group/golib/bas"
)

func main() {
	_, query, _ := bas.NewBasService().GetSecuritySignQuery(
		"12485d7307bf347ee5834f98f5465aa4",
		"3d279dde16162fcd4e173707d129d036",
		"rbac",
		3600,
	)
	fmt.Println(query)
}
