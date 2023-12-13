IFCファイルの変換をHTTPリクエストでできるようにする。

変換コマンドは公式コンテナで提供されている。これをHTTPリクエストで実行し、結果をHTTPレスポンスで取れるようにする。
```
docker run -it --rm -v $(pwd):/work aecgeeks/ifcopenshell IfcConvert work/chairs.ifc work/chairs.obj -f
```

IfcConvertコマンドが本体。Cの共有ライブラリの依存があり、このコンテナ外で実行するのは面倒そう。

## 実行方法

コンテナを立ち上げる。

```
docker-compose up -d
```

コンテナ内でテストを実行する。

```
make test
```

curlで動作を確かめる。変換後のOBJファイルのバイナリを返す。

```
curl -X POST http://localhost:8989/exec \
  -F "file=@chairs.ifc" \
  -H "Content-Type: multipart/form-data"
```

あるいはリクエストしたファイルを保存する。

```
curl -X POST http://localhost:8989/exec \
  -F "file=@chairs.ifc" \
  -H "Content-Type: multipart/form-data" > test.obj
```

## サンプルデータ元

https://www.steptools.com/docs/stpfiles/ifc/Tabel_Chairs.ifc
