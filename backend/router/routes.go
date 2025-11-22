package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yebology/skillful-certification/app/controller"
	"github.com/yebology/skillful-certification/app/repository"
	"github.com/yebology/skillful-certification/app/service"
	"github.com/yebology/skillful-certification/config"
)

func SetUp(app *fiber.App) {

	api := app.Group("/api/v1")

	db := config.GetDB()

	// CLASS LOGIC
	classRepo := repository.NewClassRepository(db)
	classService := service.NewClassService(classRepo)
	classController := controller.NewClassController(classService)

	classes := api.Group("/classes")
	classes.Post("/add", classController.CreateClass)
	classes.Patch("/:id", classController.EditClass)
	classes.Get("/", classController.GetAllClass)
	classes.Get("/:id", classController.GetClassDetail)
	classes.Delete("/:id", classController.DeleteClass)
	// ===============

	// PARTICIPANT LOGIC
	participantRepo := repository.NewParticipantRepository(db)
	participantService := service.NewParticipantService(participantRepo)
	participantController := controller.NewParticipantController(participantService)

	participants := api.Group("/participants")
	participants.Post("/add", participantController.AddParticipant)
	participants.Patch("/:id", participantController.EditParticipant)
	participants.Get("/", participantController.GetAllParticipant)
	participants.Get("/:id", participantController.GetParticipantDetail)
	participants.Delete("/:id", participantController.DeleteParticipant)
	// ===============

	// ENROLLMENT LOGIC
	enrollmentRepo := repository.NewClassParticipantRepository(db)
	enrollmentService := service.NewClassParticipantService(enrollmentRepo)
	enrollmentController := controller.NewClassParticipantController(enrollmentService)

	enrollments := api.Group("/enrollments")
	enrollments.Post("/add", enrollmentController.AssignParticipant)
	enrollments.Get("participants/:participant_id", enrollmentController.GetParticipantClass)
	enrollments.Get("classes/:class_id", enrollmentController.GetClassParticipant)
	enrollments.Delete("/:id", enrollmentController.DeleteClassParticipant)
	// ===============
}
