rm -rf ./build &&
mkdir -p ./build/static &&
cp -r ./static/* ./build/static &&
rid=$(uuidgen | md5sum | head -c 32)
bin="main_${rid}"
out="./build/${bin}"
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "$out" ./main.go &&
scp -r ./build/* root@173.212.224.156:/var/www/kilbtech/ &&
ssh root@173.212.224.156 << END_SSH
  cd /var/www/kilbtech/ &&
  mv ${bin} ./tmp_new
  rm  ./main_* &&
  mv tmp_new ${bin}
  ln -sf ./${bin} ./main &&
  systemctl stop kilbtech.service &&
  systemctl start kilbtech.service
END_SSH

