package main
import "time"
import "github.com/samuel/go-zookeeper/zk"
import (
	"./config"
	"net/http"
	"io"
)
var i int = -1
var cw int = 0
var gcd int = 2
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

func zkHandler(addresses []string,path string)  ([]string, *zk.Conn,error) {
	zkAddr := addresses
	c,_, err := zk.Connect(zkAddr,time.Second*10)
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	children,_,_,err :=  c.ChildrenW("/services/"+path)
	defer func() {
		if err := recover(); err != nil {
			return
		}
	}()
	return children,c,err;

}

func getNodeInfo(w http.ResponseWriter,r *http.Request)  {
	path := r.URL.Query().Get("service")
	nodeArr,c ,err:= zkHandler(config.ZkAddresses,path);
	defer func() {
		if err := recover(); err != nil {
			io.WriteString(w,"{\"code\":400,\"data\":\"zk error\"}")
			return
		}
	}()
	node := getNode(nodeArr)
	jsonData ,_,_,err := c.GetW("/services/"+path+"/"+node)
	if err !=nil {
	}
	defer c.Close()

	io.WriteString(w,"{\"code\":200,\"data\":"+string(jsonData)+"}")

}

func main(){
	http.HandleFunc("/zk",getNodeInfo)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}
}