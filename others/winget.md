# winget
- winget is the package manager for windows similar to homebrew for mac
- but winget is official unlike homebrew
- https://www.youtube.com/watch?v=UoxmPalvz1g
## commands
- help
```bash
winget
```
- search for a package
```bash
winget search "google chrome"
```
- for more info about a package 
```bash
winget show "google chorme"
```
- install a package
```bash
winget install "google chrome"
```
- list all packages( applications) 
```bash
winget list
```
- to list the applications which has upgrades available
```bash
winget upgrade
```
- to upgrade any application
first type winget upgrade to get the list of application which has upgrades available. you will find its corresponding id in the list
```bash
winget upgrade 
winget upgrade <the package's id>
```
- to upgrade all the application
```bash
winget upgrade --all
```
- uninstall an app
```bash
winget uninstall "google chrome"
```