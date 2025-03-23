package family

import (
	"database/sql"
	"v1/familyManager/pkg/db"
)

type FamilyRepository struct {
	Storage *db.Storage
}

func NewFamilyRepository(storage *db.Storage) *FamilyRepository {
	return &FamilyRepository{
		Storage: storage,
	}
}

func (repo *FamilyRepository) Create(family *Family) (*Family, error) {
	stmt, err := repo.Storage.DB.Prepare(`
	INSERT INTO families(
		name, creator_id
	)
	VALUES($1, $2)
	RETURNING id;
	`)
	if err != nil {
		return nil, err
	}
	var familyID int
	err = stmt.QueryRow(family.Name, family.CreatorID).Scan(&familyID)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	stmt, err = repo.Storage.DB.Prepare(`
	UPDATE users
	SET family_id = $1
	WHERE id = $2;
	`)
	if err != nil {
		return nil, err
	}
	_, err = stmt.Exec(familyID, family.CreatorID)
	if err != nil {
		return nil, err
	}
	return family, nil
}

func (repo *FamilyRepository) GetByID(id string) (*Family, error) {
	stmt, err := repo.Storage.DB.Prepare(`
		SELECT name, creator_id FROM families
		WHERE ID = $1
	`)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)
	var family Family
	err = row.Scan(&family.Name, &family.CreatorID)
	if err != nil {
		return nil, err
	}
	return &family, nil
}

func (repo *FamilyRepository) GetByCreatorID(id string) (*Family, error) {
	stmt, err := repo.Storage.DB.Prepare(`
		SELECT id, name FROM families
		WHERE creator_id = $1
	`)
	if err != nil {
		return nil, err
	}
	row := stmt.QueryRow(id)
	var family Family
	err = row.Scan(&family.ID, &family.Name)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &family, nil
}
