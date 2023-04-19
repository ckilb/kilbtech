rm -rf ./build &&
mkdir -p ./build/static &&
cp -r ./static/* ./build/static &&
cd ./src &&
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../build/main ./main.go &&
cd .. &&
ssh root@kilb.tech << END_SSH
  systemctl stop kilbtech.service
END_SSH
scp -r ./build/* root@kilb.tech:/var/www/kilbtech/ &&
ssh root@kilb.tech << END_SSH
  systemctl start kilbtech.service
END_SSH

