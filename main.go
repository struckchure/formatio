package main

import (
	"context"
	"fmt"

	_ "ariga.io/atlas-provider-gorm/gormschema"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/fx"

	"pkg.formatio/app"
	"pkg.formatio/dao"
	"pkg.formatio/handlers"
	"pkg.formatio/lib"
	"pkg.formatio/middlewares"
	"pkg.formatio/routers"
	"pkg.formatio/seeds"
	"pkg.formatio/services"
	"pkg.formatio/tasks"
)

func start(lc fx.Lifecycle, app *fiber.App, env lib.Env, scheduler *lib.Scheduler) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go app.Listen(fmt.Sprintf("0.0.0.0:%s", env.APP_PORT))
			go scheduler.Start()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			app.Shutdown()
			scheduler.Stop()

			return nil
		},
	})
}

// @title Formatio API
// @version 1.0
// @contact.name Formatio Team
// @contact.email formatio@overal-x.org
func main() {
	fx.New(
		fx.Provide(lib.NewEnv),

		fx.Provide(lib.NewDatabaseConnection),
		fx.Provide(lib.NewRabbitMQ),
		fx.Provide(lib.NewRabbitMQConnection),
		fx.Provide(lib.NewRedis),
		fx.Provide(lib.NewRedisConnection),
		fx.Provide(lib.NewHasher),
		fx.Provide(lib.NewAuth0),
		fx.Provide(lib.NewJwt),
		fx.Provide(lib.NewAblyConnection),
		fx.Provide(lib.NewAbly),
		fx.Provide(lib.NewRodelarClient),
		fx.Provide(lib.NewFlutterwaveClient),
		fx.Provide(lib.NewPayment),
		fx.Provide(lib.NewThressDSEncrypter),
		fx.Provide(lib.NewScheduler),
		fx.Provide(lib.NewSchedulerCron),

		fx.Provide(lib.NewK8SConfig),
		fx.Provide(lib.NewContainerManager),
		fx.Provide(lib.NewNetworkManager),
		fx.Provide(lib.NewInformer),

		fx.Provide(dao.NewMachineDao),
		fx.Provide(dao.NewMachinePlanDao),
		fx.Provide(dao.NewNetworkDao),
		fx.Provide(dao.NewRepoConnectionDao),
		fx.Provide(dao.NewUserDao),
		fx.Provide(dao.NewDeploymentDao),
		fx.Provide(dao.NewDeploymentLogDao),
		fx.Provide(dao.NewSocialConnectionDao),
		fx.Provide(dao.NewGithubAccountConnectionDao),
		fx.Provide(dao.NewInvoiceDao),
		fx.Provide(dao.NewCardDao),
		fx.Provide(dao.NewTransactionDao),

		fx.Provide(services.NewGithubService),
		fx.Provide(services.NewMachineService),
		fx.Provide(services.NewMachinePlanService),
		fx.Provide(services.NewNetworkService),
		fx.Provide(services.NewRepoConnectionService),
		fx.Provide(services.NewUserService),
		fx.Provide(services.NewDeploymentService),
		fx.Provide(services.NewDeploymentLogService),
		fx.Provide(services.NewInvoiceService),
		fx.Provide(services.NewCardService),

		fx.Provide(handlers.NewCallbackHandler),
		fx.Provide(handlers.NewDeploymentHandler),
		fx.Provide(handlers.NewGithubHandler),
		fx.Provide(handlers.NewMachineHandler),
		fx.Provide(handlers.NewMachinePlanHandler),
		fx.Provide(handlers.NewNetworkHandler),
		fx.Provide(handlers.NewWebhookHandler),
		fx.Provide(handlers.NewRepoConnectionHandler),
		fx.Provide(handlers.NewUserHandler),
		fx.Provide(handlers.NewCardHandler),
		fx.Provide(handlers.NewInvoiceHandler),

		fx.Provide(middlewares.NewUserMiddleware),
		fx.Provide(middlewares.NewJwtMiddleware),

		fx.Invoke(routers.NewProjectRouter),
		fx.Invoke(routers.NewMachineRouter),
		fx.Invoke(routers.NewUsersRouter),
		fx.Invoke(routers.NewPlanRouter),
		fx.Invoke(routers.NewWebhookRouter),
		fx.Invoke(routers.NewGithubRouter),
		fx.Invoke(routers.NewCallbackRouter),
		fx.Invoke(routers.NewBillingRouter),

		fx.Invoke(seeds.UserSeed),

		fx.Provide(tasks.NewDeploymentTasks),
		fx.Provide(tasks.NewGithubTasks),
		fx.Provide(tasks.NewMachineTasks),
		fx.Provide(tasks.NewBillingTasks),
		fx.Invoke(tasks.NewRootTasks),

		fx.Provide(app.NewApp),
		fx.Invoke(start),
	).Run()
}
