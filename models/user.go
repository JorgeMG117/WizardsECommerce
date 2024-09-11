package models

import "github.com/JorgeMG117/WizardsECommerce/utils"

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

var userFile string = "data/users.json"

func GetUsers() ([]User, error) {
	var users []User
	err := utils.ReadFile(userFile, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func CheckUser(username string, password string) *User {
	var users []User
	err := utils.ReadFile(userFile, &users)
	if err != nil {
		return nil
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			return &user
		}
	}

	return nil
}

// func createUser(user User) error {
// 	mutex.Lock()
// 	defer mutex.Unlock()

// 	var users []User
// 	err := readFile(userFile, &users)
// 	if err != nil && !errors.Is(err, os.ErrNotExist) {
// 		return err
// 	}

// 	user.ID = len(users) + 1
// 	users = append(users, user)

// 	return writeFile(userFile, users)
// }

// func getUser(id int) (User, error) {
// 	var users []User
// 	err := readFile(userFile, &users)
// 	if err != nil {
// 		return User{}, err
// 	}

// 	for _, user := range users {
// 		if user.ID == id {
// 			return user, nil
// 		}
// 	}
// 	return User{}, errors.New("user not found")
// }
