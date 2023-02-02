package http

import (
	"interviewDemo/internal/model"
	"interviewDemo/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler defines handlers
type Handler struct {
	service service.CoinService
}

// NewHandlers creates new HTTP handlers
func NewHandlers(service service.CoinService) *Handler {
	return &Handler{service: service}
}

// SayHello handles / request
func (h *Handler) SayHello(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "It's works!",
	})
}

// CoinInfo handles GET /coin/info/{symbol} request
// @Summary get an item
// @Description get an item from coin list by symbol ID, like "btc"
// @Tags coins
// @ID get-coin
// @Produce json
// @Param symbol path string true "coin symbol"
// @Success 200 {object} model.MessageOk
// @Failure 404 {object} model.MessageErr
// @Router /coin/info/{symbol} [get]
func (h *Handler) CoinInfo(c *gin.Context) {
	symbol := c.Param("symbol")
	coin, err := h.service.Read(c, symbol)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.MessageErr{Status: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.MessageOk{Status: true, Message: "find successfully", Data: coin})
}

// CoinAdd handles POST /coin/add/ request
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
func (h *Handler) CoinAdd(c *gin.Context) {
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

	if err := h.service.Create(c, coinBody.SymbolID, &coinItem); err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, model.MessageErr{Status: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.MessageOk{Status: true, Message: "created successfully", Data: &coinItem})
}

// CoinUpdate handles PUT /coin/change/{symbol}
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
func (h *Handler) CoinUpdate(c *gin.Context) {
	symbol := c.Param("symbol")
	coinItem := model.CoinItem{}
	if err := c.BindJSON(&coinItem); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.MessageErr{Status: false, Message: err.Error()})
		return
	}

	if err := h.service.Update(c, symbol, &coinItem); err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.MessageErr{Status: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.MessageOk{Status: true, Message: "updated successfully", Data: &coinItem})
}

// CoinDelete handles DELETE /coin/delete/{symbol}
// @Summary delete item
// @Description delete coin by symbol ID
// @Tags coins
// @ID delete-coin
// @Produce json
// @Param symbol path string true "coin symbol"
// @Success 200 {object} model.MessageOk
// @Failure 404 {object} model.MessageErr
// @Router /coin/delete/{symbol} [delete]
func (h *Handler) CoinDelete(c *gin.Context) {
	symbol := c.Param("symbol")
	coin, err := h.service.Delete(c, symbol)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, model.MessageErr{Status: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.MessageOk{Status: true, Message: "deleted successfully", Data: coin})
}
