package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/prisma/prisma-client-go/prisma"
	"github.com/prisma/prisma-client-go/prisma/db"
	"github.com/prisma/prisma-client-go/prisma/models"
	"github.com/redis/go-redis/v9"

	"data-sync-cli/models"
	"data-sync-cli/utils"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Parse command-line flags
	var (
		help bool
		configPath string
	)
	flag.BoolVar(&help, "h", false, "Print help message")
	flag.StringVar(&configPath, "c", "config.json", "Path to configuration file")
	flag.Parse()

	if help {
		printHelp()
		return
	}

	// Load configuration from file
	config, err := utils.LoadConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}

	// Create Prisma client
	prismaClient := prisma.New(config.PrismaUrl, func(client *prisma.Client) {
		client.Log = func(level prisma.LogLevel, msg string) {
			log.Printf("[Prisma] %s: %s\n", level, msg)
		}
	})

	// Create Redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.RedisUrl,
		Password: config.RedisPassword,
		DB:       0,
	})

	// Connect to PostgreSQL database
	pgDb, err := db.New(config.PostgresUrl, func(client *db.Client) {
		client.Log = func(level db.LogLevel, msg string) {
			log.Printf("[Postgres] %s: %s\n", level, msg)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	// Connect to SQLite database
	sqLiteDb, err := db.New(config.SqliteUrl, func(client *db.Client) {
		client.Log = func(level db.LogLevel, msg string) {
			log.Printf("[SQLite] %s: %s\n", level, msg)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	// Start data synchronization process
	if err := syncData(context.Background(), prismaClient, pgDb, sqLiteDb, redisClient); err != nil {
		log.Fatal(err)
	}
}

func printHelp() {
	fmt.Println("Usage: data-sync-cli [options]")
	fmt.Println("Options:")
	fmt.Println("  -h        Print help message")
	fmt.Println("  -c        Path to configuration file (default: config.json)")
}

func syncData(ctx context.Context, prismaClient *prisma.Client, pgDb *db.Client, sqLiteDb *db.Client, redisClient *redis.Client) error {
	// Fetch data from PostgreSQL database
	pgData, err := models.FetchDataFromPostgres(ctx, pgDb)
	if err != nil {
		return err
	}

	// Cache data in Redis
	if err := utils.CacheData(redisClient, pgData); err != nil {
		return err
	}

	// Fetch data from Redis cache
	cachedData, err := utils.FetchCachedData(redisClient)
	if err != nil {
		return err
	}

	// Sync data with SQLite database
	if err := models.SyncDataWithSQLite(ctx, sqLiteDb, cachedData); err != nil {
		return err
	}

	return nil
}