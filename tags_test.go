// Copyright (c) 2012-2013 Jason McVetta.  This is Free Software, released under
// the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.
// Resist intellectual serfdom - the ownership of ideas is akin to slavery.

package stags

import (
	"log"
	"reflect"
	"testing"
)

type myStruct struct {
	Foo string `mytag:"foo"`
	Bar int    `mytag:"bar"`
}

func TestTypeFields(t *testing.T) {
	p := Parser{Name: "mytag"}
	fs := p.TypeFields(reflect.TypeOf(myStruct{}))
	log.Println(fs)
}
