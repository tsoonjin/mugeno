package main

import (
  "context"
  "github.com/tsoonjin/mugeno/service/rss"
  rssSvcGrpc "github.com/tsoonjin/mugeno/service/rss/grpc"
)
// userServiceController implements the gRPC UserServiceServer interface.
type userServiceController struct {
  userService mysvc.Service
}
// NewUserServiceController instantiates a new UserServiceServer.
func NewUserServiceController(userService mysvc.Service) rssSvcGrpc.UserServiceServer {
  return &userServiceController{
    userService: userService,
  }
}
// GetUsers calls the core service's GetUsers method and maps the result to a grpc service response.
func (ctlr *userServiceController) GetUsers(ctx context.Context, req *rssSvcGrpc.GetUsersRequest) (resp *rssSvcGrpc.GetUsersResponse, err error) {
  resultMap, err := ctlr.userService.GetUsers(req.GetIds())
  if err != nil {
    return
  }
  resp := &rssSvcGrpc.GetUsersResponse{}
  for _, u := range resultMap {
    resp.Users = append(resp.Users, marshalUser(&u))
  }
  return
}
// marshalUser marshals a business object User into a gRPC layer User.
func marshalUser(u *mysvc.User) *rssSvcGrpc.User {
  return &rssSvcGrpc.User{Id: u.ID, Name: u.Name}
}
