# Bioskop

[![style: very good analysis][very_good_analysis_badge]][very_good_analysis_link]
[![License: GOOGLE][license_badge]][license_link]

Cinema Ticekt is a backend service to manage ticketing, user can create and get list of all about cinema.

### Features
- Auth System (Register, Login)
- Create, Get List Film
- Support Auth Middleware using JWT
- Password Protection using Bcrypt
- MVC

## API Endpoint
Berikut adalah daftar endpoint yang tersedia dalam API ini berserta deskripsi dan contoh permintaan.

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

### API DOCUMENTATION
-- comming soon --

[license_badge]: https://img.shields.io/badge/license-MIT-blue.svg
[license_link]: https://opensource.org/licenses/MIT
[very_good_analysis_badge]: https://img.shields.io/badge/style-very_good_analysis-B22C89.svg
[very_good_analysis_link]: https://pub.dev/packages/very_good_analysis
