# exist

windows に which コマンドがなくて、不便だったので作った。

## install

```bash
go install github.com/isksss/exist
```

## usage

パスが見つかった場合のみ、出力されます。  
エラーや、パスが通ってない場合は何も出力されません。

```bash
# 例1
$ exist go
C:\Users\isksss\scoop\shims\go.exe

# 例2
$ exist xxxx

```
