# Avalanche Runner
_quick and easy tool for running an avalanche node responsibly_

<p align="center">
  <a href="https://goreportcard.com/report/github.com/patrick-ogrady/avalanche-runner"><img src="https://goreportcard.com/badge/github.com/patrick-ogrady/avalanche-runner" /></a>
  <a href="https://github.com/patrick-ogrady/avalanche-runner/blob/master/LICENSE"><img src="https://img.shields.io/github/license/patrick-ogrady/avalanche-runner.svg" /></a>
</p>

## Usage
### Install
_must have golang installed_
```text
make install
```

_Creates in .avalanchego/staking (relative directory)_
### Create Staking Credentials
```text
avalanche-runner create
```

### Encrypt + Backup Credentials
_Make sure to set GOOGLE_APPLICATION_CREDENTIALS_
https://cloud.google.com/storage/docs/reference/libraries#setting_up_authentication
```text
export GOOGLE_APPLICATION_CREDENTIALS=blah
avalanche-runner backup [bucket]
```

### Restore + Decrypt Credentials
_Make sure to set GOOGLE_APPLICATION_CREDENTIALS_
https://cloud.google.com/storage/docs/reference/libraries#setting_up_authentication
```text
export GOOGLE_APPLICATION_CREDENTIALS=blah
avalanche-runner restore [bucket] [node ID]
```

### Start Node
_TODO: add docker cmd_
```text
make run-mainnet
```

## TODO
- [x] license generator
- [x] create staking key
- [x] backup staking key
- [x] hardcode directory name of where keys are generated to be
  .avalanchego/staking
- [x] dockerfile
- [x] run binary
- [ ] Config file with phone number + twilio tokens
- [ ] page if stops or unhealthy (only once bootstrapped has gone true)
- [ ] github workflow tester
- [ ] add sha integrity check on backed up files
- [ ] generate and host cli binaries
- [ ] setup new host script (install docker, etc)
