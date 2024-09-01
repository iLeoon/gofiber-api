
https://github.com/user-attachments/assets/93452e78-866d-4b59-a370-23448713d01f


# Description
A Simple REST API in go utilizing the repository pattern.

## Tools
- [gofiber/v3](https://github.com/gofiber/fiber) 
- [postgresql](https://github.com/lib/pq)
- [validator](https://github.com/go-playground/validator)
- [env](https://github.com/joho/godotenv)
- [live-reloading](https://github.com/air-verse/air)

Performing the crud operations on the [northwind](https://github.com/pthom/northwind_psql) database.


## Routes
| Route                | Method |  Description
|:-------------------:|:------:|:-------------------------:|
| **/customers**       | GET   | Fetches all the customers  |
| **/customers/Id**    | GET    | Fetches a customer with Id | 
| **/customers/create**| POST   | Creates a customer         | 
| **/customers/update/Id**| PATCH  | Updates a customer         | 
| **/customers/delete/Id**| DELETE | Deletes a customer         |


## Installation
**1- Clone the repo.**

**2- Navigate to the folder and run `go get`.**

**3- Replace your env variables in `.env` file.**

**4- Run `air` cmd.**

***5- Code your heart out.***
