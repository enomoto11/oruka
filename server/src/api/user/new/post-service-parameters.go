package new

import (
	"oruka/ent"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Params struct {
	first_name string
	last_name  string
	manavis_id int
	grade      int
}

func GetParameters(context *gin.Context, client *ent.Client) Params {
	first_name := context.Query("first_name")
	last_name := context.Query("last_name")
	manavis_id, _ := strconv.Atoi(context.Param("id"))
	grade, _ := strconv.Atoi(context.Query("grade"))

	params := Params{
		first_name: first_name,
		last_name:  last_name,
		manavis_id: manavis_id,
		grade:      grade,
	}

	// すでにそのマナビス生番号をもつ生徒がいればエラーを返す
	user := GetUser(client, context, params)
	if user != nil {
		apiError := UserAlreadyExistsError()
		context.JSON(apiError.Status, apiError)
	}

	return params
}
