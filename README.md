# GoFrame Template For SingleRepo

Quick Start: 
- https://goframe.org/pages/viewpage.action?pageId=1114399

/
├── api
    └─hello              // 模块
        └─v1               // 模块版本
         ├── hello.go        // 模块对象中方法的出参和入参定义
        ├── hello.go       // 模块对象定义
├── hack              // 存放项目开发工具、脚本等内容
     ├── config.yaml    // 配置文件
├── internal
│   ├── cmd   
         ├── cmd.go      // 1. 接口路由入口
│   ├── consts       // 项目所有常量定义
│   ├── controller
            └─hello                  // 2. 模块路由实现
                ├── hello_new.go         // 2.1. 模块对象创建
                ├── hello_v1_hello.go    // 2.2. 模块对象中方法的具体实现
                ├── hello.go             // 2.3. 
│   ├── dao
│   ├── logic
│   ├── model
│   |   ├── do
│   │   └── entity
│   └── service
├── manifest
├── resource
├── utility
├── go.mod  // 使用Go Module包管理的依赖描述文件
└── main.go // 程序入口文件