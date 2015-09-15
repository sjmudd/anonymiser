package anonymiser

import (
	"fmt"
	"sort"
)

/*
Designed to anonymise a group of strings.
We provide a group name and string and get back an anonymised string based on the group name.
Can be used to convert private names into a public anonymous set (e.g. ps-top)
e.g.
Anonymise( "table", "tablea" )  --> table1
Anonymise( "table", "tableb" )  --> table2
Anonymise( "table", "tablea" )  --> table1
Anonymise( "db",    "my_db" )   --> db1
Anonymise( "db",    "otherdb" ) --> db2
Anonymise( "db",    "otherdb" ) --> db2
Anonymise( "db",    "my_db" )   --> db1
*/

type anonymous struct {
	group string
	last  int
	id    map[string]int
}

var (
	groupMap map[string]anonymous
	enabled  bool
)

// initialise structures - a global per app set of string groups
func init() {
	enabled = true
	Clear()
}

// Clear removes any previous data that might have been stored.
func Clear() {
	// no need to clean up explicitly if we had old data?
	// I guess go garbage collects but might be nice to do this???
	groupMap = make(map[string]anonymous)
}

// Enable the anonymiser. We provide a boolean to signal intent
func Enable(setEnabled bool) {
	enabled = setEnabled
}

// Enabled returns if anonymising is enabled
func Enabled() bool {
	return enabled
}

// does the group name exist already?
func (a anonymous) exists(group string) bool {
	_, ok := a.id[group]
	return ok
}

// return the anonymised name
func (a *anonymous) name(orig string) string {
	if a.exists(orig) {
		return fmt.Sprintf("%s%d", a.group, a.id[orig])
	}
	return a.add(orig)
}

// add a new value and return the anonymised name
func (a *anonymous) add(orig string) string {
	a.last++
	a.id[orig] = a.last
	return a.name(orig)
}

// Anonymise takes a group and name and returns a string consisting
// of the group plus a number. If the string has been seen before
// then the same name is returned.  Use different groups if you want
// to store different sets of names.
func Anonymise(group, name string) string {
	if !enabled {
		return name
	}
	if _, ok := groupMap[group]; !ok {
		b := anonymous{group: group, id: make(map[string]int)}
		b.add(name)
		groupMap[group] = b
	}

	a := groupMap[group]
	return a.name(name)
}

// Groups returns a sorted list of known groups
func Groups() []string {
	groups := make([]string, 0)
	for grp, _ := range groupMap {
		groups = append(groups, grp)
	}
	sort.Strings(groups)

	return groups
}
