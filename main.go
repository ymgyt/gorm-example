package main

import (
	"fmt"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := dbConfig{
		user: "gopher",
		pass: "gopher",
		host: "localhost",
		port: "3333",
		db:   "app",
	}.open()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db = setup(db)

	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	db.AutoMigrate(&A{}, &B{})

	// truncate
	db.Delete(&A{})
	db.Delete(&B{})

	// insert
	createdA := time.Now()
	createdB := time.Now().AddDate(3, 0, 0)
	as := []*A{
		{ID: 1, Name: "a1", CreatedAt: createdA},
		{ID: 2, Name: "a2", CreatedAt: createdA},
		{ID: 3, Name: "a3", CreatedAt: createdA},
	}
	bs := []*B{
		{ID: 1, Name: "b1", CreatedAt: createdB},
		{ID: 2, Name: "b2", CreatedAt: createdB},
		{ID: 3, Name: "b3", CreatedAt: createdB},
	}
	for _, a := range as {
		if err := db.Create(a).Error; err != nil {
			panic(err)
		}
	}
	for _, b := range bs {
		if err := db.Create(b).Error; err != nil {
			panic(err)
		}
	}

	// patern 1
	var bind1 []*bind
	err = db.
		Table("a").
		Select("*").
		Joins("INNER JOIN b ON a.ID = b.ID").
		Find(&bind1).
		Error
	if err != nil {
		panic(err)
	}
	print(bind1)

}

type bind struct {
	A
	B
}

type A struct {
	ID        int       `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (a *A) TableName() string {
	return "a"
}

type B struct {
	ID        int       `gorm:"column:id"`
	Name      string    `gorm:"column:name"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (b *B) TableName() string {
	return "b"
}

func print(bs []*bind) {
	for _, b := range bs {
		fmt.Printf("A.ID: %d A.Name: %s A.CreatedAt: %s\n", b.A.ID, b.A.Name, b.A.CreatedAt)
		fmt.Printf("B.ID: %d B.Name: %s B.CreatedAt: %s\n", b.B.ID, b.B.Name, b.B.CreatedAt)
		fmt.Println()
	}
}

type dbConfig struct {
	user string
	pass string
	host string
	port string
	db   string
}

func setup(db *gorm.DB) *gorm.DB {
	db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	db = db.LogMode(true)
	return db
}

func (c dbConfig) open() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", c.user, c.pass, c.host, c.port, c.db)

	v := url.Values{}
	v.Add("charset", "utf8mb4")
	v.Add("parseTime", "True")
	v.Add("loc", "Asia/Tokyo")

	dsn = fmt.Sprintf("%s?%s", dsn, v.Encode())

	return gorm.Open("mysql", dsn)
}
