#   Filcoin 同步链上交易

1. 修改 conf/apptemp.ini文件重命名为app.ini
2. 修改app.ini文件 数据库配置和自己搭建的rpc节点配置
- [x] USER = root
- [x] PASSWORD =831015
- [x] HOST = 127.0.0.1:3306
- [x] NAME = filecoin_node_api
- [x] SERVER_HOST = http://xx.xx.xx.xx:1234/rpc/v0
- [x] TOKEN = xxx.xxxx

3.新建数据库 filecoin_node_api 表是自动生成

4.go run main.go 运行即可




5.表说明
表 | 字段说明
---|---
sys_wallets          | 总收款账户,付款账户，当前同步块，起始块高度
sys_fil_transations   | 交易记录,单纯的同步链上交易,不参与业务操作
sys_fil_integrals            |每同步一笔交易会关联到这个表中,主要负责处理和交易有关的业务 比如到账后加减积分等
sys_auths | 接口采用token账号授权+rsa公钥匙私钥加密方式访问 这里存放的是授权账号



6.接口安全rsa公钥加密 私钥解密


```
sequenceDiagram
A->>B: 请求参数公钥加密
B->>B: 私钥解析加密参数
B->>A: 返回结果
```





7.生成公私钥

- [x]  调用rsa.go 中的GenerateKey方法生成公私钥文件
- [x]  打开公私钥文件分别复制替换rsa.go中的privatekey 和publickey




 8.支持接口:
~~~~


授权token: ip:port/auth
1.数据库sys_auths表配置账号密码
2. 授权账号是：username : xxxx  password:xxxxxx 用这个账号去请求令牌token 拿到token放到head头部  key 是：token
3. 加解密已封装在rsa.go中: 
    加密参数：util.Decrypt
    解密参数：util.Encrypt
 


4.示例：
http://127.0.0.1:8085/auth?username=加密(xxx)&password=加密(xxxxxx)
返回结果：
{
"code": 0,
"data": {
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJwYXNzd29yZCI6InRlc3QxMjM0NTYiLCJleHAiOjE2MDY1ODcwOTUsImlzcyI6Imdpbi1ibG9nIn0.z-B-rBto3fAYzjaek4fR5D8sfRGV9HtH0LoPT4vtcKE"
},
"msg": "ok"
}
安全期间增加了token令牌请求 和rsa加密方式 调用端只需要用公钥 加密 参数即可
 
 下面这些都是要带上上面这种token放到头请求
 

 获取链记录： ip:port/api/v1/getRawTransaction/:cid 
 
 转账： ip:port/api/v1/pay/sendtoaddress/:from/:to/:amount
 获取账号余额：ip:port/api/v1/getwalletinfo/:account
 检测FIL地址有效性：ip:port/api/v1/validateAddress/:address
 创建新的FIL账号：ip:port/api/v1/newaccount    

 资金归集到入金账号：ip:port/api/v1/pay/sendtowallet/:address/:amount
 获取最新区块高度：ip:port/api/v1/getlastHeight 
 获取交易确认数：ip:port/api/v1/getComfirm/:txid 
 判断地址是不是钱包地址：ip:port/api/v1/isnodeaddres/:address

 根据状态检索一笔交易  Ip:port/api/v1/getOneTranstion/:state  
 修改一笔交易额状态  Ip:port/api/v1/changeIntegralState/:txid/:state  
（ 0:代表刚同步到的新交易，1:代表你那边已经处理完成加积分）

 提币：/api/v1/pay/sendfromoutaccount/:to/:amount 
（统一从出金账户提币,需要平台冲币到这个账户足够的钱  出金账户,入金账户数据库初始化时候自动生成的，这里只需要填to账户和金额即可）
另外转账失败等ui显示问题 请自行处理error封装   这边服务全部是原生链上操作不做error 封装处理 不然链上出错，难调试






~~~~


