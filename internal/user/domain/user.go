package domain

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Age       int
}

func NewUser(id int, firstName string, lastName string, email string, age int) *User {
	return &User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Age:       age,
	}
}

func (u *User) GetId() int {
	return u.Id
}

func (u *User) GetFirstName() string {
	return u.FirstName
}

func (u *User) GetLastName() string {
	return u.LastName
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetAge() int {
	return u.Age
}

func (u *User) SetId(id int) {
	u.Id = id
}

func (u *User) SetFirstName(firstName string) {
	u.FirstName = firstName
}

func (u *User) SetLastName(lastName string) {
	u.LastName = lastName
}

func (u *User) SetEmail(email string) {
	u.Email = email
}

func (u *User) SetAge(age int) {
	u.Age = age
}