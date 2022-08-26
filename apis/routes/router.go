package routes

import (
	"github.com/edrank/edrank_backend/apis/controllers"
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

	r.POST("/onboard/js", controllers.OnboardFromJSController)

	r.GET("/get-report-data", controllers.GetReportDataController)
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
	r.POST("/feedback-questions/:type", controllers.GetFeedbackQuestionsController)
	r.POST("/submit-feedback/:type", controllers.SubmitFeedbackController)
	r.POST("/get-feedback-teachers", controllers.GetFeedbackTeachersController)
	r.POST("/get-my-colleges-rank", controllers.GetMyCollegesRankController)
	r.GET("/get-course/:id", controllers.GetCourse)
	r.POST("/submit-gc", controllers.SubmitGCFormController)

	// college admin APIs
	r.POST("/onboard-college", controllers.OnBoardCollegeController)
	r.POST("/create-college-admin", controllers.CreateNewCollgeAdminController)
	r.GET("/my-college-entity/:entity", controllers.GetEntitiesOfMyCollegeController)
	r.POST("/toggle-feedback-drive", controllers.ToggleFeedbackDriveController)
	r.GET("/dashboard-metrics", controllers.GetDashboardMetricsController)
	r.GET("/get-college-drives", controllers.GetCollegeFeedbackDrivesController)
	r.GET("/get-teacher-linking-data", controllers.GetTeacherLinkingDataController)
	// parent APIs
	r.GET("/get-my-students", controllers.GetStudentsOfParent)

	// teacher APIs
	r.GET("/get-my-text-feedbacks", controllers.GetMyTextFeedbacksController)
	r.GET("/get-my-rank/:rank_type", controllers.GetMyRankController)

	// android specific APIs
	r.POST("/android/submit-feedback/:type", controllers.AndroidSubmitFeedbackController)

	// regulator APIs
	r.POST("/regulator/get-colleges", controllers.GetRegulatorCollegesController)
	r.POST("/detailed-feedback", controllers.GetDetailedFeedbackController)
	r.POST("/get-college", controllers.GetMyCollegeForNonCollegeController)

	// analytics APIs
	r.POST("/kbc-graph", controllers.KBCGraphController)
	r.POST("/sa-graph/:type", controllers.GetSAGraphController)
	r.GET("/progress-graph/:tid", controllers.GetProgessGraph)
}
