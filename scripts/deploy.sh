rm -rf ./build &&
mkdir -p ./build/static &&
cp -r ./static/* ./build/static &&
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ../build/main ./main.go &&
cd .. &&
ssh root@173.212.224.156 << END_SSH
  systemctl stop kilbtech.service
END_SSH
scp -r ./build/* root@173.212.224.156:/var/www/kilbtech/ &&
ssh root@173.212.224.156 << END_SSH
  systemctl start kilbtech.service
END_SSH

