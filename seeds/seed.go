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
			ImageURL:    "https://i.ibb.co.com/RGydW7xr/1695101264pcc-description.webp",
		},
		{
			Name:        "Pasir Halus",
			Price:       300000,
			Category:    "Material Bangunan",
			Stock:       20,
			Description: "Pasir halus untuk campuran plester dan cor",
			ImageURL:    "https://i.ibb.co.com/zWVrxdM2/432782f9d13f1a8cec4fe29761148dd6.jpg",
		},
		{
			Name:        "Batu Bata Merah",
			Price:       1200,
			Category:    "Material Bangunan",
			Stock:       500,
			Description: "Batu bata merah untuk dinding bangunan",
			ImageURL:    "https://i.ibb.co.com/Xk3z37mQ/2c5049ea-3537-48f6-8965-ac7421fd56caw.webp",
		},
		{
			Name:        "Besi Beton 10mm",
			Price:       85000,
			Category:    "Logam",
			Stock:       100,
			Description: "Besi beton untuk tulangan konstruksi",
			ImageURL:    "https://i.ibb.co.com/4wZRmDVZ/f8qnl-LCw-Ez-SAu-Zt-Gd-Xpw3-FZ88-SOBF1-Mj.webp",
		},
		{
			Name:        "Triplek 9mm",
			Price:       95000,
			Category:    "Kayu",
			Stock:       40,
			Description: "Papan triplek untuk furniture dan konstruksi",
			ImageURL:    "https://i.ibb.co.com/kgF5JLKZ/da2e2c67-4451-4391-aa72-f20e6e64b2c6.jpg",
		},
		{
			Name:        "Cat Tembok Putih",
			Price:       125000,
			Category:    "Finishing",
			Stock:       25,
			Description: "Cat tembok interior warna putih",
			ImageURL:    "https://i.ibb.co.com/S4mgXBWX/product-2024-decolith-1000px-c.jpg",
		},
		{
			Name:        "Keramik Lantai",
			Price:       65000,
			Category:    "Finishing",
			Stock:       80,
			Description: "Keramik lantai ukuran 40x40 cm",
			ImageURL:    "https://i.ibb.co.com/HfXDFDnW/1006-keramik-1.jpg",
		},
		{
			Name:        "Pipa PVC 3 Inch",
			Price:       45000,
			Category:    "Plumbing",
			Stock:       60,
			Description: "Pipa PVC untuk saluran air",
			ImageURL:    "https://i.ibb.co.com/PGX5h7fQ/0b980392126f792ec95d194ce21ea8.jpg",
		},
		{
			Name:        "Kabel Listrik NYA",
			Price:       150000,
			Category:    "Elektrikal",
			Stock:       30,
			Description: "Kabel listrik untuk instalasi rumah",
			ImageURL:    "https://i.ibb.co.com/pjW4DvN7/8-NYA-Multi-Kabel.jpg",
		},
		{
			Name:        "Baut Baja",
			Price:       500,
			Category:    "Hardware",
			Stock:       1000,
			Description: "Baut baja ukuran standar",
			ImageURL:    "https://i.ibb.co.com/V02Dqknc/727d5207-cf49-4029-959c-abbac41ec8bb.webp",
		},
	}
	for _, p := range products {
		config.DB.Create(&p)
	}
	log.Printf("Seed berhasil: %d produk ditambahkan", len(products))
}
