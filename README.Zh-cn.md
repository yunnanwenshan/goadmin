# go-admin
  <img align="right" width="320" src="https://doc-image.zhangwj.com/img/go-admin.svg">

[![Build Status](https://github.com/wenjianzhang/go-admin/workflows/build/badge.svg)](https://github.com/go-admin-team/go-admin)
[![Release](https://img.shields.io/github/release/go-admin-team/go-admin.svg?style=flat-square)](https://github.com/go-admin-team/go-admin/releases)
[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/go-admin-team/go-admin)
1. 接口文档：根据业务代码自动生成相关的api接口文档。
1. 代码生成：根据数据表结构生成对应的增删改查相对应业务，全程可视化操作，让基本业务可以零代码实现。
1. 表单构建：自定义页面样式，拖拉拽实现页面布局。

## 📦 本地开发

### 环境要求

go 1.18

node版本: v14.16.0

npm版本: 6.14.11

### 开发目录创建

```bash

# 创建开发目录
mkdir goadmin
cd goadmin
```

### 获取代码

> 重点注意：两个项目必须放在同一文件夹下；

```bash
# 获取后端代码
git clone https://github.com/go-admin-team/go-admin.git

# 获取前端代码
git clone https://github.com/go-admin-team/go-admin-ui.git

```

### 启动说明

#### 服务端启动说明

```bash
# 进入 go-admin 后端项目
cd ./go-admin

# 更新整理依赖
go mod tidy

# 编译项目
go build

# 修改配置 
# 文件路径  go-admin/config/settings.yml
vi ./config/settings.yml

# 1. 配置文件中修改数据库信息 
# 注意: settings.database 下对应的配置数据
# 2. 确认log路径
```

⚠️注意 在windows环境如果没有安装中CGO，会出现这个问题；

```bash
E:\go-admin>go build
# github.com/mattn/go-sqlite3
cgo: exec /missing-cc: exec: "/missing-cc": file does not exist
```

or

```bash
D:\Code\go-admin>go build
# github.com/mattn/go-sqlite3
cgo: exec gcc: exec: "gcc": executable file not found in %PATH%
```

[解决cgo问题进入](https://doc.go-admin.dev/zh-CN/guide/faq#cgo-%E7%9A%84%E9%97%AE%E9%A2%98)


#### 初始化数据库，以及服务启动

``` bash
# 首次配置需要初始化数据库资源信息
# macOS or linux 下使用
$ ./go-admin migrate -c config/settings.dev.yml

# ⚠️注意:windows 下使用
$ go-admin.exe migrate -c config/settings.dev.yml


# 启动项目，也可以用IDE进行调试
# macOS or linux 下使用
$ ./go-admin server -c config/settings.yml


# ⚠️注意:windows 下使用
$ go-admin.exe server -c config/settings.yml
```

#### sys_api 表的数据如何添加

在项目启动时，使用`-a true` 系统会自动添加缺少的接口数据
```bash
./go-admin server -c config/settings.yml -a true
```

#### 使用docker 编译启动

```shell
# 编译镜像
docker build -t go-admin .

# 启动容器，第一个go-admin是容器名字，第二个go-admin是镜像名称
# -v 映射配置文件 本地路径：容器路径
docker run --name go-admin -p 8000:8000 -v /config/settings.yml:/config/settings.yml -d go-admin-server
```

#### 文档生成

```bash
go generate
```

#### 交叉编译

```bash
# windows
env GOOS=windows GOARCH=amd64 go build main.go

# or
# linux
env GOOS=linux GOARCH=amd64 go build main.go
```
   <tr>
    <td><img src="https://raw.githubusercontent.com/wenjianzhang/image/master/img/wx.png" width="180px"></td>
    <td><img src="https://doc-image.zhangwj.com/img/qrcode_for_gh_b798dc7db30c_258.jpg" width="180px"></td>
    <td><img src="https://raw.githubusercontent.com/wenjianzhang/image/master/img/qq2.png" width="200px"></td>
    <td><a href="https://space.bilibili.com/565616721">wenjianzhang</a></td>
  </tr>
  <tr>
    <td>微信</td>
    <td>公众号🔥🔥🔥</td>
    <td><a target="_blank" href="https://shang.qq.com/wpa/qunwpa?idkey=0f2bf59f5f2edec6a4550c364242c0641f870aa328e468c4ee4b7dbfb392627b"><img border="0" src="https://pub.idqqimg.com/wpa/images/group.png" alt="go-admin技术交流乙号" title="go-admin技术交流乙号"></a></td>
    <td>哔哩哔哩🔥🔥🔥</td>
  </tr>
</table>

## 💎 贡献者


<span style="margin: 0 5px;" ><a href="https://github.com/wenjianzhang" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/3890175?v=4&h=60&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/G-Akiraka" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/45746659?s=64&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/lwnmengjing" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/12806223?s=64&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/bing127" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/31166183?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/chengxiao" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/1379545?s=64&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/NightFire0307" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/19854086?v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/appleboy" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/21979?s=64&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/ninstein" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/580303?v=4&h=60&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/kikiyou" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/17959053?s=64&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/horizonzy" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/22524871?s=64&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/Cassuis" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/48005724?s=64&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/hqcchina" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/5179057?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/nodece" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/16235121?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/stephenzhang0713" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/18169290?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/zhouxixi-dev" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/100399679?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/Jalins" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/31172582?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/wkf928592" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/6063351?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/wxxiong6" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/6983441?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/Silicon-He" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/52478309?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/GizmoOAO" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/20385106?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/bestgopher" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/36840497?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/wxb1207" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/20775558?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/misakichan" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/16569274?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/zhuxuyang" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/19301024?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/mss-boot" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/109259065?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/AuroraV" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/37330199?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/Vingurzhou" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/57127283?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/haimait" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/40926384?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/zyd" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/3446278?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/infnan" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/38274826?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/d1y" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/45585937?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/qlijin" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/515900?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/logtous
" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/88697234?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/stepway
" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/9927079?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/NaturalGao
" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/43291304?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/DemoLiang
" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/23476007?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/jfcg
" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/1410597?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>
<span style="margin: 0 5px;" ><a href="https://github.com/Nicole0724
" ><img src="https://images.weserv.nl/?url=avatars.githubusercontent.com/u/10487328?s=60&v=4&w=60&fit=cover&mask=circle&maxage=7d" /></a></span>

## JetBrains 开源证书支持

`go-admin` 项目一直以来都是在 JetBrains 公司旗下的 GoLand 集成开发环境中进行开发，基于 **free JetBrains Open Source license(s)** 正版免费授权，在此表达我的谢意。

<a href="https://www.jetbrains.com/?from=kubeadm-ha" target="_blank"><img src="https://raw.githubusercontent.com/panjf2000/illustrations/master/jetbrains/jetbrains-variant-4.png" width="250" align="middle"/></a>

## 🤝 特别感谢

1. [ant-design](https://github.com/ant-design/ant-design)
2. [ant-design-pro](https://github.com/ant-design/ant-design-pro)
2. [arco-design](https://github.com/arco-design/arco-design)
2. [arco-design-pro](https://github.com/arco-design/arco-design-pro)
4. [gin](https://github.com/gin-gonic/gin)
5. [casbin](https://github.com/casbin/casbin)
6. [spf13/viper](https://github.com/spf13/viper)
7. [gorm](https://github.com/jinzhu/gorm)
8. [gin-swagger](https://github.com/swaggo/gin-swagger)
9. [jwt-go](https://github.com/dgrijalva/jwt-go)
10. [vue-element-admin](https://github.com/PanJiaChen/vue-element-admin)
11. [ruoyi-vue](https://gitee.com/y_project/RuoYi-Vue)
12. [form-generator](https://github.com/JakHuang/form-generator)


## 🤟 打赏

> 如果你觉得这个项目帮助到了你，你可以帮作者买一杯果汁表示鼓励 :tropical_drink:

<img class="no-margin" src="https://raw.githubusercontent.com/wenjianzhang/image/master/img/pay.png"  height="200px" >

## 🤝 链接

[Go开发者成长线路图](http://www.golangroadmap.com/)

## 🔑 License

[MIT](https://github.com/go-admin-team/go-admin/blob/master/LICENSE.md)

Copyright (c) 2024 wenjianzhang
