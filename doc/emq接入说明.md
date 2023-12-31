# emq部分的接入说明

## 登录校验

> 调用接口

- 路由:/open/dm/loginAuth
- 接口定义:https://www.apifox.cn/apidoc/shared-1424696c-bc32-4678-83c0-6ff9f72c4f24/api-13581488

## 设备信息获取

- 路由:/open/dm/getDeviceInfo
- 接口定义:https://www.apifox.cn/apidoc/shared-1424696c-bc32-4678-83c0-6ff9f72c4f24/api-13581483

## mqtt client id 规则

${productID}${deviceName}  
productID为固定的11个字符,后面的字符为设备名 如: 22ARVFc8Q0gerw23 其中22ARVFc8Q0g为产品id, erw23为设备id

## 设备topic订阅校验规则

设备订阅的topic需要符合以下规则才可以通过:  
${productID} 为产品id ${deviceName} 为设备id
> 物理型topic:

- $thing/down/property/${productID}/${deviceName}, //订阅 属性下发与属性上报响应
- $thing/down/event/${productID}/${deviceName}, //订阅 事件上报响应
- $thing/down/action/${productID}/${deviceName}, //订阅 应用调用设备行为

> 系统级topic:

- $ota/update/${productID}/${deviceName} // 订阅 固件升级消息下行
- $broadcast/rxd/${productID}/${deviceName} // 订阅 广播消息下行

> 系统级广播topic:

- $broadcast/rxd/${productID}/${deviceName} //订阅 广播消息下行

> 自定义topic:

- ${productID}/${deviceName}/data //订阅和发布
- ${productID}/${deviceName}/event //发布
- ${productID}/${deviceName}/control //订阅

# EMQX配置文件

```

auth.http.auth_req = http://127.0.0.1:8080/mqtt/auth
auth.http.auth_req.method = post
auth.http.auth_req.params = username=%u,password=%P,clientid=%c,ip=%a

##--------------------------------------------------------------------

auth.http.super_req = http://127.0.0.1:8080/mqtt/superuser
auth.http.super_req.method = post
auth.http.super_req.params = username=%u,password=%P,clientid=%c,ip=%a

##--------------------------------------------------------------------

auth.http.acl_req = http://127.0.0.1:8080/mqtt/acl
auth.http.acl_req.method = get
auth.http.acl_req.params = access=%A,username=%u,clientid=%c,ip=%a,topic=%t

```

参数：

- %u: username
- %c: clientid
- %a: ipaddress
- %r: protocol
- %P: password
- %A: 1 | 2, 1 = sub, 2 = pub
- %C: common name of client TLS cert
- %d: subject of client TLS cert
- %m: mountpoint
- %t: topic

## HTTP API

成功返回**200**；失败返回**400**即可
