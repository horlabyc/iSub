package validations

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type SignupCto struct {
	Email           string `json:"email" validate:"required,email"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required,min=6,max=32,alphanum"`
	ConfirmPassword string `json:"confirm_password" validate:"eqfield=Password,required"`
}

func RegisterValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		// data, _ := ioutil.ReadAll(c.Request.Body)
		// fmt.Println(string(data))
		var newUser SignupCto
		// fmt.Println("Validator")
		// var jsonData map[string]interface{} // map[string]interface{}
		// if e := json.Unmarshal(data, &jsonData); e != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"msg": e.Error()})
		// 	return
		// }
		// fmt.Printf("%+v\n", jsonData)
		if err := c.ShouldBindBodyWith(&newUser, binding.JSON); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
		}
		validationErr := validator.New().Struct(newUser)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}
