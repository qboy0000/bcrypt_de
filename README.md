# bcrypt_de

> 对bcrypt进行密码爆破

## 编译
```
go mod tidy
go mod download
go build
```

## 运行
```
./bcrypt_de $2a$10$UN7bgaqA4F01hKV2l9Da0.T8LnnD6TXqLSvRwASd0IRI0cH6Kx58i ../password_dic.tx
```
> 第一个参数为字典，第二个参数为爆破使用的字典库
