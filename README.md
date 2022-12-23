<h1 align="center">
<img alt="Koi logo" src="asset/euler.jpeg" width="500px"/><br/>
Euler
</h1>
<p align="center">Generate dependency graph from golang struct files</p>

<p align="center">
<a href="https://pkg.go.dev/github.com/mehditeymorian/koi/v3?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.19+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>&nbsp;
<img src="https://img.shields.io/badge/license-apache_2.0-red?style=for-the-badge&logo=none" alt="license" />

<img src="https://img.shields.io/badge/Version-1.0.0-informational?style=for-the-badge&logo=none" alt="version" />
</p>

## Installation
```shell
go get github.com/mehditeymorian/euler@latest
```

## How to Use
```shell
euler -p ./structs-dir -e "a.go,b.go,c.go" -f
# output: out.svg
```

## Options
- **--path | -p (string)**: directory address of where the golang structs are
- **--exclude | -e (string, items separated by comma)**: list of files excluded from output graph separated by comma
- **--fields | -f (boolean)**: include each struct fields


## Todo List
- [ ] output path
- [ ] output format