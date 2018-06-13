package main

import "testing"

var manifesttests = []struct {
	in      map[string]interface{}
	keys    keylist
	withkey bool
	out     string
}{
	// should not display anything if keys are missing
	{map[string]interface{}{"foo": "bar"}, keylist{}, true, ""},
	{map[string]interface{}{"foo": "bar"}, keylist{}, false, ""},

	// happy path, one key defined
	{map[string]interface{}{"foo": "bar"}, keylist{"foo"}, true, "foo=bar\n"},
	{map[string]interface{}{"foo": "bar"}, keylist{"foo"}, false, "bar\n"},

	// happy path, two keys defined
	{map[string]interface{}{"foo": "bar", "baz": "test"}, keylist{"foo", "baz"}, true, "foo=bar\nbaz=test\n"},
	{map[string]interface{}{"foo": "bar", "baz": "test"}, keylist{"foo", "baz"}, false, "bar\ntest\n"},

	// should filter for the specific key
	{map[string]interface{}{"foo": "bar", "baz": "test"}, keylist{"foo"}, true, "foo=bar\n"},
	{map[string]interface{}{"foo": "bar", "baz": "test"}, keylist{"foo"}, false, "bar\n"},
}

func TestQueryManifest(t *testing.T) {
	for _, tt := range manifesttests {
		actual := queryManifest(manifest{Metadata: tt.in}, tt.keys, tt.withkey)

		if actual != tt.out {
			t.Errorf("queryManifest() expected %s got %s", tt.out, actual)
		}
	}
}
