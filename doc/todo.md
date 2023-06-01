## TODO    
- [ ] jwt，看语雀文档
- [ ] api网关
- [ ] dig换成wire
- [ ] etcd
- [ ] zap日志包
- [ ] grpc负载均衡和etcd还没完成
- [ ] 完善xzgrpc
- [ ] 完善xzetcd
- [ ] 补上xzhttp
- [ ] sessionId补上，jwt有效期内无法注销
- [ ] token验证
- [ ] grpc client
- [ ] 由于没有将sessionId存储到redis，很多中间件都没写验证逻辑



## serviceId   
ios-android-web，我这里不考虑电脑端

登录、注册时下发，将在线用户数最少的一台服务器分配给用户

10000人的群，占用redis内存大约是800-900kb


登录的时候，分配一个serverId，同时更新

