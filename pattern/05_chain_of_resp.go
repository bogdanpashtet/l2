package pattern

import "fmt"

type Site interface {
	execute(*User)
	setNext(Site)
}

type User struct {
	login string
}

type Authorization struct {
	next Site
}

func (a *Authorization) execute(user *User) {
	fmt.Printf("Authorization is running for user %s\n", user.login)
	a.next.execute(user)
}

func (a *Authorization) setNext(site Site) {
	a.next = site
}

type CheckRights struct {
	next Site
}

func (cR *CheckRights) execute(user *User) {
	fmt.Printf("Checking rights for user %s\n", user.login)
}

func (cR *CheckRights) setNext(site Site) {
	cR.next = nil
}

func ChainOfResponsabilityFunc() {
	check := &CheckRights{}
	auth := &Authorization{}
	auth.setNext(check)

	user := &User{"kriper2004"}
	auth.execute(user)
}
