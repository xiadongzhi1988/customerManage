package service

import "go_code/customerManage/model"

//完成对Customer的操作，包括
//增删改查
type CustomerService struct {
	customers []model.Customer
	//声明一个字段，表示当前切片含有多少个客户
	//该字段后面可以作为新客户的id+1
	customerNum int
}

//编写一个方法，可以返回 *CustomerService
func NewCustomerService() *CustomerService {
	//为了能看到客户在切片中，初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zs@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}
//添加客户到customers切片
func (this *CustomerService) Add(customer model.Customer) bool {
	//分配id规则
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}
//根据id删除用户
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	//index == -1 ，没有这个客户
	if index == -1 {
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true
}
func (this *CustomerService) Update(customer model.Customer) bool {
	index := this.FindById(customer.Id)
	//index == -1 ，没有这个客户
	if index == -1 {
		return false
	}
	if customer.Name != "" {
		this.customers[index].Name = customer.Name
	}
	if customer.Gender != "" {
		this.customers[index].Gender = customer.Gender
	}
	if customer.Age != 0 {
		this.customers[index].Age = customer.Age
	}
	if customer.Phone != "" {
		this.customers[index].Phone = customer.Phone
	}
	if customer.Email != "" {
		this.customers[index].Email = customer.Email
	}
	return true
}

//根据id查找客户在切片中对应的下标，如果没有，返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}