package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
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
