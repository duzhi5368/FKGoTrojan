#!/bin/bash
org_pwd=$(cd `dirname $0`; pwd)
cd $org_pwd

# convert /d/aa/bb/cc -> d:\aa\bb\cc
# /d/aa/bb/cc
pwd=${org_pwd}
# d/aa/bb/cc
pwd=${pwd:1}
# d:\aa/bb/cc
pwd=${pwd/\//:\\}
# d:\aa\bb\cc
pwd=${pwd//\//\\}
# echo $pwd
if [ "$GOPATH" = "" ] ; then
	export GOPATH="${pwd}"
else
	GOPATH=$GOPATH";${pwd}"
fi
#env | grep GOPA
cd src/FKTrojan
echo 'go build Server.go ...'
go build Server.go
echo 'go build Client.go ...'
go build Client.go

mv -f *.exe ${pwd}/bin
cd client_tools/mysql_tools
echo 'go build mysql_tools ...'
go build 
mv -f *.exe ${pwd}/bin
cd -
cd client_tools/scan_dir
go build
mv -f *.exe ${pwd}/bin
cd -
cd server_tools/command_tools
go build
mv -f *.exe ${pwd}/bin
echo "binary file in ${pwd}\\bin"
cp -f ${org_pwd}/bin/*.exe /d/bin/
