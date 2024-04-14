package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func RestHandler(ctx *gin.Context) {
	// Metadata for rest service
	var endpoint = os.Getenv("REST_REQUEST_ENDPOINT")
	var route = "/request"

	// Query string and error
	var err error
	var qs = new(CommonQuery)
	var respJSON = new(RestResponse)
	var resp *http.Response

	// Array list
	var results = []int{}

	// Bind querystring to struct
	if err = ctx.ShouldBindQuery(qs); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i := qs.From; i < qs.To; i++ {
		resp, err = http.Get(fmt.Sprintf("%v%v?to=%v", endpoint, route, i))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err = json.NewDecoder(resp.Body).Decode(respJSON); err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		results = append(results, respJSON.ResponseNumber)
		resp.Body.Close() // Response body should be close: https://pkg.go.dev/net/http
	}
	ctx.JSON(http.StatusOK, gin.H{"datas": results})
}
