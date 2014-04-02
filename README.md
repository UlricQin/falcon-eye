falcon-eye
==========

linux monitor tool. an agent running on your host collect and display performance data. just like https://github.com/afaqurk/linux-dash


### install

```
mkdir -p $GOPATH/src/github.com/ulricqin
cd $GOPATH/src/github.com/ulricqin
git clone https://github.com/UlricQin/falcon-eye.git
cd $GOPATH/src
go get github.com/ulricqin/falcon-eye/...
cd github.com/ulricqin/falcon-eye
vi package # modify GOROOT and GOPATH
./package
cd /tmp/qinxh/release_falcon_eye/falcon_eye
./startup
# default http port is: 1988. goto: http://localhost:1988
```
