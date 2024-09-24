package handler

import (
	"api/genproto/group"
	"api/genproto/notification"
	"api/genproto/user"
	"log/slog"

	"github.com/casbin/casbin/v2"
)

type Handler struct {
	User         user.UsersClient
	Notification notification.NotificationsClient
	Group        group.GroupServiceClient
	Log          *slog.Logger
	Enforcer     *casbin.Enforcer
}
