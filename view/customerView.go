package main

import (
	"fmt"
	"go_code/customerManage/model"
	"go_code/customerManage/service"
)

type customerView struct {
	key string //接收用户输入
	loop bool //表示是否循环显示
	customerService *service.CustomerService
}
//显示所有客户信息
func (this *customerView) list() {
	//获取当前所有的客户信息
	customers := this.customerService.List()
	//显示
	fmt.Println("------------客户列表------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("\n------------客户列表完成------------\n\n")
}
func (this *customerView) add() {
	fmt.Println("------------添加客户------------")
	fmt.Println("姓名： ")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别： ")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄： ")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话： ")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("电邮： ")
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customer) {
		fmt.Println("------------添加完成------------")
	} else {
		fmt.Println("------------添加失败------------")
	}
}
func (this *customerView) delete() {
	fmt.Println("------------删除客户------------")
	fmt.Println("请选择待删除客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	fmt.Println("确认是否删除(Y/N): ")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		if this.customerService.Delete(id) {
			fmt.Println("------------删除完成------------")
		} else {
			fmt.Println("-----删除失败, 输入id号不存在----")
		}
	}
}
func (this *customerView) update() {
	fmt.Println("------------修改客户------------")
	fmt.Println("请选择待修改客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}
	index := this.customerService.FindById(id)
	if index == -1 {
		return
	}
	customers := this.customerService.List()
	fmt.Printf("姓名(%v)： ", customers[index].Name)
	name := ""
	fmt.Scanln(&name)
	fmt.Printf("性别(%v)： ", customers[index].Gender)
	gender := ""
	fmt.Scanln(&gender)
	fmt.Printf("年龄(%v)： ", customers[index].Age)
	age := 0
	fmt.Scanln(&age)
	fmt.Printf("电话(%v)： ", customers[index].Phone)
	phone := ""
	fmt.Scanln(&phone)
	fmt.Printf("电邮(%v)： ", customers[index].Email)
	email := ""
	fmt.Scanln(&email)
	customer := model.NewCustomer(id,name, gender, age, phone, email)
	if this.customerService.Update(customer) {
		fmt.Println("------------修改完成------------")
	} else {
		fmt.Println("-----修改失败, 用户id不存在-----")
	}
}

func (this *customerView) exit() {
	fmt.Println("确认是否退出(Y/N): ")
	for  {
		fmt.Scanln(&this.key)
		if this.key == "Y" || this.key == "y" || this.key == "N" || this.key == "n" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出(Y/N): ")
	}
	if this.key == "Y" || this.key == "y" {
		this.loop = false
	}
}

func (this *customerView) mainMenu() {

	for  {
		fmt.Println("------------客户信息管理软件-------------")
		fmt.Println("            1 添加客户")
		fmt.Println("            2 修改客户")
		fmt.Println("            3 删除客户")
		fmt.Println("            4 客户列表")
		fmt.Println("            5 退出")
		fmt.Print("请选择(1-5): ")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.list()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !this.loop {
			break
		}
	}
	fmt.Println("你退出了客户关系管理系统...")
}

func main() {
	//在main函数中，创建一个customerView，并运行主菜单
	customerView := customerView{
		key: "",
		loop: true,
	}
	//完成customerView结构体的customerService字段的初始化
	customerView.customerService = service.NewCustomerService()
	//显示主菜单
	customerView.mainMenu()
}