| Method | Endpoint                        | Middleware        | Description                              |
| ------ | ------------------------------- | ----------------- | ---------------------------------------- |
| GET    | /                               | -                 | Root endpoint                            |
| GET    | /image/:filename                | -                 | Get image by filename                    |
| POST   | /api/authentication/register    | -                 | User registration                        |
| POST   | /api/authentication/login       | -                 | User login                               |
| GET    | /api/films                      | Authorization     | Get list of films                        |
| POST   | /api/films                      | Authorization     | Create a new film                        |
| DELETE | /api/films/remove               | Authorization     | Delete a film                            |
| GET    | /api/films/seats                | Authorization     | Get seats for a film                     |
| POST   | /api/films/booking              | Authorization     | Book a seat for a film                   |
| GET    | /api/films/actor                | Authorization     | Get list of actors                       |
| POST   | /api/films/actor                | Authorization     | Add a new actor                          |
| DELETE | /api/films/actor/delete         | Authorization     | Delete an actor                          |
| POST   | /api/films/connect              | Authorization     | Connect an actor to a film               |
