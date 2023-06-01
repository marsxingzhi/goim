# goim    


## 开发步骤     
1. 配置文件
2. 外面请求打到interfaces，然后再由interfaces进行grpc调用其他服务



## msg_gateway
1. WsServer
2. Server
3. cmd/main中启动是调用Server，而Server中又是调用WsServer