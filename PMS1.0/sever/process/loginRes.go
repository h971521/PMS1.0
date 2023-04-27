package process

import (
	"PMS/message"
	"PMS/tools"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
)

func LoginRes(data string, conn net.Conn) (err error) {
	var userMes message.UserMes
	err = json.Unmarshal([]byte(data), &userMes)
	if err != nil {
		fmt.Println("json.Unmarshal userMes fail err=", err)
		return
	}
	//fmt.Println(userMes)

	var loginresMes message.LoginResMes
	//数据库比对
	//链接数据库
	connr, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial fail err=", err)
		return
	}
	//将从数据库读取到的信息转为字符串类型
	res, err := redis.String(connr.Do("hget", userMes.UserId, "pwd"))
	if err != nil {
		//fmt.Println("get userMes.UserId fail err=", err)
		loginresMes.Code = 500
		//err = errors.New("用户名或密码错误！请重新输入！")
		loginresMes.Error = "用户名或密码错误！请重新输入！"
		//return
	} else {

		//fmt.Println(res)
		if /*userMes.UserId == 15209299216 && */ userMes.UserPwd == res {
			loginresMes.Code = 200
		} else {
			loginresMes.Code = 500
			//err = errors.New("用户名或密码错误！请重新输入！")
			loginresMes.Error = "用户名或密码错误！请重新输入！"
		}
	}
	//fmt.Println(loginresMes)

	datares, err := json.Marshal(loginresMes)
	if err != nil {
		fmt.Println("json.Marshal(loginresMes) err=", err)
		return
	}
	var mes message.Message
	mes.Type = message.LoginResMesType
	mes.Data = string(datares)
	datares, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
		return
	}
	var resMes tools.Temp
	resMes.Conn = conn
	err = resMes.WritePkg(datares)
	if err != nil {
		fmt.Println("login resMes.WritePkg fail err=", err)
		return
	}
	return
}
