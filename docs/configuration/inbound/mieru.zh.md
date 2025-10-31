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
