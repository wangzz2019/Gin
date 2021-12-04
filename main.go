package main

import (
	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	tracer.Start(
		tracer.WithEnv("democenter"),
		tracer.WithService("gin"),
		tracer.WithDebugMode(true),
		tracer.WithServiceVersion("abc123"),
	)
	defer tracer.Stop()

	// err := profiler.Start(
	// 	profiler.WithService("<SERVICE_NAME>"),
	// 	profiler.WithEnv("<ENVIRONMENT>"),
	// 	profiler.WithVersion("<APPLICATION_VERSION>"),
	// 	profiler.WithTags("<KEY1>:<VALUE1>,<KEY2>:<VALUE2>"),
	// 	profiler.WithProfileTypes(
	// 		profiler.CPUProfile,
	// 		profiler.HeapProfile,
	// 		// The profiles below are disabled by default to keep overhead
	// 		// low, but can be enabled as needed.

	// 		// profiler.BlockProfile,
	// 		// profiler.MutexProfile,
	// 		// profiler.GoroutineProfile,
	// 	),
	// )
	// if err != nil {
	// 	// log.Fatal(err)
	// }
	// defer profiler.Stop()

	r := gin.Default()
	r.Use(gintrace.Middleware("ginservice"))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Jack Wang")
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
