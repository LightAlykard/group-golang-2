package models

import "database/sql"

type MessegeItem struct {
	ID   int    `jsosn: "id"`
	Name string `jsosn: "header"`
	Text string `jsosn: "text"`
}

type MessegeItemSlice []MessegeItem

func GetAllPostItems(db *sql.DB) (MessegeItemSlice, error) {
	row, err := db.Query("SELECT ID, Name, Text FROM Messeges")
	if err != nil {
		return nil, err
	}

	Messeges := make(MessegeItemSlice, 0, 10)
	for row.Next() {
		Messege := MessegeItem{}
		if err := row.Scan(&Messege.ID, &Messege.Name, &Messege.Text); err != nil {
			return nil, err
		}
		Messeges = append(Messeges, Messege)
	}
	return Messeges, nil
}

func GetPostItem(db *sql.DB, id int) (MessegeItemSlice, error) {
	row, err := db.Query("SELECT ID, Name, Text FROM Messeges WHERE ID = ?", id)
	if err != nil {
		return nil, err
	}

	Messeges := make(MessegeItemSlice, 0, 10)
	for row.Next() {
		Messege := MessegeItem{}
		if err := row.Scan(&Messege.ID, &Messege.Name, &Messege.Text); err != nil {
			return nil, err
		}
		Messeges = append(Messeges, Messege)
	}
	return Messeges, nil
}

func (Messege *MessegeItem) Insert(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO Messeges (Name, Text) VALUES (?, ?)", Messege.Name, Messege.Text)
	return err
}
func (Messege *MessegeItem) Update(db *sql.DB) error {
	_, err := db.Exec("UPDATE Messeges SET Name=?, Text=? WHERE ID = ?", Messege.Name, Messege.Text, Messege.ID)
	return err
}
func (Messege *MessegeItem) Delete(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM Messeges WHERE ID = ?", Messege.ID)
	return err
}
