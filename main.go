package main
// import "fmt"
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Input struct {
	Input_tenant_id string 
	Input_tool string
}

type Tenant struct {
	Tenant_id string
	Api_key string
	Allowed_tools []string
	Quota int
}

var ai_agents =[]Tenant{
	{Tenant_id:"123", Api_key:"hasjcbjhgu", Allowed_tools:[]string{"getTime", "getWeather"}, Quota:100},
	{Tenant_id:"456", Api_key:"fdsgadhfd", Allowed_tools:[]string{"getTime"}, Quota:200},
	{Tenant_id:"789", Api_key:"dfgdfgdfg", Allowed_tools:[]string{"getCar", "getLocation"}, Quota:50},
	{Tenant_id:"101", Api_key:"safsdfcfg", Allowed_tools:[]string{"getWeather"}, Quota:90},
}

func getTenant(id string) *Tenant {
	for _, tenant := range ai_agents {
		if tenant.Tenant_id== id{
			return &tenant
		}
	}
	return nil
}

func validateTools(Allowed_tools []string, tool string) bool {
	for _, t := range Allowed_tools {
		if t== tool{
			return true
		}
	}
	return false
}
var tools = map[string]string{
"getTime":"11:20pm",
"getWeather":"sunny",
"getCar":"mercedes benz",
"getLocation":"Japan" ,
}

func getToolResult(tool string) string{
	return tools[tool]
}

func main() {
	myrouter := gin.Default()
	myrouter.POST("/my_mcp_gateway", func(response_context *gin.Context){
		var input Input
		if err := response_context.ShouldBindJSON(&input); err != nil {
            response_context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
		tenant := getTenant(input.Input_tenant_id)
		if tenant == nil {
			response_context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Tenant ID"})
			return
		}
		if !validateTools(tenant.Allowed_tools, input.Input_tool){
			response_context.JSON(http.StatusBadRequest, gin.H{"error": "Tool not found"})
			return
		}
		tool_res := getToolResult(input.Input_tool)
		response_context.JSON(http.StatusOK, gin.H{
			"tenant id":tenant.Tenant_id,
			"tool requested":input.Input_tool,
			"result": tool_res,
		})
	})
	myrouter.Run(":8080")

	// if tenant != nil{
	// 	fmt.Println("Tenant ID:", tenant.Tenant_id)
	// 	fmt.Println("Tenant API Key:", tenant.Api_key)
	// 	fmt.Println("Tenant Tools:", tenant.Allowed_tools)
	// 	fmt.Println("Tenant Quota:", tenant.Quota)
	// }

}