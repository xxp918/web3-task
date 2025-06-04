package main

import (
	"fmt"

	// _ "github.com/go-sql-driver/mysql"
	// "github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// var DB *sqlx.DB
var DB *gorm.DB
var mysqlloggger logger.Interface
func main() {
	DB = DB.Session(&gorm.Session{
		Logger:mysqlloggger,
	})
	// //创建表
	// //DB.AutoMigrate(&Student{})

	// //题目一：数据库使用
	// //crud()
	// //创建表
	// DB.AutoMigrate(&Account{},&Transaction{})

	// //事务操作
	// err :=transfer(1,2,100)

	// if err != nil {
	// 	fmt.Println(err)
	// }else{
	// 	fmt.Println("转账成功")
	// }
	// sqlxQuery("技术部")
	// sqlxQuery2()
	// sqlxQuery3()

	//创建表
	// DB.AutoMigrate(&User{},&Post{},&Comment{})

	// queryUserPostsAndComments(1)
	//  var user User 
	//  //查询用户1
	//  result := DB.First(&user,1)
	//  if result.Error != nil {
	// 	fmt.Println("查询失败：",result.Error)
	// 	return
	//  }

	//  fmt.Println(user)
	// //创建文章
	// post := Post{
	// 	Title: "我的第一篇文章",
	// 	Content: "这是我的第一篇文章的内容",
	// 	UserID: user.ID,
	// 	User: user,
	// }
	// DB.Create(&post)
	var comment Comment
	DB.Preload("Post").First(&comment,5).Delete(&comment)
}


//进阶gorm
//题目1：模型定义
//假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
//使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。,
//编写Go代码，使用Gorm创建这些模型对应的数据库表。

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	PostCount int
	Posts    []Post
}
type Post struct {
	gorm.Model
	Title   string
	Content string
	UserID  uint
	User    User
	CommentCount int
	CommentStatus string
	Comments []Comment
}
type Comment struct {
	gorm.Model
	Content string
	PostID  uint
	Post    Post
}

// 题目3：钩子函数
// 为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。,
// 为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	tx.Model(&p.User).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1))
	return
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("Comment:",c)
	var post Post
	tx.First(&post, c.PostID)
	fmt.Println("post:",post)
	if post.CommentCount == 0 {
		tx.Model(&post).UpdateColumn("comment_status", "无评论")
	}
	return
}



// 题目2：关联查询
//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。,
// 编写Go代码，使用Gorm查询评论数量最多的文章信息。

func queryUserPostsAndComments(userID uint) {
	var user User
	DB.Preload("Posts.Comments").First(&user, userID)
	fmt.Println("用户发布的文章及其对应的评论信息：")
	for _, post := range user.Posts {
		fmt.Printf("文章标题：%s\n", post.Title)
		fmt.Println("评论：")
		for _, comment := range post.Comments {
			fmt.Printf("评论内容：%s\n", comment.Content)
		}
		fmt.Println()
	}
}












//Sqlx入门
// 题目1：使用SQL扩展库进行查询
//假设你已经使用Sqlx连接到一个数据库，并且有一个 employees 表，包含字段 id 、 name 、 department 、 salary 。
//要求 ：编写Go代码，使用Sqlx查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中。,
//编写Go代码，使用Sqlx查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中。
// type Employee struct {
// 	ID          int
// 	Name        string
// 	Department  string
// 	Salary      float64
// }
// func sqlxQuery(department string) {
// 	//查询所有部门为"技术部"的员工信息
// 	var employees []Employee
// 	err := DB.Select(&employees, "SELECT * FROM employees WHERE department = ?", department)
// 	if err != nil {
// 		fmt.Println("查询失败：", err)
// 		return
// 	}
// 	fmt.Println("部门为技术部的员工信息：")
// 	for _, employee := range employees {
// 		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %f\n", employee.ID, employee.Name, employee.Department, employee.Salary)
// 	}
// }
// func sqlxQuery2() {
// 	//查询工资最高的员工信息
// 	var employee Employee
// 	err := DB.Get(&employee, "SELECT * FROM employees WHERE salary = (SELECT MAX(salary) FROM employees)")
// 	if err != nil {
// 		fmt.Println("查询失败：", err)
// 		return
// 	}
// 	fmt.Println("工资最高的:")
// 	fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %f\n", employee.ID, employee.Name, employee.Department, employee.Salary)
// }

// // 题目2：实现类型安全映射
// //假设有一个 books 表，包含字段 id 、 title 、 author 、 price 。
// //定义一个 Book 结构体，包含与 books 表对应的字段。,
// //编写Go代码，使用Sqlx执行一个复杂的查询，例如查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中，确保类型安全。
// type Book struct {
// 	ID     int
// 	Title  string
// 	Author string
// 	Price  float64
// }

// func sqlxQuery3() {
// 	//查询价格大于50元的书籍
// 	var books []Book
// 	err := DB.Select(&books, "SELECT * FROM books WHERE price > ?", 50)
// 	if err != nil {
// 		fmt.Println("查询失败：", err)
// 		return
// 	}
// 	fmt.Println("价格大于50元的书籍：")
// 	for _, book := range books {
// 		fmt.Printf("ID: %d, Title: %s, Author: %s, Price: %f\n", book.ID, book.Title, book.Author, book.Price)
// 	}

// }



// //使用sqlx连接数据库
// func init(){
// 		//数据库基本信息
// 	username :="root"
// 	password :="123456"
// 	host :="127.0.0.1"
// 	port :="3306"
// 	Dbname :="gotest"
// 	timeout :="10s"

// 	//连接数据库
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

// 	// 创建
// 	db, err := sqlx.Connect("mysql", dsn)

// 	if err != nil {
// 		panic("连接失败，err="+err.Error())
// 	}
// 	//连接成功
// 	DB = db
// }


func init() {
	//数据库基本信息
	username :="root"
	password :="123456"
	host :="127.0.0.1"
	port :="3306"
	Dbname :="gotest"
	timeout :="10s"

	mysqlloggger = logger.Default.LogMode(logger.Info)

	//连接信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	// 创建数据库连接
	db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{
		// Logger: mysqlloggger,
	})
	if err !=nil {
		panic("连接失败，err="+err.Error())
	}
	//连接成功
	DB =db
}

//题目2：事务语句
//假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
//要求：编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。

type Account struct {
	ID      int
	Balance float64
}
type Transaction struct {
	ID             int
	FromAccountID  int
	ToAccountID    int
	Amount         float64
}

func transfer(fromAccountID, toAccountID int, amount float64) error {
	//开启事务
	tx := DB.Begin()
	//检查账户A的余额是否足够
	var fromAccount Account
	if err := tx.First(&fromAccount, fromAccountID).Error; err != nil {
		tx.Rollback()
		return err
	}
	if fromAccount.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("insufficient balance")
	}
	//从账户A扣除100元
	fromAccount.Balance -= amount
	if err := tx.Save(&fromAccount).Error; err != nil {
		tx.Rollback()
		return err
	}
	//向账户B增加100元
	var toAccount Account
	if err := tx.First(&toAccount, toAccountID).Error; err != nil {
		tx.Rollback()
		return err
	}
	toAccount.Balance += amount
	if err := tx.Save(&toAccount).Error; err != nil {
		tx.Rollback()
		return err
	}
	//在transactions表中记录该笔转账信息
	transaction := Transaction{
		FromAccountID: fromAccountID,
		ToAccountID:   toAccountID,
		Amount:        amount,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return err
	}
	//提交事务
	return tx.Commit().Error
}


//SQL语句练习
//题目1：基本CRUD操作
//假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）要求 ：
//编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。,
//编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。,
//编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。,
//编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
type Student struct {
	ID     uint
	Name   string
	Age    int
	Grade  string
}

func crud() {
	//插入数据
	//stuCreate()
	//查询所有年龄大于18岁的学生信息
	//stuQuery(18)
	//更新数据 将学生姓名为"张三"的学生年级更新为"四年级"
	//stuUpdate("张三","四年级")
	//删除数据 删除年龄小于15岁的学生记录
	stuDelete(15)
}
///删除数据
func stuDelete(age int) {
	DB.Where("age < ?", age).Delete(&Student{})
}
//更新数据
func stuUpdate(name string, grade string) {
	DB.Model(&Student{}).Where("name = ?", name).Update("grade",grade)
}
//查询所有年龄大于18岁的学生信息
func stuQuery(age int) {
	var students18 []Student
	DB.Where("age > ?", age).Find(&students18)
	fmt.Println(students18)
}
//插入数据
func stuCreate(){
	//插入数据
	// DB.Create(&Student{
	// 	Name: "张三",
	// 	Age:  20,
	// 	Grade: "三年级",
	// })
	//批量插入数据
	students := []Student{
		{Name: "李四", Age: 19,Grade: "三年级"},
		{Name: "王五", Age: 20,Grade: "四年级"},
		{Name: "王二麻子", Age: 17,Grade: "二年级"},
		{Name: "李小狗", Age: 18,Grade: "三年级"},
		{Name: "李小三", Age: 19,Grade: "四年级"},
		{Name: "小青", Age: 12,Grade: "一年级"},
	}
	DB.Create(&students)
}
