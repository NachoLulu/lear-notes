# 简易许可证校验

## 颁发流程

### 1.生成根证书私钥（.key）
使用openssl命令生成私钥
```shell
openssl genrsa -out ca.key 2048
```

### 2.生成根证书请求文件（.csr）
证书请求文件即为为生成根证书提供的记录信息
```shell
openssl req -new -key ca.key -out ca.csr -subj "/CN=China/ST=Hubei/L=Wuhan/O=Qianxin/OU=CA"
```

### 3.生成根证书（.crt）
根据根证书和根证书请求文件生成根证书
```shell
openssl x509 -req -in ca.csr -out ca.crt -signkey ca.key -days 3650
```

### 4.生成许可证私钥（.key）
```shell
openssl genrsa -out license.key 2048
```

### 5.生成许可证请求文件（.csr）
```shell
openssl req -new -key license.key -out license.csr -subj "/CN=China/ST=Hubei/L=Wuhan/O=Qianxin/OU=License"
```

### 6.生成许可证证书（.crt）
```shell
openssl x509 -req -in license.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out license.crt -days 365
```

### 7.查看证书信息
```shell
openssl x509 -in license.crt -noout -serial -dates -subject
```