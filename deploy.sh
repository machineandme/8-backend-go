mkdir dist/

env GOOS=linux GOARCH=386 go build -o dist/backend || exit 1

cp *.html dist/ || exit 1

scp -r ./dist root@188.166.161.22:~ || exit 1

ssh root@188.166.161.22 'pkill -9 backend' || exit 1
ssh root@188.166.161.22 'mv dist/* .' || exit 1
ssh root@188.166.161.22 'bash -c "env HOST=0.0.0.0:80 ./backend & disown"'

rm -fr dist/