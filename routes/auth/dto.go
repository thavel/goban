package auth

type Credentials struct {
	User     *string `json:"user"`
	Password *string `json:"password"`
	Token    *string `json:"token"`
}

func (c *Credentials) Valid() bool {
	return (c.User != nil && c.Password != nil) || c.Token != nil
}
