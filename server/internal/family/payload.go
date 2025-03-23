package family

type FamilyInvite struct {
	FamilyName string `json:"family_name"`
	Status     string `json:"status"`
}

type FamilyCreateRequest struct {
	Name string `json:"name"`
}

type FamilyInviteRequest struct {
	Email string `json:"email"`
}

type FamilyInviteAnswerRequest struct {
	Answer   bool   `json:"answer"`
	FamilyID string `json:"family_id"`
}

type FamilyInviteAnswerResponse struct {
	Answer bool `json:"answer"`
}

type FamiliInvitesRequest struct {
	Email string `json:"email"`
}

type FamilyInvitesResponse struct {
	Invites []FamilyInvite `json:"invites"`
}
