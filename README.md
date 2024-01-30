kubectl version info

```
~/k0s/pkg/version$ k0s hub &
~/k0s/pkg/version$ k0s kubectl -s http://127.0.0.1:8000 version -o yaml
clientVersion:
  buildDate: "1970-01-01T00:00:00Z"
  compiler: gc
  gitCommit: $Format:%H$
  gitTreeState: ""
  gitVersion: v0.0.0-master+$Format:%H$
  goVersion: go1.19
  major: ""
  minor: ""
  platform: linux/amd64
kustomizeVersion: v4.5.7
serverVersion:
  buildDate: "2021-09-23T22:46:18Z"
  compiler: gc
  gitCommit: 10bca343e85180f668522fe252552da20220cb7a
  gitTreeState: clean
  gitVersion: v1.23.0
  goVersion: go1.16.8
  major: "1"
  minor: "23"
  platform: linux/amd64
```

version info offered by this package:
```
~/k0s/pkg/version$ ./demo 
{
  "major": "0",
  "minor": "0",
  "gitVersion": "0.0.0",
  "gitCommit": "34953e164e260c6d520aadf6d612a44c43fb2d21",
  "gitTreeState": "clean",
  "buildDate": "2024-01-30T17:59:40Z",
  "goVersion": "go1.21.6",
  "compiler": "gc",
  "platform": "linux/amd64"
}
```
