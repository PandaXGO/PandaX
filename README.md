# PandaX

<div align="center"><img src="https://s3.bmp.ovh/imgs/2021/12/3c6ddcf3877171c9.png"/></div>
<div align="center"><h3 align="center">PandaX 快速开发平台</h3></div>
<div align="center"><h3 align="center">基于Gin前后端分离架构，代码精简，开箱即用，前端紧随前沿 Vue3.0 + TypeScript + vite2 + Element-plus技术</h3></div>


## 🌈平台简介

* 采用前后端分离的模式，后端采用GO语言，后端集成框架gin和go-restful（k8s使用的api框架）通过对GIn和go-restful自封装框架ginx,restfulx，代码更简洁，逻辑更清晰。另外拥抱云原生后期更新会采用go-restful框架做主要开发
* 根据不同分支选择使用的框架[gin分支](https://github.com/XM-GO/PandaX/tree/ginx) [go-restful分支](https://github.com/XM-GO/PandaX/tree/go-restful)
* 前端采用VUE3.0+ TypeScript + vite2 + Element-plus：[PandaUI](https://github.com/PandaGoAdmin/PandaUi)，适配手机、平板、pc 内置多种ui功能减少开发量
* 高效率的开发，使用代码生成器可以一键生成前后端代码，可在线预览代码，减少代码开发量。。
* 完善的权限认证系统：完善的权限认证系统，包含，菜单按钮权限，api权限，部门权限。
* 多数据库：项目同时支持MySQL，PostgreSql等数据库根据自身需求更改。

## 🏭在线体验

演示地址：http://101.35.247.125:7789/  帐号：admin 密码：123456  
组态大屏：http://101.35.247.125:7790/

---
系统在线文档
---
* 文档地址 ：http://www.pandazhuan.cn/

 **> 未来会补充文档和视频，方便友友们使用！** 

## 🚧系统截图

<table>
    <tr>
        <td><img src="https://s3.bmp.ovh/imgs/2021/12/26ce3214765103e8.png"/></td>
        <td><img src="https://s3.bmp.ovh/imgs/2021/12/20e0a825d40380d3.png"/></td>
    </tr>
    <tr>
        <td><img src="https://s3.bmp.ovh/imgs/2021/12/59840c8fe6fe1493.png"/></td>
        <td><img src="https://s3.bmp.ovh/imgs/2021/12/013cf70246e96e95.png"/></td>
    </tr>
	<tr>
        <td><img src="https://s3.bmp.ovh/imgs/2021/12/c0a5fa3d0a670fa7.png"/></td>
        <td><img src="https://s3.bmp.ovh/imgs/2021/12/80192c1976a7c14f.png"/></td>
    </tr>	 
    <tr>
        <td><img src="https://s3.bmp.ovh/imgs/2021/12/2dd2a25cb0f1a1df.png"/></td>
    </tr>

</table>
更多功能请访问系统体验

#### 💒 代码仓库

- PandaX 快速开发平台 <a href="https://github.com/PandaGoAdmin/PandaX" target="_blank">https://github.com/PandaGoAdmin/PandaX</a>
- PandaUI  平台Ui <a href="https://github.com/PandaGoAdmin/PandaUi" target="_blank">https://github.com/PandaGoAdmin/PandaUi</a>

## 联系我们
 **QQ：2417920382**  <a target="_blank" href="http://wpa.qq.com/msgrd?v=3&amp;uin=2417920382&amp;site=qq&amp;menu=yes">    点击这里给我发消息</a>
 
 **QQ群：467890197**  <a target="_blank" href="https://qm.qq.com/cgi-bin/qm/qr?k=BqzPxK0qWQEyI7YhnSfxc-GsAMlAIgta&jump_from=webapi"><img border="0" src="https://pub.idqqimg.com/wpa/images/group.png" alt="PandaX快速开发交流群" title="PandaX快速开发交流群"></a>
    <div align="center"><img src="https://s3.bmp.ovh/imgs/2021/12/f0f93d5622ef9e47.jpg"/></div>

## ⚡ 内置功能

1.  _用户管理：用户是系统操作者，该功能主要完成系统用户配置。_
2.  _部门管理：配置系统组织机构（公司、部门、小组），树结构展现支持数据权限。_
3.  _岗位管理：配置系统用户所属担任职务。_
4.  _菜单管理：配置系统菜单，操作权限，按钮权限标识等。_
5.  _角色管理：角色菜单,API权限分配、设置角色按机构进行数据范围权限划分。_
6.  _字典管理：对系统中经常使用的一些较为固定的数据进行维护。_
7.  _参数管理：对系统动态配置常用参数。_
8.  通知公告：系统通知公告信息发布维护。
9.  操作日志：系统正常操作日志记录和查询；系统异常信息日志记录和查询。
10. _登录日志：系统登录日志记录查询包含登录异常。_
11. 在线用户：当前系统中活跃用户状态监控。
12. _定时任务：在线（添加、修改、删除)任务调度包含执行结果日志。_
13. _代码生成：前后端代码的生成（go、vue、sql）支持CRUD下载 。_
14. 系统接口：根据业务代码自动生成相关的api接口文档。
15. _服务监控：监视当前系统CPU、内存、磁盘、堆栈等相关信息。_

---
前端工程结构
---

```
├── src
│   ├── api                  # Api ajax 等
│   ├── assets               # 本地静态资源
│   ├── i18n                 # 国际化
│   ├── components           # 业务通用组件
│   ├── layout               # layout
│   ├── theme                # css主题样式
│   ├── router               # Vue-Router
│   ├── store                # Vuex
│   ├── utils                # 工具库
│   ├── views                # 业务页面入口和常用模板
│   ├── App.vue              # Vue 模板入口
│   └── main.ts              # Vue 入口 TS
├── README.md
└── package.json
```

## 后端工程结构

| 项目 | 说明 |
| --- | --- |
| `base` | 自封装restfulx和工具类 |
| `docs` | api接口文档 |
| `initialize` | 初始化 |
| `resource` | 文件导出目录 |
| `static` | 前端代码构建 |
| `system` | 系统模块 |

更多功能请访问系统。

## 🍉 开发计划

* :clipboard: 代码生成器
* :clipboard: 资源文件管理中心
* :clipboard: 任务调度系统
* :clipboard: 监控系统
* :clipboard: 移动开发平台-基于uniapp
* :clipboard: 工作流
* :clipboard: 大屏系统
* :clipboard: 报表系统

  
## ❤特别鸣谢

  * 感谢[VUE-NEXT-ADMIN](https://github.com/lyt-Top/vue-next-admin)

---
版权说明
---

* PandaX快速开发平台采用Apache-2.0技术协议
* 代码可用于个人项目等接私活或企业项目脚手架使用，PandaX全系开源版完全免费
* 二次开发如用于商业性质或开源竞品请先联系群主审核
* 允许进行商用，但是不允许二次开源出来并进行收费
* 请不要删除和修改PandaX源码头部的版权与作者声明及出处
* 不得进行简单修改包装声称是自己的项目
* 我们已经申请了相关的软件开发著作权和相关登记
* 如有使用我们项目功能等的扩展项目，请在项目介绍中，进行明确说明

#### 💌 支持作者

如果觉得框架不错，或者已经在使用了，希望你可以去 <a target="_blank" href="https://github.com/PandaGoAdmin/PandaX">PandaX</a> 或者
<a target="_blank" href="https://github.com/PandaGoAdmin/PandaUi">PandaUi</a> 帮我点个 ⭐ Star，这将是对我极大的鼓励与支持。
