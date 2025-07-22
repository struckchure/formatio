package tests

// import (
// 	"context"
// 	"log"
// 	"testing"

// 	"github.com/samber/lo"
// 	"github.com/shopspring/decimal"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/testcontainers/testcontainers-go/modules/postgres"

// 	"pkg.formatio/dao"
// 	"pkg.formatio/tests/config"
// 	_ "pkg.formatio/tests/config/init"
// 	"pkg.formatio/types"
// )

// func planDAORunner(t *testing.T, callback interface{}) {
// 	ctx, container := config.GetSetup()

// 	err := container.Restore(ctx, postgres.WithSnapshotName("genesis"))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	config.UseRunner(
// 		t,
// 		callback,
// 		config.NewEnv,
// 		config.NewTestDatabaseConnection,
// 		dao.NewMachinePlanDao,
// 		func() (context.Context, *postgres.PostgresContainer) {
// 			return ctx, container
// 		},
// 	)
// }

// func TestDAOListPlans(t *testing.T) {
// 	planDAORunner(
// 		t,
// 		func(machinePlanDAO dao.MachinePlanDaoInterface) {
// 			machinePlanDAO.CreateMachinePlan(types.CreateMachinePlanArgs{
// 				Name:        "test",
// 				Currency:    "NGN",
// 				MonthlyRate: decimal.NewFromInt(500),
// 				Cpu:         "256Mi",
// 				Memory:      "1024m",
// 			})

// 			result, err := machinePlanDAO.ListMachinePlans(types.ListMachinePlansArgs{})

// 			assert.Nil(t, err)
// 			assert.NotNil(t, result)

// 			assert.Equal(t, len(result), 1)
// 		})
// }

// func TestDAOCreatePlan(t *testing.T) {
// 	planDAORunner(
// 		t,
// 		func(machinePlanDAO dao.MachinePlanDaoInterface) {
// 			expected := types.CreateMachinePlanArgs{
// 				Name:        "test",
// 				Currency:    "NGN",
// 				MonthlyRate: decimal.NewFromInt(500),
// 				Cpu:         "256Mi",
// 				Memory:      "1024m",
// 			}
// 			result, err := machinePlanDAO.CreateMachinePlan(expected)

// 			assert.Nil(t, err)
// 			assert.NotNil(t, result)

// 			assert.Equal(t, expected.Name, result.Name)
// 			assert.Equal(t, expected.Currency, result.Currency)
// 			assert.Equal(t, expected.MonthlyRate, result.MonthlyRate)
// 			assert.Equal(t, expected.Cpu, result.CPU)
// 			assert.Equal(t, expected.Memory, result.Memory)
// 		})
// }

// func TestDAOGetPlan(t *testing.T) {
// 	planDAORunner(
// 		t,
// 		func(machinePlanDAO dao.MachinePlanDaoInterface) {
// 			expected := types.CreateMachinePlanArgs{
// 				Name:        "test",
// 				Currency:    "NGN",
// 				MonthlyRate: decimal.NewFromInt(500),
// 				Cpu:         "256Mi",
// 				Memory:      "1024m",
// 			}
// 			plan, err := machinePlanDAO.CreateMachinePlan(expected)

// 			assert.Nil(t, err)
// 			assert.NotNil(t, plan)

// 			result, err := machinePlanDAO.GetMachinePlan(types.GetMachinePlanArgs{Id: plan.ID})

// 			assert.Nil(t, err)
// 			assert.NotNil(t, result)

// 			assert.Equal(t, expected.Name, result.Name)
// 			assert.Equal(t, expected.Currency, result.Currency)
// 			assert.Equal(t, expected.MonthlyRate, result.MonthlyRate)
// 			assert.Equal(t, expected.Cpu, result.CPU)
// 			assert.Equal(t, expected.Memory, result.Memory)
// 		})
// }

// func TestDAOUpdatePlan(t *testing.T) {
// 	planDAORunner(
// 		t,
// 		func(machinePlanDAO dao.MachinePlanDaoInterface) {
// 			expected := types.CreateMachinePlanArgs{
// 				Name:        "test",
// 				Currency:    "NGN",
// 				MonthlyRate: decimal.NewFromInt(500),
// 				Cpu:         "256Mi",
// 				Memory:      "1024m",
// 			}
// 			plan, err := machinePlanDAO.CreateMachinePlan(expected)

// 			assert.Nil(t, err)
// 			assert.NotNil(t, plan)

// 			result, err := machinePlanDAO.UpdateMachinePlan(types.UpdateMachinePlanArgs{
// 				Id:   plan.ID,
// 				Name: lo.ToPtr("not-a-test"),
// 			})

// 			assert.Nil(t, err)
// 			assert.NotNil(t, result)
// 			assert.Equal(t, "not-a-test", result.Name)
// 		})
// }

// func TestDAODeletePlan(t *testing.T) {
// 	planDAORunner(
// 		t,
// 		func(machinePlanDAO dao.MachinePlanDaoInterface) {
// 			// TODO: write tests with soft delete considerations

// 			// expected := types.CreateMachinePlanArgs{
// 			// 	Name:        "test",
// 			// 	Currency:    "NGN",
// 			// 	MonthlyRate: 500,
// 			// 	Cpu:         "256Mi",
// 			// 	Memory:      "1024m",
// 			// }
// 			// plan, err := machinePlanDAO.CreateMachinePlan(expected)

// 			// assert.Nil(t, err)
// 			// assert.NotNil(t, plan)

// 			// err = machinePlanDAO.DeleteMachinePlan(types.DeleteMachinePlanArgs{ID: plan.ID})

// 			// assert.Nil(t, err)

// 			// result, err := machinePlanDAO.GetMachinePlan(types.GetMachinePlanArgs{ID: plan.ID})

// 			// assert.NotNil(t, err)
// 			// assert.Nil(t, result)
// 		})
// }
