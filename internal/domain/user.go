package domain

type User struct {
	ID        ID
	Document  Document
	Name      string
	Email     Email
	Password  Password
	Enabled   Enabled
	CreatedAt CreatedAt
	UpdatedAt UpdatedAt
	DeletedAt DeletedAt
}

func NewUser(document string, name string, email string, password string) (*User, error) {
	u := &User{
		ID:        NewID(),
		Name:      name,
		Email:     Email(email),
		Password:  Password(password),
		Document:  Document(document),
		Enabled:   Enabled(false),
		CreatedAt: NewCreatedAt(),
		UpdatedAt: NewUpdatedAt(),
		DeletedAt: nil,
	}

	return u, nil
}

func NewFromUser(
	id ID,
	document string,
	name string,
	email string,
	password string,
	enabled Enabled,
	createdAt CreatedAt,
	updatedAt UpdatedAt,
	deletedAt DeletedAt,
) *User {
	return &User{
		ID:        id,
		Document:  Document(document),
		Name:      name,
		Email:     Email(email),
		Password:  Password(password),
		Enabled:   enabled,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}

func (u *User) ToJSON() map[string]interface{} {
	return map[string]interface{}{
		"id":         u.ID.String(),
		"document":   u.Document.String(),
		"name":       u.Name,
		"email":      u.Email.String(),
		"enabled":    u.Enabled.Bool(),
		"created_at": u.CreatedAt.String(),
		"updated_at": u.UpdatedAt.String(),
	}
}
