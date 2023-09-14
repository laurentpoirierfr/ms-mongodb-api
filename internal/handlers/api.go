package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gookit/slog"
	"github.com/laurentpoirierfr/ms-mongodb-api/internal/core/ports"
)

type ApiHandler struct {
	repository ports.Repository
}

func NewApiHandler(repository ports.Repository) *ApiHandler {
	return &ApiHandler{
		repository: repository,
	}
}

// ListDocuments example
//
//	@Summary	List documents from mongodb
//	@Tags		api
//	@Accept		json
//	@Produce	json
//	@Param		documents	path		string	true	"Collection name"
//	@Param		offset	query		int	false	"offset for search, default 0"
//	@Param		limit	query		int	false	"offset for search, default 10"
//	@Success	200	{array}	interface{}	"ok"
//	@Failure	500	{object}	interface{}
//	@Router		/api/{documents} [get]
func (api *ApiHandler) FindDocuments(c *gin.Context) {
	Documents := c.Param("documents")
	offset, limit := api.getPaginationParams(c)
	documents, err := api.repository.GetDocuments(context.TODO(), Documents, offset, limit)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": documents})
}

// Get Document by id
//
//	@Summary	Get document by id from mongodb
//	@Tags		api
//	@Accept		json
//	@Produce	json
//	@Param		documents	path		string	true	"Collection name"
//	@Param		id	path		string	true	"Document ID"
//	@Success	200	{array}	interface{}	"ok"
//	@Failure	500	{object}	interface{}
//	@Router		/api/{documents}/{id} [get]
func (api *ApiHandler) FindOneDocument(c *gin.Context) {
	Documents := c.Param("documents")
	id := c.Param("id")
	document, err := api.repository.GetDocumentById(context.TODO(), Documents, id)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": document})
}

// Create Document
//
//	@Summary	Create documents from mongodb
//	@Tags		api
//	@Accept		json
//	@Produce	json
//	@Param		documents	path		string	true	"Collection name"
//	@Success	201	{array}	interface{}	"ok"
//	@Failure	500	{object}	interface{}
//	@Router		/api/{documents} [post]
func (api *ApiHandler) CreateDocument(c *gin.Context) {
	Documents := c.Param("documents")
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var document interface{}
	err = json.Unmarshal(jsonData, &document)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := api.repository.CreateDocument(context.TODO(), Documents, document)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Put Document by id
//
//	@Summary	Get document by id from mongodb
//	@Tags		api
//	@Accept		json
//	@Produce	json
//	@Param		documents	path		string	true	"Collection name"
//	@Param		id	path		string	true	"Document ID"
//	@Success	200	{array}	interface{}	"ok"
//	@Failure	500	{object}	interface{}
//	@Router		/api/{documents}/{id} [put]
func (api *ApiHandler) UpdateDocument(c *gin.Context) {
	Documents := c.Param("documents")
	id := c.Param("id")
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var document interface{}
	err = json.Unmarshal(jsonData, &document)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := api.repository.UpdateDocument(context.TODO(), Documents, document, id)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

// Delete a Document by id
//
//	@Summary	Get document by id from mongodb
//	@Tags		api
//	@Accept		json
//	@Produce	json
//	@Param		documents	path		string	true	"Collection name"
//	@Param		id	path		string	true	"Document ID"
//	@Success	200	{array}	interface{}	"ok"
//	@Failure	500	{object}	interface{}
//	@Router		/api/{documents}/{id} [delete]
func (api *ApiHandler) DeleteDocument(c *gin.Context) {
	Documents := c.Param("documents")
	id := c.Param("id")
	result, err := api.repository.DeleteDocument(context.TODO(), Documents, id)
	if err != nil {
		slog.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (api *ApiHandler) getPaginationParams(c *gin.Context) (int64, int64) {
	offset := int64(0)
	limit := int64(10) // Default limit

	if offsetStr := c.Query("offset"); offsetStr != "" {
		offset, _ = strconv.ParseInt(offsetStr, 10, 64)
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		limit, _ = strconv.ParseInt(limitStr, 10, 64)
	}

	return offset, limit
}
