-- File Name: itemList.elm
-- Creation Date: 03-24-2017
-- Created by: Paul Schuster
-- Notes:

port module Atelier exposing (..)

import Html exposing (Html)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick)
import Http
import Json.Decode
import Json.Decode.Pipeline as Pipe

main =
  Html.program
    { init = init
    , view = view
    , update = update
    , subscriptions = subscriptions
    }



-- MODEL

type View
  = ShowItem
  | ShowItemList

type alias Item = 
  { name : String
  , ingredients : List String
  , types : List String
  }

type alias ItemList = 
  { items : List ListItem
  }

type alias ListItem =
  { name : String
  , link : String
}

type alias Model = 
  { response : Maybe String 
  , error : Maybe Http.Error
  , item : Item
  , itemList : ItemList
  , view :  View
  , currentItem : String
  }

init : (Model, Cmd Msg)
init =
  (Model 
    Nothing 
    Nothing 
    (Item "" [] []) 
    (ItemList [])
    ShowItemList
    "Liquid Catalyst"
  , getItemList)



-- UPDATE


type Msg 
  = Response (Result Http.Error String)
  | ListResponse (Result Http.Error String)
  | ButtonItem String
  | ButtonItemList

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg |> Debug.log "msg" of
    Response dataResult ->
      case dataResult |> Debug.log "dataResult" of
        Ok data ->
          ( { model | response = Just data
            , item = cleanData data
            , view = ShowItem 
          }
            , Cmd.none
          )

        Err error ->
          ( {model | error = Just error}, Cmd.none)

    ListResponse dataResult ->
      case dataResult |> Debug.log "ListResponse: dataResult" of
        Ok data ->
          ( { model |response = Just data, itemList = cleanItemList data} 
          , Cmd.none )

        Err error ->
          ( {model | error = Just error}, Cmd.none)

    ButtonItem item->
      ( { model | currentItem = item}, getItemData 
        ({ model | currentItem = item}))

    ButtonItemList ->
      ( { model | view = ShowItemList}, getItemList)


-- SUBSCRITPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



-- VIEW


view : Model -> Html Msg
view model =
  Html.div [class "container"]
    [ Html.node "head" []
      [ Html.node "meta" [ charset "UTF-8"] []
      , Html.node "title" [] [ Html.text "Atelier" ]
      , Html.node "link" [ href "/css/normalize.css", rel "stylesheet" ]
            []
      , Html.node "link" [ href "/css/skeleton.css", rel "stylesheet" ]
            []
      , Html.node "link" [ href "/css/common.css", rel "stylesheet" ] []
      ]
      --rawView model 
      , Html.div [ class "row" ]
         [ viewItem model.item
         , viewItemList model.itemList
         ]
    ]

findView : Model -> Html Msg
findView model =
  let  
    v = case model.view of
          ShowItem ->
            viewItem model.item
          ShowItemList ->
            viewItemList model.itemList
  in
     Html.div [ class "container one-half column"] [ v ]

rawView model =
  Html.text <| toString model

viewItemList : ItemList -> Html Msg
viewItemList items =
  Html.div [ class "container one-third column" ]
    [ Html.h1 [] [Html.text "Item List"]
    , viewItemsList items.items 
    ]

viewItemsList : List ListItem -> Html Msg
viewItemsList items =
  Html.div [] 
    (List.map (\l -> Html.div [class "row" ] [Html.button 
        [class "max-full-width", onClick (ButtonItem l.name)] 
        [Html.text l.name]]) 
      (List.sortWith itemCompare items))

itemCompare : ListItem -> ListItem -> Order
itemCompare itemA itemB =
  compare itemA.name itemB.name

viewItem : Item -> Html Msg
viewItem item =
    Html.div [ class "one-half column pretty"]
    [  Html.h1 [] [Html.text item.name]
    , Html.h2 [] [Html.text "Ingredients"]
    , viewList item.ingredients
    , Html.h2 [] [Html.text "Types"]
    , viewList item.types
    ]

viewList : List String -> Html Msg
viewList str =
  Html.ul []
    (List.map (\l -> Html.li [] [Html.text l]) str)



-- HTML
getItemList : Cmd Msg
getItemList =
  let
      url = "/api/item"
  in
     Http.send ListResponse (Http.getString url)
      |> Debug.log "Http.send Item List"


decodeItemList : Json.Decode.Decoder ItemList
decodeItemList =
  Pipe.decode ItemList
    |> Pipe.required "Items" (Json.Decode.list decodeListItem)

decodeListItem : Json.Decode.Decoder ListItem
decodeListItem =
  Pipe.decode ListItem
    |> Pipe.required "Name" Json.Decode.string
    |> Pipe.required "Link" Json.Decode.string

cleanItemList : String -> ItemList
cleanItemList data =
  let 
    a = Json.Decode.decodeString decodeItemList data
  in
    case a of
      Ok data ->
        data
        |> Debug.log "CleanItemList"
      Err error ->
        ItemList []
          |> Debug.log ("cleanItemList: " ++ (toString error))

getItemData : Model -> Cmd Msg
getItemData model=
  let
      url = "http://localhost:8080/api/item/" ++ model.currentItem
  in
     Http.send Response (Http.getString url ) 
       |> Debug.log "Http.send"

cleanData : String -> Item
cleanData data =
  let
      a = Json.Decode.decodeString decodeItem data
  in
     case a of
        Ok data ->
          data
        Err _ ->
          Item "" [] []

decodeItem : Json.Decode.Decoder Item
decodeItem =
    Pipe.decode Item
        |> Pipe.required "name" (Json.Decode.string)
        |> Pipe.required "ingredients" (Json.Decode.list Json.Decode.string)
        |> Pipe.required "types" (Json.Decode.list Json.Decode.string)

