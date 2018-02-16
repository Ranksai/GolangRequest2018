package userDefinedType

type User struct {
	Id   int
	Name string
}

type Game struct {
	Id   int
	Name string
}

type UserGame struct {
	Id    int
	User  User
	Game  Game
	Score int
}

type UserGames []UserGame

type UserAllGame struct {
	Id        int
	User      User
	UserGames UserGames
}
