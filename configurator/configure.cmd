:: Make a build output directory
:: Change to sh to run on Linux
:: Runs on Windows with Bash for Linux Subsystem

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
dotnet build "../../../msbuildservice/PlatformCI.MsBuildService.Api/"
cp -a "../../../msbuildservice/PlatformCI.MsBuildService.Api/bin/Debug/netcoreapp2.1/." "."
cd "../"

echo "Done!"



