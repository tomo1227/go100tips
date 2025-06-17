# go mock

## install

```shell
go install go.uber.org/mock/mockgen@latest
```

バージョンの確認は以下。

```shell
$ mockgen -version
v0.5.2
```

## mock自動生成

全てを指定する場合は以下のコマンド。

```shell
mockgen -source=***.go
```

ファイルを指定する場合は以下

```shell
mockgen -source=main.go -destination=mock/mock_main.go -exclude_interfaces=Bird -package=mock -package=mock
```
