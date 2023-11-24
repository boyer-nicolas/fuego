package views

import (
	"github.com/go-fuego/fuego"
	"github.com/go-fuego/fuego/middleware/basicauth"
)

func (rs Ressource) Routes(s *fuego.Server) {
	// Public Pages
	fuego.GetStd(s, "/recipes-std", rs.showRecipesStd)
	fuego.Get(s, "/", rs.showIndex)
	fuego.Get(s, "/recipes", rs.showRecipes)
	fuego.Post(s, "/recipes-new", rs.addRecipe)
	fuego.Get(s, "/ingredients", rs.showIngredients)
	fuego.Get(s, "/html", rs.showHTML)
	fuego.Get(s, "/h1string", rs.showString)

	// Public Chunks
	fuego.Get(s, "/recipes-list", rs.showRecipesList)
	fuego.Get(s, "/search", rs.searchRecipes)

	// Admin Pages
	adminRoutes := fuego.Group(s, "/admin")
	fuego.UseStd(adminRoutes, basicauth.New(basicauth.Config{Username: "admin", Password: "admin"}))
	fuego.Get(adminRoutes, "/", rs.pageAdmin)
	fuego.Get(adminRoutes, "/recipes", rs.adminRecipes)
	fuego.Get(adminRoutes, "/recipes/one", rs.adminOneRecipe)
	fuego.Post(adminRoutes, "/recipes-new", rs.adminAddRecipes)
	fuego.Get(adminRoutes, "/ingredients", rs.adminIngredients)
	fuego.Post(adminRoutes, "/ingredients-new", rs.adminAddIngredient)
	fuego.Get(adminRoutes, "/users", rs.adminRecipes)

	// Admin Chunks
	fuego.Delete(adminRoutes, "/recipes/del", rs.deleteRecipe)
}