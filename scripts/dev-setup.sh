#!/bin/bash

# NestMate Development Setup Script

echo "🚀 Setting up NestMate development environment..."

# Check if Docker is installed
if ! command -v docker &> /dev/null; then
    echo "❌ Docker is not installed. Please install Docker first."
    exit 1
fi

# Check if Docker Compose is installed
if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose is not installed. Please install Docker Compose first."
    exit 1
fi

# Create necessary directories
echo "📁 Creating necessary directories..."
mkdir -p backend/data
mkdir -p web/build
mkdir -p NestMateMobile/android/app/build

# Set up backend
echo "🔧 Setting up Go backend..."
cd backend
if [ ! -f "go.mod" ]; then
    go mod init nestmate-backend
fi
go mod tidy
cd ..

# Set up web app
echo "🌐 Setting up React web app..."
cd web
if [ ! -d "node_modules" ]; then
    npm install
fi
cd ..

# Set up mobile app
echo "📱 Setting up React Native mobile app..."
cd NestMateMobile
if [ ! -d "node_modules" ]; then
    npm install
fi
cd ..

# Create environment files
echo "📝 Creating environment files..."

# Backend .env
cat > backend/.env << EOF
PORT=8080
DB_DRIVER=sqlite
DB_NAME=./data/nestmate.db
JWT_SECRET=your-development-secret-key-change-in-production
FIREBASE_PROJECT_ID=
FIREBASE_PRIVATE_KEY=
FIREBASE_CLIENT_EMAIL=
EOF

# Web .env
cat > web/.env << EOF
REACT_APP_API_URL=http://localhost:8080/api/v1
REACT_APP_ENVIRONMENT=development
EOF

echo "✅ Development environment setup complete!"
echo ""
echo "🎯 Next steps:"
echo "1. Start the development environment: docker-compose up -d"
echo "2. Access the web app at: http://localhost:3000"
echo "3. Access the API at: http://localhost:8080"
echo "4. Check API health: curl http://localhost:8080/health"
echo ""
echo "📚 Useful commands:"
echo "- View logs: docker-compose logs -f"
echo "- Stop services: docker-compose down"
echo "- Rebuild services: docker-compose up -d --build"