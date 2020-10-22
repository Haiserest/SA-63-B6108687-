package main

import (
	"context"
	"log"

	"github.com/Haiserest/app/controllers"
	_ "github.com/Haiserest/app/docs"
	"github.com/Haiserest/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
	Doctorname  string
	Doctoremail string
}

type Patients struct {
	Patient []Patient
}

type Patient struct {
	Patientname string
	Patientage  int
}

type StatusTrackings struct {
	StatusTracking []StatusTracking
}

type StatusTracking struct {
	Status string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewPatientController(v1, client)
	controllers.NewStatusTrackingController(v1, client)
	controllers.NewTreatmentController(v1, client)

	// Set Users Data
	users := Users{
		User: []User{
			User{"นพ. เอื้อสีวงศ์ มะม่วงร่วง", "Mango666@example.com"},
			User{"นพ. ตรวจโรค ไม่เจอ", "Dumpway@example.com"},
			User{"Name Surname", "me@example.com"},
		},
	}

	for _, u := range users.User {
		client.User.
			Create().
			SetDoctorEmail(u.Doctoremail).
			SetDoctorName(u.Doctorname).
			Save(context.Background())
	}

	// Set StatusTrackings Data
	statustrackings := StatusTrackings{
		StatusTracking: []StatusTracking{
			StatusTracking{"รักษาเสร็จสิ้น"},
			StatusTracking{"ติดตามอาการ"},
		},
	}

	for _, st := range statustrackings.StatusTracking {
		client.StatusTracking.
			Create().
			SetStatus(st.Status).
			Save(context.Background())
	}

	// Set Patients Data
	patients := Patients{
		Patient: []Patient{
			Patient{"นายโชคช่วย ไม่ช่วย", 42},
            Patient{"นายแกง  หม้อใหญ่", 28},
            Patient{"นายปิยพงษ์ ค้าไม้", 18},
		},
	}

	for _, p := range patients.Patient {
		client.Patient.
			Create().
			SetPatientName(p.Patientname).
			SetPatientAge(p.Patientage).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
