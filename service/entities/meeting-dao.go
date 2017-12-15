package entities

import "strings"

type meetingDao DaoSource

func (dao *meetingDao)Insert(meeting *Meeting) error {
	var meetingInsertStmt = "insert into meetings(title, sponsor, participators, start, end) values(?, ?, ?, ?, ?)"

	stmt, err := dao.Prepare(meetingInsertStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	var participants = strings.Join(meeting.Participators, "&")
	res, err := stmt.Exec(meeting.Title, meeting.Sponsor, participants, meeting.Start, meeting.End)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	meeting.Id = int(id)
	return nil
}

func (dao *meetingDao)QueryAll() []Meeting {
	var meetingQueryAll = "select * from meetings"
	rows, err := dao.Query(meetingQueryAll)
	checkErr(err)
	defer rows.Close()

	mlist := make([]Meeting, 0, 0)
	for rows.Next() {
		meeting := Meeting{}
		participants := ""
		err := rows.Scan(&meeting.Title, &meeting.Sponsor, &participants, &meeting.Start, &meeting.End)
		meeting.Participators = strings.Split(participants, "&")
		checkErr(err)
		mlist = append(mlist, meeting)
	}
	return mlist
}

func (dao *meetingDao)QueryByTitle(title string) Meeting {
	meetingQueryByTitle := "select * from meetings where title=" + title
	rows, err := dao.Query(meetingQueryByTitle)
	defer rows.Close()
	checkErr(err)
	meeting := Meeting{}
	participants := ""
	if rows.Next() {
		err = rows.Scan(&meeting.Title, &meeting.Sponsor, &participants, &meeting.Start, &meeting.End)
	}
	checkErr(err)

	return meeting
}

func (dao *meetingDao)QueryBy(key string, val string) []Meeting {
	meetingQueryBy := "select * from meetings where " + key + " = " + val
	rows, err := dao.Query(meetingQueryBy)
	defer rows.Close()
	checkErr(err)
	mlist := make([]Meeting, 0, 0)
	for rows.Next() {
		meeting := Meeting{}
		participants := ""
		err := rows.Scan(&meeting.Title, &meeting.Sponsor, &participants, &meeting.Start, &meeting.End)
		meeting.Participators = strings.Split(participants, "&")
		checkErr(err)
		mlist = append(mlist, meeting)
	}
	return mlist
}

func (dao *meetingDao)DeleteByTitle(title string) error {
	meetingDeleteByTitle := "delete from meeting where title = " + title
	_, err := dao.Exec(meetingDeleteByTitle)
	if err != nil {
		return err
	}
	return nil
}

func (dao *meetingDao)DeleteBySponsor(sponsor string) error {
	meetingDeleteBySponsor := "delete from meeting where sponsor = " + sponsor
	_, err := dao.Exec(meetingDeleteBySponsor)
	if err != nil {
		return err
	}
	return nil
}

func (dao *meetingDao)Update(meeting Meeting) error {
	var participants = strings.Join(meeting.Participators, "&")
	meetingUpdateParticipators := "update meetings set participators=? where title=?"

	stmt, err := dao.Prepare(meetingUpdateParticipators)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(participants, meeting.Title)
	if err != nil {
		return err
	}
	return nil
}