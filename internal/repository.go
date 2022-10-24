package internal

import "errors"

var ErrUserNameAlreadyExist = errors.New("UserName is already exist")

type Repository struct {
	data map[string]Membership
}

func (r *Repository) Create(user Membership) {
	r.data[user.ID] = user
}

func (r *Repository) Update(user Membership) {
	r.data[user.ID] = user
}

func (r *Repository) Delete(id string) {
	delete(r.data, id)
}

func (r *Repository) Check(id string) Membership {
	return r.data[id]
}

func NewRepository(data map[string]Membership) *Repository {
	return &Repository{data: data}
}
