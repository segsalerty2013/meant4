#### Meant4 Assessment

How to run ?

- Open Terminal and clone repository in desired directory (uses Go mod, any directory is fine)
- `cd` to the directory
- Run `docker build -t meant4 .`
- Run `docker run -d -p 8989:8989 meant4`
- Open Postman and test endpoint *POST*: `localhost:8989/calculate`

Thats all.


