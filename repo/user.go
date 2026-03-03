package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(u User) (*User, error)
	Find(email, password string) (*User, error)
	MatchUserCredentials(email, password string) (*User, error)
}

type userRepo struct {
	users []*User
}

func NewUserRepo() UserRepo {
	return &userRepo{
		users: make([]*User, 0),
	}
}

func (r *userRepo) Create(u User) (*User, error) {
	if u.ID != 0 {
		return &u, nil
	}
	u.ID = len(r.users) + 1
	r.users = append(r.users, &u)
	return &u, nil
}

func (r *userRepo) Find(email, password string) (*User, error) {
	for _, user := range r.users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}
	return nil, nil
}

func (r *userRepo) MatchUserCredentials(email, password string) (*User, error) {
	for _, user := range r.users {
		if user.Email == email && user.Password == password {
			return user, nil
		}
	}
	return nil, nil
}
