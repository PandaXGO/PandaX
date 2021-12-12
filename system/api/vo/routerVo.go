package vo

/**
 * 路由对象参数说明
 *  {
 *     component:          组件地址
 *     redirect:           重定向地址，当设置 noRedirect 的时候该路由在面包屑导航中不可被点击
 *     path:               路由地址
 *     name:               路由名字
 *     // 路由meta对象参数说明
 *     meta: {
 *         title:          菜单栏及 tagsView 栏、菜单搜索名称（国际化）
 *         isLink：        是否超链接菜单，开启外链条件，`1、isLink:true 2、链接地址不为空`
 *         isHide：        是否隐藏此路由
 *         isKeepAlive：   是否缓存组件状态
 *         isAffix：       是否固定在 tagsView 栏上
 *         isFrame：      是否内嵌窗口，，开启条件，`1、isIframe:true 2、链接地址不为空`
 *         auth：          当前路由权限标识（多个请用逗号隔开），最后转成数组格式，用于与当前用户权限进行对比，控制路由显示、隐藏
 *         icon：          菜单、tagsView 图标，阿里：加 `iconfont xxx`，fontawesome：加 `fa xxx`
 *     }
 *  }
 *
 */
type RouterVo struct {
	Name      string     `json:"name"`
	Path      string     `json:"path"`
	Redirect  string     `json:"redirect"`
	Component string     `json:"component"`
	Meta      MetaVo     `json:"meta"`
	Children  []RouterVo `json:"children"`
}
