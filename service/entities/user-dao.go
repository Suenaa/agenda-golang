package entities

import(
	
)

type userDao DaoSource

func (dao *userDao)Insert(user *User) error {
	var userInsertStmt = "insert into users(username, password, email, phone) values(?, ?, ?, ?)"

	stmt, err := dao.Prepare(userInsertStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Username, user.Password, user.Email, user.Phone)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	user.Id = int(id)
	return nil
}

func (dao *userDao)QueryAll() []User {
	var userQueryAll = "select * from users"
	rows, err := dao.Query(userQueryAll)
	checkErr(err)
	defer rows.Close()

	ulist := make([]User, 0, 0)
	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Id, &u.Username, &u.Password, &u.Email, &u.Phone)
		checkErr(err)
		ulist = append(ulist, u)
	}
	return ulist
}

func (dao *userDao)QueryByName(name string) User {
	userQueryByName := "select * from users where username=" + name
	rows, err := dao.Query(userQueryByName)
	defer rows.Close()
	checkErr(err)
	u := User{}
	if rows.Next() {
		err = rows.Scan(&u.Id, &u.Username, &u.Password, &u.Email, &u.Phone)
	}
	checkErr(err)

	return u
}

func (dao *userDao)DeleteByName(name string) error {
	userDeleteByName := "delete from users where username=" + name
	_, err := dao.Exec(userDeleteByName)
	if err != nil {
		return err
	}
	return nil
}
