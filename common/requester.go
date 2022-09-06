package common

type Requester interface {
	GetUid() int
	GetEmail() string
	GetRole() string
}
