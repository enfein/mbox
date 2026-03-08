---
icon: material/new-box
---

### Structure

```json
{
  "type": "mieru",
  "tag": "mieru-in",

  ... // Listen Fields

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

### Listen Fields

See [Listen Fields](/configuration/shared/listen/) for details.

### Fields

#### transport

==Required==

Transmission protocol. Allowed values are `TCP` and `UDP`.

#### users

==Required==

A list of mieru user name and password.

#### traffic_pattern

A base64 string to fine tune network behavior.

#### user_hint_is_mandatory

If proxy client doesn't send user hint, proxy server will refuse the connection.

#### tcp_idle_timeout

Close TCP connection if it is idle for this number of seconds. Minimum value is 60.
