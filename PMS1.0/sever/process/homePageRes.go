package process

import (
	"PMS/message"
	"PMS/tools"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"net"
)

func HomePageRes(mesint string, conn net.Conn) {
	//userId, _ := strconv.Atoi(mesint)
	var hpmes message.HomePage
	json.Unmarshal([]byte(mesint), &hpmes)
	c := tools.GetRedisPool() //链接数据库
	defer c.Close()
	data, _ := redis.Bytes(c.Do("hget", hpmes.UserId, hpmes.Day))
	var mes message.TotalMes
	json.Unmarshal(data, &mes)
	var homepageres message.HomePageRes //首页返回信息
	homepageres.RoomBalance = mes.RoomBalance
	homepageres.TotalPrice = mes.TotalPrice
	data, _ = json.Marshal(homepageres)
	var mesres message.Message
	mesres.Type = message.HomePageResType
	mesres.Data = string(data)
	data, _ = json.Marshal(mesres)
	//向客户端返回首页信息
	var tf tools.Temp
	tf.Conn = conn
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println(err)
	}
	//
	return
}
