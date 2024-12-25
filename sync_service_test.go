package main

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/HouLPcode/etcd_sync"
	"go.uber.org/mock/gomock"
	_ "trpc.group/trpc-go/trpc-go/http"
)

//go:generate go mod tidy
//go:generate mockgen -destination=stub/github.com/HouLPcode/etcd_sync/sync_mock.go -package=etcd_sync -self_package=github.com/HouLPcode/etcd_sync --source=stub/github.com/HouLPcode/etcd_sync/sync.trpc.go

func Test_syncServiceImpl_Debug(t *testing.T) {
	// Start writing mock logic.
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	syncServiceService := pb.NewMockSyncServiceService(ctrl)
	var inorderClient []*gomock.Call
	// Expected behavior.
	m := syncServiceService.EXPECT().Debug(gomock.Any(), gomock.Any()).AnyTimes()
	m.DoAndReturn(func(ctx context.Context, req *pb.SyncRequest) (*pb.SyncResponse, error) {
		s := &syncServiceImpl{}
		return s.Debug(ctx, req)
	})
	gomock.InOrder(inorderClient...)

	// Start writing unit test logic.
	type args struct {
		ctx context.Context
		req *pb.SyncRequest
		rsp *pb.SyncResponse
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var rsp *pb.SyncResponse
			var err error
			if rsp, err = syncServiceService.Debug(tt.args.ctx, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("syncServiceImpl.Debug() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(rsp, tt.args.rsp) {
				t.Errorf("syncServiceImpl.Debug() rsp got = %v, want %v", rsp, tt.args.rsp)
			}
		})
	}
}
