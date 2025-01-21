#!/bin/bash

# List of packages to downgrade with their previous versions
packages=(
  "tzdata-2024b-2"
  "libxcrypt-4.4.37-1"
  "sqlite-3.47.2-1"
  "util-linux-libs-2.40.2-3"
  "llvm-libs-18.1.8-5"
  "spirv-tools-2024.4.rc1-1"
  "systemd-libs-257.2-1"
  "audit-4.0.2-3"
  "mesa-1:24.3.3-1"
  "pugixml-1.14-1"
  "lib32-llvm-libs-1:18.1.8-1"
  "lib32-spirv-tools-2024.4.rc1-1"
  "lib32-libxcrypt-4.4.37-1"
  "lib32-mesa-1:24.3.3-1"
  "device-mapper-2.03.29-1"
  "cryptsetup-2.7.5-1"
  "util-linux-2.40.2-3"
  "systemd-257.2-1"
)

# Downgrade each package
for package in "${packages[@]}"; do
  package_path="/var/cache/pacman/pkg/${package}-*.pkg.tar.zst"
  if ls $package_path 1> /dev/null 2>&1; then
    echo "Downgrading $package..."
    sudo pacman -U --noconfirm $package_path
  else
    echo "Package $package not found in cache. Skipping..."
  fi
done

echo "Downgrade process complete."
