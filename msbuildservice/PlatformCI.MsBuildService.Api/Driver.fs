module Driver

open Suave.Successful
open Suave.Http
open System.Text
open Newtonsoft.Json

let getEnvironmentData(req : HttpRequest) = OK("Something")
let getFlattenedData(req : HttpRequest) = OK("Something")

let getInfrastructureMetadata = fun (req : Suave.Http.HttpRequest) ->
    let stringFormat = Encoding.UTF8.GetString(req.rawForm)
    let requestFilterObject = JsonConvert.DeserializeObject<InfrastructureRequestFilter>(stringFormat)
    let metadata = getEfContext.GetInfrastructureMetadata(requestFilterObject)
    let json = metadata |> JsonConvert.SerializeObject
    OK (json)


