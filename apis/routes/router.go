package routes

import (
	"github.com/edrank/edrank_backend/apis/controllers"
	"github.com/edrank/edrank_backend/apis/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Status": "OK",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Service": "Edrank Backend APIs v1.0",
		})
	})
}

func InitPublicRoutes(r *gin.RouterGroup) {
	// common APIs
	r.POST("/login", controllers.LoginController)

	// dev APIs
	r.POST("/set-onboarding-files", controllers.SetOnBoardingFileController)
	r.POST("/dev/gen-password-hash", controllers.GenPasswordHash)
}

func InitPrivateRoutes(r *gin.RouterGroup) {
	// file routes
	r.POST("/file-upload", controllers.FileUploadController)

	// common APIS
	r.POST("/change-password", controllers.ChangePasswordController)
	r.GET("/college", controllers.GetCollegeController)
	r.GET("/my-profile", controllers.GetMyProfile)
	r.POST("/top-n-teachers", controllers.TopNTeachersController)
	r.POST("/top-n-colleges", controllers.TopNCollegesController)
	r.POST("/feedback-questions/:type", middlewares.VerifyTenants([]string{"PARENT", "STUDENT", "HEI"}), controllers.GetFeedbackQuestionsController)
	r.POST("/submit-feedback/:type", middlewares.VerifyTenants([]string{"STUDENT", "PARENT", "HEI"}), controllers.SubmitFeedbackController)
	r.POST("/get-feedback-teachers", middlewares.VerifyTenants([]string{"STUDENT"}), controllers.GetFeedbackTeachersController)
	r.POST("/get-my-colleges-rank", middlewares.VerifyTenants([]string{"STUDENT", "TEACHER", "PARENT", "SUPER_ADMIN"}), controllers.GetMyCollegesRankController)
	r.GET("/get-course/:id", controllers.GetCourse)

	// college admin APIs
	r.POST("/onboard-college", middlewares.VerifyTenants([]string{"COLLEGE_ADMIN"}), controllers.OnBoardCollegeController)
	r.POST("/create-college-admin", middlewares.VerifyTenants([]string{"COLLEGE_ADMIN"}), controllers.CreateNewCollgeAdminController)
	r.GET("/my-college-entity/:entity", middlewares.VerifyTenants([]string{"COLLEGE_ADMIN"}), controllers.GetEntitiesOfMyCollegeController)
	r.POST("/toggle-feedback-drive", middlewares.VerifyTenants([]string{"COLLEGE_ADMIN"}), controllers.ToggleFeedbackDriveController)
	r.GET("/dashboard-metrics", middlewares.VerifyTenants([]string{"COLLEGE_ADMIN"}), controllers.GetDashboardMetricsController)

	// parent APIs
	r.GET("/get-my-students", middlewares.VerifyTenants([]string{"PARENT"}), controllers.GetStudentsOfParent)

	// teacher APIs
	r.GET("/get-my-text-feedbacks", middlewares.VerifyTenants([]string{"TEACHER"}), controllers.GetMyTextFeedbacksController)
	r.GET("/get-my-rank/:rank_type", middlewares.VerifyTenants([]string{"TEACHER"}), controllers.GetMyRankController)

	// android specific APIs
	r.POST("/android/submit-feedback/:type", middlewares.VerifyTenants([]string{"STUDENT", "PARENT", "HEI"}), controllers.AndroidSubmitFeedbackController)

	// regulator APIs
	r.POST("/regulator/get-colleges", middlewares.VerifyTenants([]string{"SUPER_ADMIN"}), controllers.GetRegulatorCollegesController)
	r.POST("/detailed-feedback", middlewares.VerifyTenants([]string{"SUPER_ADMIN"}), controllers.GetDetailedFeedbackController)


	// analytics APIs
	r.POST("/kbc-graph", middlewares.VerifyTenants([]string{"COLLEGE_ADMIN", "TEACHER"}), controllers.KBCGraphController)
	r.POST("/sa-graph/:type", controllers.GetSAGraphController)
}
