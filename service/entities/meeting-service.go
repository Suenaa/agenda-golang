package entities

type MeetingServiceProvider struct { }

var MeetingService = MeetingServiceProvider {}

func (*MeetingServiceProvider)Insert(meeting *Meeting) error {
	tx, err := db.Begin()
	checkErr(err)
	if err != nil {
		return err
	}

	dao := meetingDao{ tx }
	err = dao.Insert(meeting)
	checkErr(err)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func (*MeetingServiceProvider)QueryAll() []Meeting {
	dao := meetingDao{ db }
	return dao.QueryAll()
}

func (*MeetingServiceProvider)QueryByTitle(title string) Meeting {
	dao := meetingDao{ db }
	meeting := dao.QueryByTitle(title)
	return meeting
}

func (*MeetingServiceProvider)QueryBy(key string, val string) []Meeting {
	dao := meetingDao{ db }
	return dao.QueryBy(key, val)
}



func (*MeetingServiceProvider)DeleteByTitle(title string) error {
	dao := meetingDao{ db }
	err := dao.DeleteByTitle(title)
	checkErr(err)
	if err != nil {
		return err
	}
	return nil
}

func (*MeetingServiceProvider)DeleteBySponsor(sponsor string) error {
	dao := meetingDao{ db }
	err := dao.DeleteBySponsor(sponsor)
	checkErr(err)
	if err != nil {
		return err
	}
	return nil
}

func (*MeetingServiceProvider)Update(meeting Meeting) error {
	tx, err := db.Begin()
	checkErr(err)
	if err != nil {
		return err
	}

	dao := meetingDao{ tx }
	err = dao.Update(meeting)
	checkErr(err)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}