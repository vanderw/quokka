# aba module

is a configure file parsing library, i want it to support 2 types of format at most: `.toml` and `.hjson`.

`json` is critical for human reading, 

`yaml` serials are all hells for copy and paste.

`ini` is too old you know.

# Usage

```golang
var t Aba

data := ReadConfigFromFile(file)
conf := struct {
    Addr string
    LogLevel int
}
t = hjson.New() // or toml.New()
t.Unmarshal(data, conf)
```

# 