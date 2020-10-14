package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/Haiserest/app/ent"
	"github.com/Haiserest/app/ent/patient"
	"github.com/Haiserest/app/ent/statustracking"
	"github.com/Haiserest/app/ent/treatment"
	"github.com/Haiserest/app/ent/user"
	"github.com/gin-gonic/gin"
)

// TreatmentController defines the struct for the treatment controller
type TreatmentController struct {
	client *ent.Client
	router gin.IRouter
}

type Treatment struct {
	Userid    int
	Patientid int
	Status    int
	Symptom   string
	Added     string
}

// CreateTreatment handles POST requests for adding treatment entities
// @Summary Create treatment
// @Description Create treatment
// @ID create-treatment
// @Accept   json
// @Produce  json
// @Param treatment body Treatment true "Treatment entity"
// @Success 200 {object} Treatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /treatments [post]
func (ctl *TreatmentController) CreateTreatment(c *gin.Context) {
	obj := Treatment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "treatment binding failed",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.Userid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patientid))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	st, err := ctl.client.StatusTracking.
		Query().
		Where(statustracking.IDEQ(int(obj.Status))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "statustracking not found",
		})
		return
	}

	times, err := time.Parse(time.RFC3339, obj.Added)

	t, err := ctl.client.Treatment.
		Create().
		SetPatientID(p).
		SetDoctorID(u).
		SetStatusID(st).
		SetDateAppoint(times).
		SetSymptom(obj.Symptom).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, t)
}

// GetTreatment handles GET requests to retrieve a treatment entity
// @Summary Get a treatment entity by ID
// @Description get treatment by ID
// @ID get-treatment
// @Produce  json
// @Param id path int true "Treatment ID"
// @Success 200 {object} ent.Treatment
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /treatments/{id} [get]
func (ctl *TreatmentController) GetTreatment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	t, err := ctl.client.Treatment.
		Query().
		Where(treatment.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, t)
}

// ListTreatment handles request to get a list of treatment entities
// @Summary List treatment entities
// @Description list treatment entities
// @ID list-treatment
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Treatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /treatments [get]
func (ctl *TreatmentController) ListTreatment(c *gin.Context) {
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

	treatments, err := ctl.client.Treatment.
		Query().
		WithDoctorID().
		WithPatientID().
		WithStatusID().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, treatments)
}

// DeleteTreatment handles DELETE requests to delete a treatment entity
// @Summary Delete a treatment entity by ID
// @Description get treatment by ID
// @ID delete-treatment
// @Produce  json
// @Param id path int true "Treatment ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /treatments/{id} [delete]
func (ctl *TreatmentController) DeleteTreatment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Treatment.
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

// UpdateTreatment handles PUT requests to update a treatment entity
// @Summary Update a treatment entity by ID
// @Description update treatment by ID
// @ID update-treatment
// @Accept   json
// @Produce  json
// @Param id path int true "Treatment ID"
// @Param treatment body ent.Treatment true "Treatment entity"
// @Success 200 {object} ent.Treatment
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /treatments/{id} [put]
func (ctl *TreatmentController) UpdateTreatment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Treatment{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "treatment binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Treatment.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewTreatmentController creates and registers handles for the treatment controller
func NewTreatmentController(router gin.IRouter, client *ent.Client) *TreatmentController {
	uc := &TreatmentController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitTreatmentController registers routes to the main engine
func (ctl *TreatmentController) register() {
	treatments := ctl.router.Group("/treatments")

	treatments.GET("", ctl.ListTreatment)

	// CRUD
	treatments.POST("", ctl.CreateTreatment)
	treatments.GET(":id", ctl.GetTreatment)
	treatments.PUT(":id", ctl.UpdateTreatment)
	treatments.DELETE(":id", ctl.DeleteTreatment)
}
