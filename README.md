# anonymiser
Anonymise some values for showing to the "public"

# Description.

Anonymiser allows a group of strings to be anonymised: that is
converted into another set of string values.

Strings can be split into groups so that similar types of
strings are anonymised together.

An example of this might be a program which is exposing database
and table names from an internal company database server. Here
we would like to anonymise the names being used but share the
output of the screen with anonymous information.

All database names would be converted into `db1`, `db2` ... `dbn`,
and all table names converted into `table`, `table2` ... `tablen`.

Other similar use cases can be imagined.

An example: [`ps-top`](https://godoc.org/github.com/sjmudd/ps-top])

In `ps-top` I wanted to anonymise the host, database and table names
which were shown as they might expose internal information to third parties.
This package made that easy.

There is basically one routine `anonymise.Anonymise( "group", "some_name" )`.

The first parameter is the name of the group of strings to be
anonyised, The second parameter is the name to anonymise and basically
each new name gets an id starting at 1. This id is added to the end
of the group name and that's what is returned as the anonymised
name.

e.g.
To anonymise some database names:
* anonymise.Anonymise( "db",    "my_db" )   --> db1
* anonymise.Anonymise( "db",    "otherdb" ) --> db2
* anonymise.Anonymise( "db",    "otherdb" ) --> db2
* anonymise.Anonymise( "db",    "my_db" )   --> db1

To anonymise some table names:
* anonymise.Anonymise( "table", "important_name" )  --> table1
* anonymise.Anonymise( "table", "something_else" )  --> table2
* anonymise.Anonymise( "table", "important_name" )  --> table1

You can use as many prefixes as you like.

I guess in real code you'd do something like this:
```
var secret []string { ... } // holds strings of secrent information (maybe with duplicates)

... // fill secret with useful data

for i := range secret {
	fmt.Println( "secret:", secret[i], "==>", anonymise.Anonymise( "anonymised", secret[i] ) )
}
``` 

# Installation

Install by doing:
* `go get github.com/sjmudd/anonymiser`

# Documentation

Documentation can be found using `godoc` or at [https://godoc.org/github.com/sjmudd/anonymiser]
