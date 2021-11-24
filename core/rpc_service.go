package core

import "gebi/rpc"

func InitRPCService() {
	go rpc.RpcService()
}
