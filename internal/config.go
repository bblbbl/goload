package internal

type Config struct {
	Url            string
	Method         string
	Times          int
	Concurrent     int
	RequiredStatus int
}
