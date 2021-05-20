package main

type Person struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Birthday int    `json:"birthday"`
}

// 同一パッケージなので意味ないが、メソッドの試し実装
func (p Person) getId() int {
	return p.Id
}

func (p Person) getName() string {
	return p.Name
}

func (p Person) getBirthDay() int {
	return p.Birthday
}
