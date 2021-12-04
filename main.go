package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	gintrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gin-gonic/gin"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"gopkg.in/DataDog/dd-trace-go.v1/profiler"
)

func main() {
	tracer.Start(
		tracer.WithEnv("goenv"),
		tracer.WithService("ginservice"),
		tracer.WithDebugMode(true),
		tracer.WithServiceVersion("abc123"),
	)
	defer tracer.Stop()

	err := profiler.Start(
		profiler.WithService("ginserviceprofiler"),
		profiler.WithEnv("goenv"),
		profiler.WithVersion("1.0"),
		profiler.WithTags("profilertag1:value1,profilertag2:value2"),
		profiler.WithProfileTypes(
			profiler.CPUProfile,
			profiler.HeapProfile,
			// The profiles below are disabled by default to keep overhead
			// low, but can be enabled as needed.

			profiler.BlockProfile,
			profiler.MutexProfile,
			profiler.GoroutineProfile,
		),
	)
	if err != nil {
		// log.Fatal(err)
	}
	defer profiler.Stop()

	r := gin.Default()
	r.Use(gintrace.Middleware("ginservice"))
	r.GET("/test", test)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Jack Wang")
	})
	// go goroutine()
	r.Run() // listen and serve on 0.0.0.0:8080
}
func goroutine(c *gin.Context) {
	parentSpan, _ := tracer.SpanFromContext(c.Request.Context())
	span := tracer.StartSpan("goroutine", tracer.ResourceName("goroutine"), tracer.ChildOf(parentSpan.Context()))
	defer parentSpan.Finish()
	defer span.Finish()
	time.Sleep(10 * time.Second)
	fmt.Println(2)
}
func test(c *gin.Context) {
	parentSpan, _ := tracer.SpanFromContext(c.Request.Context())
	span := tracer.StartSpan("test", tracer.ResourceName("GET /test111"), tracer.ChildOf(parentSpan.Context()))
	defer parentSpan.Finish()
	defer span.Finish()
	go goroutine(c)
	span.SetTag("tag1", "value1")
	c.String(200, "This is Jack Wang's test page")
}
