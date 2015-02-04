sudo docker run -v $GOPATH/src:/src mobile /bin/bash -c 'cd /src/github.com/tbruyelle/platformer && ./make.bash'
if [[ $? -eq 0  ]] 
then
	adb uninstall com.kamosoft.platformer
	adb install bin/nativeactivity-debug.apk
fi
