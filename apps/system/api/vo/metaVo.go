package vo

/**s
 *  路由meta对象参数说明
 *  meta: {
 *      title:          菜单栏及 tagsView 栏、菜单搜索名称
 *      isLink：        是否超链接菜单，开启外链条件，`1、isLink:true 2、链接地址不为空`
 *      isHide：        是否隐藏此路由
 *      isKeepAlive：   是否缓存组件状态
 *      isAffix：       是否固定在 tagsView 栏上
 *      isFrame：      是否内嵌窗口，，开启条件，`1、isFrame:true 2、链接地址不为空`
 *      auth：          当前路由权限标识（多个请用逗号隔开），最后转成数组格式，用于与当前用户权限进行对比，控制路由显示、隐藏
 *      icon：          菜单、tagsView 图标，阿里：加 `iconfont xxx`，fontawesome：加 `fa xxx`
 *  }
 */
type MetaVo struct {
	Title       string   `json:"title"`
	IsLink      string   `json:"isLink"`
	IsHide      bool     `json:"isHide"`
	IsKeepAlive bool     `json:"isKeepAlive"`
	IsAffix     bool     `json:"isAffix"`
	IsIframe    bool     `json:"isIframe"`
	Auth        []string `json:"auth"`
	Icon        string   `json:"icon"`
}
