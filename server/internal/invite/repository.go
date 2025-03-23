package invite

import (
	"database/sql"
	"v1/familyManager/pkg/db"
)

type FamilyInviteRepository struct {
	Storage *db.Storage
}

func NewFamilyInviteRepository(storage *db.Storage) *FamilyInviteRepository {
	return &FamilyInviteRepository{
		Storage: storage,
	}
}

func (repo *FamilyInviteRepository) Create(invite *FamilyInvite) (*FamilyInvite, error) {
	stmt, err := repo.Storage.DB.Prepare(`
	INSERT INTO family_invitations(
		family_id, invited_user_id
	)
	VALUES($1, $2)
	`)
	if err != nil {
		return nil, err
	}
	err = stmt.QueryRow(invite.FamilyID, invite.InventedID).Scan()
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	invite.Status = "pending"
	return invite, nil
}

func (repo *FamilyInviteRepository) GetByID(inventedId string) (*[]FamilyInvite, error) {
	stmt, err := repo.Storage.DB.Prepare(`
		SELECT family_id, invited_user_id, status FROM family_invitations
		WHERE invited_user_id = $1
	`)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(inventedId)
	if err != nil {
		return nil, err
	}
	var invites []FamilyInvite
	for rows.Next() {
		var invite FamilyInvite
		err = rows.Scan(&invite.FamilyID, &invite.InventedID, &invite.Status)
		if err != nil {
			return nil, err
		}
		invites = append(invites, invite)
	}
	return &invites, nil
}

func (repo *FamilyInviteRepository) UpdateStatus(invite *FamilyInvite) error {
	stmt, err := repo.Storage.DB.Prepare(`
		UPDATE family_invitations
		SET status = $1, updated_at = CURRENT_TIMESTAMP
		WHERE family_id = $2 AND invited_user_id = $3
	`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(invite.Status, invite.FamilyID, invite.InventedID)
	if err != nil {
		return err
	}
	return nil
}
