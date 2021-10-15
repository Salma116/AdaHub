package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Promo struct {
	PromoId     string `json:"promo_id" binding:"required,min=2,max=20"`
	PromoName   string `json:"promo_name" binding:"required,min=2,max=120"`
	DateOfStart string `json:"date_of_start" binding:"required,min=2,max=20"`
	DateOfEnd   string `json:"date_of_end" binding:"required,min=2,max=20"`
	Biography   string `json:"biography" binding:"required"`
}

func ListPromoHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		promos := &[]Promo{}

		result := db.Find(&promos)
		if result.Error != nil {
			return
		}

		c.JSON(200, promos)
	}
}

// CreateUserHandler create a user
// func CreateUserHandler(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		user := &User{}
// 		err := c.ShouldBindJSON(user)
// 		if err != nil {
// 			httpError.Internal(c, err)
// 			return
// 		}

// 		result := db.Create(user)
// 		if result.Error != nil {
// 			httpError.Internal(c, err)
// 			return
// 		}

// 		c.JSON(200, user)
// 	}
// }

// // DeleteUserHandler delete a specific user
// func DeleteUserHandler(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
// 		id, _ := c.Params.Get("id")
// 		result := db.Delete(&User{}, id)
// 		if result.Error != nil {
// 			httpError.Internal(c, result.Error)
// 			return
// 		}

// 		c.JSON(204, nil)
// 	}
// }

// // GetUserHandler get a specific user
// func GetUserHandler(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
// 		id, _ := c.Params.Get("id")
// 		user := &User{}

// 		result := db.First(user, id)
// 		if result.Error != nil {
// 			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 				httpError.NotFound(c, "User", id, result.Error)
// 			} else {
// 				httpError.Internal(c, result.Error)
// 			}
// 			return
// 		}

// 		c.JSON(200, user)
// 	}
// }

// // UpdateUserHandler update a specific user
// func UpdateUserHandler(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		//can be c.Request.URL.Query().Get("id") but it's a shorter notation
// 		id, _ := c.Params.Get("id")
// 		user := &User{}

// 		result := db.First(user, id)
// 		if result.Error != nil {
// 			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
// 				httpError.NotFound(c, "User", id, result.Error)
// 			} else {
// 				httpError.Internal(c, result.Error)
// 			}
// 			return
// 		}

// 		err := c.ShouldBindJSON(user)
// 		if err != nil {
// 			httpError.Internal(c, err)
// 			return
// 		}

// 		db.Save(user)

// 		c.JSON(200, user)
// 	}
// }
