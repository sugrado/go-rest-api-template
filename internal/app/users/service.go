package users

type Repository interface {
	Save(firstName, lastName, email string) (int, error)
	Find(id int) (*User, error)
}

type Service interface {
	Create(user User) (int, error)
	Get(id int) (*User, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) Get(id int) (*User, error) {
	u, err := s.r.Find(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *service) Create(u User) (int, error) {
	id, err := s.r.Save(u.FirstName, u.LastName, u.Email)
	if err != nil {
		return 0, err
	}
	return id, nil
}
