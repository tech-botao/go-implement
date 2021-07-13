package memory

import (
	"github.com/tech-botao/go-implement/cleanapp/app/domain/model"
	"sync"
)

// User 这个只有自己知道, 重复的定义, 这样好吗?
type User struct {
	ID string
	Email string
}

// userRepository repository的接口实现
type userRepository struct {
	mu *sync.Mutex
	users map[string]*User
}

func NewUserRepository() *userRepository {
	return &userRepository{
		mu:    &sync.Mutex{},
		users: map[string]*User{},
	}
}

func (r *userRepository) FindAll() ([]*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	users := make([]*model.User, len(r.users))
	i := 0
	for _, user := range r.users {
		users[i] = model.NewUser(user.ID, user.Email)
	}
	return users, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, user := range r.users {
		if user.Email == email {
			return model.NewUser(user.ID, user.Email), nil
		}
	}
	return nil, nil
}

func (r *userRepository) Save(user *model.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.users[user.GetID()] = &User{
		user.GetID(), user.GetEmail(),
	}
	return nil
}
