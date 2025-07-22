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
checkstaticfiles.generate
```

---


## Dev
to generate the encoded files from the `checkstaticfiles.config.yaml` file
```sh
go run generate/main.go
```