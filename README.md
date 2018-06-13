# dep-metadata-extractor
Golang's `dep` dependency manager provides an option for metadata within its `Gopkg.toml` manifest.  `dep` designers
intentionally avoid using this metadata in any way; it claims the property is available for external tools to rely on.

This is one such tool, in the absence of a simpler TOML parsing utility like `jq` or
Python's `json.tool` for JSON.

Example `Gopkg.toml` with a metadata element describing the fully qualified package name:

```TOML
[metadata]
name = "github.com/someuser/someproject"
```

In order to use this value in scripting/automation, the following command can be used:

```commandline
$ ./dep-metadata-extractor -keys=name -withkey=false
github.com/someuser/someproject
```

Full usage information:

```commandline
Usage of dep-metadata-extractor:
  -file string
    	path to the dep manifest. (default "Gopkg.toml")
  -keys value
    	comma-separated list of keys to extract
  -withkey
    	output with keys or not (default true)
```
