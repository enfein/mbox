---
icon: material/new-box
---

### 结构

```json
{
  "type": "mieru",
  "tag": "mieru-in",

  ... // 监听字段

  "transport": "TCP",
  "users": [
    {
      "name": "asdf",
      "password": "hjkl"
    }
  ],
  "traffic_pattern": "GgQIARAK",
  "user_hint_is_mandatory": false,
  "tcp_idle_timeout": 1800,
}
```

### 监听字段

参阅 [监听字段](/zh/configuration/shared/listen/)。

### 字段

#### transport

==必填==

通信协议。可设为 `TCP` 或 `UDP`。

#### users

==必填==

一组 mieru 用户名和密码。

#### traffic_pattern

一个 base64 字符串用于微调网络行为。

#### user_hint_is_mandatory

客户端若不发送用户提示，代理服务器将拒绝连接。

#### tcp_idle_timeout

在空闲这个秒数后关闭TCP连接。最小值为60。
