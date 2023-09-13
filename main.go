package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	cluster := os.Getenv("CLUSTER_NAME")
	node := os.Getenv("NODE_NAME")
	namespace := os.Getenv("POD_NAMESPACE")
	pod := os.Getenv("POD_NAME")
	podIP := os.Getenv("POD_IP")

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"cluster_name":  cluster,
			"node_name":     node,
			"pod_namespace": namespace,
			"pod_name":      pod,
			"pod_ip":        podIP,
		})
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":        "OK",
			"node_name":     node,
			"pod_namespace": namespace,
			"pod_name":      pod,
			"pod_ip":        podIP,
		})
	})

	e.GET("/service/vessel", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from vessel endpoint !!"+" node_name : "+node+" pod_name : "+pod+" pod_ip: "+podIP)
	})

	e.GET("/service/container", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello from container endpoint !!"+" node_name : "+node+" pod_name : "+pod+" pod_ip: "+podIP)
	})

	e.Logger.Fatal(e.Start(":8080"))

}
