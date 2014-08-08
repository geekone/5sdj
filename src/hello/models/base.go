package models

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func init(){
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	dsn := dbuser + ":" + dbpassword + "@/" + dbname + "?charset=utf8"
	// orm.RegisterDataBase("default","mysql",dsn)		//orm注册数据库
	maxIdle := 30
	maxConn := 30
	//orm.RegisterDataBase("default", "mysql", "root:@/5sdj?charset=utf8", maxIdle, maxConn)
	orm.RegisterDataBase("default", "mysql", dsn, maxIdle, maxConn)
	orm.RegisterModel(new(Category))

	//在程序中直接调用自动建表：

	// 数据库别名
	name := "default"

	// drop table 后再建表 true
	force := false	

	// 打印执行过程
	verbose := true

	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
	    fmt.Println(err)
	}
}


//重新返回带前缀的表名的函数,前缀取自己配置文件 str 表名
func TableName(str string) string {
	return fmt.Sprintf("%s%s",beego.AppConfig.String("dbprefix"),str)
}