package client

import (
	"database/sql"
	"net/http"
	"strconv"

	"go-rest-api/config"
	"go-rest-api/internal/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		users = append(users, u)
	}

	c.JSON(http.StatusOK, users)
}

// GET /users/:id
func GetUser(c *gin.Context) {
	id := c.Param("id")

	var u models.User
	err := config.DB.QueryRow("SELECT id, name, email FROM users WHERE id=$1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, u)
}

// POST /users
func CreateUser(c *gin.Context) {
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := config.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, u)
}

// PUT /users/:id
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var u models.User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := config.DB.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3", u.Name, u.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	u.ID, _ = strconv.Atoi(id)
	c.JSON(http.StatusOK, u)
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := config.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
