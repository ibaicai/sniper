package userserver1

import (
	"context"

	pb "sniper/rpc/user/v1"
)

// Server 实现 /twirp/user.v1.User 服务
type Server struct{}

// Echo 实现 /twirp/user.v1.User/Echo 接口
func (s *Server) Echo(ctx context.Context, req *pb.EchoReq) (resp *pb.EchoResp, err error) {
	return &pb.EchoResp{Msg: req.Msg}, nil
}
