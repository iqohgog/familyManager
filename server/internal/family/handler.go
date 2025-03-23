package family

import (
	"net/http"
	"v1/familyManager/configs"
	"v1/familyManager/internal/invite"
	"v1/familyManager/internal/user"
	"v1/familyManager/pkg/middleware"
	"v1/familyManager/pkg/req"
	"v1/familyManager/pkg/res"
)

type FamilyHandlerDeps struct {
	FamilyRepository       *FamilyRepository
	FamilyInviteRepository *invite.FamilyInviteRepository
	UserRepository         *user.UserRepository
	Config                 *configs.Config
}

type FamilyHandler struct {
	FamilyRepository       *FamilyRepository
	FamilyInviteRepository *invite.FamilyInviteRepository
	UserRepository         *user.UserRepository
}

func NewFamilyHandler(router *http.ServeMux, deps FamilyHandlerDeps) {
	handler := &FamilyHandler{
		FamilyRepository:       deps.FamilyRepository,
		FamilyInviteRepository: deps.FamilyInviteRepository,
		UserRepository:         deps.UserRepository,
	}
	router.Handle("POST /family/create", middleware.IsAuthed(handler.CreateFamily(), deps.Config))
}

func (handler *FamilyHandler) CreateFamily() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[FamilyCreateRequest](w, r)
		if err != nil {
			return
		}
		email := r.Context().Value(middleware.ContextEmailKey).(string)
		user, err := handler.UserRepository.GetByEmail(email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if user.FamilyID != nil {
			http.Error(w, "You already in a family", http.StatusBadRequest)
			return
		}
		createdFamily, err := handler.FamilyRepository.Create(&Family{
			Name:      body.Name,
			CreatorID: user.ID,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdFamily, 201)
	}
}

/*

func NewLinkHandler(router *http.ServeMux, deps LinkHandlerDeps) {
	handler := &LinkHandler{LinkRepository: deps.LinkRepository, EventBus: deps.EventBus}
	router.Handle("POST /link", middleware.IsAuthed(handler.Create(), deps.Config))
	router.Handle("PATCH /link/{id}", middleware.IsAuthed(handler.Update(), deps.Config))
	router.Handle("DELETE /link/{id}", middleware.IsAuthed(handler.Delete(), deps.Config))
	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.Handle("GET /link", middleware.IsAuthed(handler.GetAll(), deps.Config))

}

func (handler *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LinkCreateRequest](&w, r)
		if err != nil {
			return
		}
		link := NewLink(body.Url)
		for {
			if link, _ := handler.LinkRepository.GetByHash(link.Hash); link != nil {
				link.GenereateHash()
				continue
			}
			break
		}
		createdLink, err := handler.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, createdLink, 201)
	}
}

func (handler *LinkHandler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := handler.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		go handler.EventBus.Publish(event.Event{
			Type: event.EventLinkVisited,
			Data: link.ID,
		})
		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)

	}
}

func (handler *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email, ok := r.Context().Value(middleware.ContextEmailKey).(string)
		if ok {
			fmt.Println(email)
		}
		body, err := req.HandleBody[LinkUpdateRequest](&w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := handler.LinkRepository.Update(&Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		res.Json(w, link, 201)
	}
}
func (handler *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err = handler.LinkRepository.GetById(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		err = handler.LinkRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Json(w, nil, 200)

	}
}

func (handler *LinkHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, "Invalid Limit", http.StatusBadRequest)
			return
		}
		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, "Invalid offset", http.StatusBadRequest)
			return
		}
		links := handler.LinkRepository.GetALl(limit, offset)
		count := handler.LinkRepository.Count()
		res.Json(w, GetAllLinksResponse{
			Links: links,
			Count: count,
		}, 200)

	}
}


*/
