# Final Project


# Package That Used in this Project

- [goFiber](https://github.com/gofiber/fiber/v2) 
- [gORM](https://gorm.io/docs/index.html)
- [driver Gorm Postgresql](https://github.com/go-gorm/postgres)
- [Crypto](https://pkg.go.dev/golang.org/x/crypto)
- [Golang jwt](github.com/golang-jwt/jwt)
- [validation](github.com/go-playground/validator/v10)

# Configuration

First step is make `config.yaml` in config folder. The default for migration is true because when auto migrate is on gorm will make table social_media. which is not correct based specification that shared in google drive.

```yaml
DB_AUTO_MIGRATE: 

ENV: 

DB_HOST: 
DB_PORT: 
DB_USERNAME: 
DB_PASSWORD: 
DB_DATABASE: 

JWT_SECRET_KEY: 
```

1. DB_AUTO_MIGRATE (true or false)  
    Set true to turn on the migration.
2. ENV (TESTING, DEVELOPMENT or PRODUCTION)  
   Set "TESTING" for the configuration read from `config.yaml`.  
    Set "DEVELOPMENT" or "PRODUCTION" to load the configuration from Environment Variables
3. DB host, port, username, password and database.  
    You know what is this.
4. JWT_SECRET_KEY  
   For this you must generate your own secret key. There is an function to generate secure key but i did't call this function. THe function are located at `./helper/jwt` with function name are `GenerateSecretKey`.


# No Deploy
Yes, i was trying to deploy it at railways but not work. This are the [Discussions](https://help.railway.app/questions/golang-fiber-deploy-unsuccessful-87b709c5)  
Nothing wrong with my app. The error says `undifined NewController` which means A function not found in this app. But when i build and run locally it works. 


