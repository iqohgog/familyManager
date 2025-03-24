package task

import (
	"net/http"
	"v1/familyManager/configs"
	"v1/familyManager/internal/family"
	"v1/familyManager/internal/user"
	"v1/familyManager/pkg/middleware"
	"v1/familyManager/pkg/req"
	"v1/familyManager/pkg/res"
)

type TaskHandlerDeps struct {
	FamilyRepository *family.FamilyRepository
	UserRepository   *user.UserRepository
	TaskRepository   *TaskRepository
	Config           *configs.Config
}

type TaskHandler struct {
	FamilyRepository *family.FamilyRepository
	UserRepository   *user.UserRepository
	TaskRepository   *TaskRepository
}

func NewTaskHandler(router *http.ServeMux, deps TaskHandlerDeps) {
	handler := &TaskHandler{
		FamilyRepository: deps.FamilyRepository,
		UserRepository:   deps.UserRepository,
		TaskRepository:   deps.TaskRepository,
	}
	router.Handle("POST /task/create", middleware.IsAuthed(handler.CreateTask(), deps.Config))
}

func (handler *TaskHandler) CreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[TaskCreateRequest](w, r)
		if err != nil {
			return
		}
		email := r.Context().Value(middleware.ContextEmailKey).(string)
		user, err := handler.UserRepository.GetByEmail(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if user.FamilyID == nil {
			http.Error(w, "You are not in a family", http.StatusBadRequest)
			return
		}
		familyId, _ := user.FamilyID.(int64)
		_, err = handler.TaskRepository.Create(&Task{
			Name:        body.Name,
			Description: body.Description,
			AssigneeID:  body.AssigneeID,
			FamilyID:    string(familyId),
			CreatorID:   user.ID,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, nil, http.StatusCreated)
	}
}


