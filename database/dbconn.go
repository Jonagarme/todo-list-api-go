

type Task struct {
	gorm.Model
	Name   string `json:"name"`
	Status string `json:"status"`
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "<your-password>"
	dbname   = "todo-list-api"
)

var dsn string = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
	host, port, user, password, dbname)