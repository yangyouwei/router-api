package models

import (
	db "github.com/yangyouwei/router-api/databases"
	"log"
)

//定义person类型结构
type Line struct {
	Id      int    `json:"rowid"`
	Ipaddr  string `json:"ipaddr"`
	Port    int    `json:"port"`
	Comment string `json:"comment"`
}

func (l *Line) AddLines() (id int64, err error) {
	rs, err := db.SqlDB.Exec("INSERT INTO lines(ipaddr, port, comment) VALUES (?, ?, ?)", l.Ipaddr, l.Port,l.Comment)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}

func (l *Line) GetLines() (lines []Line, err error) {
	lines = make([]Line, 0)
	rows, err := db.SqlDB.Query("SELECT rowid, * FROM lines")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var line Line
		rows.Scan(&line.Id, &line.Ipaddr, &line.Port,&line.Comment)
		lines = append(lines, line)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

func (l *Line) GetLine() (lines Line, err error) {
	err = db.SqlDB.QueryRow("SELECT rowid, * FROM lines WHERE rowid=?", l.Id).Scan(&lines.Id, &lines.Ipaddr, &lines.Port,&lines.Comment )
	return
}

func (l *Line) ModLine() (ra int64, err error) {
	stmt, err := db.SqlDB.Prepare("UPDATE lines SET ipaddr=?, port=? , comment=? WHERE rowid=?")
	defer stmt.Close()
	if err != nil {
		return
	}
	rs, err := stmt.Exec(l.Ipaddr, l.Port, l.Comment,l.Id)
	if err != nil {
		return
	}
	ra, err = rs.RowsAffected()
	return
}

func (l *Line) DelLine() (ra int64, err error) {
	rs, err := db.SqlDB.Exec("DELETE FROM lines WHERE rowid=?", l.Id)
	if err != nil {
		log.Fatalln(err)
	}
	ra, err = rs.RowsAffected()
	return
}
