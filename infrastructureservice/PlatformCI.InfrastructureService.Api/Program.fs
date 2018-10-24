open System
open System.Threading

open Suave
open Suave.Successful
open Suave.Writers
open Suave.Web
open Suave.Filters  
open Suave.Operators

open PlatformCI.InfrastructureService.Driver
open PlatformCI.InfrastructureService.Driver.Implementation
open Microsoft.EntityFrameworkCore
open PlatformCI.InfrastructureService.Entities.EfContextBuilder
open PlatformCI.InfrastructureService.Entities
open PlatformCI.InfrastructureService.Models.Implementation

open Newtonsoft.Json

[<EntryPoint>]
let main _argv = 
  let cts = new CancellationTokenSource()
  let conf = { 
    defaultConfig with
      bindings = [ HttpBinding.createSimple HTTP "127.0.0.1" 7999 ];
      cancellationToken = cts.Token
    }
  
  let dbContextOptions = new DbContextOptionsBuilder<BuildSystemContext>()
  dbContextOptions.UseSqlServer("Data Source=***REMOVED***;Initial Catalog=***REMOVED***;Integrated Security=True;MultipleActiveResultSets=True;") |> ignore
  let efContextBuilder = new EfContextBuilder(dbContextOptions.Options)
  let driver = new DefaultInfrastructureStore(efContextBuilder)
  let infrastructureRequestFilter = new InfrastructureRequestFilter()

  let app = 
    choose 
        [ GET >=> choose 
            [ path "/Daemon/GetInfrastructureMetadata" >=> 
                request (fun _ -> OK (driver.GetInfrastructureMetadata(infrastructureRequestFilter) |> JsonConvert.SerializeObject)) >=> 
                setMimeType "application/json;charset=utf-8";
              path "/Daemon/GetFlattenedInfrastructureMetadata" >=> 
                request (fun _ -> OK (driver.GetInfrastructureMetadata(infrastructureRequestFilter).GetFlattenedData(infrastructureRequestFilter) |> JsonConvert.SerializeObject)) >=> 
                setMimeType "application/json;charset=utf-8";
              path "/Daemon/GetEnvironmentData" >=> 
                request (fun _ -> OK (driver.GetInfrastructureMetadata(infrastructureRequestFilter).GetEnvironmentData() |> JsonConvert.SerializeObject)) >=> 
                setMimeType "application/json;charset=utf-8";
            ]
        ]
        
  let _listening, server = startWebServerAsync conf app
    
  Async.Start(server, cts.Token)
  printfn "Make requests now"
  Console.ReadKey true |> ignore
    
  cts.Cancel()

  0
