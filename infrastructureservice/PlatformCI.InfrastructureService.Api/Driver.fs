module Driver

open System
open System.Threading
open System.Text

open Microsoft.EntityFrameworkCore
open Newtonsoft.Json

open Suave.Successful

open PlatformCI.InfrastructureService.Driver.Implementation
open PlatformCI.InfrastructureService.Models.Implementation
open PlatformCI.InfrastructureService.Entities.EfContextBuilder
open PlatformCI.InfrastructureService.Entities

let getEfContext = 
    let dbContextOptions = new DbContextOptionsBuilder<BuildSystemContext>()
    dbContextOptions.UseSqlServer("Data Source=***REMOVED***;Initial Catalog=BuildSystem;Integrated Security=True;MultipleActiveResultSets=True;") |> ignore
    let efContextBuilder = new EfContextBuilder(dbContextOptions.Options)
    let driver = new DefaultInfrastructureStore(efContextBuilder)
    driver

let getInfrastructureMetadata = fun (req : Suave.Http.HttpRequest) ->
    let stringFormat = Encoding.UTF8.GetString(req.rawForm)
    let requestFilterObject = JsonConvert.DeserializeObject<InfrastructureRequestFilter>(stringFormat)
    let metadata = getEfContext.GetInfrastructureMetadata(requestFilterObject)
    let json = metadata |> JsonConvert.SerializeObject
    OK (json)

let getFlattenedData = fun (req : Suave.Http.HttpRequest) ->
    let stringFormat = Encoding.UTF8.GetString(req.rawForm)
    let requestFilterObject = JsonConvert.DeserializeObject<InfrastructureRequestFilter>(stringFormat)
    let metadata = getEfContext.GetInfrastructureMetadata(requestFilterObject).GetFlattenedData(requestFilterObject)
    let json = metadata |> JsonConvert.SerializeObject
    OK (json)

let getEnvironmentData = fun (req : Suave.Http.HttpRequest) ->
    let stringFormat = Encoding.UTF8.GetString(req.rawForm)
    let requestFilterObject = JsonConvert.DeserializeObject<InfrastructureRequestFilter>(stringFormat)
    let metadata = getEfContext.GetInfrastructureMetadata(requestFilterObject).GetEnvironmentData() 
    let json = metadata |> JsonConvert.SerializeObject
    OK (json)
