package budget

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Budget struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{budgetID}", GetBudget)
	router.Delete("/{budgetID}", DeleteBudget)
	router.Post("/", CreateBudget)
	router.Get("/", GetAllBudgets)
	return router
}

func GetBudget(w http.ResponseWriter, r *http.Request) {
	budgetID := chi.URLParam(r, "budgetID")
	budgets := Budget{
		Slug:  budgetID,
		Title: "Budget Manager",
		Body:  "Heloo world from Budget Manager",
	}
	render.JSON(w, r, budgets) // A chi router helper for serializing and returning json
}

func DeleteBudget(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted Budget successfully"
	render.JSON(w, r, response) // Return some demo response
}

func CreateBudget(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Budget successfully"
	render.JSON(w, r, response) // Return some demo response
}

func GetAllBudgets(w http.ResponseWriter, r *http.Request) {
	budgets := []Budget{
		{
			Slug:  "slug",
			Title: "Hello world",
			Body:  "Heloo world from planet earth",
		},
	}
	render.JSON(w, r, budgets) // A chi router helper for serializing and returning json
}
