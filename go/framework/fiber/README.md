```bash
# init module
go mod init github.com/naufalihsan/academy

# install third party packages
go get github.com/gofiber/fiber/v2
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

go mod tidy

# init psql
createdb go-db
```