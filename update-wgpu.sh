#!/usr/bin/env bash

set -ex -o pipefail

# gets the prebuilt release files from wgpu-native and installs
# to wgpu/lib

BASE="$PWD"

rm -rf wgpu/lib

rm -rf release
mkdir release

pushd release
  rm -rf wgpu-native zips
  mkdir zips

  git clone https://github.com/gfx-rs/wgpu-native wgpu-native
  pushd wgpu-native
    gh release download --dir ../zips -p "*-release.zip"
  popd

  pushd zips

    for ZIP in *.zip ; do
      if [[ "$ZIP" = *-x86_64-msvc-* || "$ZIP" = *-aarch64-simulator-* ]] ; then
        continue
      fi

      ARCH=$(cut -d- -f2-3 <<< "$ZIP")
      mkdir $ARCH
      pushd $ARCH
        unzip ../"$ZIP"
      popd
    done
  popd
popd

function copy-to-target() {
    local ZIP="$1"
    local TARGET="$2"
    local ARCH="$3"
    local LIB="$4"

    local DEST="wgpu/lib/$TARGET/$ARCH"
    mkdir -p $DEST

    # TODO do we really need this?
    echo "package vendor" > $DEST/vendor.go
    echo "package vendor" > wgpu/lib/vendor.go

    cp release/zips/$ZIP/lib/$LIB $DEST
    cp release/zips/$ZIP/include/webgpu/*.h $DEST
}

copy-to-target "macos-aarch64"    "darwin"    "arm64"   libwgpu_native.a
copy-to-target "macos-x86_64"     "darwin"    "amd64"   libwgpu_native.a
copy-to-target "ios-aarch64"      "ios"       "arm64"   libwgpu_native.a
copy-to-target "ios-x86_64"       "ios"       "amd64"   libwgpu_native.a
copy-to-target "windows-aarch64"  "windows"   "arm64"   wgpu_native.lib
copy-to-target "windows-x86_64"   "windows"   "amd64"   libwgpu_native.a
copy-to-target "windows-i686"     "windows"   "386"     wgpu_native.lib
copy-to-target "linux-aarch64"    "linux"     "arm64"   libwgpu_native.a
copy-to-target "linux-x86_64"     "linux"     "amd64"   libwgpu_native.a
copy-to-target "android-aarch64"  "android"   "arm64"   libwgpu_native.a
copy-to-target "android-x86_64"   "android"   "amd64"   libwgpu_native.a
copy-to-target "android-armv7"    "android"   "arm"     libwgpu_native.a
copy-to-target "android-i686"     "android"   "386"     libwgpu_native.a

rm -rf release

cat > wgpu/lib/.gitattributes <<EOF
# See
# https://github.com/github/linguist/blob/249bbd1c2ffc631ca2ec628da26be5800eec3d48/docs/overrides.md#vendored-code

webgpu.h linguist-vendored
wgpu.h linguist-vendored
EOF
