# mbox

A reference implementation of [mieru](https://github.com/enfein/mieru) protocol
in sing-box.

We don't guarantee all sing-box features are working. We recommend you maintain
your own fork.

## Example Configuration with mieru Outbound (Proxy Client)

```js
{
    "inbounds": [
        {
            "type": "mixed",
            "tag": "mixed-in",
            "listen": "0.0.0.0",
            "listen_port": 1080
        }
    ],
    "outbounds": [
        {
            "type": "mieru",
            "tag": "mieru-out",
            "server": "127.0.0.1",
            "server_port": 8964,
            "transport": "TCP",
            "username": "baozi",
            "password": "manlianpenfen"
        }
    ],
    "route": {
        "rules": [
            {
                "inbound": ["mixed-in"],
                "action": "route",
                "outbound": "mieru-out"
            }
        ]
    },
    "log": {
        "level": "warn"
    }
}
```

You can also use `server_ports` to set a list of port ranges.

## Example Configuration with mieru Inbound (Proxy Server)

```js
{
    "inbounds": [
        {
            "type": "mieru",
            "tag": "mieru-tcp",
            "listen": "0.0.0.0",
            "listen_port": 8964,
            "transport": "TCP",
            "users": [
                {
                    "name": "baozi",
                    "password": "manlianpenfen"
                }
            ]
        }
    ],
    "outbounds": [],
    "route": {},
    "log": {
        "level": "warn"
    }
}
```

## Branches

- `main` branch tracks the upstream release.
- `mieru` branch has the latest code based on upstream `main` branch. It may not be stable.
- `release-*` branches record the references where binaries are published. Those branches don't change after creation.

## Limitations

In 1.12.15 release:

- UDP outbound can't use domain name as server address.

## License

```
Copyright (C) 2022 by nekohasekai <contact-sagernet@sekai.icu>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.

In addition, no derivative work may use the name or imply association
with this application without prior consent.
```
