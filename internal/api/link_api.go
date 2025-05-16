package api

import (
	"demo/dto/responses"
	"demo/internal/interfaces"
	"demo/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LinksHandler struct {
	service interfaces.LinkService
}

func NewLinkHandler(s interfaces.LinkService) *LinksHandler {
	return &LinksHandler{
		service: s,
	}
}

func (lh *LinksHandler) GetAll(c *gin.Context) {
	links, err := lh.service.GetAll()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[[]*models.Link](err))
		return
	}

	c.IndentedJSON(http.StatusOK, responses.SuccessResponse(links))
}

func (lh *LinksHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	link, err := lh.service.GetByID(uuid.MustParse(id))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[*models.Link](err))
		return
	}

	c.IndentedJSON(http.StatusOK, responses.SuccessResponse(link))
}

func (lh *LinksHandler) GetByCode(c *gin.Context) {
	code := c.Param("code")

	link, err := lh.service.GetByCode(code)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[*models.Link](err))
		return
	}

	c.Redirect(http.StatusMovedPermanently, link.Url)
}

func (lh *LinksHandler) CreateLink(c *gin.Context) {
	var newLink models.Link
	if err := c.BindJSON(&newLink); err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[*models.Link](err))
		return
	}

	link, err := lh.service.CreateNewLink(newLink.Url)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, responses.ErrorResponse[*models.Link](err))
		return
	}

	c.IndentedJSON(http.StatusOK, responses.SuccessResponse(link))
}
