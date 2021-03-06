// Copyright © 2019 Mephis Pheies <mephistommm@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package main

import (
	"fmt"
	"github.com/MephistoMMM/gAlgorithmAndDesignPattern/design_patterns/Builder/builder"
	"github.com/MephistoMMM/gAlgorithmAndDesignPattern/design_patterns/Builder/director"
	"os"
)

func usage() {
	fmt.Println("Usage: go run main.go plain  // write document in plain")
	fmt.Println("Usage: go run main.go html  // write document in html")
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	if os.Args[1] == "plain" {
		bdr := builder.NewTextBuilder()
		drctr := director.NewDirector(bdr)
		drctr.Construct()
		fmt.Println(bdr.Result())
	} else {
		bdr := builder.NewHTMLBuilder("./test.html")
		drctr := director.NewDirector(bdr)
		drctr.Construct()
		fmt.Println(bdr.Result())
	}

}
