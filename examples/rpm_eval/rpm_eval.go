/* -*- mode:go; coding:utf-8; -*-
 * author: Eugene Zamriy <eugene@zamriy.info>
 * created: 2015-12-18
 * description: "rpm --eval" implementation example.
 */

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ezamriy/gorpm"
)

func main() {
	expr := flag.String("eval", "", "print macro expansion of specified expression")

	flag.Parse()
	if flag.NFlag() < 1 {
		flag.Usage()
		os.Exit(1)
	}

	_, err := rpm.ReadConfigFiles(nil, nil)
	if err != nil {
		panic(fmt.Sprintf("cannot read RPM config files: %+v", err))
	}
	println(rpm.Expand(*expr))
}
