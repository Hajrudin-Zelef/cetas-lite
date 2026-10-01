---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/mat-ice-qc-mat-ice-hub-blob-head-docs-scripting-powershell-unifi-manual-unifi-ba-046e9ca9
title: "SPDX-License-Identifier: AGPL-3.0-or-later"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "copyright", "licenses"]
source: docs/RAG/collect-261001-unifi-ubiquiti/mat-ice-qc-mat-ice-hub-blob-head-docs-scripting-powershell-unifi-manual-unifi-ba-046e9ca9.md
source_anchor: ""
source_lines: [1, 172]
sha256: 9fe0b360cefe4740a2105ef4f4a6b37ce4b112d07a066dc548d5de19c9865f56
---

# SPDX-License-Identifier: AGPL-3.0-or-later

| title | Unifi Backup Extraction Script | 
|---|---|
| description | Explanation and source code for the PowerShell script that extracts Unifi backups over SSH with Putty. | 
| license | AGPL-3.0-or-later | 
| author | Mat Ice | 
This document explains and includes the PowerShell script used to extract Unifi backups over SSH with Putty.
Basically, if you want a security system with no WAN access, this makes it possible to back up Unifi data to a LAN server.
The script connects to a Unifi device via SSH, retrieves the list of backup files, compares them with the local directory, and automatically downloads any that are missing.
It uses PuTTY utilities (plink and pscp) for SSH communication and file transfers.
| Parameter | Description | 
|---|---|
| unifiUserRoot | SSH username for the Unifi device (default: root ). | 
| unifiDevice | IP address of the Unifi device. | 
| unifiBackupDir | Remote directory path containing Unifi backups. | 
| password | SSH password for the Unifi user. | 
| localDir | Local directory where backups will be stored. | 
The script ensures the local directory exists before proceeding.
If not, it exits to avoid errors.
if (-not (Test-Path $localDir)) {
    Write-Host "Local directory does not exist: $localDir"
    exit 1
}
To verify the Unifi device identity, the script retrieves its SSH fingerprint and ensures the connection is secure.
$hostKeyScan = & ssh-keyscan -t ed25519 $unifiDevice 2>$null
$hostKeyFingerprint = $hostKeyScan | ssh-keygen -lf - | ForEach-Object {
    $_.ToString().Split(" ") | Select-Object -Index 1
}
This prevents man-in-the-middle attacks by confirming that the SSH key matches the expected device.
The script connects using plink and lists all files in the Unifi backup directory.
$unifiBackups = & "$plinkPath" `
    -ssh `
    -pw $password `
    -batch `
    -noagent `
    -hostkey "$fullHostKey" `
    "$unifiUserRoot@$unifiDevice" `
    "ls -1 $unifiBackupDir"
It compares the list of remote backups with local files to find any missing ones.
$localFiles = Get-ChildItem -Path $localDir -File | Select-Object -ExpandProperty Name
$missingBackups = $unifiBackups | Where-Object { $_ -and ($_ -notin $localFiles) }
If there are no missing files, the script exits gracefully.
For each missing file, the script downloads it using pscp over SSH.
foreach ($file in $missingBackups) {
    $remoteFullPath = "$unifiUserRoot@$unifiDevice:$unifiBackupDir/$file"
    & "$pscpPath" -pw $password -batch -noagent -hostkey $fullHostKey $remoteFullPath $localDir
}
Once complete, it reports that all missing files have been downloaded.
To run the script:
.\BackupUnifi.ps1
Before running:
- Ensure plink.exe andpscp.exe (from PuTTY) are installed.
- Verify the Unifi IP address and SSH credentials.
- Make sure the local backup directory exists and has write permissions.
# SPDX-License-Identifier: AGPL-3.0-or-later
#
# This file is part of the "Unifi Automatic Backup Script via SSH".
#
# Copyright (C) 2025  Mat Ice
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as published by
# the Free Software Foundation, either version 3 of the License, or (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.
#
<#
.SYNOPSIS
Automatically backs up Unifi configuration files over SSH.
.DESCRIPTION
It connects to a Unifi device via SSH, retrieves the list of backup files,
compares them with the local ones, and automatically downloads any that are missing.
.PARAMETER unifiUserRoot
SSH username for connecting to the Unifi device (default: "root").
.PARAMETER unifiDevice
IP address of the Unifi device.
.PARAMETER unifiBackupDir
Remote directory containing Unifi backups.
.PARAMETER password
SSH password for the Unifi user.
.PARAMETER localDir
Local directory where backups will be stored.
.EXAMPLE
.\BackupUnifi.ps1
.NOTES
Author  : Mat Ice
#>
$unifiUserRoot = "root"
$unifiDevice = "192.168.2.1"
$unifiBackupDir = "/data/unifi/data/backup/autobackup"
$password = "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
$localDir = "G:\unifi"
# Path to PuTTY tools (adjust if needed)
$pscpPath = "C:\Program Files\PuTTY\pscp.exe"
$plinkPath = "C:\Program Files\PuTTY\plink.exe"
# === STEP 0: Verify local directory ===
if (-not (Test-Path $localDir)) {
    Write-Host "Local directory does not exist: $localDir"
    exit 1
}
# === STEP 1: Retrieve SSH host key fingerprint ===
Write-Host "Retrieving SSH host key fingerprint..."
$hostKeyScan = & ssh-keyscan -t ed25519 $unifiDevice 2>$null
if (-not $hostKeyScan) {
    Write-Host "Failed to retrieve SSH key from $unifiDevice"
    exit 1
}
# Generate SHA256 fingerprint (requires OpenSSH)
$hostKeyFingerprint = $hostKeyScan | ssh-keygen -lf - | ForEach-Object {
    $_.ToString().Split(" ") | Select-Object -Index 1
}
if (-not $hostKeyFingerprint) {
    Write-Host "Failed to parse fingerprint."
    exit 1
}
$fullHostKey = "ssh-ed25519 255 $hostKeyFingerprint"
Write-Host "Using host key fingerprint: $fullHostKey"
# === STEP 2: Retrieve remote backup list ===
Write-Host "Fetching file list from remote Unifi device..."
try {
    $unifiBackups = & "$plinkPath" `
        -ssh `
        -pw $password `
        -batch `
        -noagent `
        -hostkey "$fullHostKey" `
        "$unifiUserRoot@$unifiDevice" `
        "ls -1 $unifiBackupDir"
} catch {
    Write-Host "Failed to retrieve remote file list: $_"
    exit 1
}
if (-not $unifiBackups) {
    Write-Host "No backup files found or SSH command failed."
    exit 1
}
# === STEP 3: Retrieve local files ===
$localFiles = Get-ChildItem -Path $localDir -File | Select-Object -ExpandProperty Name
# === STEP 4: Compare and download missing backups ===
$missingBackups = $unifiBackups | Where-Object { $_ -and ($_ -notin $localFiles) }
if ($missingBackups.Count -eq 0) {
    Write-Host "All backups are already present locally. Nothing to download."
    exit 0
}
foreach ($file in $missingBackups) {
    $remoteFilePath = "$unifiBackupDir/$file"
    $localFilePath = Join-Path $localDir $file
    Write-Host "Downloading: $file"
    try {
        $remoteFullPath = "$unifiUserRoot@$unifiDevice" + ":$remoteFilePath"
        $arguments = @(
            '-pw', $password,
            '-batch',
            '-noagent',
            '-hostkey', $fullHostKey,
            $remoteFullPath,
            $localFilePath
        )
        & "$pscpPath" @arguments
    }
    catch {
        Write-Host "Failed to download $file : $_"
    }
}
Write-Host "All missing backup files have been downloaded successfully."
This PowerShell script automates Unifi configuration backups through SSH.
It ensures your Unifi device’s backups are securely mirrored on a lan server, minimizing the risk of configuration loss.
