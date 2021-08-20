# Sample Golang API Server

Sample REST API build using echo server.

The code implementation was inspired by Model View Controller [MVC](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controller):

-   **Models**<br/>Contains implementation model database and quary with ORM.
-   **View**<br/>Not implemented in this code because not required User Interview (View)
-   **Contoller (API)**<br/>API http handler or controller

# How To Run Apps

1. Create Database Based on Config
2. Execute command:
```console
./run.sh
```
3. Read Documentation in folder `openapi`

# How To Test Apps
1. Create Database Based on Config
2. Execute command:
```console
./test.sh
```
3. See Report Coverage
![alt text](https://github.com/iswanulumam/book-api/blob/main/coverage-test.png?raw=true)