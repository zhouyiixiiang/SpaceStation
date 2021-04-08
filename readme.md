# 报到系统流程梳理

## 会议信息管理

管理员

* 用户角色
* 密码
* 

大会信息

* 编号
* 会议名称
* 开始时间
* 结束时间
* 管理员组编号
* 备注

会议信息

* 编号

* 大会编号
* 子会议信息
  * 序号
  * 子会议名称
  * 是否判断迟到
  * 迟到时间
* 报到规则
  * 是否设置会议人数线
  * 规则人员类型
    * 委员
    * 所有类型
  * 规则（>或>=）
  * 规则按人数设置或比例设置
  * 会议人数线
  * 按比例判断
    * 按照委员或所有类型的
    * 比例
* 达标文字

分组信息

* 编号

* 大会编号
* 子会议编号
* 分组详细信息

人员信息

* 人员基本信息
  * 编号
  * 大会编号
  * 子会议编号
  * 姓名
  * 界别
  * 类型
    * 工作人员
    * 委员
    * 陪同人员...
  * 级别
  * 性别
  * 民族
  * 坐席
  * 单位
  * 联系方式
  * 照片名称
  * 主席团成员
    * 是\否
  * 职务
  * 卡片编号
  * 备注
  * 关注
  * 分组信息
    * 分组信息下标
    * 分组表编号

写卡软件（卡片信息）

* 编号
* 人员编号
* 卡片信息
  * 主卡号
  * 备卡号
  * 测试卡号

会议人员信息表

* 编号
* 子会议编号
* 人员编号
* 座位编号
  * 行
  * 列

## 会议进行中

会议报到信息表

* 编号
* 子会议编号
* 人员编号
* 报到时间
* 报到状态

# 数据中心







# PCMDataCenter

## GO Mod使用

```
cd project
go env -w GO111MODULE=on
go mod init project
```

此时会在project目录下生成go.mod，文件内容以project为根目录

go build 会自动查找文件依赖并下载

## 环境搭建

安装mysql：

解压mysql安装包到想要安装的路径下

```
# 以管理员形式打开cmd，cd到mysql的安装路径下的bin文件下
# 初始化数据库
mysqld --initialize --console
# 记录下初始化数据库后的随机初始密码，我是 XSMOd)5DwuHY
# 安装mysql为windows的服务
mysqld -install
# 启动mysql服务
net start mysql
# 在mysql\bin 目录下，使用以下指令登录mysql
mysql -u root -p
# 成功登录后修改密码
alter user 'root'@'localhost' identified by '设置密码';
commit;
# 将mysql\bin目录添加到电脑系统环境变量，方便下次使用
```

## 安装vue

### 安装nodejs

https://www.cnblogs.com/coder-lzh/p/9232192.html

### 安装vue

https://www.jianshu.com/p/e215a33aed40

## Vue使用

```
# 新建一个项目
vue init webpack my-project
# 开始
npm install
npm run dev
# vue格式化全文
Alt+shift+F
```

## Vue安装依赖

* 修改config/index.js目录下，配置跨域

```
dev: {
    // Various Dev Server settings
    host: 'localhost', // can be overwritten by process.env.HOST
    port: 8080, // can be overwritten by process.env.PORT, if port is in use, a free one will be determined

    // Paths
    assetsSubDirectory: 'static',
    assetsPublicPath: '/',
    proxyTable: {
      '/backend':{
        target: "http://127.0.0.1:8000",
        ws: true,
        changeOrigin: true,
        pathRewrite: {
          '^/backend': '/backend' //这里理解成用‘/api’代替target里面的地址，组件中我们调接口时直接用/api代替
              // 比如我要调用'http://0.0:300/user/add'，直接写‘/api/user/add’即可 代理后地址栏显示
        }
      }
    },
```

* 安装axios与element

```
npm install --save axios
npm install element-ui -S
```

* 关闭eslint检查，在build/webpack.base.conf.js，注释掉下述最后一行

```
 module: {
    rules: [
      //...(config.dev.useEslint ? [createLintingRule()] : []),
```



## Vue结构目录

* build 项目构建（webpack）相关代码
* config ：配置目录
  * 端口号
* node_modules
  * npm加载的项目依赖块
* src：开发目录
  * assets：放置图片
  * components：组件文件，可以不用
  * App.vue：项目入口文件，可以把组件写在这里，就不用components目录
  * main.js 项目核心文件
* static 静态资源目录：图片、字体
* test： 初始化目录，可删除
* .xxxx文件：配置文件，包括语法配置、git配置等
* index.html 首页入口文件，可添加一些meta信息或同级代码
* package.json 项目配置文件

### vue后台管理模板

标星最多：https://github.com/herozhou/vue-framework-wz

可以参考：https://github.com/bailicangdu/vue2-manage

具有所有组件参考：https://github.com/lin-xin/vue-manage-system

# golang配置

一些中间件

* 保持登录状态
  * 查看singo项目中的r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
* 跨域Cores
* 验证登录中间件
* 

# rfid

### 结构：

* 读卡器硬件层
  * usb、rs232、lan
* 设备驱动层
  * rfidlib_G301.dll、rfidlib_G302.dll
* 设备通用函数库
  * rfidlib_reader.dll
* 会议签到门应用

### 总体调用：

* 设备驱动库文件拷贝到与rfidlib_reader.dll同级目录
* 调用RDR_LoadReaderDrivers加载驱动到应用程序
  * 加载目录下的所有驱动
* 使用RDR_Open打开设备驱动
  * 传入该驱动的连接字符串
  * 打开通信接口并创建对应驱动实例
    * 然后所有操作都需传入该实例句柄

### 获取人员进出：

* 原理
  * 在库里面有创建线程，使用RDR_BuffMode_FetchRecords，取到记录则发送到RFIDLIB_EVNT_GETBUF_REPORT_FOUND事件，如有错误，发送错误事件
* 调用方法
  * 创建线程，循环调用RDR_BuffMode_FetchRecords
    * 如果成功
      * RDR_GetTagDataReport(SEEK_FIRST)读取第一条进出报告
      * 读取成功：RDR_ParseTagDataReportRow解析报告数据
      * RDR_GetTagDataReport(SEEK_NEXT)
    * 读完之后，返回

### 关键点解释：签到数据解析

RDR_ParseTagDataReportRow返回记录的原始数据，包含多个字段

* 记录头部
  * 事件类型，BYTE，保留未用
  * 进出方向，BYTE，1、进；2、出
  * 时间，BYTE[6]年月日时分秒的BCD码
    * B[0]：2位年，如2016年是0x16
    * B[1]：月......B[5]秒
  * 卡片数据长度，BYTE，后面包含的卡片内容长度
* 卡片数据
  * 卡片数据内容，BYTE，目前只包含8个字节的UID

### 读取已有配置

```
	iret  = RDR_ConfigBlockRead(hr,0,cfgblock,m_size)  ;
	if(iret != 0) {
		LogCfgMsg(_T("failed in reading block 0")) ;
		goto exit_func;
	}
	m_cfgIPAddr.SetAddress(cfgblock[0],cfgblock[1],cfgblock[2],cfgblock[3]) ;
	m_cfgIPMask.SetAddress(cfgblock[4] ,cfgblock[5] ,cfgblock[6],cfgblock[7]) ;
	iret  = RDR_ConfigBlockRead(hr,1,cfgblock,m_size)  ;
	if(iret != 0) {
		LogCfgMsg(_T("failed in reading block 1")) ;
		goto exit_func;
	}
	m_cfgIPGW.SetAddress(cfgblock[0],cfgblock[1],cfgblock[2],cfgblock[3]) ;
```



### 关闭设备驱动

RDR_Close

### 串口连接串说明

```
"RDType=RD201;
CommType=COM;COMName=COM1;BaudRate=38400;Frame=8E1;BusAddr=255"
```

* RDType：设备驱动的类型，可传入名称或型号，通过查看相关设备驱动说明
* CommType：通信接口类型
  * COM，USB，NET，BLUETOOTH
* COMName：串口名称
  * COM1，COM2
* BaudRate：波特率
* Frame：帧结构
  * 8E1，8位数据位，偶校验，1位停止位
* BusAddr：RS485总线地址
  * 1-254，255位广播地址

### 网络接口连接串说明



# 程序整体架构

语言：c++、Golang、VUE

结构：

* 数据中心
  * BS
  * CS
* 主控
  * 主控前端
  * 数据中间件
  * 会议信息配置
* 报到机
  * RFID
  * 人脸识别

流程：

* 数据中心：
  * 配置会议基本信息、配置人员基本信息、配置参会信息
* 主控：
  * 从数据中心获取会议、人员配置信息。绘制编辑模板
* 数据中间件：
  * 控制下游报到机、将报到结果显示至上游主控界面
  * 在报到过程中，将报道数据进行缓存和整合处理
* 报到机：
  * 记录每个站点的报到数据，配置签到门的工作模式

优点：

* 导入数据只需从前端一次配置，不需要反复从一个软件导入另一个软件
* 支持多链路主控，提高报到数据的安全性、报到过程中的稳定性，任意台主控失效，其他台主控可以继续进行报到
* 报到机工作性能大幅提高
* 支持签退、人脸识别、网络摄像头等功能

应用技术：

* 微服务集群式管理，打破BS、CS架构的局限性，多个集群服务间通过socket通信进行数据同步，有断点异常处理，即使socket通信中断，通信恢复后可以及时同步报到数据，形成安全可靠的可扩展式报到服务架构。

## mysql导入导出

```
mysqldump -h localhost -uroot -p12345678  -d data_center sign_in_info > D:\database\dump.sql;
source d:\database\dump.sql
```

## mysql导入整个数据库

```
mysqldump -h localhost -uroot -p data_center > D:\database\data_center.sql
source d:\database\dump.sql
```

## mysql安装

```
mysqld --install
mysqld --initialize --console
net start mysql

mysql -u root -p
粘贴生成的密码
alter user 'root'@'localhost' identified by '12345678';
```



## mysql允许用户远程连接

```
CREATE USER 'root'@'%' IDENTIFIED BY '12345678';
grant all privileges on *.* to 'root'@'%' with grant option;
flush privileges;
```



## 工作人员联系方式

电话：13255668132朱经理 Y型签到门设备 旧款报到门

微信：Liang 何工   A型签到门设备 新型报到门

电话：莫工 A型报到门

## opencv

较全的编译教程：

https://blog.csdn.net/qq_29191321/article/details/88699933

动态库调用教程：

https://blog.csdn.net/zhjinw/article/details/79628632

cgo:

https://blog.csdn.net/u010884123/article/details/60872980

# vs2010

$(SolutionDir)$(Configuration)\