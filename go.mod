module data-sync-cli

go 1.19

require (
	github.com/lib/pq v1.10.7
	github.com/prisma/prisma-client-go v1.14.0
	gorm.io/driver/postgres v1.3.8
	gorm.io/driver/sqlite v1.3.8
	gorm.io/gorm v1.9.16
	github.com/go-redis/redis/v9 v9.6.0
	github.com/spf13/viper v1.13.0
)

replace github.com/prisma/prisma-client-go => github.com/prisma/prisma-client-go v1.14.0
