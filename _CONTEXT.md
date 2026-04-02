# Prototype for a microservice using Gin

**Status:** 🟢 Active
**Objective:** A proof of concept (POC) about running a microservice in `Golang`, using `Gin` framework for `Inventory Management`, including integration with `Python` for data analysis.

- Add and Retrieve items from a persistent SQLite database.
- List all items recorded in the inventory.
- **Python Integration**: Execute on-demand product-specific analysis using a dedicated Python engine.
- **Microservices Deployment**: Orchestrated to communicate with the `PROTOTYPE-WebApp` over a shared Docker network.

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
- [x] Run and test all endpoints (v1/items, v1/items:id/analysis, v1/items/count)
- [x] **Add endpoint for listing all items** (v1/items)
- [x] **Network Integration**:
  - [x] Standardize Docker network name: `prototype-webapp-service`
  - [x] Ensure external accessibility for downstream applications
- [x] **Final Documentation and Verification**

## 📂 Directory Structure

```text
PROTOTYPE-GinServices/
├── main.go               # Server initialization
├── Dockerfile            # Multi-stage container definition
├── docker-compose.yml    # Service orchestration
├── inventory.db          (local storage)
├── ...
```

## 🔗 Connectivity
- **Upstream:** `PROTOTYPE-WebApp` (FastAPI backend)
- **Downstream:** Python Interpreter (Internal Analysis Engine)
- **Shared Network:** `prototype-webapp-service`

## 🏷️ Semantic Hooks
#golang #gin #microservices #poc #sqlite #python #os-exec