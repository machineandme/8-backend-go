mkdir dist/;

echo "*" > dist/.gitignore

env GOOS=linux GOARCH=386 go build -o dist/backend || exit 1

cp *.html dist/ || exit 1