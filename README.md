# Agenda

> 课程《服务计算》作业七：用 Go 实现 Agenda 客户端和服务端 RESTful API 交互

[![Build Status](https://travis-ci.org/Mensu/Agenda-cli-service-Go.svg?branch=master)](https://travis-ci.org/Mensu/Agenda-cli-service-Go)

## 安装运行

```
docker build . -t agenda
# 服务端
docker run -dit --name agenda-sevice -v $PATH_TO_SERVER_DATA:/data -p 8080:8080 agenda service
# 客户端
docker run --rm --network host -v $PATH_TO_CLI_DATA:/data agenda cli help
```

## 注意

- 由于使用了 [cobra](github.com/spf13/cobra) 需要番羽~土啬才能编译运行
- 如果希望在屏幕 log，请**设置环境变量 DEBUG** 或**设置配置文件的 log 路径**
- 调试时通过设置环境变量 ``MOCK`` 访问 mock 服务器

## 客户端配置文件

默认使用 ``$HOME/.agenda-go.yaml``

如果找不到的话，将使用如下的默认设置

```yaml
# 工作目录。其他配置如果使用相对路径，则相对该工作目录
cwd: .
# log 的路径。如果环境变量里有 DEBUG 变量，则 log 会输出到 stderr
log: /dev/null
# 会话数据的路径。JSON 格式
curUser: data/curUser.json

```

## 项目管理与团队协作

### 团队协作

- 团队成员从 master 的 ``master`` 分支 fork 出[新的仓库](https://github.com/pfjhyyj/Agenda-cli-service-go)进行需求开发
- 团队成员开发完毕，向 master 的 ``master`` 分支发起 ``Pull Request``，并邀请 master ``review``
- master ``review`` 完觉得可以，且 ``CI`` 通过，方可确认归并代码
- master 作为 master 开发时，不得直接向 ``master`` 分支 push commit。而应该同样通过另开分支的方式进行需求开发。开发完毕后，向 master 的 ``master`` 分支发起 ``Pull Request``，并邀请团队成员 ``review``。同样，团队成员 ``review`` 完觉得可以，且 ``CI`` 通过，方可确认归并代码
- 以上限制通过设置 Github 完成，无需由团队成员假装限制

### 持续集成

- 使用 [``Travis CI``](https://travis-ci.org/Mensu/Agenda-Go)，通过执行 go test 命令运行编写好的测试文件进行持续集成
- 从最开始的开发开始，**边开发边写对应的测试**，在一次次提交的过程中不断集成，减少新的改动破坏原有功能的可能性，为项目功能的稳定提供有力保障

### TODO

## 架构设计与实现细节

### 客户端

学习初级实训 Agenda 的设计思路，我们使用的是三层架构

#### 表示层 ``cmd``

- 负责接受用户输入，交给*业务逻辑层*提供的业务逻辑服务，得到结果并展示给用户
- 使用 ``fmt`` 包向屏幕打印信息

#### 业务逻辑层 ``service``

- 负责简单的表单验证，并调用*实体层*提供的接口发送请求，获得需要的数据

#### 实体层 ``entity``

- 负责和服务端的进行交互，维护交互所需的数据结构，返回上层需要的数据

### 服务端

使用课程提供的服务端框架，分为 service - dao - orm (entity) 三层

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
- 会话的有效性由服务端控制。客户端将 ``openid`` 放到 ``cookie`` 里发到服务端来维持一段会话
