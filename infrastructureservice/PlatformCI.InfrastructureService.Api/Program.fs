open System
open System.Threading

open Suave
open Suave.Successful
open Suave.Writers
open Suave.Filters  
open Suave.Operators

[<EntryPoint>]
let main _argv = 
  let jsonMimeType = "application/json;charset=utf-8"
  // let appSettingsFilePath = "./appsettings.json"

  let cts = new CancellationTokenSource()
  let conf = { 
    defaultConfig with
      bindings = [ HttpBinding.createSimple HTTP "127.0.0.1" 7999 ];
      cancellationToken = cts.Token
    }

  let app = 
    choose 
        [ POST >=> choose 
            [ path "/Daemon/GetInfrastructureMetadata" >=> 
                request (fun req -> Driver.getInfrastructureMetadata(req)) >=> 
                setMimeType jsonMimeType;
              path "/Daemon/GetFlattenedInfrastructureMetadata" >=> 
                request (fun req -> Driver.getFlattenedData(req)) >=> 
                setMimeType jsonMimeType;
              path "/Daemon/GetEnvironmentData" >=> 
                request (fun req -> Driver.getEnvironmentData(req)) >=> 
                setMimeType jsonMimeType;
            ]
        ]
        
  let _listening, server = startWebServerAsync conf app
    
  Async.Start(server, cts.Token)
  printfn "Make requests now"
  Console.ReadKey true |> ignore
    
  cts.Cancel()

  0
