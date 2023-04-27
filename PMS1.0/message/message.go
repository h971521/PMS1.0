package message

const (
	LoginMesType      = "loginMes"     //登录信息
	LoginResMesType   = "longinResMes" //登录返回信息
	RegisterMesType   = "registerMes"  //注册信息
	RoomUpdateMesType = "roomUpdateMes"
	HomePageType      = "homePage"      //首页信息
	HomePageResType   = "homePageRes"   //首页返回信息
	DayRoomMes        = "dayRoomMess"   //当日房间信息
	DayRoomMesRes     = "dayRoomMesRes" //当日房间信息返回给客户端
	AddOrdersType     = "addOrders"     //对某个房间添加订单
	AddOrderResType   = "addOrderRes"   //添加订单返回消息
	DelOrdersType     = "delOrders"     //删除指定日期的指定房间信息
	DelOrdersResType  = "delOrdersRes"  //删除订单返回信息
	ReadOrderType     = "readOrder"     //查看指定日期的指定房间信息
	ReadOrderResType  = "readOrderRes"  //查看订单返回信息
	TurnOrderType     = "turnOrder"     //修改订单信息
	TurnOrderResType  = "turnOrderRes"  //修改订单返回信息
)

// 定义一个信息的结构体，方便判断信息的类型
type Message struct {
	Type string //信息的类型
	Data string //信息的内容
}
type UserMes struct { //登录信息
	UserId  int    `json:"userId"`
	UserPwd string `json:"userPwd"`
}
type LoginResMes struct {
	Code  int    //状态码为200表示登录成功，状态码为500表示登录失败
	Error string //返回错误提示信息
}

type RegisterMes struct { //注册信息
	UserId  int
	UserPwd string
	RoomNum int
}
type RegistrResMes struct { //反馈注册信息
	Code  int
	Error string
}

type RoomMes struct {
	CustomersName string  //客人姓名
	CustomersTel  int     //联系方式
	CheckInTime   string  //入住时间（后续换为time,直接导入入住时间）
	Price         float32 //房间价格
	// 后续可补充字段...
}
type RoomUpdate struct { //对应房间的房间信息更新
	Roomnum int //房间号
	RoomMes
	Update
}
type Update struct {
	Day    string
	UserId int
}
type TotalMes struct {
	Rooms       []RoomMes
	TotalPrice  float32
	RoomBalance int
}
type HomePageRes struct {
	//Rooms []RoomMes
	TotalPrice  float32
	RoomBalance int
}
type HomePage struct {
	Update
}

//var DayMes struct {
//	Day    string
//	UserId int
//}
