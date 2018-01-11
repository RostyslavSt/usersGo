
#From golang:latest
#ADD . /go/src/github.com/RostyslavSt/usersGo

# Build the contact_registry command inside the container.

#RUN go install github.com/RostyslavSt/usersGo

# Run the contact_registry command when the container starts.

#ENTRYPOINT /go/bin/usersGo

# http server listens on port 8080.

#EXPOSE 8080

FROM golang:onbuild
EXPOSE 8080