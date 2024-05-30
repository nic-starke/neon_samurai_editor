# NEON_SAMURAI Editor

## Development Environment

### Dependencies

- [Wails](https://wails.io/docs/gettingstarted/installation)

```bash
# install wails
go install github.com/wailsapp/wails/v2/cmd/wails@latest

# run wails doctor to check for dependency issues
wails doctor
```

- [NodeJS (via node version manager)](https://github.com/nvm-sh/nvm)

I have found that node version manager (nvm) is the easiest way to install an upto date version of node on your system:

```bash
# cd to home directory
cd ~

# download and run nvm install script
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash

#################################################
# CLOSE THE TERMINAL HERE - OPEN A NEW TERMINAL #
#################################################

# Insall nodejs/npm
nvm install node

# .. wait for the install, then run npm to check its installed:
npm
```

### Errors
