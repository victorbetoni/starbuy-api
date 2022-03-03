package database

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Schema   string `json:"schema"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func RetrieveCredentials() Credentials {
	/*info := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}*/
}
