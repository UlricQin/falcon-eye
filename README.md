falcon-eye
==========

linux monitor tool. an agent running on your host collect and display performance data. just like https://github.com/afaqurk/linux-dash


### install

```
mkdir -p $GOPATH/src/github.com/ulricqin
cd $GOPATH/src/github.com/ulricqin && git clone https://github.com/UlricQin/falcon-eye.git
go get github.com/ulricqin/falcon-eye/...

cd falcon-eye && go build
./control start
# default http port is: 1988. goto: http://localhost:1988
# control usage: ./control start|stop|restart|status
```

*too complicated?* use gopm(https://github.com/gpmgo/gopm) and:
```
git clone https://github.com/UlricQin/falcon-eye.git
cd falcon-eye && gopm build && ./falcon-eye
```
