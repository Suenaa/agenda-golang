package entities

type UserServiceProvider struct {}

var UserService = UserServiceProvider {}

func (*UserServiceProvider) Insert(user *User) error {
	tx, err := db.Begin()
	checkErr(err)
	if err != nil {
		return err
	}
	

	dao := userDao{ tx }
	err = dao.Insert(user)
	checkErr(err)
	if err != nil {
		return err
	}
	
	tx.Commit()
	return nil
}

func (*UserServiceProvider) QueryAll() []User {
	dao := userDao{db}
	return dao.QueryAll()
}

func (*UserServiceProvider) QueryByName(name string) User {
	dao := userDao{ db }
	user := dao.QueryByName(name)
	return user
}

func (*UserServiceProvider)DeleteByName(name string) error {
	dao := userDao{ db }
	err := dao.DeleteByName(name)
	checkErr(err)
	if err != nil {
		return err
	}
	
	return nil
}