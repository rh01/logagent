package main

import (
	etcd_client "github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
	"github.com/astaxie/beego/logs"
	"context"
	"strings"
	"encoding/json"
)

type EtcdClient struct {
	client *etcd_client.Client
}

var (
	etcdClient *EtcdClient
)

func initEtcd(addr string, key string) (collectConf []CollectConf,err error) {

	// create client using clientv3.New()
	// refer: https://github.com/coreos/etcd/tree/master/clientv3
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		logs.Error("connect failed.. err: %v", err)
		return
	}

	etcdClient = &EtcdClient{
		client: cli,
	}



	if strings.HasSuffix(key, "/") == false {
		key += "/"
	}

	//var collectConf []CollectConf
	for _, ip := range localIPArray {

		etcdKey := fmt.Sprintf("%s%s", key, ip)
		// 设置超时时间
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		response, err := cli.Get(ctx, etcdKey)

		if err != nil {
			logs.Error("client get from etcd failed, err: %v", err)
			continue
		}

		cancel()
		logs.Debug("resp from etcd, resp: %v", response.Kvs)
		for _, v := range response.Kvs {
			if string(v.Key) == etcdKey {
				err = json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("unmarshal failed, err: %v", err)
					continue

				}
				logs.Debug("log config is %v", collectConf )
			}
		}
	}


	return
}
