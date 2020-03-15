package customer

type Customer struct {
	Name    string
	Surname string
	Age     int
}

func New(n, s string, a int) Customer {
	return Customer{
		Name:    n,
		Surname: s,
		Age:     a,
	}
}
