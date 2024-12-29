package routes

import (
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/logger"
	// "github.com/kgermando/mspos-api/controllers/area"
	// "github.com/kgermando/mspos-api/controllers/asm"
	// authentification "github.com/kgermando/mspos-api/controllers/auth"
	// "github.com/kgermando/mspos-api/controllers/dashboard"
	// "github.com/kgermando/mspos-api/controllers/manager"
	// poss "github.com/kgermando/mspos-api/controllers/pos"
	// "github.com/kgermando/mspos-api/controllers/posform"
	// "github.com/kgermando/mspos-api/controllers/province"
	// "github.com/kgermando/mspos-api/controllers/sup"
	// userlogs "github.com/kgermando/mspos-api/controllers/user_logs"
	// "github.com/kgermando/mspos-api/controllers/userr"
	// "github.com/kgermando/mspos-api/middlewares"
)

func Setup(app *fiber.App) {

	// api := app.Group("/api", logger.New())

	
	// Authentification controller
	// auth := api.Group("/auth")
	// auth.Post("/register", authentification.Register)
	// auth.Post("/login", authentification.Login) 
	// auth.Post("/forgot-password", authentification.Forgot)
	// auth.Post("/reset/:token", authentification.ResetPassword) 

	// app.Use(middlewares.IsAuthenticated)

	// auth.Get("/user", authentification.AuthUser)
	// auth.Put("/profil/info", authentification.UpdateInfo)
	// auth.Put("/change-password", authentification.ChangePassword)
	// auth.Post("/logout", authentification.Logout)

	// // Users controller
	// user := api.Group("/users")
	// user.Get("/all", userr.GetAllUsers)
	// user.Get("/all/paginate", userr.GetPaginatedUsers)
	// user.Get("/all/paginate/province/:province_id", userr.GetUserByProvinceID)
	// user.Get("/all/paginate/sup/:sup_id", userr.GetUserBySupID)
	// user.Get("/all/:id", userr.GetUserByID)
	// user.Post("/create", userr.CreateUser)
	// user.Get("/get/:id", userr.GetUser)
	// user.Put("/update/:id", userr.UpdateUser)
	// user.Delete("/delete/:id", userr.DeleteUser) 

	// // Province controller
	// prov := api.Group("/provinces")
	// prov.Get("/all", province.GetAllProvinces)
	// prov.Get("/all/paginate", province.GetPaginatedProvince)
	// prov.Get("/all/dropdown", province.GetProvinceDropdown)
	// prov.Get("/all/:id", province.GetProvinceByID)
	// prov.Post("/create", province.CreateProvince)
	// prov.Get("/get/:id", province.GetProvince)
	// prov.Put("/update/:id", province.UpdateProvince)
	// prov.Delete("/delete/:id", province.DeleteProvince)

	// // Areas controller
	// ar := api.Group("/areas")
	// ar.Get("/all", area.GetAllAreas)
	// ar.Get("/all/paginate", area.GetPaginatedAreas)
	// ar.Get("/all/dropdown", area.GetAreaDropdown) 
	// ar.Get("/all/paginate/province/:province_id", area.GetAreaByProvinceID)
	// ar.Get("/all/:id", area.GetAreaByID)
	// ar.Get("/all-area/:id", area.GetSupAreaByID)
	// ar.Post("/create", area.CreateArea)
	// ar.Get("/get/:id", area.GetArea)
	// ar.Put("/update/:id", area.UpdateArea)
	// ar.Delete("/delete/:id", area.DeleteArea)

	// // ASM controller
	// as := api.Group("/asms")
	// as.Get("/all", asm.GetAllAsms)
	// as.Get("/all/paginate", asm.GetPaginatedASM)
	// // as.Get("/all/:id", asm.GetAsmByID)
	// as.Post("/create", asm.CreateAsm)
	// as.Get("/get/:id", asm.GetAsm)
	// as.Put("/update/:id", asm.UpdateAsm)
	// as.Delete("/delete/:id", asm.DeleteAsm)

	// // Manager controller
	// ma := api.Group("/managers")
	// ma.Get("/all", manager.GetAllManagers)
	// ma.Get("/all/paginate", manager.GetPaginatedManager)
	// // ma.Get("/all/:id", manager.GetManagerByID)
	// ma.Post("/create", manager.CreateManager)
	// ma.Get("/get/:id", manager.GetManager)
	// ma.Put("/update/:id", manager.UpdateManager)
	// ma.Delete("/delete/:id", manager.DeleteManager)

	// // Posforms controller
	// posf := api.Group("/posforms")
	// posf.Get("/all", posform.GetAllPosforms)
	// posf.Get("/all/paginate", posform.GetPaginatedPosForm)
	// posf.Get("/all/paginate/:user_id", posform.GetPosformByID)
	// posf.Get("/all/paginate/province/:province_id", posform.GetPosformByProvinceID)
	// posf.Get("/all/paginate/sup/:area_id", posform.GetPosformBySupID)
	// posf.Post("/create", posform.CreatePosform)
	// posf.Get("/get/:id", posform.GetPosform)
	// posf.Put("/update/:id", posform.UpdatePosform)
	// posf.Delete("/delete/:id", posform.DeletePosform)

	// // Sup controller
	// su := api.Group("/sups")
	// su.Get("/all", sup.GetAllSups)
	// su.Get("/all/paginate", sup.GetPaginatedSups)
	// su.Get("/all/paginate/province/:province_id", sup.GetSupByProvinceID) 
	// su.Get("/all-asm/:asm_id", sup.GetSupASMByID)
	// su.Post("/create", sup.CreateSup)
	// su.Get("/get/:id", sup.GetSup)
	// su.Put("/update/:id", sup.UpdateSup)
	// su.Delete("/delete/:id", sup.DeleteSup)

	// // Pos controller
	// po := api.Group("/pos")
	// po.Get("/all", poss.GetAllPoss)
	// po.Get("/all/paginate", poss.GetPaginatedPos)
	// po.Get("/all/paginate/:user_id", poss.GetPosPaginateByID)
	// po.Get("/all/paginate/province/:province_id", poss.GetPosByProvinceID)
	// po.Get("/all/paginate/sup/:sup_id", poss.GetPosBySupID)
	// po.Get("/all/search/:search", poss.GetAllPosSearch)
	// po.Get("/all/:id", poss.GetPosByID)
	// po.Get("/all-area/:id", poss.GetPosAreaByID)
	// po.Post("/create", poss.CreatePos)
	// po.Get("/get/:id", poss.GetPos)
	// po.Put("/update/:id", poss.UpdatePos)
	// po.Delete("/delete/:id", poss.DeletePos)

	// // UserLogs controller
	// userLog := api.Group("/users-logs")
	// userLog.Get("/all", userlogs.GetUserLogs)
	// userLog.Get("/all/paginate", userlogs.GetPaginatedUserLogs)
	// userLog.Get("/all/paginate/:user_id", userlogs.GetUserLogByID)
	// userLog.Post("/create", userlogs.CreateUserLog)
	// userLog.Get("/get/:id", userlogs.GetUserLog)
	// userLog.Put("/update/:id", userlogs.UpdateUserLog)
	// userLog.Delete("/delete/:id", userlogs.DeleteUserLog)

	// dash := api.Group("/dashboard")

	// nd := dash.Group("/numeric-distribution")
	// nd.Get("/table-view/:province/:start_date/:end_date", dashboard.NdTableView)
	// nd.Get("/nd-year/:province", dashboard.NdByYear)

	// sum := dash.Group("/sammury")
	// sum.Get("/dr-count", dashboard.DrCount)
	// sum.Get("/pos-count", dashboard.POSCount)
	// sum.Get("/province-count", dashboard.ProvinceCount)
	// sum.Get("/area-count", dashboard.AreaCount)
	// sum.Get("/sos-pie/:start_date/:end_date", dashboard.SOSPie)
	// sum.Get("/tracking-visit-dr/:days/:start_date/:end_date", dashboard.TrackingVisitDRS)
	// sum.Get("/summary-chart-bar/:start_date/:end_date", dashboard.SummaryChartBar)
	// sum.Get("/better-dr/:start_date/:end_date", dashboard.BetterDR)
	// sum.Get("/better-supervisor/:start_date/:end_date", dashboard.BetterSup)
	// sum.Get("/status-equements/:start_date/:end_date", dashboard.StatusEquipement)
	// sum.Get("/google-maps/:start_date/:end_date", dashboard.GoogleMaps)
	// sum.Get("/price-sales/:start_date/:end_date", dashboard.PriceSale)

	// sos := dash.Group("/share-of-stock")
	// sos.Get("/sos-pie/:province/:start_date/:end_date", dashboard.SOSPieByArea)
	// sos.Get("/sos-year/:province", dashboard.SOSByYear)
	// sos.Get("/table-view/:province/:start_date/:end_date", dashboard.SOSTableView)

}
