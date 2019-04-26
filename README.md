# go-api-scaffolding
学习是昂贵的，本项目的宗旨就是，会CTRL+C和CTRL+V就能写代码，哈哈哈。

## 项目简介
1. 基于bilibili开源的Kratos进行开发
2. 目标是开箱即用（拿来直接写业务逻辑）

## 运行环境
1. OS: WIN7 + 64位
2. Go: go1.12.4
3. 运行的相关依赖包200多MB，如果需要，请联系作者

## 执行步骤
1. 导入数据库文件 doc/api_scaffolding_v1_0_0.sql
2. 修改数据库用户名和密码 configs/mysql.toml
3. cd go-api-scaffolding/cmd
4. go build
5. cmd --conf ../configs
6. 打开浏览器，输入地址http://localhost:8000/go-api-scaffolding/start，会输出"Golang 大法好！！！"，说明成功！
7. postman发送POST请求![image](https://github.com/xiaohei2015/go-api-scaffolding/blob/master/doc/screenshot001.PNG)
    

## 性能比较
1. vs PHP7

## 联系作者
1. QQ: 306539332
2. 微信：xiaoheihacker