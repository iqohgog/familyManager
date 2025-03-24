package task

import (
	"net/http"
	"strconv"
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
	router.Handle("POST /task", middleware.IsAuthed(handler.CreateTask(), deps.Config))
	router.Handle("GET /task", middleware.IsAuthed(handler.Tasks(), deps.Config))
	router.Handle("DELETE /task/{id}", middleware.IsAuthed(handler.DeleteTask(), deps.Config))
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
		familyID := strconv.Itoa(int(familyId))
		_, err = handler.TaskRepository.Create(&Task{
			Name:        body.Name,
			Description: body.Description,
			AssigneeID:  body.AssigneeID,
			Priority:    body.Priority,
			FamilyID:    familyID,
			CreatorID:   user.ID,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, nil, http.StatusCreated)
	}
}

func (handler *TaskHandler) Tasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		tasks, err := handler.TaskRepository.GetTaskByFamilyID(familyId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, tasks, http.StatusOK)
	}
}

func (handler *TaskHandler) DeleteTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
		taskID := r.PathValue("id")
		err = handler.TaskRepository.DeleteTaskByID(taskID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, nil, http.StatusOK)
	}
}
