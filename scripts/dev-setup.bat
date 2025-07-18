@echo off
echo 🚀 Setting up NestMate development environment...

REM Check if Docker is installed
docker --version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Docker is not installed. Please install Docker first.
    exit /b 1
)

REM Check if Docker Compose is installed
docker-compose --version >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Docker Compose is not installed. Please install Docker Compose first.
    exit /b 1
)

REM Create necessary directories
echo 📁 Creating necessary directories...
if not exist "backend\data" mkdir backend\data
if not exist "web\build" mkdir web\build
if not exist "NestMateMobile\android\app\build" mkdir NestMateMobile\android\app\build

REM Create backend clean architecture folders
echo �  Creating backend clean architecture folders...
if not exist "backend\internal\domain\entities" mkdir backend\internal\domain\entities
if not exist "backend\internal\domain\repositories" mkdir backend\internal\domain\repositories
if not exist "backend\internal\application\services" mkdir backend\internal\application\services
if not exist "backend\internal\infrastructure\config" mkdir backend\internal\infrastructure\config
if not exist "backend\internal\infrastructure\repositories" mkdir backend\internal\infrastructure\repositories
if not exist "backend\internal\infrastructure\repositories\sqlite" mkdir backend\internal\infrastructure\repositories\sqlite
if not exist "backend\internal\infrastructure\repositories\firestore" mkdir backend\internal\infrastructure\repositories\firestore
if not exist "backend\internal\infrastructure\sync" mkdir backend\internal\infrastructure\sync
if not exist "backend\internal\interfaces\http" mkdir backend\internal\interfaces\http
if not exist "backend\cmd\server" mkdir backend\cmd\server

REM Create web app folders
echo 📁 Creating web app folders...
if not exist "web\src\components" mkdir web\src\components
if not exist "web\src\hooks" mkdir web\src\hooks
if not exist "web\src\services" mkdir web\src\services
if not exist "web\src\types" mkdir web\src\types
if not exist "web\src\utils" mkdir web\src\utils

REM Create mobile app folders
echo 📁 Creating mobile app folders...
if not exist "NestMateMobile\src\screens" mkdir NestMateMobile\src\screens
if not exist "NestMateMobile\src\services" mkdir NestMateMobile\src\services
if not exist "NestMateMobile\src\types" mkdir NestMateMobile\src\types
if not exist "NestMateMobile\src\components" mkdir NestMateMobile\src\components
if not exist "NestMateMobile\src\utils" mkdir NestMateMobile\src\utils

REM Set up backend
echo 🔧 Setting up Go backend...
cd backend
if not exist "go.mod" (
    go mod init nestmate-backend
)
go mod tidy
cd ..

REM Set up web app
echo 🌐 Setting up React web app...
cd web
if not exist "node_modules" (
    npm install
)
cd ..

REM Set up mobile app
echo 📱 Setting up React Native mobile app...
cd NestMateMobile
if not exist "node_modules" (
    npm install
)
cd ..

REM Create environment files
echo 📝 Creating environment files...

REM Backend .env
(
echo PORT=8080
echo DB_DRIVER=sqlite
echo DB_NAME=./data/nestmate.db
echo JWT_SECRET=your-development-secret-key-change-in-production
echo FIREBASE_PROJECT_ID=
echo FIREBASE_PRIVATE_KEY=
echo FIREBASE_CLIENT_EMAIL=
) > backend\.env

REM Web .env
(
echo REACT_APP_API_URL=http://localhost:8080/api/v1
echo REACT_APP_ENVIRONMENT=development
) > web\.env

echo ✅ Development environment setup complete!
echo.
echo 🎯 Next steps:
echo 1. Start the development environment: docker-compose up -d
echo 2. Access the web app at: http://localhost:3000
echo 3. Access the API at: http://localhost:8080
echo 4. Check API health: curl http://localhost:8080/health
echo.
echo 📚 Useful commands:
echo - View logs: docker-compose logs -f
echo - Stop services: docker-compose down
echo - Rebuild services: docker-compose up -d --build

pause