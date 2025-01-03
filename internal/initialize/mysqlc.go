package initialize

import (
	"database/sql"
	"fmt"

	"github.com/Youknow2509/go-ecommerce/global"
	"go.uber.org/zap"
)

// Handle err panic
func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

// Initial my sql
func InitMysqlC() {
	m := global.Config.MySQL

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)

	db, err := sql.Open("mysql", s)

	checkErrorPanicC(err, "SQLC connection failed")

	global.Logger.Info("SQLC connection successful")
	global.Mdbc = db

	// set Pool
	SetPoolC()

	// migrate tables
	// MigrateTables()

	// genTableDAO
	// genTableDAO()
}

// InitMysql().SetPool()
func SetPoolC() {
	// m := global.Config.MySQL

	// sqlDb, err := global.Mdb.DB()
	// if err != nil {
	// 	fmt.Println("MySQL error: %s::", err)
	// 	global.Logger.Error("SetPool error", zap.Error(err))
	// }

	// sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))    // Thoi gian toi da ket noi nhan doi -> Phuc hoi ket noi
	// sqlDb.SetMaxOpenConns(m.MaxOpenConns)                      // Gioi han so luong ket noi toi da
	// sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime)) // Gioi han thoi gian toi da cua ket noi

	// TODO
}

// genTableDAO
func genTableDAOC() {
	// g := gen.NewGenerator(gen.Config{
	// 	OutPath: "./internal/model",                                                 // output path
	// 	Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	// })

	// // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	// g.UseDB(global.Mdb) // reuse your gorm db

	// // Generate all table
	// g.GenerateAllTable()

	// // Generate the code
	// g.Execute()

	// TODO

}

// migrate tables
func MigrateTablesC() {
	// err := global.Mdb.AutoMigrate(
	// 	&po.User{},
	// 	&po.Role{},
	// )
	// checkErrorPanic(err, "AutoMigrate tables failed")

	// TODO

}
