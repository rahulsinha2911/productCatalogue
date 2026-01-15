package connection

import (
	"fmt"
	"highlevel/structs"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// WriteDB is the primary database connection for write operations
	WriteDB *gorm.DB

	// ReadDB is the read replica database connection for read operations
	ReadDB *gorm.DB
)

// InitDatabase initializes both read and write database connections with connection pooling
func InitDatabase() {
	// Get database configuration from environment variables
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "rootpassword")
	dbName := getEnv("DB_NAME", "highlevel")

	// Build DSN for write connection
	writeDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Build DSN for read connection (can be same or different for read replicas)
	readDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Configure MySQL driver with connection pool settings
	mysqlConfig := mysql.Config{
		DSN:                       writeDSN,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}

	// Initialize write database connection
	writeDB, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to write database: %v", err))
	}

	// Configure connection pool for write database
	sqlDB, err := writeDB.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to get write database instance: %v", err))
	}

	// Set connection pool settings for write DB
	sqlDB.SetMaxIdleConns(10)                  // Maximum idle connections
	sqlDB.SetMaxOpenConns(100)                 // Maximum open connections
	sqlDB.SetConnMaxLifetime(time.Hour)        // Maximum connection lifetime
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // Maximum idle time

	WriteDB = writeDB
	fmt.Println("Write database connected successfully")

	// Auto-migrate User model
	if err := WriteDB.AutoMigrate(&structs.User{}); err != nil {
		panic(fmt.Sprintf("failed to migrate User model: %v", err))
	}
	fmt.Println("User model migrated successfully")

	// Initialize read database connection
	readMysqlConfig := mysql.Config{
		DSN:                       readDSN,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}

	readDB, err := gorm.Open(mysql.New(readMysqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to read database: %v", err))
	}

	// Configure connection pool for read database
	readSqlDB, err := readDB.DB()
	if err != nil {
		panic(fmt.Sprintf("failed to get read database instance: %v", err))
	}

	// Set connection pool settings for read DB
	readSqlDB.SetMaxIdleConns(10)                  // Maximum idle connections
	readSqlDB.SetMaxOpenConns(100)                 // Maximum open connections
	readSqlDB.SetConnMaxLifetime(time.Hour)        // Maximum connection lifetime
	readSqlDB.SetConnMaxIdleTime(10 * time.Minute) // Maximum idle time

	ReadDB = readDB
	fmt.Println("Read database connected successfully")
}

// getEnv gets environment variable or returns default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
