package service

import (
  "errors"
  "io/ioutil"
  "github.com/Suenaa/agenda-golang/service/entities"
)

//用户注册
func UserRegister(userName string, password string, email string, phone string) error {
  allUsers := entities.UserService.QueryAll()
  isNameOk := true
  for i:= 0; i < len(allUsers); i ++ {
    if userName == allUsers[i].GetUsername() {
      isNameOk = false
    }
  }
  if isNameOk == true {
    var user entities.User
    user.Init(userName, password, email, phone)
    return entities.UserService.Insert(&user)
  } else {
    return errors.New("this username is aleardy exist")
  }
  return nil
}

//用户登录
func UserLogin(userName string, password string) error {
  currentUsername := GetCurrentUser()
  if currentUsername != "" {
    return errors.New("login failed, exist user login")
  }
  allUsers := entities.UserService.QueryAll()
  for i := 0; i < len(allUsers); i ++ {
    if userName == allUsers[i].GetUsername() {
      if password == allUsers[i].GetPassword() {
        currentUsername = userName
        return SetCurrentUser(userName)
      } else {
        return errors.New("incorrect password")
      }
    }
  }
  return errors.New("username not exist")
}

//用户登出
func UserLogout() error {
  currentUsername := GetCurrentUser()
  if currentUsername == "" {
    return errors.New("no user login")
  } else {
    return SetCurrentUser("")
  }
}

//查询所有用户
func QueryAllUsers() []entities.User {
  return entities.UserService.QueryAll()
}


//删除用户
//删除该用户的用户信息以及参加和发起的会议，如果该用户是某一个会议的发起者，则删除该会议，
//如果由于删除该用户造成某一个会议的参与者变成0，则删除该会议
func DeleteUser(password string) error {
  currentUsername := GetCurrentUser()
  allUsers := entities.UserService.QueryAll()
  if currentUsername == "" {
    return errors.New("no user login")
  } else {
    for i := 0; i < len(allUsers); i ++ {
      if currentUsername == allUsers[i].GetUsername() {
        if allUsers[i].GetPassword() == password {
          err := entities.MeetingService.DeleteBySponsor(currentUsername)
          if err != nil {
            return err
          }
          err = entities.UserService.DeleteByName(currentUsername)
          if err != nil {
            return err
          }
          return SetCurrentUser("")
        } else {
          return errors.New("incorrect password")
        }
      }
    }
  }
  return nil
}


//创建会议
func CreateMeeting(title string, startDate string, endDate string, participators []string) error {
  currentUsername := GetCurrentUser()
  if currentUsername == "" {
    return errors.New("no user login")
  }
  start := entities.StringToDate(startDate)
  end := entities.StringToDate(endDate)
  if start.IsAfter(end) {
    return errors.New("start time is after end time")
  }
  if entities.MeetingService.QueryByTitle(title).Title != "" {
    return errors.New("title aleardy exist")
  }
  if !IsUserSpace(currentUsername, startDate, endDate) {
    return errors.New("the sponsor at that time is busy")
  }
  for i := 0; i < len(participators); i ++ {
    if !IsUser(participators[i]) {
      return errors.New("participator "+participators[i]+" is not exist")
    }
    if !IsUserSpace(participators[i], startDate, endDate) {
      return errors.New("participator "+participators[i]+" is busy at that time")
    }
  }
  meeting := entities.Meeting{}
  meeting.Init(title, currentUsername, participators, startDate, endDate)
  return entities.MeetingService.Insert(&meeting)
}

//添加自己发起的某一会议的一个参与者
func AddParticipator(title string, participator string) error {
  currentUsername := GetCurrentUser()
  tMeeting := entities.MeetingService.QueryByTitle(title)
  startTime := (tMeeting).GetStart()
  endTime := (tMeeting).GetEnd()
  if currentUsername == "" {
    return errors.New("no user login")
  }
  if tMeeting.Title == "" {
    return errors.New("the title not exist")
  }
  if currentUsername != (tMeeting).GetSponsor() {
    return errors.New("current user is not the sponsor of the meeting")
  }
  if !IsUser(participator) {
    return errors.New("participator "+participator+" is not exist")
  }
  if !IsUserSpace(participator, startTime, endTime) {
    return errors.New("participator "+participator+" is busy at that time")
  }
  if tMeeting.AddParticipator(participator) == false {
    return errors.New("the new participator is aleardy in the participators")
  } else {
    return entities.MeetingService.Update(tMeeting)
  }
}

//删除自己创建的某一会议的一个参与者
//如果会议的参与者因此变成0，则删除该会议
func DeleteParticipator(title string, participator string) error {
  currentUsername := GetCurrentUser()
  tMeeting := entities.MeetingService.QueryByTitle(title)
  if currentUsername == "" {
    return errors.New("no user login")
  }
  if tMeeting.Title == "" {
    return errors.New("the title not exist")
  }
  if currentUsername != tMeeting.GetSponsor() {
    return errors.New("current user is not the sponsor of the meeting")
  }
  if !IsUser(participator) {
    return errors.New("participator "+participator+" is not exist")
  }
  if !tMeeting.IsParticipator(participator) {
    return errors.New("participator is not in the participators")
  }
  tMeeting.DeleteParticipator(participator)
  entities.MeetingService.Update(tMeeting)
  if len(tMeeting.GetParticipators()) == 0 {
    return entities.MeetingService.DeleteByTitle(title)
  }
  return nil
}

//查询会议，通过开始时间和结束时间查询当前用户需要参加的所有会议
func QueryMeeting(startDate string, endDate string) ([]entities.Meeting, error) {
  currentUsername := GetCurrentUser()
  if currentUsername == "" {
    return nil,errors.New("no user login")
  }
  start := entities.StringToDate(startDate)
  end := entities.StringToDate(endDate)
  if start.IsAfter(end) {
    return nil,errors.New("start time is after end time")
  }
  allMeeting := entities.MeetingService.QueryAll()
  returnMeeting := []entities.Meeting{}
  for i := 0; i < len(allMeeting); i ++ {
    if allMeeting[i].IsSponsor(currentUsername) {
      returnMeeting = append(returnMeeting, allMeeting[i])
    }
    if allMeeting[i].IsParticipator(currentUsername) {
      returnMeeting = append(returnMeeting, allMeeting[i])
    }
  }
  return returnMeeting,nil
}

//取消会议
func DeleteMeeting(title string) error {
  currentUsername := GetCurrentUser()
  tMeeting := entities.MeetingService.QueryByTitle(title)
  if currentUsername == "" {
    return errors.New("no user login")
  }
  if tMeeting.Title == "" {
    return errors.New("the title not exist")
  }
  if currentUsername != tMeeting.GetSponsor() {
    return errors.New("current user is not the sponsor of the meeting")
  }
  return entities.MeetingService.DeleteByTitle(title)
}

//退出会议
//如果退出会议之后的参与者为0的会议将会被删除
func QuitMeeting(title string) error {
  currentUsername := GetCurrentUser()
  tMeeting := entities.MeetingService.QueryByTitle(title)
  if currentUsername == "" {
    return errors.New("no user login")
  }
  if tMeeting.Title == "" {
    return errors.New("the title not exist")
  }
  if tMeeting.IsSponsor(currentUsername) {
    return errors.New("current user is the sponsor of the meeting, can't quit")
  }
  if !tMeeting.IsParticipator(currentUsername) {
    return errors.New("the user is not participators")
  }
  tMeeting.DeleteParticipator(currentUsername)
  if len(tMeeting.GetParticipators()) == 0 {
    return entities.MeetingService.DeleteByTitle(title)
  }
  return entities.MeetingService.Update(tMeeting)
}

//清空当前用户发起的所有会议安排
func DeleteAllMeeting() error {
  currentUsername := GetCurrentUser()
  if currentUsername == "" {
    return errors.New("no user login")
  }
  allMeeting := entities.MeetingService.QueryAll()
  if allMeeting[0].Title == "" {
    return nil
  }
  for i := 0; i < len(allMeeting); i ++ {
    if allMeeting[i].IsSponsor(currentUsername) {
      err := entities.MeetingService.DeleteByTitle(allMeeting[i].GetTitle())
      return err
    }
  }
  return nil
}

//--------------------非接口函数---------------------------

//判断一个用户在某个时间段是否空闲
func IsUserSpace(userName string, startDate string, endDate string) bool {
  start := entities.StringToDate(startDate)
  end := entities.StringToDate(endDate)
  allMeeting := entities.MeetingService.QueryAll()
  for i := 0; i < len(allMeeting); i ++ {
    if allMeeting[i].IsSponsor(userName) {
      startTime := entities.StringToDate(allMeeting[i].GetStart())
      endTime := entities.StringToDate(allMeeting[i].GetEnd())
      if !(!endTime.IsAfter(start) || !end.IsAfter(startTime)) {
        return false
      }
    }
    if allMeeting[i].IsParticipator(userName) {
      startTime := entities.StringToDate(allMeeting[i].GetStart())
      endTime := entities.StringToDate(allMeeting[i].GetEnd())
      if !(!endTime.IsAfter(start) || !end.IsAfter(startTime)) {
        return false
      }
    }
  }
  return true
}

//删除该用户发起的所有会议以及从参与会议中的参与者名单中删除
func DeleteMeetingByUserName(userName string) {
  allMeeting := entities.MeetingService.QueryAll()
  for i := 0; i < len(allMeeting); i ++ {
    if allMeeting[i].IsSponsor(userName) {
      entities.MeetingService.DeleteByTitle(allMeeting[i].GetTitle())
      continue
    }
    if allMeeting[i].IsParticipator(userName) {
      tMeeting := &allMeeting[i]
      tMeeting.DeleteParticipator(userName)
      if len((*tMeeting).GetParticipators()) == 0 {
        entities.MeetingService.DeleteByTitle(allMeeting[i].GetTitle())
      } else {
        entities.MeetingService.Update(allMeeting[i])
      }
    }
  }
}

//判断传入的名字是否是已注册用户
func IsUser(name string) bool {
  allUser := entities.UserService.QueryAll()
  for i := 0; i < len(allUser); i ++ {
    if name == allUser[i].GetUsername() {
      return true
    }
  }
  return false
}

func SetCurrentUser(name string) error {
  err := ioutil.WriteFile("./current.dat", []byte(name), 0644)
  return err
}

func GetCurrentUser() string {
  content, err := ioutil.ReadFile("./current.dat")
  if err != nil {
    return ""
  }
  return string(content)
}
