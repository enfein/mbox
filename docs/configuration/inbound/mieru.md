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
