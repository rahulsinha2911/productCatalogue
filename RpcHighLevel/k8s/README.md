# Kubernetes Deployment for HighLevel

Minimal Kubernetes manifests for deploying the HighLevel application.

## Files

- `mysql.yaml` - MySQL Secret, PVC, Deployment, and Service
- `app.yaml` - Application Secret, ConfigMap, Deployment, and Service
- `Makefile` - Make targets for easy deployment (Linux/Mac)
- `deploy-minikube.ps1` - PowerShell script for automated Minikube deployment (Windows)
- `deploy-minikube.sh` - Bash script for automated Minikube deployment (Linux/Mac)
- `cleanup.ps1` - PowerShell script to cleanup deployment (Windows)
- `cleanup.sh` - Bash script to cleanup deployment (Linux/Mac)
- `README.md` - This file

## Prerequisites

### For Minikube Deployment

1. **Minikube** installed ([Installation Guide](https://minikube.sigs.k8s.io/docs/start/))
2. **kubectl** installed and configured
3. **Docker** installed (or another container runtime supported by minikube)

### For Docker Desktop Deployment

1. **Docker Desktop** with Kubernetes enabled
2. **kubectl** installed and configured
3. Docker installed

### Enable Kubernetes in Docker Desktop

1. Open Docker Desktop → Settings → Kubernetes
2. Check "Enable Kubernetes"
3. Click "Apply & Restart"

## Quick Start (Automated Deployment)

### Option A: Using Makefile (Linux/Mac - Recommended)

```bash
# Navigate to k8s folder
cd k8s

# Deploy everything
make deploy-minikube

# Check status
make status

# View logs
make logs

# Access application
make access

# Cleanup
make cleanup
```

### Option B: Using PowerShell Script (Windows)

```powershell
# Navigate to k8s folder
cd k8s

# Run the deployment script
.\deploy-minikube.ps1
```

### Option C: Using Bash Script (Linux/Mac)

```bash
# Navigate to k8s folder
cd k8s

# Make script executable (first time only)
chmod +x deploy-minikube.sh

# Run the deployment script
./deploy-minikube.sh
```

The scripts will automatically:
1. Check prerequisites
2. Start Minikube (if not running)
3. Configure Docker environment
4. Build the Docker image
5. Deploy MySQL and wait for it to be ready
6. Deploy the application and wait for it to be ready
7. Display deployment status

## Deployment Options

### Option 1: Deploy to Minikube (Recommended for Local Development)

#### Step 1: Start Minikube

```powershell
# Start minikube cluster
minikube start

# Verify minikube is running
minikube status

# Configure Docker to use minikube's Docker daemon
minikube docker-env | Invoke-Expression
```

#### Step 2: Build Docker Image

```powershell
# Make sure you're in the project root directory (not k8s folder)
# The image will be built in minikube's Docker environment
docker build -t rpc-highlevel-service:latest .
```

#### Step 3: Deploy to Minikube

```powershell
# Navigate to k8s folder for deployment
cd k8s

# Deploy MySQL
kubectl apply -f mysql.yaml

# Wait for MySQL to be ready
kubectl wait --for=condition=ready pod -l app=mysql --timeout=120s

# Deploy Application
kubectl apply -f app.yaml

# Wait for Application to be ready
kubectl wait --for=condition=ready pod -l app=rpc-highlevel-service --timeout=120s
```

#### Step 4: Access the Application

```powershell
# Option A: Use minikube service (opens browser automatically)
minikube service rpc-highlevel-service

# Option B: Port forward (manual)
kubectl port-forward service/rpc-highlevel-service 8080:8080
```

Then open: **http://localhost:8080** (if using port-forward)

### Option 2: Deploy to Docker Desktop Kubernetes

```powershell
# Build Docker image
# Make sure you're in the project root directory (not k8s folder)
docker build -t rpc-highlevel-service:latest .

# Navigate to k8s folder for deployment
cd k8s

# Deploy MySQL
kubectl apply -f mysql.yaml

# Wait for MySQL to be ready
kubectl wait --for=condition=ready pod -l app=mysql --timeout=120s

# Deploy Application
kubectl apply -f app.yaml

# Wait for Application to be ready
kubectl wait --for=condition=ready pod -l app=rpc-highlevel-service --timeout=120s
```

## Access the Application

### For Minikube

```powershell
# Method 1: Use minikube service (automatically opens browser)
minikube service rpc-highlevel-service

# Method 2: Port forward
kubectl port-forward service/rpc-highlevel-service 8080:8080
```

Then open: **http://localhost:8080** (if using port-forward)

### For Docker Desktop

```powershell
# Port forward
kubectl port-forward service/rpc-highlevel-service 8080:8080
```

Then open: **http://localhost:8080**

### Test API

```powershell
Invoke-WebRequest -Uri "http://localhost:8080/v1/user?user_id=123"
```

## Useful Commands

### General Kubernetes Commands

```powershell
# Check status
kubectl get pods
kubectl get services
kubectl get pods -l app=rpc-highlevel-service

# View logs
kubectl logs -l app=rpc-highlevel-service
kubectl logs -l app=mysql

# Delete everything
kubectl delete -f mysql.yaml -f app.yaml

# Or use the cleanup script
# Windows: .\cleanup.ps1
# Linux/Mac: ./cleanup.sh
```

### Minikube-Specific Commands

```powershell
# Check minikube status
minikube status

# View minikube dashboard
minikube dashboard

# Stop minikube
minikube stop

# Delete minikube cluster
minikube delete

# Get minikube IP
minikube ip

# List all services in minikube
minikube service list

# SSH into minikube VM
minikube ssh
```

## Configuration

- **Database**: MySQL 8.0 with persistent storage (10Gi)
- **Application**: 2 replicas with health checks
- **Service**: LoadBalancer type
  - **Minikube**: Use `minikube service` or port-forward to access
  - **Docker Desktop**: Maps to localhost automatically
