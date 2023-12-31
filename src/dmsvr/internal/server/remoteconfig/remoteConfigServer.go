// Code generated by goctl. DO NOT EDIT.
// Source: dm.proto

package server

import (
	"context"

	"github.com/i-Things/things/src/dmsvr/internal/logic/remoteconfig"
	"github.com/i-Things/things/src/dmsvr/internal/svc"
	"github.com/i-Things/things/src/dmsvr/pb/dm"
)

type RemoteConfigServer struct {
	svcCtx *svc.ServiceContext
	dm.UnimplementedRemoteConfigServer
}

func NewRemoteConfigServer(svcCtx *svc.ServiceContext) *RemoteConfigServer {
	return &RemoteConfigServer{
		svcCtx: svcCtx,
	}
}

func (s *RemoteConfigServer) RemoteConfigCreate(ctx context.Context, in *dm.RemoteConfigCreateReq) (*dm.Response, error) {
	l := remoteconfiglogic.NewRemoteConfigCreateLogic(ctx, s.svcCtx)
	return l.RemoteConfigCreate(in)
}

func (s *RemoteConfigServer) RemoteConfigIndex(ctx context.Context, in *dm.RemoteConfigIndexReq) (*dm.RemoteConfigIndexResp, error) {
	l := remoteconfiglogic.NewRemoteConfigIndexLogic(ctx, s.svcCtx)
	return l.RemoteConfigIndex(in)
}

func (s *RemoteConfigServer) RemoteConfigPushAll(ctx context.Context, in *dm.RemoteConfigPushAllReq) (*dm.Response, error) {
	l := remoteconfiglogic.NewRemoteConfigPushAllLogic(ctx, s.svcCtx)
	return l.RemoteConfigPushAll(in)
}

func (s *RemoteConfigServer) RemoteConfigLastRead(ctx context.Context, in *dm.RemoteConfigLastReadReq) (*dm.RemoteConfigLastReadResp, error) {
	l := remoteconfiglogic.NewRemoteConfigLastReadLogic(ctx, s.svcCtx)
	return l.RemoteConfigLastRead(in)
}
