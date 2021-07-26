#!/bin/bash
echo "Building Go Project"

rm main
go build main.go banners.go sqlinjection.go commandinjection.go urlredirection.go socket.go fileaccess.go miscoptions.go userdata.go

echo "Finished building!"