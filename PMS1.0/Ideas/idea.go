package Ideas

//打印输出商家登录界面，用户名和密码
//1. 登录进入酒店管理界面
//2. 管理界面菜单，包括
//		首页 	房态		订单		营收		等...（可扩展）
//  2.1 首页内容
//		当日营收		房间余量		等...
//	2.2 房态 --> 选择进入下级页面
//		2.2.1 房态显示房间号 --> 选择房间查看房间状态
//			2.2.1.1 房间状态包括订单信息、房间状态（空脏、空净、住脏、住净、维修等）--> 选择对房间的操作
//				2.2.1.1.1 录入客人信息、调整房间状态... -->进入下一级菜单输入
//
//数据库存储 :(目前的存储形式不好，存储和读取的数据量大，存在无用数据)
//		   key:用户Id
//		   field:用户密码、房间数量、时间（以天为单位）
//		   value:-月-日 房间总信息（如果一天有多个人入住，则以链表形式存储）
//		      		--->通过房间号进行增删查改操作-->姓名、手机号、入住时间（时/分）、房间价格
//

//小结：
	//功能：
	// 登录、注册、查询当日总收入及房间余量、订单的增删查改
	//可以对指定日期的订单进行修改
	//
	//不足：
	//1.同一天同一个房间如果有多个订单信息，只能显示最新的订单信息
	//解决方法：房间的订单以链表的形式存储
	//2.在获取指定房间信息时，需要先读取所有房间的信息
	//原因：在数据库中存储的方式有问题，将房间信息（结构体）存放到了一个切片中，需要先读取切片再读取指定房间的信息
	//解决办法：每个房间的信息单独存储...
	//3.每次向服务器发送数据时，都需要建立新的链接（net.Conn）
	//解决方法：统一用一个链接并发处理？
	//4.订单功能没有加入
	//.只有基础功能
	//解决方法：时间允许的条件下扩展功能...
