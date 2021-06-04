package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"os"
	"strings"
	"time"
)

var timeout time.Duration
var server string
var id string
var node bool
var zabbixdiscovery bool
var zabbixnodediscovery bool
var address []string

type zabbixdata struct {
	Data []map[string]string `json:"data"`
}

func init() {
	flag.StringVar(&server, "server", "192.168.1.1:2181", "zk地址")
	flag.DurationVar(&timeout, "timeout", time.Second * 10, "设置连接超时时间")
	flag.StringVar(&id, "id", "0", "channel ID")
	flag.BoolVar(&node, "node", false, "启用node监控")
	flag.BoolVar(&zabbixdiscovery, "zabbixdiscovery", false, "返回zabbix自动发现channel json数据")
	flag.BoolVar(&zabbixnodediscovery, "zabbixnodediscovery", false, "返回zabbix自动发现node json数据")
}

func GetChannel(conn zk.Conn) {
	if id == "0"{
		fmt.Println("请使用-id参数")
		os.Exit(2)
	}
	list, _, err := conn.Children("/otter/channel")
	if err != nil {
		fmt.Println("zk连接异常,获取不到节点信息")
		os.Exit(2)
	}
	for _, cn := range list {
		if cn == id {
			res, _, _ := conn.Get("/otter/channel/" + cn)
			status := string((res))
			if len(status) == 0 {
				fmt.Println("NULL")
			} else {
				fmt.Println(strings.Replace(status, "\"", "", 2))
			}
			return
		}
	}
	fmt.Println("NONE")
}

// node状态
func GetNode(conn zk.Conn, node bool) {
	list, _, err := conn.Children("/otter/node")
	if err != nil {
		fmt.Println("zk连接异常,获取不到节点信息")
		os.Exit(2)
	}
	if node {
		res := make([]map[string]string, 0)
		for _, node := range list {
			data := make(map[string]string)
			data["{#NODE}"] = node
			res = append(res, data)
		}
		json_res, _ := json.Marshal(res)
		fmt.Println(string(json_res))
	} else {
		if id == "0"{
			fmt.Println("请使用-node -id 参数")
			os.Exit(2)
		}
		for _, node := range list {
			if node == id {
				fmt.Println(1)
				return
			}
		}
		fmt.Println(0)
	}
}
// zabbix
func zabbix(conn zk.Conn) {
	var zabbix zabbixdata
	list := make([]map[string]string, 0)
	res, _, _ := conn.Children("/otter/channel")
	for _, id := range res {
		data := make(map[string]string)
		data["{#CHANNEL_ID}"] = id
		list = append(list, data)
	}
	zabbix.Data = list
	r, _ := json.Marshal(list)
	fmt.Println(string(r))

}

func main() {
	flag.Usage = func() {
		fmt.Println("参数说明:")
		flag.PrintDefaults()
		fmt.Println("默认channel返回值说明:")
		fmt.Printf("  START 启动\n  STOP  停止\n  PAUSE 挂起\n  NONE  不存在\n")
		fmt.Println("otter node返回值说明:")
		fmt.Printf("  1 正常\n  0 停止\n")
	}
	flag.Parse()
	address = append(address, server)
	conn, _, _ := zk.Connect(address, time.Second*timeout, zk.WithLogInfo(false))
	defer conn.Close()
	if node {
		GetNode(*conn, false)
		return
	}
	if zabbixnodediscovery {
		GetNode(*conn,true)
		return
	}
	if zabbixdiscovery {
		zabbix(*conn)
		return
	}
	GetChannel(*conn)
}
