package main
//import "fmt"
import "time"
import "github.com/samuel/go-zookeeper/zk"
import (
	//"encoding/json"
	"net/http"
	"io"
	"fmt"
)


type zkAddrAndPort struct {
	Addr interface{}
	Tport interface{}
}


func zkHandler(w http.ResponseWriter,r *http.Request)  {
	zkAddr := []string{"192.168.100.24"}
	//dateTimetime :=time.Now()
	var d zk.Dialer
	c,_, err := zk.ConnectWithDialer(zkAddr,time.Second,d)
	fmt.Println(d)
	if err != nil {
		panic(err)
	}
	children,_,_,err :=  c.ChildrenW("/services/cart")
	if err != nil {
		panic(err)
	}
	jsonData ,_,_,err := c.GetW("/services/cart/"+children[0])
	jsonStr :="";
	for  key:=  range jsonData{
		jsonStr += string(jsonData[key]);
	}
/*	var res map[string] interface{}
	json.Unmarshal(jsonData,&res)
	var address interface{}
	var tport interface{}
	for key,value:= range res {
		if key == "address" {
			address = value
		}
		if key == "port" {
			tport = value
		}
	}
	zkInfo := &zkAddrAndPort{address,tport}
	//fmt.Println(json.Marshal(zkInfo))
	result, err := json.Marshal(zkInfo)
	if err != nil{

	}*/
	 io.WriteString(w,string(jsonStr))

}

func main(){


	http.HandleFunc("/func",zkHandler)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		panic(err)
	}



}