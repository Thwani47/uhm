$owner = "Thwani47"
$repo = "uhm"


$latestReleaseResponse = Invoke-RestMethod -Uri "https://api.github.com/repos/$owner/$repo/releases/latest"
$latestVersion = $latestReleaseResponse.tag_name
$latestVersion = $latestVersion -replace 'v', ''

Write-Output "Downloading uhm v$latestVersion...."

$downloadUrl = "https://github.com/$owner/$repo/releases/download/v$latestVersion/uhm_$latestVersion" + "_windows_amd64.tar.gz"
$downloadPath =  "$HOME\tools\uhm_v" +$latestVersion + "_windows_amd64.tar.gz"

Write-Output "Downloading uhm v$latestVersion from $downloadUrl to $downloadPath"
Invoke-WebRequest -Uri $downloadUrl -OutFile (New-Item -Path $downloadPath -Force)

$unzipPath = "$HOME\tools\uhm_v$latestVersion"+"_windows_amd64"
Write-Output "Unzipping $downloadPath to $unzipPath"

Write-Output $downloadPath
Write-Output $unzipPath

tar -zxvf $downloadPath -C $unzipPath

$env:PATH += ";$unzipPath"

try {
    [System.Environment]::SetEnvironmentVariable("Path", $env:Path, [System.EnvironmentVariableTarget]::User)
    Write-Output "uhm v$latestVersion has been installed successfully"
} catch {
    Write-Output "Failed to add uhm to the PATH. Please add $unzipPath to your PATH manually."
}