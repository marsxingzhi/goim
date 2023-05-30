package xzetcd

import (
	"context"
	"fmt"
	"github.com/marsxingzhi/goim/pkg/config"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"strings"
	"sync"
	"time"
)

type Resolver struct {
	grpcClientConn     *grpc.ClientConn
	cli                *clientv3.Client
	resolverClientConn resolver.ClientConn
	schema             string
	serviceName        string
	startWatchRevision int64
}

func (rv *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {

	rv.resolverClientConn = cc

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	prefix := fmt.Sprintf("%s:///%s/", rv.schema, rv.serviceName)
	resp, err := rv.cli.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, fmt.Errorf("faield to get value from etcd, and prefix: %s", prefix)
	} else {
		// 将addr转换成resolver.Address
		addrs := make([]resolver.Address, 0)
		for _, kv := range resp.Kvs {
			addrs = append(addrs, resolver.Address{Addr: string(kv.Value)})
		}
		rv.resolverClientConn.UpdateState(resolver.State{Addresses: addrs})
		rv.startWatchRevision = resp.Header.Revision + 1

		go rv.watch(prefix, addrs)
	}

	return rv, nil
}

func (rv *Resolver) Scheme() string {
	return rv.schema
}

func (rv *Resolver) ResolveNow(rn resolver.ResolveNowOptions) {
	return
}

// Close closes the resolver.
func (rv *Resolver) Close() {

}

func (rv *Resolver) watch(prefix string, addrs []resolver.Address) {

	var exists func([]resolver.Address, string) bool
	exists = func(addresses []resolver.Address, s string) bool {
		for _, v := range addresses {
			if v.Addr == s {
				return true
			}
		}
		return false
	}

	var remove func([]resolver.Address, string) ([]resolver.Address, bool)
	remove = func(addresses []resolver.Address, s string) ([]resolver.Address, bool) {
		for i, _ := range addresses {
			if addresses[i].Addr == s {
				// 删除第i个
				addresses[i] = addresses[len(addresses)-1]
				return addresses[:len(addresses)-1], true
			}
		}
		return nil, false
	}

	watchChan := rv.cli.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for resp := range watchChan {
		update := false
		for _, ev := range resp.Events {
			if ev.Type == mvccpb.PUT {
				if !exists(addrs, string(ev.Kv.Value)) {
					// 新加的
					update = true
					addrs = append(addrs, resolver.Address{Addr: string(ev.Kv.Value)})
				}
			} else if ev.Type == mvccpb.DELETE {
				index := strings.LastIndexAny(string(ev.Kv.Key), "/")
				if index < 0 {
					return
				}
				t := string(ev.Kv.Key)[index+1:]
				newAddrs, ok := remove(addrs, t)
				if ok {
					update = true
					addrs = newAddrs
				}
			}
		}
		if update {
			rv.resolverClientConn.UpdateState(resolver.State{Addresses: addrs})
		}
	}

}

var (
	resolvers = make(map[string]*Resolver)
	mu        sync.RWMutex
)

// GetClientConn 主要功能：查找ClientConn
// 1. 加锁读map，如果有，直接返回
// 2. 创建，加入到map，并返回
func GetClientConn(opt *config.GrpcDialOption) *grpc.ClientConn {
	key := opt.Etcd.Schema + "/" + opt.ServiceName

	mu.RLock()
	resolver, ok := resolvers[key]
	mu.RUnlock()
	if ok {
		return resolver.grpcClientConn
	}

	rv, err := newResolver(opt)
	if err != nil {
		fmt.Println("[xzetcd] failed to create resolver: ", err)
		return nil
	}

	mu.Lock()
	resolvers[key] = rv
	mu.Unlock()
	return rv.grpcClientConn
}

// 解释
func newResolver(opt *config.GrpcDialOption) (*Resolver, error) {
	var cli *clientv3.Client

	cli, err := clientv3.New(clientv3.Config{
		Endpoints: opt.Etcd.Endpoints,
		Username:  opt.Etcd.Username,
		Password:  opt.Etcd.Password,
	})
	if err != nil {
		fmt.Println("[xzetcd] failed to create cli client: ", err)
		return nil, err
	}

	rv := new(Resolver)
	rv.cli = cli
	rv.schema = opt.Etcd.Schema
	rv.serviceName = opt.ServiceName

	mu.Lock()
	resolver.Register(rv)

	return nil, nil
}
