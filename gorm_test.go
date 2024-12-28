package learngorm

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func Connection() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/perpustakaan?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
	if err != nil {
		panic(err)
	}
	return db
}
var db = Connection()
func TestConnectionToDB(t *testing.T) {
	assert.NotNil(t, db)
}

func TestExecuteSQL(t *testing.T) {
	err := db.Exec("INSERT INTO buku(title, genre, publication_year, author_id) VALUE (?, ?, ?, ?)", "Hanyut di Kesunyian", "Galau", 2019, 3).Error
	assert.Nil(t, err)
	err = db.Exec("INSERT INTO buku(title, genre, publication_year, author_id) VALUE (?, ?, ?, ?)", "Kemana Harus Pergi", "Tak Tahu", 2019, 5).Error
	assert.Nil(t, err)
	err = db.Exec("INSERT INTO buku(title, genre, publication_year, author_id) VALUE (?, ?, ?, ?)", "Malin Kundang", "Kemana Saja", 2019, 5).Error
	assert.Nil(t, err)
}

type Data struct {
	Id int
	Title string
	Genre string
	PublicationYear int
	AuthorId int
}
func TestRawSQL(t *testing.T) {
	var data Data
	err := db.Raw("SELECT id, title, genre, publication_year, author_id FROM buku WHERE ?", 1).Scan(&data).Error
	assert.Nil(t, err)
	assert.Equal(t, "Fantasy", data.Genre)

	var databanyak []Data
	err = db.Raw("SELECT id, title, genre, publication_year, author_id FROM buku").Scan(&databanyak).Error
	assert.Nil(t, err)
	assert.Equal(t, 8, len(databanyak))
}

func TestRowSQL(t *testing.T) {
	var samples []Data
	rows, err := db.Raw("SELECT id, title, genre, publication_year, author_id FROM buku").Rows()
	assert.Nil(t, err)

	defer rows.Close()
	for rows.Next() {
		var buku Data
		err := rows.Scan(&buku.AuthorId, &buku.Title, &buku.Genre, &buku.PublicationYear, &buku.Id)
		assert.Nil(t, err)
		samples = append(samples, Data{
			AuthorId: buku.AuthorId,
			Title: buku.Title,
			Genre: buku.Genre,
			PublicationYear: buku.PublicationYear,
			Id: buku.Id,
		})
	}
	assert.Equal(t, 8, len(samples))
}

func TestScanRow(t *testing.T) {
	rows, err := db.Raw("SELECT * FROM author").Rows()
	assert.Nil(t, err)
	var data []Data
	for rows.Next() {
		err := db.ScanRows(rows, &data)
		assert.Nil(t, err)
	}
	assert.Equal(t, 5, len(data))
}

func TestCreateUser(t *testing.T) {
    user := User{
        ID: "7",
        Password: "admin123",
        Name: Name{
            FirstName: "Erna",
            MiddleName: "Dama",
            LastName: "yanti",
        },
        Information: "kaga ngaruh ke database",
    }
    response := db.Create(&user)
    assert.Nil(t, response.Error)
    assert.Equal(t, int64(1), response.RowsAffected)
}

func TestBatchUser(t *testing.T) {
    var users []User
    for i := 22; i <= 23; i++ {
        users = append(users, User{
            ID: strconv.Itoa(i),
            Password: "rahasia",
            Name: Name{
                FirstName: "User " + strconv.Itoa(i),
            },
        })
    }
    response := db.Create(&users)
    assert.Nil(t, response.Error)
    println(response.RowsAffected)
    assert.Equal(t, int64(2), response.RowsAffected)
}

func TestTransactionSuccess(t *testing.T) {
    err := db.Transaction(func(tx *gorm.DB) error {
        err := tx.Create(&User{ID: "122", Password: "admin123", Name: Name{FirstName: "Erna"}}).Error
        if err != nil {
            return err
        }
        return nil
    })

    assert.NotNil(t, err)
}

func TestTransactionGagal(t *testing.T) {
    err := db.Transaction(func(tx *gorm.DB) error {
        err := tx.Create(&User{ID: "127", Password: "admin123", Name: Name{FirstName: "Erna"}}).Error
        if err != nil {
            return err
        }
        err = tx.Create(&User{ID: "122", Password: "admin123", Name: Name{FirstName: "Erna"}}).Error
        if err != nil {
            return err
        }
        return nil
    })

    assert.NotNil(t, err)
}

func TestQuerySingleObject(t *testing.T) {
    user := User{}
    err := db.First(&user).Error
    assert.Nil(t, err)
    assert.Equal(t, strconv.Itoa(10), user.ID)
    err = db.Last(&user).Error
    assert.Nil(t, err)
    assert.Equal(t, strconv.Itoa(10), user.ID)
}

func TestQueryInlineCondition(t *testing.T) {
    user := User{}
    err := db.Take(&user, "id = ?", "122").Error
    assert.Nil(t, err)
    assert.Equal(t, strconv.Itoa(122), user.ID)
}

func TestQueryAllObjects(t *testing.T) {
    var users []User
    err := db.Find(&users, "id in ?", []string{"122","20", "23"}).Error
    assert.Nil(t, err)
    assert.Equal(t, 3, len(users))
}

func TestQueryCondition(t *testing.T) {
    var user []User
    err := db.Where("first_name like ?", "%Erna%").Where("password = ?", "admin123").Find(&user).Error
    assert.Nil(t, err)
    assert.Equal(t, 3, len(user))
}

func TestOrCondition(t *testing.T) {
    var user []User
    err := db.Where("first_name like ?", "%Erna%").Or("password = ?", "rahasia").Find(&user).Error
    println(len(user))
    assert.Nil(t, err)
    assert.Equal(t, 19, len(user))
}
func TestNotCondition(t *testing.T) {
    var user []User
    err := db.Not("first_name like ?", "%Erna%").Or("password = ?", "rahasia").Find(&user).Error
    println(len(user))
    assert.Nil(t, err)
    assert.Equal(t, 16, len(user))
}
func TestSelectCondition(t *testing.T) {
    var users []User
    err := db.Select("id", "first_name").Find(&users).Error
    assert.Nil(t, err)
    for _, user := range users {
        assert.NotNil(t, user.ID)
        assert.NotEqual(t, "", user.Name.FirstName)
    }
    println(len(users))
    assert.Equal(t, 19, len(users))
}
func TestStructCondition(t *testing.T) {
    userCondition := User{
        Name: Name{
            FirstName: "User 17",
        },
        Password: "rahasia",
    }
    var user []User
    err := db.Where(userCondition).Find(&user).Error
    fmt.Println(user)

    assert.Nil(t, err)
    assert.Equal(t, 1, len(user))
}
func TestMapCondition(t *testing.T) {
    mapCondition := map[string]interface{}{
        "middle_name":"",
    }
    var user []User
    err := db.Where(mapCondition).Find(&user).Error
    assert.Nil(t, err)
    assert.Equal(t, 18, len(user))
}
func TestOrderLimitOffset(t *testing.T) {
    var user []User
    err := db.Order("id asc, first_name desc").Limit(5).Offset(5).Find(&user).Error
    fmt.Println(user)
    assert.Nil(t, err)
    assert.Equal(t, 5, len(user))
}

type UserResponse struct {
    ID string
    FirstName string
    LastName string
}

func TestQueryNonModel(t *testing.T) {
    var user []UserResponse
    err := db.Model(&User{}).Select("id", "first_name", "last_name").Find(&user).Error
    fmt.Println(user)
    assert.Nil(t, err)
}

func TestUpdate(t *testing.T){
    user := User{}
    err := db.Take(&user, "id = ?", "15").Error
    assert.Nil(t, err)

    user.Name.FirstName = "Abdi"
    user.Name.LastName = "Setiawan"
    user.Password = "rahasia123"
    err = db.Save(&user).Error
    assert.Nil(t, err)
}

func TestSelectedColumns(t *testing.T) {
    err := db.Model(&User{}).Where("id = ?", "122").Updates(map[string]interface{}{
        "first_name":"Kakashi",
        "last_name":"Uchiha",
    }).Error
    assert.Nil(t, err)
    err = db.Model(&User{}).Where("id = ?", "122").Update("password", "passwordyangtelahdiubah").Error
    assert.Nil(t, err)
    err = db.Where("id = ?", "122").Updates(
        User{
            Name: Name{
                FirstName: "Abdi",
                LastName: "Setiawan",
            },
            }).Error
    assert.Nil(t, err)
}

func TestAutoIncrement(t *testing.T) {
    for i := 0; i < 10; i++ {
        userLog := UserLog{
            UserId: "1",
            Action: "Test Action",
        }
        err := db.Create(&userLog).Error
        fmt.Println(userLog)
        assert.Nil(t, err)
    }
}

func TestSaveAtauUpdate(t *testing.T) {
    userLog := UserLog{
        UserId: "satu",
        Action: "Halah Sia BOYYYYY",
    }
    err := db.Save(&userLog).Error
    assert.Nil(t, err)
    userLog.UserId = "satusatu"
    err = db.Save(&userLog).Error
    assert.Nil(t, err)
}


func TestSaveAtauUpdateNonIncrement(t *testing.T) {
    userLog := User{
        ID: "12221",
        Name: Name{
            FirstName: "kacauuuuuu",
        },
    }
    err := db.Save(&userLog).Error
    assert.Nil(t, err)
    userLog.Name.FirstName = "kacauuuuu v2"
    err = db.Save(&userLog).Error
    assert.Nil(t, err)

}

func TestConflict(t *testing.T) {
    user := User{
        ID: "99",
        Name: Name{
            FirstName: "kacau v2",
        },
    }
    err := db.Clauses(clause.OnConflict{
        UpdateAll: true,
    }).Create(&user).Error
    assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
    var user User
    err := db.Take(&user, "id = ?", "23").Error
    assert.Nil(t, err)

    err = db.Delete(&user).Error
    assert.Nil(t, err)

    err = db.Delete(&User{}, "id = ?", "7").Error
    assert.Nil(t, err)

    err = db.Where("id = ?", "8").Delete(&User{}).Error
    assert.Nil(t, err)
}

func TestSoftDelete(t *testing.T) {
    todo := Todo{
        UserId: "Erna Damayanti",
        Title: "Abdi iniiii",
        Description: "lorem ipsum bla bla bla bla bla bla",
    }

    err := db.Create(&todo).Error
    assert.Nil(t, err)
    
    err = db.Delete(&todo).Error
    assert.Nil(t, err)
    assert.NotNil(t, todo.DeletedAt)

    var todos []Todo
    err = db.Find(&todos).Error
    assert.Nil(t, err)
    assert.Equal(t, 0, len(todos))
}

func TestUnscoped(t *testing.T) {
    var todo Todo
    err := db.Unscoped().Take(&todo, "id = 2").Error
    assert.Nil(t, err)

    err = db.Unscoped().Delete(&todo).Error
    assert.Nil(t, err)

    var todos []Todo
    err = db.Unscoped().Find(&todos).Error
    assert.Nil(t, err)
}

func TestLock(t *testing.T) {
    err := db.Transaction(func(tx *gorm.DB) error {
        var user User
        err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&user, "id = ?", "10").Error
        if err != nil {
            return err
        }

        user.Password = "Sudah digantin dengan Locking"
        user.Name.FirstName = "Abdi Locking"
        err = tx.Save(&user).Error
        return err
    })
    assert.Nil(t, err)
}

func TestCreateWallet(t *testing.T) {
    wallet := Wallet{
        ID: "1",
        UserId: "10",
        Balance: 10000,
    }
    err := db.Create(&wallet).Error
    assert.Nil(t, err)
}

func TestRetrieveRelation(t *testing.T) {
    var user User
    err := db.Model(&User{}).Preload("Wallet").Take(&user, "id = ?", "10").Error
    assert.Nil(t, err)

    assert.Equal(t, "10", user.ID)
    assert.Equal(t, "10", user.Wallet.UserId)
}

func TestRetrieveRelationJoin(t *testing.T) {
    var user User
    err := db.Model(&User{}).Joins("Wallet").Take(&user, "users.id = ?", "10").Error
    assert.Nil(t, err)

    assert.Equal(t, "10", user.ID)
    assert.Equal(t, "10", user.Wallet.UserId)
}