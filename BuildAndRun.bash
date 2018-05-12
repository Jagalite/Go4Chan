docker build -t my-golang-app .
docker run -it --rm -p 0.0.0.0:80:1337 --name my-running-app my-golang-app
