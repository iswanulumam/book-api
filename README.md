# Sample Golang API Server

Sample REST API build using echo server.

The code implementation was inspired by Model View Controller [MVC](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller):

-   **Models**<br/>Contains implementation model database and quary to database with raw query or ORM.
-   **View**<br/>Not implemented in this code because not required User Interview (View)
-   **Contoller (API)**<br/>API http handler or controller

# How To Run Server

Just execute code below in your console

```console
./run.sh
```

# How To Consume The API

There are 4 availables API that ready to use:

-   GET `/users`
-   POST `/users`