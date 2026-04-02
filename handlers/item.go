package handlers

import (
	"fmt"
	"net/http"

	"strconv"

	"github.com/ducphu/inventory-service/models"
	"github.com/ducphu/inventory-service/services"
	"github.com/ducphu/inventory-service/storage"
	"github.com/gin-gonic/gin"
)

// ItemHandler contains the dependencies needed for item-related HTTP operations.
// It acts as the bridge between user requests and database operations.
type ItemHandler struct {
	storage *storage.SQLiteStorage
}

// NewItemHandler initializes a new handler with the given database storage.
func NewItemHandler(storage *storage.SQLiteStorage) *ItemHandler {
	return &ItemHandler{storage: storage}
}

// AddItem is the HTTP handler for processing POST /items requests.
// It parses the incoming JSON, validates it, and saves it that the inventory.
//
// Responses:
// - 201 Created: Successfully saved the item.
// - 400 Bad Request: Failed to parse or validate JSON input.
// - 500 Internal Server Error: Error communicating with the database.
func (handler *ItemHandler) AddItem(context *gin.Context) {
	var item models.Item
	// Try to bind JSON to the Item model struct
	if err := context.ShouldBindJSON(&item); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Persist the item that the database
	id, err := handler.storage.AddItem(item)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to database"})
		return
	}

	// Respond with the newly created item, including its ID
	item.ID = id
	context.JSON(http.StatusCreated, item)
}

// GetTotal is the HTTP handler for processing GET /items/count requests.
// It summarizes the inventory by returning the total number of items recorded.
//
// Responses:
// - 200 OK: Returns the total count as a JSON object.
// - 500 Internal Server Error: Error querying the database.
func (handler *ItemHandler) GetTotal(context *gin.Context) {
	count, err := handler.storage.GetTotalCount()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve total item count"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"total_items": count})
}

// GetAnalysis is the HTTP handler for processing GET /items/:id/analysis requests.
// It retrieves the item from the DB and then executes a Python script for further analysis.
//
// Responses:
// - 200 OK: Combined DB object and Python-provided analysis.
// - 404 Not Found: Item with the specified ID doesn't exist.
// - 500 Internal Server Error: Database or Python execution failure.
func (handler *ItemHandler) GetAnalysis(context *gin.Context) {
	// 1. Extract item ID from parameter and convert it to integer.
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid item ID, must be a number"})
		return
	}

	// 2. Fetch the item details from the database storage.
	item, err := handler.storage.GetItemByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Database error fetching item info"})
		return
	}
	if item == nil {
		context.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Item with ID %d not found", id)})
		return
	}

	// 3. Call the Python-based analysis service for this specific product name.
	analysis, err := services.AnalyzeWithPython(item.ProductName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Python analysis failed: %v", err)})
		return
	}

	// 4. Respond with a combined perspective.
	context.JSON(http.StatusOK, gin.H{
		"item_info":       item,
		"python_analysis": analysis,
	})
}

// ListItems is the HTTP handler for processing GET /items requests.
// It retrieves all items from the database storage.
func (handler *ItemHandler) ListItems(context *gin.Context) {
	items, err := handler.storage.GetAllItems()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Database error fetching items"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"items": items})
}
