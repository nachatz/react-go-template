package handlers

import (
	"github.com/nachatz/React-Go-Template/app/internal/models"
	"github.com/nachatz/React-Go-Template/app/internal/utils"

	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	/* HellowHandler - Returns hello world for a authenticated endpoint
	   @Param w - The http.ResponseWriter to write the response to.
	   @Param r - The *http.Request representing the incoming request.
	*/

	var response models.Response
	response.Message = "Hello World!"
	response.StatusCode = http.StatusOK

	utils.WriteResponse(w, response)
}
