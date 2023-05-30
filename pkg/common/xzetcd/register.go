package xzetcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
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
	fmt.Printf("[xzetcd] register serviceKey: %v, serviceValue: %v\n", serviceKey, serviceValue)

	etcdRegister := &EtcdRegister{
		endpoints:    endpoints,
		ttl:          ttl,
		host:         host,
		port:         port,
		schema:       schema,
		serviceName:  serviceName,
		serviceKey:   serviceKey,
		serviceValue: serviceValue,
	}

	etcdRegister.register()
	return nil
}

func (er *EtcdRegister) register() {
	// 1. 创建etcd client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   er.endpoints,
		DialTimeout: 5 * time.Second,
	})
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
					return
				}
			}
		}
	}()
	return
}
