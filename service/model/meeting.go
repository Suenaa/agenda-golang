package model

import (
	"fmt"
	"time"
)

type Meeting struct {
	Title string
	Sponsor string
	Participators []string
	Start string
	End string
}

//the examples of start and end: 2017-10-20T19:00 2017-01-01T0:00
func (meeting *Meeting) Init(title string, sponsor string, 
	participators []string, start string, end string) {
	meeting.Title = title
	meeting.Sponsor = sponsor
	meeting.Participators = participators
	meeting.Start = start
	meeting.End = end
}

func (meeting Meeting) GetTitle() string {
	return meeting.Title
}

func (meeting *Meeting) SetTitle(title string) {
	meeting.Title = title
}

func (meeting Meeting) GetSponsor() string {
	return meeting.Sponsor
}

func (meeting *Meeting) SetSponsor(sponsor string) {
	meeting.Sponsor = sponsor
}

func (meeting Meeting) GetParticipators() []string {
	return meeting.Participators
}

func (meeting *Meeting) SetParticipators(participators []string) {
	meeting.Participators = participators
}

func (meeting Meeting) GetStart() string {
	t, _ := time.Parse(time.RFC3339, meeting.Start + ":00Z")
	start := t.String()[0 : 16]
	return start
}

func (meeting *Meeting) SetStart(start string) {
	meeting.Start = start
}

func (meeting Meeting) GetEnd() string {
	t, _ := time.Parse(time.RFC3339, meeting.End + ":00Z")
	end := t.String()[0 : 16]
	return end
}

func (meeting *Meeting) SetEnd(end string) {
	meeting.End = end
}

func (meeting Meeting) IsSponsor(username string) bool {
	if meeting.Sponsor == username {
		return true
	}
	return false
}

func (meeting Meeting) IsParticipator(username string) bool {
	for _, participator := range meeting.GetParticipators() {
		if participator == username {
			return true
		}
	}
	return false
}

//bool好像该改成err
func (meeting *Meeting) AddParticipator(username string) bool {
	if meeting.IsParticipator(username) { //判断是否已在参与者当中
		return false
	}
	meeting.SetParticipators(append(meeting.GetParticipators(), username))
	return true
}

//bool好像该改成err
func (meeting *Meeting) DeleteParticipator(username string) bool {
	for i := 0; i < len(meeting.GetParticipators()); i++ {
		if meeting.GetParticipators()[i] == username {
			meeting.SetParticipators(append(meeting.GetParticipators()[: i], meeting.GetParticipators()[i + 1:]...))
			return true
		}
	}
	return false
}

func (meeting Meeting) GetParticipatorsLength() int {
	return len(meeting.GetParticipators())
}

func (meeting Meeting) String() {
	fmt.Println(meeting.GetTitle())
	fmt.Println("　- sponsor: " + meeting.GetSponsor())
	fmt.Println("　- time: " + meeting.GetStart() + " - " + meeting.GetEnd())
	fmt.Print("　- participators: ")
	length := len(meeting.GetParticipators())
	for i := 0; i < length - 1; i++ {
		fmt.Print(meeting.GetParticipators()[i] + ", ")
	}
	fmt.Println(meeting.GetParticipators()[length - 1])
	fmt.Println()
}