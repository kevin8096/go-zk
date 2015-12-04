package main
import "time"
import "github.com/samuel/go-zookeeper/zk"
import (
	"./config"
	"net/http"
	"io"
	"os"
)
var i int = -1
var cw int = 0
var gcd int = 2
var ZkConn,_,_ = zk.Connect(config.ZkAddresses,time.Second*10)


//轮询节点
func getNode(NodeString []string) string {
	for {
		i = (i + 1) % len(NodeString)
		if i == 0 {
			cw = cw - gcd
			if cw <= 0 {
				cw = 1
				if cw == 0 {
					return ""
				}
			}
		}
		if weight:= 1; weight >= cw {
			return string(NodeString[i])
		}
	}
}
func zkHandler(addresses []string,path string)  ([]string,error) {

	children,_,err :=  ZkConn.Children("/services/"+path)
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	return children,err;

}
func getNodeInfo(w http.ResponseWriter,r *http.Request)  {
	path := r.URL.Query().Get("service")
	nodeArr,err:= zkHandler(config.ZkAddresses,path);
	defer func() {
		if err := recover(); err != nil {
			io.WriteString(w,"{\"code\":400,\"data\":\"zk error\"}")
			return
		}
	}()
	node := getNode(nodeArr)
	jsonData ,_,err := ZkConn.Get("/services/"+path+"/"+node)
	if err !=nil {

	}

	io.WriteString(w,"{\"code\":200,\"data\":"+string(jsonData)+"}")

}


func main(){
	http.HandleFunc("/zk",getNodeInfo)
	os.Hostname();
	hostName,err := os.Hostname();
	if err != nil {
		panic(err)
	}
	errr := http.ListenAndServe(hostName+":"+config.HttpPort, nil)
	if errr != nil {
		panic(errr)
	}
}