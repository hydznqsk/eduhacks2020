package database

// 定义路由
const (
	APILogin                = "/api/login"                //用户登录的接口
	APIManagerStudentGet    = "/api/manager/student/get"  //管理端获取学生的接口
	APILogout               = "/api/logout"               //用户退出登录的接口
	APIManagerCollegeGet    = "/api/manager/college/get"  //管理端获取学院的接口
	APIManagerCollegeView   = "/api/manager/college/html" //管理端获取学院的 html 渲染
	APIManagerCollegeAdd    = "/api/manager/college/add"  //管理端添加学院的接口
	APIManagerCollegeDelete = "/api/manager/college/del"  //管理端删除学院的接口
	APIManagerCollegeEdit   = "/api/manager/college/edit" //管理端更新学院的接口
	APIManagerMajorGet      = "/api/manager/major/get"    //管理端获取专业的接口
	APIManagerMajorAdd      = "/api/manager/major/add"    //管理端添加专业的接口
	APIManagerMajorDelete   = "/api/manager/major/del"    //管理端删除专业的接口
	APIManagerMajorEdit     = "/api/manager/major/edit"   //管理端编辑专业的接口
)