package controllers

import (
	"bjb-crud/initializers"
	"bjb-crud/models"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type suratRequest struct {
	Tanggal string `json:"tanggal"`
	NoMemo  string `json:"no_memo"`
	Perihal string `json:"perihal"`
	Pic     string `json:"pic"`
}

func SuratCreate(c *gin.Context) {

	// Get data off req body
	var requestBody suratRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.Status(400)
		c.Error(err) // log the error
		return
	}

	// Add some logging to see what's being received
	log.Println("Received request body:", requestBody)

	// Parse the date string
	tanggalString := requestBody.Tanggal
	tanggal, err := time.Parse("2006-01-02", tanggalString)
	if err != nil {
		c.Status(400)
		c.JSON(400, gin.H{"error": "Invalid date format: " + err.Error()})
		return
	}

	surat := models.Surat{
		Tanggal: tanggal,
		NoMemo:  requestBody.NoMemo,
		Perihal: requestBody.Perihal,
		Pic:     requestBody.Pic,
	}

	result := initializers.DB.Create(&surat)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"Surat": surat,
	})

}

func SuratIndex(c *gin.Context) {

	// Get models from DB
	var surat []models.Surat
	initializers.DB.Find(&surat)

	//Respond with them
	c.JSON(200, gin.H{
		"Surat": surat,
	})
}

func SuratShow(c *gin.Context) {

	id := c.Params.ByName("id")
	// Get models from DB
	var surat models.Surat

	initializers.DB.First(&surat, id)

	//Respond with them
	c.JSON(200, gin.H{
		"Surat": surat,
	})
}

func SuratUpdate(c *gin.Context) {

	var requestBody suratRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.Status(400)
		c.Error(err) // log the error
		return
	}

    id := c.Params.ByName("id")

    var surat models.Surat
    initializers.DB.First(&surat, id)

    if err := initializers.DB.First(&surat, id).Error; err != nil {
        c.JSON(404, gin.H{"error": "surat tidak ditemukan"})
        return
    }

    if requestBody.Tanggal != "" {
        tanggal, err := time.Parse("2006-01-02", requestBody.Tanggal)
        if err != nil {
            c.JSON(400, gin.H{"error": "Format tanggal tidak valid: " + err.Error()})
            return
        }
        surat.Tanggal = tanggal
    }

    if requestBody.NoMemo != "" {
        surat.NoMemo = requestBody.NoMemo
    } else {
        surat.NoMemo = surat.NoMemo // gunakan nilai yang ada dari database
    }

	if requestBody.Perihal != "" {
		surat.Perihal = requestBody.Perihal
	} else {
		surat.Perihal = surat.Perihal // gunakan nilai yang ada dari database
	}
	
	if requestBody.Pic != "" {
		surat.Pic = requestBody.Pic
	} else {
		surat.Pic = surat.Pic // gunakan nilai yang ada dari database
	}

    initializers.DB.Model(&surat).Updates(surat)

    c.JSON(200, gin.H{
        "surat": surat,
    })

}

func SuratDelete(c *gin.Context) {

	//get id
	id := c.Params.ByName("id")

	// find the surat
	var surat models.Surat

	if err := initializers.DB.First(&surat, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Surat not found"})
		return
	}

	/// delete it
	if err := initializers.DB.Delete(&surat).Error; err != nil {
		c.JSON(404, gin.H{"error": "Surat Failed to Delete"})
		return
	}

	c.JSON(200, gin.H{
		"Surat": "Surat deleted",
	})
}