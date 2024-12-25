package main

import (
	"context"

	pb "github.com/HouLPcode/etcd_sync"
)

type syncServiceImpl struct {
	pb.UnimplementedSyncService
}

// Debug 执行切换操作
func (s *syncServiceImpl) Debug(
	ctx context.Context,
	req *pb.SyncRequest,
) (*pb.SyncResponse, error) {
	rsp := &pb.SyncResponse{}
	return rsp, nil
}
