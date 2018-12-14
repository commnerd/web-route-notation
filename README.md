# Web Route Notation 
Create web routes for your app in a short-hand, readable notation

## Example 1:
The following Web Route Notation:

`+GPUADH/commnerd/web-route-notation[middleware1,middleware2]*route.name@Handler#method`

will translate to:

```json
[
  {
    "name": "route.name",
    "controller": "Handler",
    "method": "method",
    "route": "/commnerd/web-route-notation",
    "verbs": [
      "GET",
      "POST",
      "PUT",
      "PATCH",
      "DELETE",
      "HEAD"
    ],
    "middlewares": [
      "middleware1",
      "middleware2"
    ]
  }
]
```

### Delimiter Definitions:
- "+" Delimits the verbs and a new route, if immediately followed by one of the following delimiters, all routes will map to this -definition
- "/" Initiates the start of the path
- "[]" Delimits middlewares
- "*" Is optional and delimits the route name
- "@" Delimits the Controller for a route
- "#" Delimits the Method within the controller
- "()" Delimits a group

Note: Order does not matter.

### Http Verb Mapping:
- G = GET
- P = POST
- U = PUT
- A = PATCH
- D = DELETE
- H = HEAD

## Example 2:

Web Route Notation:

```
+/@HomeController#index
+GP/route1@RouteController
+D#delete@SomeController/tested
+/subroute/*subrouteName(
  +/awesome[middleware]*subrouteName
)
```
Will generate the following translation:
```json
[
  {
    "controller": "HomeController",
    "method": "index",
    "route": "/"
  },
  {
    "controller": "RouteController",
    "route": "/route1",
    "verbs": [
      "GET",
      "POST"
    ],
  },
  {
    "name": "route.name",
    "controller": "Handler",
    "method": "delete",
    "route": "/commnerd/web-route-notation",
    "verbs": [
      "DELETE"
    ],
  },
  {
    "name": "subrouteName",
    "route": "/subroute/",
    "group": [
      {
        "name": "subrouteName",
        "route": "/awesome",
        "middlewares": [
          "middleware"
        ],
      }
    ]
  }
]
```
