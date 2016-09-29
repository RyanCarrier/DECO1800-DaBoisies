#!/bin/bash
set -e
OSS=("darwin" "windows" "solaris" "linux")
GOARCH="amd64"

for OS in ${OSS[@]};do
		OUT="$OS""Server"
		if [ "$OS" == "windows" ];then
			OUT="$OUT"".exe"
		fi
		echo -ne "Building for $OS...\t"
		env GOOS=$OS go build -o "$OUT"
		echo "OK!"
done
echo -e "Done!\nAll ARCH=amd64"
echo "Zipping..."
env GZIP=-9 tar cvzf servers.tar.gz *Server*
