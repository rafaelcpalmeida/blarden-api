package actions

import (
	"blarden/models"
	"fmt"
)

func (as *ActionSuite) Test_UsersResource_List() {
	users := models.Users{
		{
			PhoneNumber: "911111111",
		},
		{
			PhoneNumber: "912222222",
		},
		{
			PhoneNumber: "913333333",
		},
	}
	for _, user := range users {
		err := as.DB.Create(&user)
		as.NoError(err)
	}

	res := as.JSON("/users").Get()
	body := res.Body.String()
	for _, user := range users {
		as.Contains(body, user.PhoneNumber)
	}
}

func (as *ActionSuite) Test_UsersResource_Show() {
	w := &models.User{PhoneNumber: "911111111"}
	_ = as.JSON("/users").Post(w)
	err := as.DB.First(w)
	as.NoError(err)
	as.NotZero(w.ID)
	res := as.JSON(fmt.Sprintf("/users/%s", w.ID)).Get()
	user := &models.User{}
	res.Bind(user)
	as.Equal(user.PhoneNumber, "911111111")
	as.Equal(user.HasPermission, false)
}

func (as *ActionSuite) Test_UsersResource_Create() {
	// Successfully create a user
	w := &models.User{PhoneNumber: "911111111"}
	_ = as.JSON("/users").Post(w)
	err := as.DB.First(w)
	as.NoError(err)
	as.NotZero(w.ID)
	as.Equal("911111111", w.PhoneNumber)

	// Unsuccessfully create a user due to duplicate phone number
	w = &models.User{PhoneNumber: "911111111"}
	res := as.JSON("/users").Post(w)
	as.Equal(res.Code, 400)
	as.Contains(res.Body.String(), "user is already registered")

	// Unsuccessfully create a user due to invalid phone number
	w = &models.User{PhoneNumber: "91111111"}
	res = as.JSON("/users").Post(w)
	err = as.DB.First(w)
	as.Equal(res.Code, 422)
	as.Contains(res.Body.String(), "Phone number doesn't appear to be valid")

	// Unsuccessfully create a user due to empty phone number
	w = &models.User{}
	res = as.JSON("/users").Post(w)
	err = as.DB.First(w)
	as.Equal(res.Code, 422)
	as.Contains(res.Body.String(), "Phone number can not be blank")
}

func (as *ActionSuite) Test_UsersResource_Update() {
	// Successfully update a user
	w := &models.User{PhoneNumber: "911111111"}
	_ = as.JSON("/users").Post(w)
	err := as.DB.First(w)
	as.NoError(err)
	as.NotZero(w.ID)
	as.Equal("911111111", w.PhoneNumber)
	as.Equal(false, w.HasPermission)
	w.PhoneNumber = "922222222"
	w.HasPermission = true
	_ = as.JSON(fmt.Sprintf("/users/%s", w.ID)).Put(w)
	res := as.JSON(fmt.Sprintf("/users/%s", w.ID)).Get()
	user := &models.User{}
	res.Bind(user)
	as.Equal(user.PhoneNumber, "922222222")
	as.Equal(user.HasPermission, true)

	// Unsuccessfully update a user due to duplicate phone number
	w.PhoneNumber = "922222222"
	res = as.JSON(fmt.Sprintf("/users/%s", w.ID)).Put(w)
	as.Equal(res.Code, 400)
	as.Contains(res.Body.String(), "user is already registered")

	// Unsuccessfully update a user due to invalid phone number
	w.PhoneNumber = "92222222"
	res = as.JSON(fmt.Sprintf("/users/%s", w.ID)).Put(w)
	as.Equal(res.Code, 422)
	as.Contains(res.Body.String(), "Phone number doesn't appear to be valid")

	// Unsuccessfully update a user due to empty phone number
	w.PhoneNumber = ""
	res = as.JSON(fmt.Sprintf("/users/%s", w.ID)).Put(w)
	as.Equal(res.Code, 422)
	as.Contains(res.Body.String(), "Phone number can not be blank")
}

func (as *ActionSuite) Test_UsersResource_Destroy() {
	w := &models.User{PhoneNumber: "911111111"}
	_ = as.JSON("/users").Post(w)
	err := as.DB.First(w)
	as.NoError(err)
	as.NotZero(w.ID)
	res := as.JSON(fmt.Sprintf("/users/%s", w.ID)).Delete()
	as.Equal(res.Code, 200)
}
