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
csf.generate
  --package=main
  --output=checkstaticfiles.data.go
  --variable=CheckstaticfilesOutputJSONGz
```
(*csf* should a shortcut for checkstaticfiles)

- package is the name of you go package
- output the name of the output file
- variable, the variable the data will be saved

add this to your git ignore to ignore the useless files
```sh
checkstaticfiles.output.json
checkstaticfiles.output.json.gz
checkstaticfiles.data.go          # if you dont changed the name
```


## Use the data

then import in your code
```sh
"github.com/shadowdara/checkstaticfiles"
```

and run
```sh
go get github.com/shadowdara/checkstaticfiles
```

and then run on your programm start to create the files
```go
shadowdara_checkstaticfiles.Checkfiles(CheckstaticfilesOutputJSONGz)
```


## Disclaimer

this module is develepment andNOT stable yet


## TODO
- [ ] add option for ignore paths
- [ ] add option to add files via the extension `(*.html)`
