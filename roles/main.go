package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"net/http"
	"roles/api"
	"roles/db"
)

var err error

func main() {
	// set up database
	db.DB, err = gorm.Open("mysql", db.DbURL(db.BuildDBConfig()))

	if err != nil {
		fmt.Println("status: ", err)
	}

	// defer db.DB.Close()
	db.DB.AutoMigrate(&api.NameRole{}, &api.RolePermissions{})

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		// get name/role/permission
		v1.GET("/", GetAllNames)
		v1.GET("/name/:name", GetRoleByName)
		//v1.GET("/role/:role", GetNamesByRole)
		v1.GET("/permission/:role", GetPermissionsByRole)

		// create name
		v1.POST("/name/:name/:role", CreateName)
		v1.POST("/role/:role/:permission", CreatePermission)

		// delete name/role/permission
		v1.DELETE("/name/:name", DeleteName)
		//v1.DELETE("/role/:role", DeleteRole)
		v1.DELETE("/permission/:role/:permission", DeletePermission)

		// update name/role
		v1.PUT("/:name/:role", UpdateRole)
	}
	r.Run(":8080")
}

func GetAllNames(ctx *gin.Context) {
	res, err := api.GetAllNames(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func GetRoleByName(ctx *gin.Context) {
	name := ctx.Param("name")
	res, err := api.GetRoleByName(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func GetPermissionsByRole(ctx *gin.Context) {
	role := ctx.Param("role")
	res, err := api.GetPermissionsByRole(ctx, role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func CreateName(ctx *gin.Context) {
	name := ctx.Param("name")
	role := ctx.Param("role")
	req := &api.NameRole{
		ID:   uuid.New().String(),
		Name: name,
		Role: role,
	}
	res, err := api.CreateName(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func CreatePermission(ctx *gin.Context) {
	role := ctx.Param("role")
	permission := ctx.Param("permission")
	req := &api.RolePermissions{
		ID:         uuid.New().String(),
		Role:       role,
		Permission: permission,
	}
	res, err := api.CreatePermission(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func DeleteName(ctx *gin.Context) {
	name := ctx.Param("name")
	res, err := api.DeleteName(ctx, name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func DeletePermission(ctx *gin.Context) {
	role := ctx.Param("role")
	permission := ctx.Param("permission")
	res, err := api.DeletePermission(ctx, role, permission)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}

func UpdateRole(ctx *gin.Context) {
	name := ctx.Param("name")
	role := ctx.Param("role")
	res, err := api.UpdateRole(ctx, name, role)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"result": res,
	})
}
