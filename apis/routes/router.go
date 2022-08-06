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
	r.POST("/top-3-teachers", controllers.Top3TeachersController)

	// college admin APIs
	r.POST("/onboard-college", middlewares.VerifyTenants([]string{"COLLEGE_ADMIN"}), controllers.OnBoardCollegeController)
	r.POST("/create-college-admin", middlewares.VerifyTenants([]string{"COLLEGE_ADMIN"}), controllers.CreateNewCollgeAdminController)
	r.GET("/my-college-teachers", middlewares.VerifyTenants([]string{"COLLEGE_ADMIN"}), controllers.GetTeachersOfMyCollegeController)
	r.GET("/my-college-parents", middlewares.VerifyTenants([]string{"COLLEGE_ADMIN"}), controllers.GetParentsOfMyCollegeController)
	r.GET("/my-college-students", middlewares.VerifyTenants([]string{"COLLEGE_ADMIN"}), controllers.GetStudentsOfMyCollegeController)
	r.GET("/my-college-college-admins", middlewares.VerifyTenants([]string{"COLLEGE_ADMIN"}), controllers.GetAdminsOfMyCollegeController)
}
