sudo docker build -t acrmmdebug.azurecr.io/workshop/aryan-go-web .
sudo docker push acrmmdebug.azurecr.io/workshop/aryan-go-web
# sudo docker run --name go-webserver --rm -p 8080:8080 -d acrmmdebug.azurecr.io/workshop/aryan-go-web:latest
