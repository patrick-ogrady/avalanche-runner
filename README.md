# Avalanche Runner
_quick and easy tool for running an avalanche node responsibly_

<p align="center">
  <a href="https://goreportcard.com/report/github.com/patrick-ogrady/avalanche-runner"><img src="https://goreportcard.com/badge/github.com/patrick-ogrady/avalanche-runner" /></a>
  <a href="https://github.com/patrick-ogrady/avalanche-runner/blob/master/LICENSE"><img src="https://img.shields.io/github/license/patrick-ogrady/avalanche-runner.svg" /></a>
</p>

## Install
To download a binary for the latest release, run:
```
curl -sSfL https://raw.githubusercontent.com/patrick-ogrady/avalanche-runner/master/scripts/install.sh | sh -s
```

The binary will be installed inside the `./bin` directory (relative to where the install command was run).

_Downloading binaries from the Github UI will cause permission errors on Mac._

### Installing in Custom Location
To download the binary into a specific directory, run:
```
curl -sSfL https://raw.githubusercontent.com/patrick-ogrady/avalanche-runner/master/scripts/install.sh | sh -s -- -b <relative directory>
```

## Usage
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
export TWILIO_TOKEN=twilio_token
avalanche-runner run
```

## Development
### Compile a Release
```text
`make compile version=RELEASE_TAG`
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
