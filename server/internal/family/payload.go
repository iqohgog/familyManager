package family

type FamilyInvite struct {
	Email string `json:"email"`
}

type FamilyCreateRequest struct {
	Name string `json:"name"`
}

type FamilyInviteRequest struct {
	Email string `json:"email"`
}

type FamilyInviteAnswerRequest struct {
	FamilyID string `json:"family_id"`
}


type FamilyInvitesResponse struct {
	Invites []FamilyInvite `json:"invites"`
}
