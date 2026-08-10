$dest = "C:\Program Files\sca"
New-Item -ItemType Directory -Path $dest -Force | Out-Null
Copy-Item "sca.exe" "$dest\" -Force
$env:Path += ";$dest"
[Environment]::SetEnvironmentVariable("Path", [Environment]::GetEnvironmentVariable("Path", "Machine") + ";$dest", "Machine")
