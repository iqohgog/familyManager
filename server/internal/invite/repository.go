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

func (repo *FamilyInviteRepository) GetByID(invitedId string) (*[]FamilyInvites, error) {
	stmt, err := repo.Storage.DB.Prepare(`
		SELECT f.name, fi.invited_user_id, fi.status 
		FROM family_invitations fi
		JOIN families f ON fi.family_id = f.id
		WHERE fi.invited_user_id = $1
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(invitedId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invites []FamilyInvites
	for rows.Next() {
		var invite FamilyInvites
		err = rows.Scan(&invite.FamilyName, &invite.InventedID, &invite.Status)
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
	if invite.Status == "accepted" {
		stmt, err = repo.Storage.DB.Prepare(`
			UPDATE users
			SET family_id = $1
			WHERE id = $2
		`)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(invite.FamilyID, invite.InventedID)
		if err != nil {
			return err
		}
		stmt, err = repo.Storage.DB.Prepare(`
		UPDATE family_invitations
		SET status = 'declined', updated_at = CURRENT_TIMESTAMP
		WHERE family_id != $1 AND invited_user_id = $2
		`)
		if err != nil {
			return err
		}
		_, err = stmt.Exec(invite.FamilyID, invite.InventedID)
		if err != nil {
			return err
		}
	}
	return nil
}
