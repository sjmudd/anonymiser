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

var (
	prefixMap map[string]anonymous
	enabled bool
)

// initinitialise structures
func init() {
	enabled = true
	Clear()
}

// Clear removes any previous data that might have been stored.
func Clear() {
	// no need to clean up explicitly if we had old data?
	// I guess go garbage collects but might be nice to do this???
	prefixMap = make(map[string]anonymous)
}

// Enable the anonymiser. We provide a boolean to signal intent
func Enable(set_enabled bool) {
	enabled = set_enabled
}

// Enabled returns if anonymising is enabled
func Enabled() bool {
	return enabled
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

// Anonymise takes a prefix and name and returns a string consisting
// of the prefix plus a number. If the string has been seen before
// then the same name is returned.  Use different prefixes if you want
// to store different sets of names.
func Anonymise(prefix, name string) string {
	if ! enabled {
		return name
	}
	if _, ok := prefixMap[prefix]; !ok {
		b := anonymous{prefix: prefix, id: make(map[string]int)}
		b.add(name)
		prefixMap[prefix] = b
	}

	a := prefixMap[prefix]
	return a.name(name)
}
