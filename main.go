package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	tempclass "assisthan/class"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"192.168.1.0/24"})
	router.GET("/temp", getTemp)
	//router.StaticFile("/weather", "./yr-response.json")
	router.GET("/weather", func(c *gin.Context) {
		//data, err := os.ReadFile("./yr-response.json")
		data, err := os.ReadFile("./yr_no.json")
		if err != nil {
			// error handler
		}

		c.Header("Content-Type", "application/json")

		_, _ = c.Writer.Write(data)
	})
	router.Run("0.0.0.0:9090")
}

func getWeather(context *gin.Context) {

	context.IndentedJSON(http.StatusOK, "")
}

func getTemp(context *gin.Context) {

	// construct `go version` command
	cmd := exec.Command("vcgencmd", "measure_temp")

	// run command
	if output, err := cmd.Output(); err != nil {
		fmt.Println("Error:", err)
		context.IndentedJSON(http.StatusBadRequest, err)
	} else {
		fmt.Printf("Output: %s\n", output)

		context.IndentedJSON(http.StatusOK, formatTemp(string(output)))
	}
}

func formatTemp(cmdOutput string) *tempclass.Temp {
	tempSplit := strings.Split(cmdOutput, "=")[1]
	//tempPart := strings.Split(tempSplit, "'")[0]
	actualTemp, _ := strconv.ParseFloat(strings.Split(tempSplit, "'")[0], 32)
	var f32Temp float32 = float32(actualTemp)
	temp := tempclass.Temp{Temp: f32Temp}

	return &temp
}
