package user

type UserFormatter struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"email"`
	Token    string `json:"token"`
	Uuid     string `json:"uuid"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		UserName: user.UserName,
		Token:    token,
		Uuid:     user.Uuid,
	}

	return formatter
}

func FormatUsers(user User) UserFormatter {
	formatter := UserFormatter{
		ID:       user.ID,
		Name:     user.Name,
		UserName: user.UserName,
		Uuid:     user.Uuid,
	}
	return formatter
}

func FormatUsersAll(users []User) []UserFormatter {
	if len(users) == 0 {
		return []UserFormatter{}
	}
	var output []UserFormatter

	for _, value := range users {
		users := FormatUsers(value)
		if users.ID != 0 {
			output = append(output, users)
		}
	}
	return output
}
