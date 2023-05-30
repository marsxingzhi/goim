package xzetcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"log"
	"net"
	"strconv"
	"time"
)

type EtcdRegister struct {
	cli          *clientv3.Client
	endpoints    []string
	ttl          int
	host         string
	port         int
	schema       string
	serviceName  string
	serviceKey   string
	serviceValue string
	connected    bool
	reChan       chan struct{}
}

// Register schema：以etcd///api为例，schema为etcd
// endpoints：节点
// host：注册到etcd的服务的host
// port：注册到etcd服务的port
// serverName: 注册到etcd的服务的名字
// ttl：有效时间
func Register(schema string, endpoints []string, host string, port int, serviceName string, ttl int) error {
	serviceValue := net.JoinHostPort(host, strconv.Itoa(port))
	serviceKey := fmt.Sprintf("%s:///%s/", schema, serviceName) + serviceValue
	fmt.Printf("[xzetcd] register serviceKey: %v, serviceValue: %v, serviceName: %v, ttl: %v\n", serviceKey, serviceValue, serviceName, ttl)

	etcdRegister := &EtcdRegister{
		endpoints:    endpoints,
		ttl:          ttl,
		host:         host,
		port:         port,
		schema:       schema,
		serviceName:  serviceName,
		serviceKey:   serviceKey,
		serviceValue: serviceValue,
		reChan:       make(chan struct{}),
	}
	//etcdRegister.reRegister()
	etcdRegister.register()
	return nil
}

func (er *EtcdRegister) register() (err error) {
	defer func() {
		if err != nil {
			er.reChan <- struct{}{}
		}
	}()

	// 1. 创建etcd client
	// @xz 注意使用下面这个方式创建的话，请求不到，我理解应该是跟Endpoints有关，这里的Endpoints是client的，naming下面也有Endpoints，这两者是不同的
	// 用这个方法创建的时候，Endpoints需要时可访问的。这里我遇到一个问题，最开始endpoints在config中配置的都是字符串内容，并不是host+port，因此导致一直访问不到etcd
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   er.endpoints,
		DialTimeout: 5 * time.Second,
	})

	//cli, err := clientv3.NewFromURL(etcd_url)

	if err != nil {
		fmt.Println("[xzetcd] failed to register: ", err)
		return
	}

	// 2. 创建lease
	ctx, cancel := context.WithCancel(context.Background())
	lease, err := cli.Grant(ctx, int64(er.ttl))
	if err != nil {
		fmt.Println("[xzetcd] failed to grant lease: ", err)
		return
	}

	// 3. 注册服务。       clientv3.WithLease将租约ID与key绑定
	_, err = cli.Put(ctx, er.serviceKey, er.serviceValue, clientv3.WithLease(lease.ID))
	if err != nil {
		fmt.Println("[xzetcd] failed to put k-v to etcd: ", err)
		return
	}

	// 4. 发送心跳
	alive, err := cli.KeepAlive(ctx, lease.ID)
	if err != nil {
		fmt.Println("[xzetcd] failed to keep alive: ", err)
		return
	}
	er.cli = cli

	go func() {
		for {
			select {
			case _, ok := <-alive:
				er.connected = ok
				if !ok {
					fmt.Println("[xzetcd] failed to continue lease")
					cancel()
					er.reChan <- struct{}{}
					return
				} else {
					fmt.Println("[xzetcd] keep alive")
				}
			}
		}
	}()
	return
}

func (r *EtcdRegister) reRegister() {
	go func() {
		var (
			ok bool
		)
		for {
			select {
			case _, ok = <-r.reChan:
				if ok == false {
					return
				}
				time.Sleep(1 * time.Second)
				r.register()
			}
		}
	}()
}

//-----------

const (
	etcd_url    = "http://127.0.0.1:2379"
	server_name = "goim_auth_server"
	ttl         = 10
)

func RegisterV2(addr string) error {
	etcdClient, err := clientv3.NewFromURL(etcd_url)
	if err != nil {
		return err
	}

	em, err := endpoints.NewManager(etcdClient, server_name)
	if err != nil {
		return err
	}

	// TTL 单位是秒
	lease, err := etcdClient.Grant(context.Background(), ttl)
	if err != nil {
		return err
	}

	// 注册服务的地址
	err = em.AddEndpoint(etcdClient.Ctx(), fmt.Sprintf("%s/%s", server_name, addr),
		endpoints.Endpoint{
			Addr: addr,
		},
		clientv3.WithLease(lease.ID),
	)
	if err != nil {
		return err
	}

	// keep alive
	alive, err := etcdClient.KeepAlive(context.Background(), lease.ID)
	if err != nil {
		return err
	}

	go func() {
		for {
			<-alive
			log.Println("[etcd-server] keep alive!")
		}
	}()

	return nil
}
