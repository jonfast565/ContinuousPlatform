:: Make a build output directory
:: Change to sh to run on Linux
:: Runs on Windows with Bash for Linux Subsystem

:: go get all libs
go get "github.com/ahmetb/go-linq"
go get "github.com/go-errors/errors"
go get "github.com/jinzhu/gorm"
go get "github.com/jinzhu/gorm/dialects/postgres"
go get "github.com/lib/pq"
go get "github.com/satori/go.uuid"
go get "gopkg.in/natefinch/lumberjack.v2"
go get "github.com/aymerick/raymond"
go get "github.com/gorilla/mux"
go get "github.com/yosssi/gohtml"
go get "github.com/patrickmn/go-cache"
go get "golang.org/x/time"

:: output folder
mkdir output
cd "./output"

:: Build data loader
mkdir "dataloader"
cd "./dataloader"
go build "../../../dataloader"
cp "../../../dataloader/appsettings.json" .
cp "../../../dataloader/data.json" .
cd "../"

:: Build jenkins service
mkdir "jenkinsservice"
cd "./jenkinsservice"
go build "../../../jenkinsservice"
cp "../../../jenkinsservice/appsettings.json" .
mkdir "Templates"
cp -r "../../../jenkinsservice/Templates" .
cd "../"

:: Build job builder batch
mkdir "jobbuilder"
cd "./jobbuilder"
go build "../../../jobbuilder"
cp "../../../jobbuilder/appsettings.json" .
cp "../../../jobbuilder/jenkinsclient-settings.json" .
cp "../../../jobbuilder/msbuildclient-settings.json" .
cp "../../../jobbuilder/persistenceclient-settings.json" .
cp "../../../jobbuilder/repoclient-settings.json" .
cp "../../../jobbuilder/scripttemplates.json" .
mkdir "Templates"
cp -r "../../../jobbuilder/Templates" .
cd "../"

:: Build persistence service
mkdir "persistenceservice"
cd "./persistenceservice"
go build "../../../persistenceservice"
cp "../../../persistenceservice/appsettings.json" .
cd "../"

:: Build repo service
mkdir "reposervice"
cd "./reposervice"
go build "../../../reposervice"
cp "../../../reposervice/appsettings.json" .
cd "../"

:: Build msbuild service
mkdir "msbuildservice"
cd "./msbuildservice"
dotnet publish "../../../msbuildservice/PlatformCI.MsBuildService.Api/" --self-contained
:: use dotnet build with the line below for just a build
cp -a "../../../msbuildservice/PlatformCI.MsBuildService.Api/bin/Debug/netcoreapp2.1/win10-x64/publish/." "."
cd "../"

:: Build teams notifier service
mkdir "teamsnotifier"
cd "./teamsnotifier"
go build "../../../teamsnotifier"
cp "../../../teamsnotifier/appsettings.json" .
cd "../"

echo "Done!"



