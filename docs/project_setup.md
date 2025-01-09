## **1. Directory Structure**
Your project directory should look like this:
```
go-server/
├── Dockerfile
├── docker-compose.yml
├── main.go
├── static/
│   ├── index.html
│   └── other.html
```

- **Dockerfile**: Defines the Docker image for your Go web server.
- **docker-compose.yml**: Simplifies building and running the Docker container.
- **main.go**: Your Go web server code.
- **static/**: A directory containing your `.html` files.

---

## **2. Dockerfile**
Create a `Dockerfile` to set up the Go environment and build your web server:

```Dockerfile
# Use the official Golang image as the base image
FROM golang:1.21-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Download and install Go dependencies
RUN go mod init go-server
RUN go mod tidy

# Build the Go application
RUN go build -o go-server .

# Expose the port your application will run on
EXPOSE 8080

# Command to run the application
CMD ["./go-server"]
```

---

## **3. Go Web Server (main.go)**
Here’s an example of a Go web server that serves static files from the `static/` directory:

```go
package main

import (
	"net/http"
)

func main() {
	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Start the web server on port 8080
	http.ListenAndServe(":8080", nil)
}
```

---

## **4. Static HTML Files**
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

### **Explanation**:
- **build**: Builds the Docker image using the `Dockerfile` in the current directory.
- **ports**: Maps port `8080` on your host to port `8080` in the container.
- **volumes**: Mounts your local project directory into the container for live code updates.

---

## **6. Build and Run the Container**
1. Open a terminal in your project directory.
2. Run the following command to build and start the container:
   ```bash
   docker-compose up
   ```
3. The Go web server will start, and you’ll see logs in the terminal.

---

## **7. Access the Web Server**
Once the container is running, open your browser and navigate to:
- **Home Page**: `http://localhost:8080/`
- **Other Page**: `http://localhost:8080/other.html`

You should see the HTML content served by your Go web server.

---

## **8. Live Code Updates**
Since the project directory is mounted as a volume, you can make changes to your Go code or HTML files, and the changes will be reflected immediately in the running container. For example:
- Edit `main.go` or any `.html` file.
- Restart the server by stopping (`Ctrl+C`) and restarting the container:
  ```bash
  docker-compose up
  ```

---

## **9. Stop and Remove the Container**
To stop and remove the container, run:
```bash
docker-compose down
```

---

## **10. Clean Up**
To remove the Docker image after testing, run:
```bash
docker rmi go-server
```

---

## **11. Using Docker Desktop**
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