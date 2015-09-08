package anonymiser

import (
	"fmt"
)

// looks like a key cache but when we get a new name we return the prefix plus a number.
// each new name has the number increased.
// Can be used to convert private names into a public anonymous set (e.g. ps-top)
// e.g.
// Anonymise( "table", "tablea" )  --> table1
// Anonymise( "table", "tableb" )  --> table2
// Anonymise( "table", "tablea" )  --> table1
// Anonymise( "db",    "my_db" )   --> db1
// Anonymise( "db",    "otherdb" ) --> db2
// Anonymise( "db",    "otherdb" ) --> db2
// Anonymise( "db",    "my_db" )   --> db1

type anonymous struct {
	prefix string
	last   int
	id     map[string]int
}

var prefixMap map[string]anonymous

// initinitialise structures
func init() {
	prefixMap = make(map[string]anonymous)
}

// does the name exist already
func (a anonymous) exists(name string) bool {
	_, ok := a.id[name]
	return ok
}

// return the anonymised name
func (a *anonymous) name(orig string) string {
	if a.exists(orig) {
		return fmt.Sprintf("%s%d", a.prefix, a.id[orig])
	}
	return a.add(orig)
}

// add a new value and return the anonymised name
func (a *anonymous) add(orig string) string {
	a.last++
	a.id[orig] = a.last
	return a.name(orig)
}

func Anonymise(prefix, name string) string {
	if _, ok := prefixMap[prefix]; !ok {
		b := anonymous{prefix: prefix, id: make(map[string]int)}
		b.add(name)
		prefixMap[prefix] = b
	}

	a := prefixMap[prefix]
	return a.name(name)
}
