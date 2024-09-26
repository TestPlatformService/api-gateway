package casbin

import (
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
)

const (
	host     = "postgres-db-casbin"
	port     = "5432"
	dbname   = "casbin"
	username = "postgres"
	password = "1234"
)

func CasbinEnforcer(logger *slog.Logger) (*casbin.Enforcer, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, username, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Error connecting to database", "error", err.Error())
		return nil, err
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE IF EXISTS casbin")
	if err != nil {
		logger.Error("Error dropping Casbin database", "error", err.Error())
		return nil, err
	}

	adapter, err := xormadapter.NewAdapter("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbname, password))
	if err != nil {
		logger.Error("Error creating Casbin adapter", "error", err.Error())
		return nil, err
	}

	enforcer, err := casbin.NewEnforcer("casbin/model.conf", adapter)
	if err != nil {
		logger.Error("Error creating Casbin enforcer", "error", err.Error())
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		logger.Error("Error loading Casbin policy", "error", err.Error())
		return nil, err
	}

	policies := [][]string{
		//user
		{"admin", "/api/user/register", "POST"},
		{"admin", "/api/user/all", "GET"},
		{"admin", "/api/user/updateprofile", "PUT"},
		{"admin", "/api/user/update", "PUT"},

		{"student", "/api/user/getprofile", "GET"},
		{"student", "/api/user/updateprofile", "PUT"},
		{"student", "/api/user/photo", "POST"},
		{"student", "/api/user/photo", "DELETE"},

		{"teacher", "/api/user/getprofile", "GET"},
		{"teacher", "/api/user/updateprofile", "PUT"},

		{"support", "/api/user/getprofile", "GET"},
		{"support", "/api/user/updateprofile", "PUT"},

		//group
		{"admin", "/api/groups/create", "POST"},
		{"admin", "/api/groups/update", "PUT"},
		{"admin", "/api/groups/delete", "DELETE"},
		{"admin", "/api/groups/getById", "GET"},
		{"admin", "/api/groups/getAll", "GET"},
		{"admin", "/api/groups/add-student", "POST"},
		{"admin", "/api/groups/delete-student", "DELETE"},
		{"admin", "/api/groups/add-teacher", "POST"},
		{"admin", "/api/groups/delete-teacher", "DELETE"},
		{"admin", "/api/groups/student-groups", "GET"},
		{"admin", "/api/groups/teacher-groups", "GET"},
		{"admin", "/api/group-students", "GET"},
	}

	_, err = enforcer.AddPolicies(policies)
	if err != nil {
		logger.Error("Error adding Casbin policy", "error", err.Error())
		return nil, err
	}

	err = enforcer.SavePolicy()
	if err != nil {
		logger.Error("Error saving Casbin policy", "error", err.Error())
		return nil, err
	}

	return enforcer, nil
}
