## word  

```
# go run main.go help word
该子命令支持各种单词格式转换，模式如下：
1：全部转大写
2：全部转小写
3：下划线转大写驼峰
4：下划线转小写驼峰
5：驼峰转下划线

Usage:
   word [flags]

Flags:
  -h, --help         help for word
  -m, --mode int8    请输入单词转换的模式
  -s, --str string   请输入单词内容

```

## time  

```
# go run main.go time now
2020/07/17 17:44:23 输出结果:2020-07-17 17:44:23,1594979063

```

## sql2struct  

`struct`拼写有误,需要修改.  

```
# go run main.go sql struct --username 数据库的账号 --password 数据库的密码 --db=数据库的名称 --table "需要转换的表名"

```