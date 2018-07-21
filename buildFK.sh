export GOPATH="D:\\Work\\GIT\\Trojan"
env | grep GOPA
cd src/FKTrojan

go build Server.go
go build Client.go

mv -f *.exe /d/Work/GIT/Trojan/bin
read -p "Press any key to continue." var