package Config

import (
	"example/auth-services/model"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func Connect() (*gorm.DB, error) {
	err := godotenv.Load("app.env")
	if err != nil {
		return nil, err
	}
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp("+ os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

if (err != nil && err.Error() == "Error 1049 (42000): Unknown database '"+ os.Getenv("DB_NAME") + "'"){
	
		dsn1 := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASS") + "@tcp("+ os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn1), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		db.Exec("CREATE DATABASE " + os.Getenv("DB_NAME"))
		db.Exec("USE " + os.Getenv("DB_NAME"))
		
		err = db.AutoMigrate(model.User{}, model.Token{}, model.Vertify{})
		if err != nil {
			return nil, err
		}
		db.Exec("ALTER TABLE user ADD UNIQUE INDEX email (email)")
		db.Exec("ALTER TABLE tokens ADD UNIQUE INDEX token (token)")
		db.Exec("Alter TABLE vertifies ADD UNIQUE INDEX email_code (email, code)")
		db.Exec("Alter TABLE user ADD CONSTRAINT FK_User_Token  FOREGIN KEY (id) REFERENCES tokens (user_id) ON DELETE CASCADE")
		fmt.Println("Database created and migrated")
		
		db, err = Connect()
}



	return db, err
}