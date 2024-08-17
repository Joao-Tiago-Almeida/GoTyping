FROM golang:1.22.6-bullseye

# Set the working directory
WORKDIR /app

COPY . /app/

ARG IMAGE_NAME
RUN echo "IMAGE_NAME=${IMAGE_NAME:-'local'}"

RUN str1='export PS1="\e[1;34m' && str2='@\w> \e[m"' && echo "$str1${IMAGE_NAME##*/}$str2" >> ~/.bashrc

RUN  go build -o exec ./src/main.go