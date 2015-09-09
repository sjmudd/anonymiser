package anonymiser

import (
	"testing"
)

func TestAnonymise(t *testing.T) {
	cases := []struct{ prefix, name, want string }{
		{"a", "b", "a1"},
		{"table", "tablea", "table1"},
		{"table", "tableb", "table2"},
		{"table", "tablea", "table1"},
		{"db", "my_db", "db1"},
		{"db", "otherdb", "db2"},
		{"db", "otherdb", "db2"},
		{"db", "my_db", "db1"},
	}

	for _, c := range cases {
		got := Anonymise(c.prefix, c.name)
		if got != c.want {
			t.Errorf("Anonymise(%s,%s) => %s, want %s", c.prefix, c.name, got, c.want)
		}

	}
}

func TestDisabled(t *testing.T) {
	cases := []struct{ prefix, name, want string }{
		{"a", "b", "b"},
		{"table", "tablea", "tablea"},
		{"table", "tableb", "tableb"},
		{"table", "tablea", "tablea"},
		{"db", "my_db", "my_db"},
		{"db", "otherdb", "otherdb"},
		{"db", "otherdb", "otherdb"},
		{"db", "my_db", "my_db"},
	}

	Enable(false)

	for _, c := range cases {
		got := Anonymise(c.prefix, c.name)
		if got != c.want {
			t.Errorf("Anonymise(%s,%s) => %s, want %s", c.prefix, c.name, got, c.want)
		}

	}
}

func TestClear(t *testing.T) {
	Enable(true)

	cases := []struct{ prefix, name, want string }{
		{"prefix", "valueXX", "prefix1"},
		{"prefix", "valueZZ", "prefix2"},
	}

	for _, c := range cases {
		got := Anonymise(c.prefix, c.name)
		if got != c.want {
			t.Errorf("Anonymise(%s,%s) => %s, want %s", c.prefix, c.name, got, c.want)
		}

	}
	// Reset data
	Clear()

	cases2 := []struct{ prefix, name, want string }{
		{"prefix", "valueZZ", "prefix1"},
		{"prefix", "valueXX", "prefix2"},
	}
	for _, c := range cases2 {
		got := Anonymise(c.prefix, c.name)
		if got != c.want {
			t.Errorf("Anonymise(%s,%s) => %s, want %s", c.prefix, c.name, got, c.want)
		}

	}
}

func BenchmarkEnabled(b *testing.B) {
	Enable(true)
	for i := 0; i < b.N; i++ {
		Anonymise("prefix", "some_name")
	}
}
