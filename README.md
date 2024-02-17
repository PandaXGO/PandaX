# PandaX

<div align="center"><img src="https://s3.bmp.ovh/imgs/2021/12/86b15968432fc6ea.png" width="100"/></div>
<div align="center"><h3 align="center">PandaX企业级物联网平台快速开发框架</h3></div>
<div align="center"><h3 align="center">基于Go 1.20前后端分离架构，代码精简，开箱即用，前端紧随前沿 Vue3.0 + TypeScript + vite3 + Element-plus技术</h3></div>


## 🌈平台简介

* 对前后端进行了大部分功能的封装，后端自封装go-restful，使用起来更加简洁，功能逻辑清晰，能快速上手学习，并用在生产中。
* 报表大屏设计器: 我们只需要拖拉拽即可绑定数据库，完成组态，报表和炫酷大屏的制作，无需要单独开发大屏。
* 成熟的规则引擎: 项目针对数据处理使用了规则链进行处理，简化开发及配置。
* 前端采用VUE3.0+ TypeScript + vite3 + Element-plus：[PandaUI](https://gitee.com/XM-GO/PandaUi)，适配手机、平板、pc 内置多种ui功能减少开发量
* 高效率的开发，使用代码生成器可以一键生成前后端代码，可在线预览代码，减少代码开发量。。
* 完善的权限认证系统：完善的权限认证系统，包含，菜单按钮权限，api权限，组织权限。
* 多数据库：项目同时支持MySQL，PostgreSql等数据库根据自身需求更改。

## 🏭在线体验

演示地址：http://101.35.247.125:7789/  帐号：admin 密码：123456  
组态大屏：http://101.35.247.125:7790/  
规则引擎：http://101.35.247.125:7791/

## Debian/Ubuntu系统快速部署测试环境

``` sh
git clone https://gitee.com/XM-GO/PandaX.git

cd PandaX

sudo ./startup.sh

```

---
系统在线文档
---
* 文档地址 ：http://101.35.247.125

 **> 未来会补充文档和视频，方便友友们使用！** 

## 🚧系统截图

<table>
    <tr>
        <td><img src="https://s3.bmp.ovh/imgs/2023/08/22/9b285c377717adc7.png"/></td>
        <td><img src="https://s3.bmp.ovh/imgs/2023/10/09/34ffa64e871f5264.png"/></td>
    </tr>
    <tr>
        <td><img src="https://s3.bmp.ovh/imgs/2023/10/09/4cea91a8e1dfe99b.png"/></td>
        <td><img src="https://s3.bmp.ovh/imgs/2023/10/09/0369e8ca0e71f0bb.png"/></td>
    </tr>
    <tr>
        <td><img src="https://s3.bmp.ovh/imgs/2023/10/09/c0a18770afc652c3.png"/></td>
        <td><img src="https://s3.bmp.ovh/imgs/2023/10/09/b8cf369ea64daf52.png"/></td>
    </tr>
    <tr>
        <td><img src="https://s3.bmp.ovh/imgs/2023/03/24/0f9a87733b5fe8da.png"/></td>
         <td><img src="https://s3.bmp.ovh/imgs/2023/08/22/58dda6cddceba5da.png"/></td>
    </tr>
</table>
更多功能请访问系统体验

## 联系我们
 **QQ：2417920382**  <a target="_blank" href="http://wpa.qq.com/msgrd?v=3&amp;uin=2417920382&amp;site=qq&amp;menu=yes">    点击这里给我发消息</a>
 
 **QQ群：467890197**  <a target="_blank" href="https://qm.qq.com/cgi-bin/qm/qr?k=BqzPxK0qWQEyI7YhnSfxc-GsAMlAIgta&jump_from=webapi"><img border="0" src="https://pub.idqqimg.com/wpa/images/group.png" alt="PandaX快速开发交流群" title="PandaX快速开发交流群"></a>

## ⚡ 内置功能

- <span class="tag done-tag">✔</span> **`用户管理`** - _用户是系统操作者，该功能主要完成系统用户配置。._
- <span class="tag done-tag">✔</span> **`组织管理`** - _配置系统组织机构（公司、组织、小组），树结构展现支持数据权限。_
- <span class="tag done-tag">✔</span> **`岗位管理`** - _配置系统用户所属担任职务。_
- <span class="tag done-tag">✔</span> **`菜单管理`** - _配置系统菜单，操作权限，按钮权限标识等。_
- <span class="tag done-tag">✔</span> **`角色管理`** - _角色菜单,API权限分配、设置角色按机构进行数据范围权限划分。_
- <span class="tag done-tag">✔</span> **`字典管理`** - _对系统中经常使用的一些较为固定的数据进行维护。_
- <span class="tag done-tag">✔</span> **`参数管理`** - _对系统动态配置常用参数。_
- <span class="tag done-tag">✔</span> **`通知公告`** - _系统通知公告信息发布维护_
- <span class="tag done-tag">✔</span> **`日志系统`** - _记录日志，更直观浏览_
- <span class="tag done-tag">✔</span> **`系统接口`** - _根据业务代码自动生成相关的api接口文档。_
- <span class="tag done-tag">✔</span> **`服务监控`** - _监视当前系统CPU、内存、磁盘、堆栈等相关信息。_
- <span class="tag done-tag">✔</span> **`代码生成`** - _可直接通过框架生成前后端基础业务代码（go、vue），减少开发时间。_
- <span class="tag done-tag">✔</span> **`组态大屏设计器`** - _通过拖拉拽直接生成组态、大屏。_
- <span class="tag done-tag">✔</span> **`规则链设计`** - _物联网规则链过滤_
- <span class="tag done-tag">✔</span> **`表单设计`** - _表单设计_
- <span class="tag done-tag">✔</span> **`报表设计`** - _数据报表设计_
- <span class="tag done-tag">✔</span> **`产品管理`** - _设备的产品管理_
- <span class="tag done-tag">✔</span> **`设备管理`** - _设备的管理,支持多协议接入，MQTT,TCP,UDP,COAP,Modbus,Opcua,S7,HL7等_

## 🛠 以后可能会有什么NB功能？
- <span class="tag wip-tag">开发中</span> **`3D组态(2023-Q4)`** - _根据2d组态自动生成3D组态_
- <span class="tag wip-tag">开发中</span> **`报表设计器（2024-Q1）`** - _excel的低代码报表设计_

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

|     目录     | 功能                                   |
|:----------:|:-------------------------------------|
|  `deploy`  | 部署文件，本项目部署是利用`K3S`进行部署的，因此里面的文档为部署文档 |
|   `apps`   | 基本功能,所有功能模块全在这里面                     |
|  `iothub`  | 设备接入层，设备数据上报在这里处理，使用emqx的hook模式      |
| `resource` | 项目启动或生成的资源文件存放目录。                    |
|   `pkg`    | 所有开发过程中的全局通用代码。                      |
| `uploads`  | 存储上传的文件的地方                           |

更多功能请访问系统。

---
版权说明
---

* PandaX物联网低代码开发基座采用AGPL-3.0技术协议
* PandaX代码完全开源，可用于个人学习交流使用,
* 不允许商业使用,如需商业使用请联系作者。
* 不得进行简单修改包装声称是自己的项目
* 我们已经申请了相关的软件开发著作权和相关登记
* 如有使用我们项目功能等的扩展项目，请在项目介绍中，进行明确说明

#### 💌 支持作者

如果觉得框架不错，或者已经在使用了，希望你可以去 <a target="_blank" href="https://gitee.com/XM-GO/PandaX">PandaX</a> 或者
<a target="_blank" href="https://gitee.com/XM-GO/PandaUi">PandaUi</a> 帮我点个 ⭐ Star，这将是对我极大的鼓励与支持。
