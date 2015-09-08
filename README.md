# anonymiser
Anonymise some values for showing to the "public"

# Description.

Looks like a key cache but when we get a new name we return the prefix plus a number each new name has the number increased.

Can be used to convert private names into a public anonymous set (e.g. ps-top)

There is basically one routine `anonymise.Anonymise( "prefix", "some_name" )`.

I wanted to anonymise the database and table names this is the prefix in the Anonymise() function.
The second parameter is the name to anonymise and basically each
new name gets an id starting at 1. This id is added to the end of
the prefix and that's what is returned as the anonymised name.

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
var hide_me []string { ... } // holds important strings (maybe with duplicates)

... // fill hide_me with useful data

for i := range hide_me {
	fmt.Println( "secret:", hide_me[i], "==>", anonymise.Anonymise( "anonymised", hide_me[i] ) )
}
``` 

# Installation

Install by doing:
* `go get github.com/sjmudd/ps-top/anonymiser`

