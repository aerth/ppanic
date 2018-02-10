/*
 * Copyright (c) 2017 aerth <aerth@riseup.net>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

// ppanic - loud and quick
package ppanic

import (
	"fmt"
	"io"
	"runtime"

	"github.com/davecgh/go-spew/spew"
	"github.com/kr/pretty"
)

const author = "aerth"
const version = "0.0.1"

func Panic(i ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	println(fmt.Sprintf("%s:%d", fn, line))
	panic(pretty.Sprint(i))
}

func Spew(i ...interface{}) {
	_, fn, line, _ := runtime.Caller(1)
	println(fmt.Sprintf("%s:%d", fn, line))
	panic(spew.Sdump(i...))
}

func Panicf(f string, i ...interface{}) {
	panic(fmt.Sprintf(f, i...))
}

func Fdump(out io.Writer, i ...interface{}) (int, error) {
	return fmt.Fprintf(out, spew.Sdump(i...))
}
