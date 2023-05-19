# exist

## 概要(Overview)

`exist` は、Windows に `which` コマンドがないことを解決するために作られたツールです。パスが見つかった場合のみ出力され、エラーやパスが通ってない場合は何も出力されません。  
exist is a tool created to solve the absence of the which command in Windows. It only outputs if a path is found, and does not output anything in case of an error or if the path is not found.

## インストール(Installation)

Go 言語を使用していますので、以下のコマンドでインストールすることができます：  
As this is implemented in Go, you can install it with the following command:

```bash
go install github.com/isksss/exist
```

## 使い方(usage)

コマンドラインから exist コマンドを使用します。例えば：  
You can use the exist command from the command line. For example:

```bash
# 例1
$ exist go
C:\Users\isksss\scoop\shims\go.exe

# 例2
$ exist xxxx

```
