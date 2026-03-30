# 📦 Inventory Management Microservice (POC)

A high-performance **Golang** microservice for managing inventory items, featuring a persistent **SQLite** database and integrated **Python** analytical tools.

> [!NOTE]
> **Disclaimer**: This project is a personal **prototype developed for learning and experimentation purposes**. It is intended as a proof-of-concept (POC) to explore Go microservices, CGO-free SQLite integration, and Python interoperability.

## 🚀 Key Features
- **Fast & Reliable**: Built with the [Gin Web Framework](https://github.com/gin-gonic/gin) for high-performance HTTP routing.
- **Embedded Persistence**: Uses a CGO-free **SQLite** driver (`modernc.org/sqlite`) for zero-dependency database storage.
- **Go-Python Hybrid**: Seamlessly executes Python scripts from Go via `os/exec` to power secondary logic like data analysis.
- **Modern Python Tooling**: Python scripts and dependencies are managed through **uv** for isolation and speed.

## 🛠️ Technology Stack
- **Languages**: Golang, Python 3
- **Framework**: [Gin](https://gin-gonic.com/)
- **Database**: SQLite 3
- **Environment Management**: [uv](https://github.com/astral-sh/uv)

---

## 🏃 Getting Started

### Prerequisites
1.  **Go** (1.20+) installed.
2.  **Python 3** installed.
3.  **uv** (Python profile manager) installed.

### Installation
Clone the repository and install the Go dependencies:
```bash
go mod tidy
```

### Running the Microservice
Start the server directly using Go:
```bash
go run main.go
```
The service will start on `http://localhost:8080`.

---

## 📡 API Endpoints

| Method | Endpoint | Description |
| :--- | :--- | :--- |
| `POST` | `/v1/items` | Add a new product to the inventory. |
| `GET` | `/v1/items/count` | Retrieve the total number of products in stock. |
| `GET` | `/v1/items/:id/analysis` | Fetch item info from DB + extended Python analysis. |
| `GET` | `/health` | Service health check. |

### Sample JSON Payload (for POST /v1/items)
```json
{
  "product_name": "Premium Laptop",
  "stock_count": 10,
  "item_details": "Next-gen creative workstation"
}
```

---

## 🧪 Testing the Service

A comprehensive test script is included to quickly verify all endpoints:
```bash
chmod +x test_endpoints.sh
./test_endpoints.sh
```

## 🏗️ Project Structure
```text
PROTOTYPE-GinServices/
├── main.go               # Server initialization & routing
├── models/               # Data structures (Item)
├── handlers/             # HTTP logic & input validation
├── storage/              # SQLite database layer
├── services/             # Go/Python orchestration (using uv)
├── scripts/              # Python analytical engine
├── inventory.db          # SQLite database (created on first run)
├── pyproject.toml        # uv configuration for Python
├── test_endpoints.sh     # Automation for verification
└── README.md             # This file!
```

## 🏷️ Tags
#golang #gin #sqlite #python #uv #microservices

---

## 📜 License & Credits
- **Designed & Implemented with [Antigravity](https://google.com)**: A collaborative POC built with an AI-first coding assistant.
- This project is released under the [MIT License](LICENSE.md).
