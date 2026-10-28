[CmdletBinding()]
param(
    [switch]$Check,
    [Parameter(ValueFromRemainingArguments = $true)]
    [string[]]$Files
)

$ErrorActionPreference = 'Stop'
$moduleRoot = [IO.Path]::GetFullPath($PSScriptRoot)
$isWindowsHost = [Environment]::OSVersion.Platform -eq [PlatformID]::Win32NT
$comparison = if ($isWindowsHost) { [StringComparison]::OrdinalIgnoreCase } else { [StringComparison]::Ordinal }

if (-not $Files -or $Files.Count -eq 0) {
    $selected = @(Get-ChildItem -LiteralPath $moduleRoot -File -Filter '*.go' |
            Where-Object { $_.Extension.Equals('.go', [StringComparison]::Ordinal) } |
            Sort-Object Name | ForEach-Object FullName)
}
else {
    $selected = @(
        foreach ($file in $Files) {
            if ([IO.Path]::IsPathRooted($file)) { throw "Use a module-relative Go path: $file" }
            $path = [IO.Path]::GetFullPath((Join-Path $moduleRoot $file))
            if (-not [IO.Path]::GetDirectoryName($path).Equals($moduleRoot, $comparison) -or
                -not [IO.Path]::GetExtension($path).Equals('.go', [StringComparison]::Ordinal) -or
                -not (Test-Path -LiteralPath $path -PathType Leaf)) {
                throw "Path is not a Go file at this module root: $file"
            }
            $path
        }
    )
}
if ($selected.Count -eq 0) { throw 'No Go files selected.' }
foreach ($path in $selected) {
    if ((Get-Item -LiteralPath $path).LinkType) {
        throw "Refusing to format a linked Go file: $path"
    }
}

Push-Location -LiteralPath $moduleRoot
try {
    $version = & go version
    if ($LASTEXITCODE -ne 0 -or $version -notmatch '^go version go1\.27\.1 ') {
        throw "Go 1.27.1 is required for formatting; found: $version"
    }
    $goRoot = & go env GOROOT
    if ($LASTEXITCODE -ne 0) { throw 'Could not resolve GOROOT.' }
    $formatter = Join-Path $goRoot 'bin/gofmt'
    if ($isWindowsHost) { $formatter += '.exe' }
    if (-not (Test-Path -LiteralPath $formatter -PathType Leaf)) { throw "gofmt not found: $formatter" }

    if ($Check) {
        $unformatted = @(& $formatter -l @selected)
        if ($LASTEXITCODE -ne 0) { throw 'gofmt check failed.' }
        if ($unformatted.Count -gt 0) {
            $unformatted | Write-Output
            throw 'Go formatting check failed.'
        }
        Write-Output "Formatting OK: $($selected.Count) Go file(s)."
    }
    else {
        & $formatter -w @selected
        if ($LASTEXITCODE -ne 0) { throw 'gofmt failed.' }
        Write-Output "Formatted $($selected.Count) Go file(s)."
    }
}
finally {
    Pop-Location
}
