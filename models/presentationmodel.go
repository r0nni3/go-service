package models

import (
	"database/sql"
	"time"
)

type Presentation struct {
	Id      int64          `json:"id"`
	UserId  sql.NullInt64  `json:"user_id"`
	Title   string         `json:"title"`
	Created sql.NullString `json:"created_at"`
	Deleted sql.NullString `json:"deleted_at"`
	Updated sql.NullString `json:"updated_at"`
	Active  bool           `json:"active"`
}

func (conObj ConnectionObj) GetPresentation(id int64) (*Presentation, error) {
	con, err := conObj.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	query := "SELECT * FROM Presentation"
	query += " WHERE id=?"
	query += " and active=TRUE"
	query += " and deleted_at IS NULL"
	row := con.QueryRow(query, id)

	presentation := new(Presentation)
	err = row.Scan(&presentation.Id, &presentation.UserId,
		&presentation.Title, &presentation.Created, &presentation.Updated,
		&presentation.Deleted, &presentation.Active)
	if err != nil {
		return nil, err
	}

	return presentation, nil
}

func (conObj ConnectionObj) GetPresentations() ([]*Presentation, error) {
	con, err := conObj.OpenConnection()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	query := "SELECT * FROM Presentation"
	query += " WHERE active=TRUE and deleted_at IS NULL"
	query += " ORDER BY id ASC"
	rows, err := con.Query(query)
	if err != nil {
		return nil, err
	}

	presentations := make([]*Presentation, 0, 10)
	for rows.Next() {
		presentation := new(Presentation)
		err := rows.Scan(&presentation.Id, &presentation.UserId,
			&presentation.Title, &presentation.Created,
			&presentation.Updated, &presentation.Deleted,
			&presentation.Active)
		if err != nil {
			return nil, err
		}
		presentations = append(presentations, presentation)
	}

	return presentations, nil
}

func (conObj ConnectionObj) InsertPresentation(p *Presentation) (int64, error) {
	con, err := conObj.OpenConnection()
	if err != nil {
		return -1, err
	}
	defer con.Close()

	queryStr := "INSERT INTO Presentation (title) VALUES (?)"
	result, err := con.Exec(queryStr, p.Title)
	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func (conObj ConnectionObj) UpdatePresentation(p *Presentation) error {
	con, err := conObj.OpenConnection()
	if err != nil {
		return err
	}
	defer con.Close()

	queryStr := "UPDATE Presentation "
	queryStr += "SET title=? "
	queryStr += ", updated_at=?"
	queryStr += " WHERE id=?"
	_, err = con.Exec(queryStr, p.Title, time.Now(), p.Id)
	if err != nil {
		return err
	}

	return nil
}

func (conObj ConnectionObj) DeletePresentation(id int64) error {
	con, err := conObj.OpenConnection()
	if err != nil {
		return err
	}
	defer con.Close()

	queryStr := "UPDATE Presentation "
	queryStr += "SET deleted_at=?"
	queryStr += " WHERE id=?"
	_, err = con.Exec(queryStr, time.Now(), id)

	if err != nil {
		return err
	}

	return nil
}
