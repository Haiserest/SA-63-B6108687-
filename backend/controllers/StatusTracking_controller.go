package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Haiserest/app/ent"
	"github.com/Haiserest/app/ent/statustracking"
	"github.com/gin-gonic/gin"
)

// StatusTrackingController defines the struct for the statustracking controller
type StatusTrackingController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateStatusTracking handles POST requests for adding statustracking entities
// @Summary Create statustracking
// @Description Create statustracking
// @ID create-statustracking
// @Accept   json
// @Produce  json
// @Param statustracking body ent.StatusTracking true "StatusTracking entity"
// @Success 200 {object} ent.StatusTracking
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statustrackings [post]
func (ctl *StatusTrackingController) CreateStatusTracking(c *gin.Context) {
	obj := ent.StatusTracking{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "statustracking binding failed",
		})
		return
	}

	u, err := ctl.client.StatusTracking.
		Create().
		SetStatus(obj.Status).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetStatusTracking handles GET requests to retrieve a statustracking entity
// @Summary Get a statustracking entity by ID
// @Description get statustracking by ID
// @ID get-statustracking
// @Produce  json
// @Param id path int true "StatusTracking ID"
// @Success 200 {object} ent.StatusTracking
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statustrackings/{id} [get]
func (ctl *StatusTrackingController) GetStatusTracking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := ctl.client.StatusTracking.
		Query().
		Where(statustracking.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, u)
}

// ListStatusTracking handles request to get a list of statustracking entities
// @Summary List statustracking entities
// @Description list statustracking entities
// @ID list-statustracking
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.StatusTracking
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statustrackings [get]
func (ctl *StatusTrackingController) ListStatusTracking(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	statustrackings, err := ctl.client.StatusTracking.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, statustrackings)
}

// DeleteStatusTracking handles DELETE requests to delete a statustracking entity
// @Summary Delete a statustracking entity by ID
// @Description get statustracking by ID
// @ID delete-statustracking
// @Produce  json
// @Param id path int true "StatusTracking ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statustrackings/{id} [delete]
func (ctl *StatusTrackingController) DeleteStatusTracking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.StatusTracking.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateStatusTracking handles PUT requests to update a statustracking entity
// @Summary Update a statustracking entity by ID
// @Description update statustracking by ID
// @ID update-statustracking
// @Accept   json
// @Produce  json
// @Param id path int true "StatusTracking ID"
// @Param statustracking body ent.StatusTracking true "StatusTracking entity"
// @Success 200 {object} ent.StatusTracking
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /statustrackings/{id} [put]
func (ctl *StatusTrackingController) UpdateStatusTracking(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.StatusTracking{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "statustracking binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.StatusTracking.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewStatusTrackingController creates and registers handles for the statustracking controller
func NewStatusTrackingController(router gin.IRouter, client *ent.Client) *StatusTrackingController {
	uc := &StatusTrackingController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitStatusTrackingController registers routes to the main engine
func (ctl *StatusTrackingController) register() {
	statustrackings := ctl.router.Group("/statustrackings")

	statustrackings.GET("", ctl.ListStatusTracking)

	// CRUD
	statustrackings.POST("", ctl.CreateStatusTracking)
	statustrackings.GET(":id", ctl.GetStatusTracking)
	statustrackings.PUT(":id", ctl.UpdateStatusTracking)
	statustrackings.DELETE(":id", ctl.DeleteStatusTracking)
}
