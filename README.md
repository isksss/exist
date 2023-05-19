# which
windowsにwhichコマンドがなくて、不便だったので作った。

## install
```bash
go install github.com/isksss/which
```

## usage
パスが見つかった場合のみ、出力されます。  
エラーや、パスが通ってない場合は何も出力されません。  
```bash
# 例1
$ which go
C:\Users\isksss\scoop\shims\go.exe

# 例2
$ which xxxx

```