// Copyright 2013 Jongmin Kim. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package gabric

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func Local(cmd string) {
	c := strings.Split(cmd, " ")
	out, err := exec.Command(c[0], c[1:]...).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", out)
}
