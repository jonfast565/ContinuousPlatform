# ContinuousPlatform
A CI system that can be used to instrument builds. Written in Go as an open source port of BuildSystem,
an internal app used at ACC to instrument Jenkins builds from source control metadata.

## Usage
 
#### Run the following to get everything built:

~~~~Batchfile
:: TODO: fix this, not all copy steps for config files are here
cd ./directory/where/this/repo/is/cloned

:: Build all the Go stuff
go build -o build ./dataloader/main.go 
go build -o build ./jenkinsservice/main.go
go build -o build ./jobbuilder/main.go
go build -o build ./jobdashboard/main.go
go build -o build ./persistenceservice/main.go
go build -o build ./reposervice/main.go
go build -o build ./teamsnotifier/main.go

:: Build all the .NET Core stuff
cd ./msbuildservice/PlatformCI.MsBuildService.Api
dotnet build
~~~~

#### Run the following to run everything:

~~~~Batchfile
cd ./directory/where/this/repo/is/cloned

:: Run all the Go stuff
start ./build/dataloader.exe
start ./build/jenkinsservice.exe
start ./build/jobbuilder.exe
start ./build/jobdashboard.exe
start ./build/persistenceservice.exe
start ./build/reposervice.exe
start ./build/teamnotifier.exe

:: Run all the .NET Core stuff
cd ./msbuildservice/PlatformCI.MsBuildService.Api
dotnet run
~~~~

## Library Credits
1. Gorm - https://github.com/jinzhu/gorm
1. Lumberjack - https://gopkg.in/natefinch/lumberjack.v2	
1. Go-Errors - https://github.com/go-errors/errors
1. Raymond - https://github.com/aymerick/raymond