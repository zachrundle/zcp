// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of zcp
package main

import (
	"github.com/zachrundle/zcp/cmd"

)


func main() {
//	f, err := tea.LogToFile("/var/log/zcp/debug.log", "debug")
	//if err != nil {
		  //fmt.Println("fatal:", err)
		  //os.Exit(1)
	//}
	//defer f.Close()
	cmd.Execute()
	
}
