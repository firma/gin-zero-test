package etcd

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
)

var serviceTarget = "user.rpc"

type remoteService struct {
	name  string
	nodes map[string]string
	mutex sync.Mutex
}

// 获取服务目录下所有key初始化到服务的可用节点列表中
func GetService(etcdClient *clientv3.Client) *remoteService {
	service := &remoteService{
		name: serviceTarget,
	}
	kv := clientv3.NewKV(etcdClient)
	rangeResp, err := kv.Get(context.TODO(), service.name, clientv3.WithPrefix())
	if err != nil {
		panic(err)
	}

	service.mutex.Lock()
	for _, kv := range rangeResp.Kvs {
		service.nodes[string(kv.Key)] = string(kv.Value)
	}
	service.mutex.Unlock()

	go WatchServiceUpdate(etcdClient, service)
	return service
}

// 监控服务目录下的事件
func WatchServiceUpdate(etcdClient *clientv3.Client, service *remoteService) {
	watcher := clientv3.NewWatcher(etcdClient)
	// Watch 服务目录下的更新
	watchChan := watcher.Watch(context.TODO(), service.name, clientv3.WithPrefix())
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			service.mutex.Lock()
			switch event.Type {
			case mvccpb.PUT: //PUT事件，目录下有了新key
				service.nodes[string(event.Kv.Key)] = string(event.Kv.Value)
			case mvccpb.DELETE: //DELETE事件，目录中有key被删掉(Lease过期，key 也会被删掉)
				delete(service.nodes, string(event.Kv.Key))
			}
			service.mutex.Unlock()
		}
	}
}
