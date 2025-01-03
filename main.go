package main

import (
	pb "github.com/HouLPcode/etcd_sync"
	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
	trpc "trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/log"
)

func main() {
	s := trpc.NewServer()
	pb.RegisterSyncServiceService(s.Service("etcd_sync.SyncService"), &syncServiceImpl{})
	if err := s.Serve(); err != nil {
		log.Fatal(err)
	}
}
