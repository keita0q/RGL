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
  "save_path" :"himatch.db"
}
```

## Usage

```
./himatch < -c [ path to config.json ] >
```

## REST API

### Get

#### Search Spare Time
ndde body
```
{
  "time": "2013-09-29T17:21:12Z",
  "tags": [
    "student",
  ]
}
```


### POST

#### Save User

need body

```
{
  "id": "hoge@gmail.com",
  "password": "hgoehgoehgoehgoe",
  "name": "onishi",
  "age": 23,
  "sns": {
    "email": "hoge@yahoo.jp"
  },
  "tags": [
    "student",
  ]
}
```

#### Save Spare Time

need body

```
{
  "id": "hoge",
  "user_id": "hgoe@gmail.com",
  "start": "2013-09-29T17:21:12Z",
  "end": "2013-09-29T17:41:12Z",
  "Comment": "develop shimane",
  "Tags": [
    "shimane",
    "develop"
  ]
}
```

time format is RFC 3339

### PUT

#### Edit User

need body

```
{
  "id": "hoge@gmail.com",
  "name": "onishi",
  "age": 23,
  "sns": {
    "email": "hoge@yahoo.jp"
  },
  "tags": [
    "student",
  ]
}
```

#### Edit Spare Time

need body
```
{
  "id": "hoge",
  "user_id": "hgoe@gmail.com",
  "start": "2013-09-29T17:21:12Z",
  "end": "2013-09-29T17:41:12Z",
  "Comment": "develop shimane",
  "Tags": [
    "shimane",
    "develop"
  ]
}
```

time format is RFC 3339
