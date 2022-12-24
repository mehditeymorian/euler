<h1 align="center">
<img alt="Koi logo" src="asset/euler.jpeg" width="500px"/><br/>
Euler
</h1>
<p align="center">Generate dependency graph from golang struct files</p>

<p align="center">
<a href="https://pkg.go.dev/github.com/mehditeymorian/euler/v1?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.19+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>&nbsp;
<img src="https://img.shields.io/badge/license-MIT-red?style=for-the-badge&logo=none" alt="license" />

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
- **--path | -p (string)**: Directory path of the golang models
- **--exclude | -e (string)**: List of files excluded from output graph separated by comma
- **--fields | -f (boolean)**: Include struct fields in the output graph


## Todo List
- [ ] output path
- [ ] output format
- [ ] improve error handling