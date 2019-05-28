// Copyright (c) 2019. Tom Harris <mrtom.h.84@gmail.com>. All rights reserved.
// Use of this source code is governed by the MIT LICENSE
// license that can be found in the LICENSE.txt file.

package main

import (
	"bufio"
	"fmt"
	"github.com/tomeh/caseconv"
	"io"
	"os"
	"strings"
)

func main() {
	input := strings.Join(os.Args[1:], " ")

	if len(input) == 0 {
		input = fromStdIn()
	}

	fmt.Println(caseconv.ToSnake(input))
}

func fromStdIn() string {
	reader := bufio.NewReader(os.Stdin)
	var input []rune

	for {
		r, _, err := reader.ReadRune()
		if err != nil && err == io.EOF {
			break
		}
		input = append(input, r)
	}

	return string(input)
}
