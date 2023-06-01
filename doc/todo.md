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
- [ ] 学习websocket的 golang库

## 需要学习的
- [ ] etcd
- [ ] grpc证书，客户端拨号
- [ ] gorilla/websocket

## 未来规划     
1. 添加chatgpt机器人
2. OAuth2，实现第三方登录功能
3. LBS服务，基于位置信息的服务（MongoDB比较有优势）
4. 加入QUIC（快速 UDP网络连接），在消息转发的效率上会有提升
5. k8s进行服务管理
6. 现有业务优化


## serviceId   
ios-android-web，我这里不考虑电脑端

登录、注册时下发，将在线用户数最少的一台服务器分配给用户

10000人的群，占用redis内存大约是800-900kb


登录的时候，分配一个serverId，同时更新

