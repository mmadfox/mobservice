
DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
echo "Install mobservice into $DIR"

go get github.com/Shaked/gomobiledetect
go get github.com/gorilla/mux
go get github.com/pmylund/go-cache
go build $DIR/mobservice.go
chmod a+x $DIR/mobservice
cp $DIR/mobservice $DIR/../../bin/mobservice
cp $DIR/rc.d/mobservicectl $DIR/../../bin/mobservicectl




