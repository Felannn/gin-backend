package main

import (
	"github.com/Felannn/gin-backend.git/config"
	"github.com/Felannn/gin-backend.git/models"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	godotenv.Load()
	config.InitDatabase()
	products := []models.Product{
		{
			Name:        "Semen Portland",
			Price:       75000,
			Category:    "Material Bangunan",
			Stock:       50,
			Description: "Semen untuk kebutuhan konstruksi bangunan",
			ImageURL:    "https://example.com/semen-portland.jpg",
		},
		{
			Name:        "Pasir Halus",
			Price:       300000,
			Category:    "Material Bangunan",
			Stock:       20,
			Description: "Pasir halus untuk campuran plester dan cor",
			ImageURL:    "https://example.com/pasir-halus.jpg",
		},
		{
			Name:        "Batu Bata Merah",
			Price:       1200,
			Category:    "Material Bangunan",
			Stock:       500,
			Description: "Batu bata merah untuk dinding bangunan",
			ImageURL:    "https://example.com/batu-bata.jpg",
		},
		{
			Name:        "Besi Beton 10mm",
			Price:       85000,
			Category:    "Logam",
			Stock:       100,
			Description: "Besi beton untuk tulangan konstruksi",
			ImageURL:    "https://example.com/besi-beton.jpg",
		},
		{
			Name:        "Triplek 9mm",
			Price:       95000,
			Category:    "Kayu",
			Stock:       40,
			Description: "Papan triplek untuk furniture dan konstruksi",
			ImageURL:    "https://example.com/triplek.jpg",
		},
		{
			Name:        "Cat Tembok Putih",
			Price:       125000,
			Category:    "Finishing",
			Stock:       25,
			Description: "Cat tembok interior warna putih",
			ImageURL:    "https://example.com/cat-putih.jpg",
		},
		{
			Name:        "Keramik Lantai",
			Price:       65000,
			Category:    "Finishing",
			Stock:       80,
			Description: "Keramik lantai ukuran 40x40 cm",
			ImageURL:    "https://example.com/keramik.jpg",
		},
		{
			Name:        "Pipa PVC 3 Inch",
			Price:       45000,
			Category:    "Plumbing",
			Stock:       60,
			Description: "Pipa PVC untuk saluran air",
			ImageURL:    "https://example.com/pipa-pvc.jpg",
		},
		{
			Name:        "Kabel Listrik NYA",
			Price:       150000,
			Category:    "Elektrikal",
			Stock:       30,
			Description: "Kabel listrik untuk instalasi rumah",
			ImageURL:    "https://example.com/kabel-listrik.jpg",
		},
		{
			Name:        "Baut Baja",
			Price:       500,
			Category:    "Hardware",
			Stock:       1000,
			Description: "Baut baja ukuran standar",
			ImageURL:    "https://example.com/baut-baja.jpg",
		},
	}
	for _, p := range products {
		config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}
