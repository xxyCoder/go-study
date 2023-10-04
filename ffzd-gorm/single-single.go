package main

func ss() {
	DB.AutoMigrate(&User{}, &UserInfo{})
	DB.Create(&User{
		Name: "xxyCoder",
		UserInfo: &UserInfo{
			Like: "Code",
			Age:  21,
		},
	})
}
