package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDb() {
	// Charger les variables d'environnement à partir du fichier .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erreur de chargement du fichier .env")
		panic(err)
	}

	// Récupérer la chaîne de connexion à la base de données
	dsn := os.Getenv("DB")
	if dsn == "" {
		fmt.Println("La variable d'environnement DB n'est pas définie")
		panic("La variable d'environnement DB n'est pas définie")
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println("Erreur lors de la connexion à la base de données:", err)
		panic(err)
	}

	fmt.Println("Connexion à la base de données réussie")
}
