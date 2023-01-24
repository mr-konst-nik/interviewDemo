package server

import (
	"interviewDemo/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "It's works!",
	})
}

// @Summary get an item
// @Description get an item from coin list by symbol ID, like "btc"
// @Tags coins
// @ID get-coin
// @Produce json
// @Param symbol path string true "coin symbol"
// @Success 200 {object} model.MessageOk
// @Failure 404 {object} model.MessageErr
// @Router /coin/info/{symbol} [get]
func coinInfo(c *gin.Context) {
	symbol := c.Param("symbol")
	coin, err := coinService.Read(symbol)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.MessageErr{Status: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.MessageOk{Status: true, Message: "find successfully", Data: coin})
}

// @Summary create a new item
// @Description create a new coin
// @Tags coins
// @ID add-coin
// @Produce json
// @Param data body model.CoinItemRequestBody true "coin data"
// @Success 201 {object} model.MessageOk
// @Failure 400 {object} model.MessageErr
// @Failure 409 {object} model.MessageErr
// @Router /coin/add/ [post]
func coinAdd(c *gin.Context) {
	coinBody := model.CoinItemRequestBody{}

	if err := c.BindJSON(&coinBody); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.MessageErr{Status: false, Message: err.Error()})
		return
	}

	coinItem := model.CoinItem{
		Name:        coinBody.Name,
		TypeID:      coinBody.TypeID,
		ExchangesID: coinBody.ExchangesID,
	}

	if err := coinService.Create(coinBody.SymbolID, &coinItem); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, model.MessageErr{Status: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, model.MessageOk{Status: true, Message: "created successfully", Data: &coinItem})
}

// @Summary update item
// @Description update coin by symbol ID
// @Tags coins
// @ID update-coin
// @Produce json
// @Param symbol path string true "coin symbol"
// @Param data body model.CoinItem true "coin data"
// @Success 200 {object} model.MessageOk
// @Failure 400 {object} model.MessageErr
// @Failure 404 {object} model.MessageErr
// @Router /coin/change/{symbol} [put]
func coinUpdate(c *gin.Context) {
	symbol := c.Param("symbol")
	coinItem := model.CoinItem{}
	if err := c.BindJSON(&coinItem); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.MessageErr{Status: false, Message: err.Error()})
		return
	}

	if err := coinService.Update(symbol, &coinItem); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.MessageErr{Status: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.MessageOk{Status: true, Message: "updated successfully", Data: &coinItem})
}

// @Summary delete item
// @Description delete coin by symbol ID
// @Tags coins
// @ID delete-coin
// @Produce json
// @Param symbol path string true "coin symbol"
// @Success 200 {object} model.MessageOk
// @Failure 404 {object} model.MessageErr
// @Router /coin/delete/{symbol} [delete]
func coinDelete(c *gin.Context) {
	symbol := c.Param("symbol")
	coin, err := coinService.Delete(symbol)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.MessageErr{Status: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.MessageOk{Status: true, Message: "deleted successfully", Data: coin})
}
