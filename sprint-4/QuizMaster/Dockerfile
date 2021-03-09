#FROM baseImage:tag
FROM golang:latest

#make a directory in the image called sprint-4
RUN mkdir /360-RedDragons/sprint-4

#Copy all files from sprint-4 folder into the image
ADD . /360-RedDragons/sprint-4

#Set the working directory to build executable
WORKDIR /360-RedDragons/sprint-4/QuizMaster/httpserver

#Build executable
RUN go build -o main .

#Expose port 9000
EXPOSE 9000

#Run httpServer executable
CMD ["360-RedDragons/sprint-4/QuizMaster/httpserver/main"]