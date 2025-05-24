module github.com/yourorg/aurelia-backend

go 1.20

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/jinzhu/gorm v1.9.16
    github.com/joho/godotenv v1.5.1
    github.com/dgrijalva/jwt-go v3.2.0+incompatible
    golang.org/x/crypto v0.22.0
    github.com/stretchr/testify v1.8.4
)

replace golang.org/x/crypto => golang.org/x/crypto v0.22.0
