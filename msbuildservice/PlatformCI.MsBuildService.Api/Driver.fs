module Driver

open Suave.Successful
open Suave.Http
open System.Text
open Newtonsoft.Json

open PlatformCI.MsBuildService.Driver.Implementation
open PlatformCI.MsBuildService.Models.Implementation

let filesystemImpl = new BasicFilesystemProvider()
let projectResolver = new HackedMicrosoftProjectResolver()
let microsoftBuildEndpoint = new DefaultMicrosoftBuildProviderEndpoint(filesystemImpl, projectResolver)

let getProjectFromFileBytes = fun (req : Suave.Http.HttpRequest) ->
    let stringFormat = Encoding.UTF8.GetString(req.rawForm)
    let payloadObject = JsonConvert.DeserializeObject<FilePayload>(stringFormat)
    let metadata = microsoftBuildEndpoint.GetProjectFromFileBytes(payloadObject)
    let json = metadata |> JsonConvert.SerializeObject
    OK (json)

let getSolutionFromFileBytes = fun (req : Suave.Http.HttpRequest) ->
    let stringFormat = Encoding.UTF8.GetString(req.rawForm)
    let payloadObject = JsonConvert.DeserializeObject<FilePayload>(stringFormat)
    let metadata = microsoftBuildEndpoint.GetSolutionFromFileBytes(payloadObject)
    let json = metadata |> JsonConvert.SerializeObject
    OK (json)

let getPublishProfileFromFileBytes = fun (req : Suave.Http.HttpRequest) ->
    let stringFormat = Encoding.UTF8.GetString(req.rawForm)
    let payloadObject = JsonConvert.DeserializeObject<FilePayload>(stringFormat)
    let metadata = microsoftBuildEndpoint.GetPublishProfileFromFileBytes(payloadObject)
    let json = metadata |> JsonConvert.SerializeObject
    OK (json)

let getProjectFromLocalPath = fun (req : Suave.Http.HttpRequest) ->
    let stringFormat = Encoding.UTF8.GetString(req.rawForm)
    let metadata = microsoftBuildEndpoint.GetProjectFromLocalPath(stringFormat)
    let json = metadata |> JsonConvert.SerializeObject
    OK (json)

let getSolutionFromLocalPath = fun (req : Suave.Http.HttpRequest) ->
    let stringFormat = Encoding.UTF8.GetString(req.rawForm)
    let metadata = microsoftBuildEndpoint.GetSolutionFromLocalPath(stringFormat)
    let json = metadata |> JsonConvert.SerializeObject
    OK (json)

let getPublishProfileFromLocalPath = fun (req : Suave.Http.HttpRequest) ->
    let stringFormat = Encoding.UTF8.GetString(req.rawForm)
    let metadata = microsoftBuildEndpoint.GetPublishProfileFromLocalPath(stringFormat)
    let json = metadata |> JsonConvert.SerializeObject
    OK (json)


