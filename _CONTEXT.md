# Prototype for a microservice using Gin

**Status:** 🟢 Active
**Objective:** A proof of concept (POC) about running a microservice in `Golang`, using `Gin` framework for `Inventory Management`, including integration with `Python` for data analysis.

- Add new items into database.
- Read the total amount of items in the database.
- **Python Integration**: Call external Python scripts for product-specific analysis.

## Engineer Decisions:

- Try to use `basic` Go packages.
- Use `SQLite` (via `modernc.org/sqlite`) for persistence to remain CGO-free.
- Use `Gin` for the HTTP framework.
- Use `os/exec` for **Go-Python inter-process communication**, enabling analytical tasks in Python while retaining Go's high-performance API layer.
- Use **descriptive naming conventions**
- Standardize on **Go documentation comments** for all public constructs.

## 🎯 Next Actions:

- [x] Initialize Go modules and install dependencies
- [x] Create directory structure (models, handlers, storage)
- [x] Implement Item model and SQLite storage (CGO-free)
- [x] Descriptive renaming and documentation
- [x] **Go/Python Integration**:
  - [x] Python: `scripts/analyze_product.py`
  - [x] Go: `services/python_service.go` using `os/exec`
- [x] Run and test all endpoints (v1/items, v1/items/count, v1/items/:id/analysis)

## 📂 Directory Structure

```text
PROTOTYPE-GinServices/
├── go.mod
├── go.sum
├── main.go
├── inventory.db
├── models/
│   └── item.go
├── handlers/
│   └── item.go
├── storage/
│   └── sqlite.go
├── services/
│   └── python_service.go (Python Orchestrator)
├── scripts/
│   └── analyze_product.py (Analytical Engine)
└── test_endpoints.sh
```

## 🔗 Connectivity
- **Upstream:** None
- **Downstream:** Python Interpreter (Analysis Engine)
- **Links:** None

## 🏷️ Semantic Hooks
#golang #gin #microservices #poc #sqlite #python #os-exec