# anonymiser
Anonymise some values for showing to the "public"

# Description.

Looks like a key cache but when we get a new name we return the prefix plus a number each new name has the number increased.

Can be used to convert private names into a public anonymous set (e.g. ps-top)
e.g.
* anonymise.Anonymise( "table", "tablea" )  --> table1
* anonymise.Anonymise( "table", "tableb" )  --> table2
* anonymise.Anonymise( "table", "tablea" )  --> table1
* anonymise.Anonymise( "db",    "my_db" )   --> db1
* anonymise.Anonymise( "db",    "otherdb" ) --> db2
* anonymise.Anonymise( "db",    "otherdb" ) --> db2
* anonymise.Anonymise( "db",    "my_db" )   --> db1

There is basically one routine anonymise.Anonymise( "prefix", "some_name" )

# Installation

Install each library by doing:
`go get github.com/sjmudd/ps-top/anonymiser`

