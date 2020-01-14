# Mineswipper-API development steps

## Requirements analysis and priorization

1. Tasks and priorities seems to be clear and priorities are given. As
   suggested, I will put a limit of 5 hours to spent in the project,
   measuring development time with Pomodoro technique.

## Development

### Design and implement a documented RESTful API for the game (think of a mobile app for your API)

For this step I will use a Swagger plugin since it helps me to document the API
along with drafts of structs and methods. There is a useful plugin for Golang
to do this: [swag][1].

The web framework chosen is [gin-gonic][2] since it is lightweight and easy to
use, also is very well supported and have a solid community behind.

The project layout will reflect functionality grouped by package. Each package
should take responsibility over clear scopes and communicate to other objects
about what is only necessary.

* _games_: configuration, game status and progression, time measuring, other
  business rules. 
* _users_: registration, list of games created/participating.
* _boards_: Game rules, object representations and relations between them.
* _storages_: Client implementations to persist/load previous states.
* _status_: provides a simple report about the application health.

### Problems faced

* Initializing Swagger docs: Had some trouble loading correctly the `docs`
  package until I realised that the module need to be imported as anonymous
  package.

[1]: https://github.com/swaggo/swag
[2]: https://github.com/gin-gonic/gin
