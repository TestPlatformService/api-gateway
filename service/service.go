package service

import (
	"api/config"
	"api/genproto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManager interface {
	UserService() user.UsersClient
}

type serviceManagerImpl struct {
	userClient    user.UsersClient
}

func (s *serviceManagerImpl) UserService() user.UsersClient {
	return s.userClient
}


func NewServiceManager() (ServiceManager, error) {
	connUser, err := grpc.Dial(
		config.Load().USER_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}


	return &serviceManagerImpl{
		userClient:         user.NewUsersClient(connUser),
	}, nil
}