package main

import "os"

var (
	reportRepository repository.ReportRepository = repository.NewSQLiteRepository()
	reportService    service.ReportService       = service.NewPostService(postRepository)
	reportController controller.ReportController = controller.NewPostController(postService)
	httpRouter       router.Router               = router.NewMuxRouter()
)

func main() {
	httpRouter.GET("/posts", postController.GetPosts)

	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(os.Getenv("PORT"))
}
