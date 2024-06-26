[![Actions Status](https://github.com/tmdtky/go-fizzbuzz/workflows/Go/badge.svg)](https://github.com/tmdtky/go-fizzbuzz/actions)

# go-fizzbuzz
Go 1.18を用いてfizzbuzz問題を実行
- 3の倍数の場合は「Fizz」、5の場合は「Buzz」、3の倍数かつ5の倍数は「FizzBuzz」と出力する。また、いずれでもない場合はその数を出力する
- 実行したら1から100までの数字において、上記条件で標準出力に改行つきで出力するプログラムを作成

# Features
- testingパッケージによるテストコードの作成
- Docker、docker-composeによる環境構築
- GitHub Actionによる自動テスト

# Usage
- Docker Desktopを事前に起動してください。
- 起動方法
```bash
git clone https://github.com/tmdtky/go-fizzbuzz.git
cd <git cloneしたディレクトリ>
docker-compose up --build
```
- 単体テストの実施方法
```bash
cd <git cloneしたディレクトリ>
docker-compose run fizzbuzz go test
```
- 起動後のdockerプロセス削除方法
```bash
docker-compose down
```

# Note
docker-composeを起動した際、ビルド及びテストが成功した場合、
アプリケーションが実行されます。