# go-movie
Golangで動画を処理する

```bash
# OpenCVの依存パッケージをインストール
sudo apt-get update
sudo apt-get install -y libopencv-dev

# Goのパッケージ管理ツールをインストール
sudo apt-get install -y golang

# gocvをインストール
go get -u -d gocv.io/x/gocv
cd $GOPATH/pkg/mod/gocv.io/x/gocv@<latest-version>
make install
```