package api

// func SetupAdminRoutes(app *fiber.App, authMiddleware *auth.AuthMiddleware) {
// 	app.Post("/admin/create", authMiddleware.AdminMiddleware, func(c *fiber.Ctx) error {
// 		var request struct {
// 			APIKey string `json:"apiKey"`
// 			Quota  int    `json:"quota"`
// 			Admin  bool   `json:"admin"`
// 		}

// 		if err := c.BodyParser(&request); err != nil {
// 			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid request body"})
// 		}

// 		if request.APIKey == "" || request.Quota == 0 {
// 			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Missing API key or quota"})
// 		}

// 		err := authMiddleware.Storage().CreateUser(request.APIKey, request.Quota, request.Admin)
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
// 		}

// 		return c.Status(201).JSON(fiber.Map{"status": "success"})
// 	})

// 	app.Put("/admin/update-quota", authMiddleware.AdminMiddleware, func(c *fiber.Ctx) error {
// 		var request struct {
// 			APIKey string `json:"apiKey"`
// 			Quota  int    `json:"quota"`
// 		}

// 		if err := c.BodyParser(&request); err != nil {
// 			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Invalid request body"})
// 		}

// 		if request.APIKey == "" || request.Quota == 0 {
// 			return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Missing API key or quota"})
// 		}

// 		err := authMiddleware.Storage().UpdateQuota(request.APIKey, request.Quota)
// 		if err != nil {
// 			return c.Status(500).JSON(fiber.Map{"status": "error", "message": err.Error()})
// 		}

// 		return c.Status(200).JSON(fiber.Map{"status": "success"})
// 	})
// }
