Himatch
====

## Overview

Himatch - Spare Time Sharing Service

## Build

### require

- go 1.5 or later

```
$ go get -u github.com/keita0q/RGL
$ cd /path/to/RGL
$ go build
```

## SetUp

以下のような設定ファイルを作成してください。

```
{
  "context_path":"himatch",
  "port":8080,
  "save_path" :"out"
}
```

## Usage

```
./himatch < -c [ path to config.json ] >
```
