# Agenda

> 课程《服务计算》作业三：用 Go 实现命令行 Agenda

[![Build Status](https://travis-ci.org/Mensu/Agenda-Go.svg?branch=master)](https://travis-ci.org/Mensu/Agenda-Go)

## 安装运行

```
go get github.com/mensu/Agenda-Go
$GOPATH/bin/Agenda-Go
```

## 注意

- 由于使用了 [cobra](github.com/spf13/cobra) 需要番羽~土啬才能编译运行
- 如果希望在屏幕 log，请**设置环境变量 DEBUG** 或**设置配置文件的 log 路径**

## 配置文件

默认使用 ``$HOME/.agenda-go.yaml``

如果找不到的话，将使用如下的默认设置

```yaml
# 工作目录。其他配置如果使用相对路径，则相对该工作目录
cwd: .
# log 的路径。如果环境变量里有 DEBUG 变量，则 log 会输出到 stderr
log: /dev/null
# 用户数据的路径。JSON 格式
user_data: data/user_data.json
# 会议数据的路径。JSON 格式
meeting_data: data/meeting_data.json
# 会话数据的路径。JSON 格式
curUser: data/curUser.json

```

## 项目管理与团队协作

### 团队成员

- 陈宇翔 15331042 (master)
  + 项目框架和架构设计
  + 负责 JSON 读写的 ``storage``、日志 ``log``、异常处理 ``error`` 等工具
  + 用户实体、服务、UI
  + 需求：用户注册、登录、登出
  + 接入持续集成
  + README 工程师
- 杨奕嘉 15331366
  + 命令设计
  + 需求：用户查询、删除
  + 需求：会议创建、增删参与者、查询
  + 需求：会议取消、退出、清空
  + 会议实体、服务、UI
  + 用户和会议相关的测试代码

### 团队协作

- 杨奕嘉从陈宇翔的 ``master`` 分支 fork 出[新的仓库](https://github.com/pfjhyyj/Agenda-Go)进行需求开发
- 杨奕嘉开发完毕，向陈宇翔的 ``master`` 分支发起 ``Pull Request``，并邀请陈宇翔 ``review``
- 陈宇翔 ``review`` 完觉得可以，且 ``CI`` 通过，方可确认归并代码
- 陈宇翔作为 master 开发时，不得直接向 ``master`` 分支 push commit。而应该同样通过另开分支的方式进行需求开发。开发完毕后，向陈宇翔的 ``master`` 分支发起 ``Pull Request``，并邀请杨奕嘉 ``review``。同样，杨奕嘉 ``review`` 完觉得可以，且 ``CI`` 通过，方可确认归并代码
- 以上限制通过设置 Github 完成，无需由团队成员假装限制

### 持续集成

- 使用 [``Travis CI``](https://travis-ci.org/Mensu/Agenda-Go)，通过执行 go test 命令运行编写好的测试文件进行持续集成
- 从最开始的开发开始，**边开发边写对应的测试**，在一次次提交的过程中不断集成，减少新的改动破坏原有功能的可能性，为项目功能的稳定提供有力保障

### TODO

#### 整体需求

- [x] 业务需求：用户注册 ``register``
- [x] 业务需求：用户登录 ``login``
- [x] 业务需求：用户登出 ``logout``
- [x] 业务需求：用户查询 ``list``
- [x] 业务需求：用户删除 ``delete-account``
- [x] 业务需求：创建会议 ``add-meeting``
- [x] 业务需求：增删会议参与者 ``add-participator`` ``delete-participator``
- [x] 业务需求：查询会议 ``show``
- [x] 业务需求：取消会议 ``cancel-meeting``
- [x] 业务需求：退出会议 ``quit-meeting``
- [x] 业务需求：清空会议 ``clear``
- [x] 功能需求： 设计一组命令完成 agenda 的管理
- [x] 持久化要求：使用 json 存储 User 和 Meeting 实体
- [x] 持久化要求：当前用户信息存储在 curUser.txt 中
- [x] 开发需求：团队 **2-3 人**，一人作为 master 创建程序框架，其他人 fork 该项目，所有人同时开发。团队**不能少于 2 人**
- [x] 日志服务：使用 [log](https://go-zh.org/pkg/log/) 包记录命令执行情况

#### 第一周任务

- [x] 按 3.3 安装 cobra 并完成小案例
- [x] 按需求设计 agenda 的命令与参数（制品 cmd-design.md）
- [x] master 创建项目，提交到 github， 其他人 fork 该项目
- [x] 每人分别创建属于自己的命令（命令实现 Print 读取的参数即可），提交并归并。确保不同人管理不同文件，以便于协作
- [x] 如时间富余，请完成 User 和 Meeting 实体 json 文件读写

#### 第二周任务

- [x] 在项目中添加 ``.travis.yml`` 文件，并添加测试程序。让你的项目“持续集成” -- “CI” 了！
- [x] 添加 log 服务，记录用户的操作过程，以及关键的输出
- [x] 约定 ``entity`` 和 ``cmd`` 之间的接口服务，实现 agenda，并在 ``README.md`` 文件中给出简要使用说明和测试结果
- [ ] （如果你有兴趣），使用 ``pflag`` 包，自己实现一个简版 ``Command.go`` 取代 ``cobra``。必须使用组合设计模式

## 架构设计与实现细节

### 三层架构

学习初级实训 Agenda 的设计思路，我们使用的是三层架构

#### 表示层 ``cmd``

- 负责接受用户输入，交给*业务逻辑层*提供的业务逻辑服务，得到结果并展示给用户
- 使用 ``fmt`` 包向屏幕打印信息

#### 业务逻辑层 ``service``

- 负责具体的业务逻辑，通过调用*实体层*提供的接口操纵数据，完成业务逻辑
- 负责数据的逻辑合法性验证

#### 实体层 ``entity``

- 负责提供直接操纵数据实体的接口以及持久化储存
- 不负责数据的逻辑合法性，认为上层所有的数据操作都是合法的，从而专注于数据实体的操作。这也是分层的意义所在

### 实体加载

- 实体使用 ``JSON`` 格式储存
- 数据的加载和储存由专门的 ``storage`` 结构完成，被负责数据操纵的 ``*Model`` 结构使用
- 各 ``*Model`` 应该只有一个实例，因此考虑使用**单例模式**。这里使用包全局变量的方式实现
- 各 ``*Model`` 的加载涉及 IO 操作，数据量大相对会比较耗时。又考虑到各 ``*Model`` 加载的独立性，于是我们通过 ``goroutine`` 实现并行加载

### log

- 每个层通过**工厂模式**产生并使用不同的 log 实例，方便统一设定 log 需要的属性，如输出的目的地、输出格式等。以后换 [``logrus``](https://github.com/Sirupsen/logrus) 包时也方便添加前缀
- 在表示层、业务逻辑层、实体层的关键地方记录用户的操作过程和输出

### 异常处理

- 对于因用户输入不当的逻辑错误，使用 ``go`` 语言标准的返回错误的模式，并提示用户
- 对于其他不可抗力产生的异常，如文件写入失败等，使用 ``Panic`` 函数生成函数栈并抛出，在每个 ``goroutine`` 最外面的恢复并记录至 ``log`` 中，为调试提供有力的线索

### 登录控制

- 使用 ``session`` 表示会话，控制登录状态。持久化于 ``curUser.json`` 文件中
- 目前会话不过期，除非用户手动登出。以后可以考虑像浏览器的会话控制一样，增加会话过期时间
