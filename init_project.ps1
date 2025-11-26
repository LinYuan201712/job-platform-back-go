# 1. Init Go Module
if (-not (Test-Path go.mod)) {
    go mod init job-platform-go
}

# 2. Create directories
$dirs = @(
    "cmd/api",
    "internal/config",
    "internal/controller",
    "internal/service",
    "internal/repository",
    "internal/model",
    "internal/middleware",
    "internal/pkg/utils",
    "configs"
)
foreach ($dir in $dirs) { 
    New-Item -ItemType Directory -Force -Path $dir | Out-Null 
    Write-Host "Created directory: $dir"
}

# 3. Create core files
$files = @(
    "cmd/api/main.go",
    "internal/config/config.go",
    "configs/application.yaml",
    "README.md",
    ".gitignore"
)
foreach ($file in $files) { 
    if (-not (Test-Path $file)) {
        New-Item -ItemType File -Force -Path $file | Out-Null
        Write-Host "Created file: $file"
    }
}

# 4. Install dependencies
Write-Host "Downloading dependencies..." -ForegroundColor Cyan
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u github.com/spf13/viper
go get -u github.com/golang-jwt/jwt/v5

Write-Host "Project initialization finished! Please open with VS Code." -ForegroundColor Green