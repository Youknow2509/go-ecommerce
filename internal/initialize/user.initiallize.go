package initialize

import (
	"log"

	"github.com/Youknow2509/go-ecommerce/global"
	domainCache "github.com/Youknow2509/go-ecommerce/internal/user/domain/cache"
	"github.com/Youknow2509/go-ecommerce/internal/user/application/services"
	"github.com/Youknow2509/go-ecommerce/internal/user/controller/http"
	domainRepository "github.com/Youknow2509/go-ecommerce/internal/user/domain/responsitory"
	"github.com/Youknow2509/go-ecommerce/internal/user/infrastructure/database"
	infMq "github.com/Youknow2509/go-ecommerce/internal/user/infrastructure/mq"
	"github.com/Youknow2509/go-ecommerce/internal/user/infrastructure/mq/impl"
	infRespon "github.com/Youknow2509/go-ecommerce/internal/user/infrastructure/persistence/responsitory"
	infCacheDis "github.com/Youknow2509/go-ecommerce/internal/user/infrastructure/cache/distributed"
	infCacheLo "github.com/Youknow2509/go-ecommerce/internal/user/infrastructure/cache/local"
	"github.com/gin-gonic/gin"
)

func InitializeUser(routerGroup *gin.RouterGroup) {
	// initialize user service
	userService := services.NewUserServiceImpl(
		getUserResponsitory(),
		getUserDistributedCache(),
		getUserLocalCache(),
		getKafkaSendOtpService(),
	)

	// initialize user handler http
	userRouterHandler := http.NewUserHandlerHttp(userService)

	// define user router manager
	routerManager := http.InitUserRouter(userRouterHandler)
	{
		routerManager.UserPublicRouterGroup.InitUserPublicRouter(routerGroup)
		routerManager.UserPrivateRouterGroup.InitUserPrivateRouter(routerGroup)
	}
}

func getKafkaSendOtpService() infMq.IKafkaService {
	infMq.InitKafkaService(impl.NewSendOtpService(global.KafkaSendOtp))
	kafkaService, err := infMq.GetKafkaService()
	if err != nil {
		log.Fatalf("KafkaService is not initialized: %v", err)
	}
	return kafkaService
}

func getUserResponsitory() domainRepository.IUserResponsitory {
	dbQueries := database.New(global.Mdbc) // global.Mdbc là *sql.DB
	return infRespon.NewUserResponsitory(dbQueries)
}

func getUserDistributedCache() domainCache.ICacheService {
	domainCache.SetDistributedCacheService(
		infCacheDis.NewRedisDistributedCache(global.Rdb),
	)
	cacheService, err := domainCache.GetDistributedCacheService()
	if err != nil {
		log.Fatalf("DistributedCache is not initialized: %v", err)
	}
	return cacheService
}

func getUserLocalCache() domainCache.ICacheService {
	domainCache.SetLocalCacheService(
		infCacheLo.NewReistrettoLocalCache(),
	)
	cacheService, err := domainCache.GetLocalCacheService()
	if err != nil {
		log.Fatalf("LocalCache is not initialized: %v", err)
	}
	return cacheService
}
