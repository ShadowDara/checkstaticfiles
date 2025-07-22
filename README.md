# checkstaticfiles

This is a go module to include static files to binary (e.g. html, or
other) by encoding them and adding the encoded value to the binary.

The binary will then check on startup if these files are existing, if
not the executable will create them.


## Start

First create a `checkstaticfiles.config.yaml` file in the root of your
project

You can add Folders and Files to `paths`

```yaml
# Example Content
paths:
  - .gitignore
  - q.exe
```

then you need to installer the generator module

```sh
go install github.com/shadowdara/checkstaticfiles/generate
```

and run it with
```sh
csf.generate --package=main --output=checkstaticfiles.data.go --variable=CheckstaticfilesOutputJSONGz
```

add this to your git ignore to ignore the useless files
```sh
checkstaticfiles.output.json
checkstaticfiles.output.json.gz
checkstaticfiles.data.go        # if you dont changed the name
```


## Compile

Then compile your go binary with

**`shell`**
```sh
# Read the file
data=$(cat checkstaticfiles.output.json | tr -d '\n')

# package name and var name
PKG="main"
VAR="checkstaticfiles_data"

# add variable to the build
go build -ldflags "-X '$PKG.$VAR=$data'" -o myapp main.go
```

```sh
xxd -i checkstaticfiles.output.json.gz | \
sed -e '1s/.*/package main\n\nvar checkstaticfilesOutputJSONGz = []byte{/' \
    -e '$s/};/}/' > data.go
```


to compile with this command you need the include
```go
// will be overwritten at build
var checkstaticfiles_data = ""
```


## Use the data

add this to your code to execute the file creation on start

```go
func main() {
	checkfiles()
}
```


## Develepment

to generate the encoded files from the `checkstaticfiles.config.yaml` file
```sh
cd generate
go run csf.generate.go
```

(*csf* should a shortcut for checkstaticfiles)
