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
		{"admin", "/api/user/photo", "POST"},

		{"student", "/api/user/getprofile", "GET"},
		{"student", "/api/user/updateprofile", "PUT"},
		{"student", "/api/user/photo", "POST"},
		{"student", "/api/user/photo", "DELETE"},

		{"teacher", "/api/user/getprofile", "GET"},
		{"teacher", "/api/user/updateprofile", "PUT"},
		{"teacher", "/api/user/photo", "POST"},

		{"support", "/api/user/getprofile", "GET"},
		{"support", "/api/user/updateprofile", "PUT"},
		{"support", "/api/user/photo", "POST"},

		//group
		{"admin", "/api/groups/create", "POST"},
		{"admin", "/api/groups/update", "PUT"},
		{"admin", "/api/groups/delete", "DELETE"},
		{"admin", "/api/groups/getById/:group_id", "GET"},
		{"admin", "/api/groups/getAll", "GET"},
		{"admin", "/api/groups/add-student", "POST"},
		{"admin", "/api/groups/delete-student", "DELETE"},
		{"admin", "/api/groups/add-teacher", "POST"},
		{"admin", "/api/groups/delete-teacher", "DELETE"},
		{"admin", "/api/groups/student-groups/:hh_id", "GET"},
		{"admin", "/api/groups/teacher-groups/:id", "GET"},
		{"admin", "/api/groups/students/:group_id", "GET"},
		{"student", "/api/groups/student-groups/:hh_id", "GET"},
		{"teacher", "/api/groups/teacher-groups/:id", "GET"},


		//topic
		{"admin", "/api/topics/create", "POST"},
		{"admin", "/api/topics/update", "PUT"},
		{"admin", "/api/topics/delete/:topic_id", "DELETE"},
		{"admin", "/api/topics/getAll", "GET"},

		{"teacher", "/api/topics/create", "POST"},
		{"teacher", "/api/topics/update", "PUT"},

		{"teacher", "/api/topics/delete/:topic_id", "DELETE"},
		{"teacher", "/api/topics/getAll", "GET"},

		{"student", "/api/topics/getAll", "GET"},

		//subject
		{"admin", "/api/subjects/create", "POST"},
		{"admin", "/api/subjects/get/:id", "GET"},
		{"admin", "/api/subjects/getall", "GET"},
		{"admin", "/api/subjects/update/:id", "PUT"},
		{"admin", "/api/subjects/delete/:id", "DELETE"},

		{"student", "/api/subjects/get/:id", "GET"},
		{"student", "/api/subjects/getall", "GET"},

		//question
		{"admin", "/api/questions/create", "POST"},
		{"admin", "/api/questions/:id", "GET"},
		{"admin", "/api/questions/update/:id", "PUT"},
		{"admin", "/api/questions/delete/:id", "DELETE"},
		{"admin", "/api/questions/getAll", "GET"},
		{"admin", "/api/questions/upload-image/:id", "POST"},
		{"admin", "/api/questions/delete-image/:id", "DELETE"},

		{"teacher", "/api/questions/create", "POST"},
		{"teacher", "/api/questions/:id", "GET"},
		{"teacher", "/api/questions/update/:id", "PUT"},
		{"teacher", "/api/questions/delete/:id", "DELETE"},
		{"teacher", "/api/questions/getAll", "GET"},
		{"teacher", "/api/questions/upload-image/:id", "POST"},
		{"teacher", "/api/questions/delete-image/:id", "DELETE"},

		//question output
		{"admin", "/api/question-outputs/create", "POST"},
		{"admin", "/api/question-outputs/:id", "GET"},
		{"admin", "/api/question-outputs/question/:question_id", "GET"},
		{"admin", "/api/question-outputs/delete/:id", "DELETE"},

		{"teacher", "/api/question-outputs/create", "POST"},
		{"teacher", "/api/question-outputs/:id", "GET"},
		{"teacher", "/api/question-outputs/question/:question_id", "GET"},
		{"teacher", "/api/question-outputs/delete/:id", "DELETE"},

		//question input
		{"admin", "/api/question-inputs/create", "POST"},
		{"admin", "/api/question-inputs/:id", "GET"},
		{"admin", "/api/question-inputs/question/:question_id", "GET"},
		{"admin", "/api/question-inputs/delete/:id", "DELETE"},

		{"teacher", "/api/question-inputs/create", "POST"},
		{"teacher", "/api/question-inputs/:id", "GET"},
		{"teacher", "/api/question-inputs/question/:question_id", "GET"},
		{"teacher", "/api/question-inputs/delete/:id", "DELETE"},

		//test case
		{"admin", "/api/test-cases/create", "POST"},
		{"admin", "/api/test-cases/:id", "GET"},
		{"admin", "/api/test-cases/question/:question_id", "GET"},
		{"admin", "/api/test-cases/delete/:id", "DELETE"},

		{"teacher", "/api/test-cases/create", "POST"},
		{"teacher", "/api/test-cases/:id", "GET"},
		{"teacher", "/api/test-cases/question/:question_id", "GET"},
		{"teacher", "/api/test-cases/delete/:id", "DELETE"},

		//task
		{"teacher", "/api/task/create", "POST"},
		{"teacher", "/api/task/delete", "DELETE"},
		{"teacher", "api/task/get", "GET"},
		{"student", "api/task/get", "GET"},

		// admin
		{"admin", "/api/task/create", "POST"},
		{"admin", "/api/task/delete", "DELETE"},
		{"admin", "api/task/get", "GET"},
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
