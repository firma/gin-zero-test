package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"log"
	"miya/api/internal"
	"miya/api/internal/config"
	"miya/api/internal/query"
	"miya/api/internal/srv"
	"miya/api/internal/svc"
	"miya/common/stores/gormx"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var configFile = flag.String("f", "etc/api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	router := gin.Default()

	internal.Setup(internal.Build())
	internal.RegisterRoutes(router)

	svcCtx := svc.NewServiceContext(c)

	srv.BuildAndSetupSrvHub(svcCtx)
	query.SetUpDBManager(gormx.MustBuildGormDB(c.DB))

	server := &http.Server{Addr: fmt.Sprintf(":%d", c.Port), Handler: router}

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
