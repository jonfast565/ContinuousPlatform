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

  let cts = new CancellationTokenSource()
  let conf = { 
    defaultConfig with
      bindings = [ HttpBinding.createSimple HTTP "127.0.0.1" 7999 ];
      cancellationToken = cts.Token
    }

  let app = 
    choose 
        [ POST >=> choose 
            [ path "/Daemon/GetProject/Bytes" >=> 
                request (fun req -> Driver.getProjectFromFileBytes(req)) >=> 
                setMimeType jsonMimeType;
              path "/Daemon/GetSolution/Bytes" >=> 
                request (fun req -> Driver.getSolutionFromFileBytes(req)) >=> 
                setMimeType jsonMimeType;
              path "/Daemon/GetPublishProfile/Bytes" >=> 
                request (fun req -> Driver.getPublishProfileFromFileBytes(req)) >=> 
                setMimeType jsonMimeType;
              path "/Daemon/GetProject/LocalPath" >=> 
                request (fun req -> Driver.getProjectFromLocalPath(req)) >=> 
                setMimeType jsonMimeType;
              path "/Daemon/GetSolution/LocalPath" >=> 
                request (fun req -> Driver.getSolutionFromLocalPath(req)) >=> 
                setMimeType jsonMimeType;
              path "/Daemon/GetPublishProfile/LocalPath" >=> 
                request (fun req -> Driver.getPublishProfileFromLocalPath(req)) >=> 
                setMimeType jsonMimeType;
            ]
        ]
        
  let _listening, server = startWebServerAsync conf app
    
  Async.Start(server, cts.Token)
  printfn "Make requests now"
  Console.ReadKey true |> ignore
    
  cts.Cancel()

  0
