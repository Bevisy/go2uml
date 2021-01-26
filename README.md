# go2uml

## 命令详情
```shell script
$ ./go2uml --help
Usage:
  go2uml [OPTIONS]

Application Options:
      --codedir=   要扫描的代码目录
      --gopath=    GOPATH directory
      --out=       解析结果保存到该文件中
      --ignoredir= 需要排除的目录,不需要扫描和解析

Help Options:
  -h, --help       Show this help message
```

## 配合 Goland 使用
新增 Tool :
```shell script
Program: go2uml
Parameters: --gopath $GOPATH$ --codedir $FileDir$ --outputfile $FileDir$.puml
```

# 备注
Fork from [go-package-plantuml](http://git.oschina.net/jscode/go-package-plantuml)
