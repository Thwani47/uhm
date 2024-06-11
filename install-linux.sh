#!/bin/bash

# Define the GitHub repository
owner="Thwani47"
repo="uhm"

# Call the GitHub API to get the latest release
latest_version=$(curl --silent "https://api.github.com/repos/$owner/$repo/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/' | cut -c 2-)

# Define the URL of the release
url="https://github.com/$owner/$repo/releases/download/v$latest_version/uhm_$latest_version""_linux_amd64.tar.gz"

# Define the path where the release will be downloaded
download_path="$HOME/tools/uhm_v$latest_version""_linux_amd64.tar.gz"

# Download the release
curl -L $url -o $download_path

# Unzip the release
unzip_path="$HOME/tools/uhm_v$latest_version""_linux_amd64"
mkdir -p $unzip_path
tar -xzf $download_path -C $unzip_path

# Add the unzipped directory to your PATH
echo 'export PATH=$PATH:'$unzip_path >> $HOME/.bashrc

# Source the .bashrc to update the PATH
source $HOME/.bashrc

echo "uhm v$latest_version has been installed successfully"