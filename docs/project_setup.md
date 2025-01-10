## **1. Directory Structure**
Your project directory should look like this:
```
go-server/
├── cmd/
│   └── go-server/
│       └── main.go
├── internal/
│   └── server/
│       ├── server.go
│       └── handlers.go
├── pkg/
│   └── utils/
│       └── utils.go
├── static/
│   ├── index.html
│   └── other.html
├── go.mod
├── go.sum
├── Dockerfile
└── docker-compose.yml
```

---

## **2. Dockerfile**
Create a `Dockerfile` to set up the Go environment and build your web server:

```Dockerfile
# Use the official Golang image as the base image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download and install Go dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o go-server ./cmd/go-server

# Expose the port your application will run on
EXPOSE 8080

# Command to run the application
CMD ["./go-server"]
```

---

## **3. Go Web Server (main.go)**
This is the entry point for your Go web server, located in `cmd/go-server/main.go`:

```go
package main

import (
	"log"
	"go-server/internal/server"
)

func main() {
	// Initialize the server
	srv := server.NewServer()

	// Start the server
	log.Println("Starting server on :8080...")
	if err := srv.Start(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
```

---

## **4. Web Server Logic (internal/server/server.go)**
This file contains the web server logic:

```go
package server

import (
	"net/http"
)

type Server struct {
	router *http.ServeMux
}

func NewServer() *Server {
	srv := &Server{
		router: http.NewServeMux(),
	}
	srv.routes()
	return srv
}

func (s *Server) Start() error {
	return http.ListenAndServe(":8080", s.router)
}

func (s *Server) routes() {
	// Serve static files
	fs := http.FileServer(http.Dir("static"))
	s.router.Handle("/", fs)
}
```

---

## **5. Static HTML Files**
Create a `static/` directory and add your `.html` files. For example:

### **static/index.html**
```html
<!DOCTYPE html>
<html>
<head>
    <title>Go Server</title>
</head>
<body>
    <h1>Welcome to the Go Web Server!</h1>
</body>
</html>
```

### **static/other.html**
```html
<!DOCTYPE html>
<html>
<head>
    <title>Other Page</title>
</head>
<body>
    <h1>This is another page!</h1>
</body>
</html>
```

---

## **5. Docker Compose (docker-compose.yml)**
Create a `docker-compose.yml` file to simplify building and running your Docker container:

```yaml
version: "3.8"
services:
  go-server:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - .:/app
```

---

## **7. Build and Run the Container**
1. Open a terminal in your project directory.
2. Run the following command to build and start the container:
   ```bash
   docker-compose up
   ```
3. The Go web server will start, and you’ll see logs in the terminal.

---

## **8. Access the Web Server**
Once the container is running, open your browser and navigate to:
- **Home Page**: `http://localhost:8080/`
- **Other Page**: `http://localhost:8080/other.html`

---

## **9. Live Code Updates**
Since the project directory is mounted as a volume, you can make changes to your Go code or HTML files, and the changes will be reflected immediately in the running container. For example:
- Edit `main.go` or any `.html` file.
- Restart the server by stopping (`Ctrl+C`) and restarting the container:
  ```bash
  docker-compose up
  ```

---

## **10. Stop and Remove the Container**
To stop and remove the container, run:
```bash
docker-compose down
```

---

## **11. Clean Up**
To remove the Docker image after testing, run:
```bash
docker rmi go-server
```

---

## **12. Using Docker Desktop**
If you prefer using Docker Desktop:
1. Open Docker Desktop.
2. Go to the **Images** tab.
3. Click **Build** and select your project directory (`go-server`).
4. Name the image `go-server` and click **Build**.
5. Go to the **Containers** tab.
6. Click **Add Container**.
7. Select the `go-server` image.
8. Map port `8080` (host) to `8080` (container).
9. Click **Run**.

---