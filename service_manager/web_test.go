package service_manager

import (
	"fmt"
	"github.com/jamestack/snow"
	"testing"
)

type User struct {
	*snow.Node
}

func TestServiceManager(t *testing.T) {
	cluster := snow.NewClusterWithLocal()
	done,err := cluster.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		<-done
	}()

	_,err = cluster.MountRandNode("SnowNodes", &ServiceManager{
		Service:       []ServiceInfo{
			{
				Name:    "Gate",
				Remark:  "网关组件",
				Methods: nil,
				Fields:  nil,
			},
		},
		WebListenAddr: "127.0.0.1:8080",
	})

	_,_ = cluster.MountRandNode("Gate", &User{})
	_,_ = cluster.MountRandNode("Gate", &User{})
	_,_ = cluster.MountRandNode("Gate", &User{})
	_,_ = cluster.MountRandNode("Gate", &User{})
	_,_ = cluster.MountRandNode("Gate", &User{})



}
