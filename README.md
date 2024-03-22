# Final Project


# Package That Used in this Project

- [goFiber](https://github.com/gofiber/fiber/v2) 
- [gORM](https://gorm.io/docs/index.html)
- [driver Gorm Postgresql](https://github.com/go-gorm/postgres)
- [Crypto](https://pkg.go.dev/golang.org/x/crypto)
- [jwt-go](github.com/dgrijalva/jwt-go) Deprecated
- [validation](github.com/go-playground/validator/v10)

# Need to ask 
- [jwt-go](github.com/dgrijalva/jwt-go) 
- Relation in user or photo. I use FK in gorm so if user deleted then social media, comments and photo will be deleted -> yes kedelete semua
- comment ID in PUT comment/:commentID. bisa comment photo id yang mana aja
- any url need validation, photo and social media
- Get all -> no data, array kosong


# Configuration

First step is make `config.yaml` in config folder. The default for migration is true because when auto migrate is on gorm will make table social_media. which is not correct based specification that shared in google drive.

```yaml
DB_AUTO_MIGRATE: true

DB_HOST: localhost
DB_PORT: 5432
DB_USERNAME: postgres
DB_PASSWORD: root
DB_DATABASE: my_gram

JWT_SECRET_KEY: 
```

Also you need to create jwt secret key. you can fill it with any secret key (there is no limitation). Or you can blank it and run `go run main.go`. And then program will give you the secret key and closed


