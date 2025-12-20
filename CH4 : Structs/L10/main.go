package main

type Membership struct {
	Type             string
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType string) User {
	var charLimit int = 100

	if membershipType == "premium" {
		charLimit = 1000
	}

	var u = User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: charLimit,
		},
	}

	return u
}
