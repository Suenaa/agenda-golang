package service

import (
  "errors"
  model "github.com/Suenaa/agenda-go/model"
  storage "github.com/Suenaa/agenda-go/storage"
)

//用户注册
func UserRegister(userName string, password string, email string, phone string) error {
  allUsers := storage.GetAllUser()
  isNameOk := true
  for i:= 0; i < len(allUsers); i ++ {
    if userName == allUsers[i].GetUsername() {
      isNameOk = false
    }
  }
  if isNameOk == true {
    storage.AddUser(model.User{Username:userName, Password:password, Email:email, Phone:phone})
  } else {
    return errors.New("this username is aleardy exist")
  }
  return nil
}

//用户登录
func UserLogin(userName string, password string) error {
  if storage.GetCurrentUser() != "" {
    return errors.New("login failed, exist user login")
  }
  allUsers := storage.GetAllUser()
  for i := 0; i < len(allUsers); i ++ {
    if userName == allUsers[i].GetUsername() {
      if password == allUsers[i].GetPassword() {
        storage.SetCurrentUser(userName)
        return nil
      } else {
        return errors.New("incorrect password")
      }
    }
  }
  return errors.New("username not exist")
}

//用户登出
func UserLogout() error {
  if storage.GetCurrentUser() == "" {
    return errors.New("no user login")
  } else {
    storage.SetCurrentUser("")
    return nil
  }
}

//查询所有用户
func QueryAllUsers() []model.User {
  return storage.GetAllUser()
}


//删除用户
//删除该用户的用户信息以及参加和发起的会议，如果该用户是某一个会议的发起者，则删除该会议，
//如果由于删除该用户造成某一个会议的参与者变成0，则删除该会议
func DeleteUser(password string) error {
  currentUser := storage.GetCurrentUser()
  allUsers := storage.GetAllUser()
  if currentUser == "" {
    return errors.New("no user login")
  } else {
    for i := 0; i < len(allUsers); i ++ {
      if currentUser == allUsers[i].GetUsername() {
        if allUsers[i].GetPassword() == password {
          DeleteMeetingByUserName(currentUser)
          storage.DeleteUser(currentUser)
          storage.SetCurrentUser("")
          return nil
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
  currentUser := storage.GetCurrentUser()
  if currentUser == "" {
    return errors.New("no user login")
  }
  start := model.StringToDate(startDate)
  end := model.StringToDate(endDate)
  if start.IsAfter(end) {
    return errors.New("start time is after end time")
  }
  if storage.GetMeeting(title) != nil {
    return errors.New("title aleardy exist")
  }
  if !IsUserSpace(currentUser, startDate, endDate) {
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
  storage.AddMeeting(model.Meeting{Title:title, Sponsor:currentUser, Participators:participators, Start:startDate, End:endDate})
  return nil
}

//添加自己发起的某一会议的一个参与者
func AddParticipator(title string, participator string) error {
  tMeeting := storage.GetMeeting(title)
  startTime := (*tMeeting).GetStart()
  endTime := (*tMeeting).GetEnd()
  currentUser := storage.GetCurrentUser()
  if currentUser == "" {
    return errors.New("no user login")
  }
  if tMeeting == nil {
    return errors.New("the title not exist")
  }
  if currentUser != (*tMeeting).GetSponsor() {
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
    storage.UpdateMeeting(*tMeeting)
    return nil
  }
}

//删除自己创建的某一会议的一个参与者
//如果会议的参与者因此变成0，则删除该会议
func DeleteParticipator(title string, participator string) error {
  tMeeting := storage.GetMeeting(title)
  currentUser := storage.GetCurrentUser()
  if currentUser == "" {
    return errors.New("no user login")
  }
  if tMeeting == nil {
    return errors.New("the title not exist")
  }
  if currentUser != (*tMeeting).GetSponsor() {
    return errors.New("current user is not the sponsor of the meeting")
  }
  if !IsUser(participator) {
    return errors.New("participator "+participator+" is not exist")
  }
  if !(*tMeeting).IsParticipator(participator) {
    return errors.New("participator is not in the participators")
  }
  tMeeting.DeleteParticipator(participator)
  storage.UpdateMeeting(*tMeeting)
  if len((*tMeeting).GetParticipators()) == 0 {
    storage.DeleteMeeting(title)
  }
  return nil
}

//查询会议，通过开始时间和结束时间查询当前用户需要参加的所有会议
func QueryMeeting(startDate string, endDate string) ([]model.Meeting, error) {
  currentUser := storage.GetCurrentUser()
  if currentUser == "" {
    return nil,errors.New("no user login")
  }
  start := model.StringToDate(startDate)
  end := model.StringToDate(endDate)
  if start.IsAfter(end) {
    return nil,errors.New("start time is after end time")
  }
  allMeeting := storage.GetALLMeeting()
  returnMeeting := []model.Meeting{}
  for i := 0; i < len(allMeeting); i ++ {
    if allMeeting[i].IsSponsor(currentUser) {
      returnMeeting = append(returnMeeting, allMeeting[i])
    }
    if allMeeting[i].IsParticipator(currentUser) {
      returnMeeting = append(returnMeeting, allMeeting[i])
    }
  }
  return returnMeeting,nil
}

//取消会议
func DeleteMeeting(title string) error {
  tMeeting := storage.GetMeeting(title)
  currentUser := storage.GetCurrentUser()
  if currentUser == "" {
    return errors.New("no user login")
  }
  if tMeeting == nil {
    return errors.New("the title not exist")
  }
  if currentUser != (*tMeeting).GetSponsor() {
    return errors.New("current user is not the sponsor of the meeting")
  }
  storage.DeleteMeeting(title)
  return nil
}

//退出会议
//如果退出会议之后的参与者为0的会议将会被删除
func QuitMeeting(title string) error {
  tMeeting := storage.GetMeeting(title)
  currentUser := storage.GetCurrentUser()
  if currentUser == "" {
    return errors.New("no user login")
  }
  if tMeeting == nil {
    return errors.New("the title not exist")
  }
  if (*tMeeting).IsSponsor(currentUser) {
    return errors.New("current user is the sponsor of the meeting, can't quit")
  }
  if !(*tMeeting).IsParticipator(currentUser) {
    return errors.New("the user is not participators")
  }
  (*tMeeting).DeleteParticipator(currentUser)
  if len((*tMeeting).GetParticipators()) == 0 {
    storage.DeleteMeeting(title)
    return nil
  }
  return nil
}

//清空当前用户发起的所有会议安排
func DeleteAllMeeting() error {
  currentUser := storage.GetCurrentUser()
  if currentUser == "" {
    return errors.New("no user login")
  }
  allMeeting := storage.GetALLMeeting()
  for i := 0; i < len(allMeeting); i ++ {
    if allMeeting[i].IsSponsor(currentUser) {
      storage.DeleteMeeting(allMeeting[i].GetTitle())
    }
  }
  return nil
}

//--------------------非接口函数---------------------------

//判断一个用户在某个时间段是否空闲
func IsUserSpace(userName string, startDate string, endDate string) bool {
  start := model.StringToDate(startDate)
  end := model.StringToDate(endDate)
  allMeeting := storage.GetALLMeeting()
  for i := 0; i < len(allMeeting); i ++ {
    if allMeeting[i].IsSponsor(userName) {
      startTime := model.StringToDate(allMeeting[i].GetStart())
      endTime := model.StringToDate(allMeeting[i].GetEnd())
      if !(!endTime.IsAfter(start) || !end.IsAfter(startTime)) {
        return false
      }
    }
    if allMeeting[i].IsParticipator(userName) {
      startTime := model.StringToDate(allMeeting[i].GetStart())
      endTime := model.StringToDate(allMeeting[i].GetEnd())
      if !(!endTime.IsAfter(start) || !end.IsAfter(startTime)) {
        return false
      }
    }
  }
  return true
}

//删除该用户发起的所有会议以及从参与会议中的参与者名单中删除
func DeleteMeetingByUserName(userName string) {
  allMeeting := storage.GetALLMeeting()
  for i := 0; i < len(allMeeting); i ++ {
    if allMeeting[i].IsSponsor(userName) {
      storage.DeleteMeeting(allMeeting[i].GetTitle())
      continue
    }
    if allMeeting[i].IsParticipator(userName) {
      tMeeting := &allMeeting[i]
      tMeeting.DeleteParticipator(userName)
      if len((*tMeeting).GetParticipators()) == 0 {
        storage.DeleteMeeting(allMeeting[i].GetTitle())
      } else {
        storage.UpdateMeeting(*tMeeting)
      }
    }
  }
}

//判断传入的名字是否是已注册用户
func IsUser(name string) bool {
  allUser := storage.GetAllUser()
  for i := 0; i < len(allUser); i ++ {
    if name == allUser[i].GetUsername() {
      return true
    }
  }
  return false
}
