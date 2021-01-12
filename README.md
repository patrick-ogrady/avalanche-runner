<p align="center">
  <a href="https://www.avax.network">
    <img width="25%" alt="snowplow" src="assets/logo.png?raw=true">
  </a>
</p>
<h3 align="center">
  avalanche-runner
</h3>
<p align="center">
quick and easy tool for running and monitoring an <a href="https://www.avax.network">avalanche</a> validator
</p>
<p align="center">
  <a href="https://goreportcard.com/report/github.com/patrick-ogrady/avalanche-runner"><img src="https://goreportcard.com/badge/github.com/patrick-ogrady/avalanche-runner" /></a>
  <a href="https://github.com/patrick-ogrady/avalanche-runner/blob/master/LICENSE"><img src="https://img.shields.io/github/license/patrick-ogrady/avalanche-runner.svg" /></a>
  <a href="https://github.com/patrick-ogrady/avalanche-runner/actions"><img src="https://github.com/patrick-ogrady/avalanche-runner/workflows/go/badge.svg?branch=master" /></a>
  <a href="https://github.com/patrick-ogrady/avalanche-runner/actions"><img src="https://github.com/patrick-ogrady/avalanche-runner/workflows/golangci-lint/badge.svg?branch=master" /></a>
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

#### Twilio Notifications
_add .avalanchego/.avalanche-runner.yaml_
```yaml
twilio:
  accountSid: "<accountSid>"
  authToken: "<authToken>"
  sender: "<sender phone number>"
  recipient: "<your phone number>"
```

## TODO
* change name to snowplow
* setup new host on google cloud script (install docker, etc)
* add sha integrity check on backed up files
