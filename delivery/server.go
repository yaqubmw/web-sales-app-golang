package delivery

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/yaqubmw/web-sales-app-golang/config"
	"github.com/yaqubmw/web-sales-app-golang/delivery/controller"
	"github.com/yaqubmw/web-sales-app-golang/manager"
	"github.com/yaqubmw/web-sales-app-golang/utils/checker"
)

type Server struct {
	useCaseManager manager.UsecaseManager
	engine         *gin.Engine
	host           string
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	checker.CheckErr(err)
}

func (s *Server) initController() {
	controller.NewAuthController(s.engine, s.useCaseManager.AuthUsecase())
	controller.NewSaleController(s.engine, s.useCaseManager.SaleUsecase())
	controller.NewReportController(s.engine, s.useCaseManager.ReportUsecase())
	controller.NewUserController(s.engine, s.useCaseManager.UserUsecase())
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	checker.CheckErr(err)
	infaManager, _ := manager.NewInfraManager(cfg)
	repoManager := manager.NewRepoManager(infaManager)
	useCaseManager := manager.NewUsecaseManager(repoManager)

	engine := gin.Default()
	host := fmt.Sprintf("%s:%s", cfg.ApiHost, cfg.ApiPort)
	return &Server{
		useCaseManager: useCaseManager,
		engine:         engine,
		host:           host,
	}
}
