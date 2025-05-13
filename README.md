# zcp
A tui/cli for building install-config.yaml files used to create ocp clusters

# commands

### get list of commands, brief summary, how to use, etc
1. help
### get information of zts version
2. version
### get list of supported platforms (aws, vsphere, azure, gcp)
3. platform
### show required/expected details to interact with that provider (like having ~/.aws/config present)
- platform show
### test permissions to ensure an account can run install-config.yaml (using api calls to create vpc and such)
- platform test
