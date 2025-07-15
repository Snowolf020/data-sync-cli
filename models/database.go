package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/prisma/prisma-client-go/prisma"
)

// Database holds the database connection and Prisma client
类型 Database struct {
	DB        *sql.DB
	PrismaClient *prisma.Client
}

// NewDatabase returns a new Database instance
func NewDatabase(db *sql.DB, prismaClient *prisma.Client) *Database {
	return &Database{
		DB:        db,
		PrismaClient: prismaClient,
	}
}

// Connect establishes a connection to the PostgreSQL database
func (d *Database) Connect(ctx context.Context) error {
	var err error
	d.DB, err = sql.Open("postgres", "user=myuser password=mypass dbname=mydb sslmode=disable")
	if err != nil {
		return err
	}
	return d.DB.PingContext(ctx)
}

// Disconnect closes the database connection
func (d *Database) Disconnect(ctx context.Context) error {
	return d.DB.Close()
}

// GetPrismaClient returns the Prisma client
func (d *Database) GetPrismaClient() *prisma.Client {
	return d.PrismaClient
}

// SyncData synchronizes data between PostgreSQL and SQLite databases
func (d *Database) SyncData(ctx context.Context) error {
	// Use Prisma client to fetch data from PostgreSQL database
	data, err := d.PrismaClient.Data.FindMany().Exec(ctx)
	if err != nil {
		return err
	}
	// Use SQLite database connection to store the fetched data
	_, err = d.DB.ExecContext(ctx, "INSERT INTO data (id, name) VALUES ($1, $2)", data[0].ID, data[0].Name)
	if err != nil {
		return err
	}
	return nil
}
