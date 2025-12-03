package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/siddhesht795/studentApiGo/internal/types"
	"github.com/siddhesht795/studentApiGo/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var student types.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		fmt.Println(err)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
			return
		}

		if err != nil {
			response.WriteJson(w, http.StatusBadRequest, response.GeneralError(err))
		}

		// validate the request
		if err := validator.New().Struct(student); err != nil {
			validatedErrs := err.(validator.ValidationErrors)
			response.WriteJson(w, http.StatusBadRequest, response.ValidationErrors(validatedErrs))
			return
		}

		slog.Info("Creating a student...")

		// w.Write([]byte("welcome to students api"))

		response.WriteJson(w, http.StatusCreated, map[string]string{
			"success": "OK",
		})
	}
}
