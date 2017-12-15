package entities

type MeetingServiceProvider struct { }

var MeetingService = MeetingServiceProvider {}

var meetingInsertStmt = "INSERT meetings SET title=?,sponsor=?,participants=?,start=?,end=?"
var meetingQueryAll = "SELECT * FROM meetings"
var meetingQueryByTitle = "SELECT * FROM meetings where title = ?"

func (*MeetingServiceProvider)Insert(meeting *Meeting) error {
	tx, err := db.Begin()
	checkErr(err)

	dao := meetingDao{ tx }
	err := dao.Insert(meeting)
	if err != nil {
		return err
	}
	checkErr(err)
	tx.Commit()
	return nil
}

func (*MeetingServiceProvider)QueryAll() []Meeting {
	dao := meetingDao{ db }
	meetings := dao.QueryAll()
	return meetings
}

func (*MeetingServiceProvider)QueryByTitle(title string) Meeting {
	dao := meetingDao{ db }
	meeting := dao.QueryByTitle(title)
	return meeting
}

func (*MeetingServiceProvider)QueryBy(key string, val string) []Meeting {
	dao := meetingDao{ db }
	meetings := dao.QueryBy(key, val)
	return meetings
}



func (*MeetingServiceProvider)DeleteByTitle(title string) error {
	dao := meetingDao{ db }
	err := dao.DeleteByTitle(title)
	checkErr(err)
	return nil
}

func (*MeetingServiceProvider)DeleteBySponsor(sponsor string) error {
	dao := meetingDao{ db }
	err := dao.DeleteBySponsor(sponsor)
	checkErr(err)
	return nil
}

func (*MeetingServiceProvider)Update(meeting Meeting) error {
	tx, err := db.Begin()
	checkErr(err)

	dao := meetingDao{ tx }
	err := dao.Update(meeting)
	if err != nil {
		return err
	}
	checkErr(err)
	tx.Commit()
	return nil
}