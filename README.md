# ContinuousPlatform
A CI system that can be used to instrument builds. Written in Go as an open source port of BuildSystem,
an internal app used at ACC to instrument Jenkins builds from source control metadata.

## Usage
 
#### Run the following to get everything built:

##### Windows
~~~~Batchfile
:: TODO: fix this, not all copy steps for config files are here
cd ./directory/where/this/repo/is/cloned

:: Build all the things
./build/build.cmd
~~~~

##### Linux/OSX
~~~~bash
# TODO: fix this, not all copy steps for config files are here
cd ./directory/where/this/repo/is/cloned

# Build all the things
./build/build.sh
~~~~

#### Run the following to run everything:

##### Windows
~~~~Batchfile
cd ./directory/where/this/repo/is/cloned

:: Run all the Go stuff
start ./build/output/dataloader/dataloader.exe
start ./build/output/jenkinsservice/jenkinsservice.exe
start ./build/output/jobbuilder/jobbuilder.exe
start ./build/output/jobdashboard/jobdashboard.exe
start ./build/output/persistenceservice/persistenceservice.exe
start ./build/output/reposervice/reposervice.exe
start ./build/output/teamsnotifier/teamsnotifier.exe

:: Run all the .NET Core stuff
cd ./output/msbuildservice
start ./PlatformCI.MsBuildService.Api.exe
~~~~

##### Linux/OSX
~~~~bash
cd ./directory/where/this/repo/is/cloned

# Run all the Go stuff
bg ./build/output/dataloader/dataloader.exe
bg ./build/output/jenkinsservice/jenkinsservice.exe
bg ./build/output/jobbuilder/jobbuilder.exe
bg ./build/output/jobdashboard/jobdashboard.exe
bg ./build/output/persistenceservice/persistenceservice.exe
bg ./build/output/reposervice/reposervice.exe
bg ./build/output/teamsnotifier/teamsnotifier.exe

# Run all the .NET Core stuff
cd ./output/msbuildservice
bg dotnet ./PlatformCI.MsBuildService.Api.dll # ??
~~~~

### Libraries
1. Gorm - https://github.com/jinzhu/gorm
1. Lumberjack - https://gopkg.in/natefinch/lumberjack.v2	
1. Go-Errors - https://github.com/go-errors/errors
1. Raymond - https://github.com/aymerick/raymond
1. Go-Cache - https://github.com/patrickmn/go-cache