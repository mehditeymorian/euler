<h1 align="center">
Euler
</h1>
<p align="center">Generate dependency graph from golang structs and components</p>

<p align="center">
<a href="https://pkg.go.dev/github.com/mehditeymorian/euler/v1?tab=doc"target="_blank">
    <img src="https://img.shields.io/badge/Go-1.19+-00ADD8?style=for-the-badge&logo=go" alt="go version" />
</a>&nbsp;
<img src="https://img.shields.io/badge/license-MIT-red?style=for-the-badge&logo=none" alt="license" />

<img src="https://img.shields.io/badge/Version-1.1.0-informational?style=for-the-badge&logo=none" alt="version" />
</p>

## Installation
```shell
go install github.com/mehditeymorian/euler@latest
```

## How to Use

### Generate for Components
Generate graph for component
```shell
euler component -p GIT_REPO|LOCAL_DIR
```

For example the following is dependency graph is for [Docker Compose](https://github.com/docker/compose) project.

![Docker Compose Project Dependency Graph](assets/compose-repo.svg)


### Generate for Struct Models
```shell
euler struct -p ./structs-dir -e "a.go,b.go,c.go" -f
# output: out.svg
```