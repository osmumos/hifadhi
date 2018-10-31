package expense

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Expense struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{expenseID}", GetExpense)
	router.Delete("/{expenseID}", DeleteExpense)
	router.Post("/", CreateExpense)
	router.Get("/", GetAllExpenses)
	return router
}

func GetExpense(w http.ResponseWriter, r *http.Request) {
	expenseID := chi.URLParam(r, "expenseID")
	expenses := Expense{
		Slug:  expenseID,
		Title: "Expense Manager",
		Body:  "Heloo world from Expense Manager",
	}
	render.JSON(w, r, expenses) // A chi router helper for serializing and returning json
}

func DeleteExpense(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted Expense successfully"
	render.JSON(w, r, response) // Return some demo response
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Expense successfully"
	render.JSON(w, r, response) // Return some demo response
}

func GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	expenses := []Expense{
		{
			Slug:  "slug",
			Title: "Hello world",
			Body:  "Heloo world from planet earth",
		},
	}
	render.JSON(w, r, expenses) // A chi router helper for serializing and returning json
}
