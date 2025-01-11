package controller

import 
    "github.com/gin-gonic/gin"

)

func createUser(c *gin.Context) {

    var userRequest request.UserRequest
    
    
    if err := c.ShouldBindJSON(&userRequest); err!= nil {
        restErr := NewBadRequestError(
            fmt.Sprintf("Invalid JSON body, there are some incorrect fields: %s \n", err.))
            c.JSON(restErr.Code, restErr)
            return
    }
}

