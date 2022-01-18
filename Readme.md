# Banner Printer
This library is to simply print a ASCII banner when you start the application. 

# Steps
1. Generate an ASCII banner here,
https://manytools.org/hacker-tools/ascii-banner/
2. Copy the banner to file named `banner.txt` in the root folder of you application.
3. Copy the following line in your `main.go`

``_ = banner.Print(nil)``
4. Banner will get printed along with some useful information when you start the application.

Example,

````
  ______                   _                             
 / _____)                 | |                            
( (____  _____ ____  ____ | | _____    _____ ____  ____  
 \____ \(____ |    \|  _ \| || ___ |  (____ |  _ \|  _ \ 
 _____) ) ___ | | | | |_| | || ____|  / ___ | |_| | |_| |
(______/\_____|_|_|_|  __/ \_)_____)  \_____|  __/|  __/ 
                    |_|                     |_|   |_|    
                    
GOPATH: /Users/ssiva17/Documents/Codebase/go
HOST: C02Z13W3LVDR
GO Version: go1.17.2
GO Root: /usr/local/go
OS: darwin
Architecture: amd64
App Path: /private/var/folders/ts/85zlxc8j7154df7cctny7tb40000gn/T/GoLand
