Write-Output "build..."
$Env:GOOS = "linux"
$Env:GOARCH = "amd64"
go build
Write-Output "finished."