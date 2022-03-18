docker build -t ingress:1.0 --force-rm=true ./build/
docker run -d -p 8080:8080 -p 8443:8443  --name ingress ingress:1.0