package main

import (
	"PMS/business/operates"
	"fmt"
	"os"
)

func main() {
	var userId int
	var userPwd string
	//初始页面
lable1:
	for {
		fmt.Println("----------欢迎使用LLD酒店管理系统----------")
		fmt.Println("1.登录用户")
		fmt.Println("2.注册用户")
		fmt.Println("3.退出系统")
		fmt.Println("请选择（1-3）")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			break lable1
		case 2:
			register, err := operates.Register()
			if err != nil {
				fmt.Println("注册失败 err=", err)
				goto lable1
			} else {
				if register.Code == 200 {
					fmt.Println("注册成功，请登录！")
					goto lable1
				} else {
					fmt.Println("注册失败 err=", err)
					goto lable1
				}
			}

		case 3:
			fmt.Println("退出系统")
			os.Exit(0)
		default:
			fmt.Println("输入有误！请重新输入！")
		}
	}
	//商家登录界面
	for {
		fmt.Println("----------欢迎使用LLD酒店管理系统----------")

		fmt.Println("请输入用户ID")
		fmt.Scanln(&userId)
		fmt.Println("请输入密码")
		fmt.Scanln(&userPwd)
		//确认用户ID和密码是否正确
		loginmes, err := operates.Login(userId, userPwd)
		if err == nil {
			if loginmes.Code == 200 {
				fmt.Println("登陆成功！")
				break
			} else {
				fmt.Println(loginmes.Error)
				//
				//continue
				goto lable1
			}
		} else {
			fmt.Println("login err=", err)
			return
		}
	}

	//登录成功后，进入管理界面
lable2:
	fmt.Println("----------菜单---------")
	fmt.Println("1.首页")
	fmt.Println("2.房态")
	fmt.Println("3.订单")
	fmt.Println("4.营收")
	fmt.Println("~~~请选择您的操作（1-4）~~~")
	var key int //选择进行的操作序号
	fmt.Scanln(&key)
	switch key {
	case 1: //进入首页
		totalPrice, roomBalance := operates.HomePage(userId)
		fmt.Printf("当日营收：%0.2f\n", totalPrice) //显示当日的营收
		fmt.Printf("房间余量：%v\n", roomBalance)   //显示当日房间余量
		fmt.Println()
		fmt.Println("输入 0 返回上一级")
		fmt.Println("输入 1 退出系统")
		var num int
		fmt.Scanln(&num)
		switch num {
		case 0:
			goto lable2
		case 1:
			os.Exit(0)
		}

	case 2: //进入房态页面

		operates.RoomState(userId)
		goto lable2
	case 3: //进入订单页面（查看过往订单）
	case 4: //进入营收页面（查看自定义时间段的营收）
	default:
		fmt.Println("输入有误请重新输入！")
		goto lable2
	}

}
