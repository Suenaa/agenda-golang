package entities

import (
	"encoding/json"
	"io/ioutil"

	"github.com/Suenaa/agenda-golang/service/tools"
)

//saveUser 将所有用户保存
func saveUser(users []User) {
	jzon, err := json.Marshal(users)
	tools.Report(err)
	ioutil.WriteFile("datas/users.dat", jzon, 0644)
}

//readUser 读取所有用户
func readUser() (res []User) {
	content, err := ioutil.ReadFile("datas/users.dat")
	if err != nil {
		return []User{}
	}
	json.Unmarshal(content, &res)
	return
}

//saveMeeting 将所有会议保存
func saveMeeting(meetings []Meeting) {
	jzon, err := json.Marshal(meetings)
	tools.Report(err)
	err = ioutil.WriteFile("datas/meetings.dat", jzon, 0644)
	tools.Report(err)
}

//readMeeting 读取所有会议
func readMeeting() (res []Meeting) {
	content, err := ioutil.ReadFile("datas/meetings.dat")
	if err != nil {
		return []Meeting{}
	}
	err = json.Unmarshal(content, &res)
	tools.Report(err)
	return
}

//AddUser 添加用户
func AddUser(user User) {
	users := readUser()
	users = append(users, user)
	saveUser(users)
}

//DeleteUser 删除用户
func DeleteUser(username string) {
	users := readUser()
	index := -1
	for i, user := range users {
		if user.GetUsername() == username {
			index = i
			break
		}
	}
	if index != -1 {
		users = append(users[:index], users[index+1:]...)
		saveUser(users)
	}
}

//SetCurrentUser 设置当前登录用户
//空字符串表示 退出登录
func SetCurrentUser(username string) {
	err := ioutil.WriteFile("datas/current.dat", []byte(username), 0644)
	tools.Report(err)
}

//GetCurrentUser 获取当前登录用户的用户名
//空字符串表示 未登录
func GetCurrentUser() string {
	content, err := ioutil.ReadFile("datas/current.dat")
	if err != nil {
		return ""
	}
	return string(content)
}

//GetAllUser 获取所有注册用户
func GetAllUser() []User {
	return readUser()
}

//AddMeeting 增加会议
func AddMeeting(meeting Meeting) {
	meetings := readMeeting()
	meetings = append(meetings, meeting)
	saveMeeting(meetings)
}

//GetMeeting 获取会议
func GetMeeting(title string) *Meeting {
	meetings := readMeeting()
	for _, meeting := range meetings {
		if meeting.GetTitle() == title {
			return &meeting
		}
	}
	return nil
}

//UpdateMeeting 修改会议
func UpdateMeeting(meeting Meeting) {
	meetings := readMeeting()
	for index, m := range meetings {
		if meeting.GetTitle() == m.GetTitle() {
			meetings[index] = meeting
			saveMeeting(meetings)
			break
		}
	}
}

//GetALLMeeting 获取所有会议
func GetALLMeeting() []Meeting {
	return readMeeting()
}

//DeleteMeeting 删除指定会议
func DeleteMeeting(title string) {
	meetings := readMeeting()
	index := -1
	for i, meeting := range meetings {
		if meeting.GetTitle() == title {
			index = i
			break
		}
	}
	if index != -1 {
		meetings = append(meetings[:index], meetings[index+1:]...)
		saveMeeting(meetings)
	}
}
