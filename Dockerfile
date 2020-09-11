FROM ubuntu:20.10
RUN mkdir /home/rjar
COPY . /home/rjar
#This is not in the git repo, a certificate will be needed in the local machine location
ADD ./resources/my_Cert.crt usr/local/share/ca-certificates/my.crt
#This is not in the git repo, a certificate will be needed in the local machine location
ADD ./resources/rootcaCert.crt usr/local/share/ca-certificates/root.crt
RUN apt-get update
RUN apt-get -y install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common
RUN apt-get -y install vim
RUN apt-get -y install gcc
RUN cd /home/rjar \ 
    && curl -O https://storage.googleapis.com/golang/go1.13.5.linux-amd64.tar.gz \ 
    && tar -xvf go1.13.5.linux-amd64.tar.gz \ 
    && mv go /usr/local
RUN mv /home/rjar/resources/.profile ~/.profile
RUN apt-get -y install default-jre
RUN mkdir /home/rjar/kafka \
    && cd /home/rjar/kafka \
    && curl "https://downloads.apache.org/kafka/2.6.0/kafka_2.12-2.6.0.tgz" -o ./kafka.tgz \
    && tar -xvzf kafka.tgz
RUN export GOPATH=$HOME/work
RUN export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
EXPOSE 4000
CMD ./home/rjar/start-postback-in-image.sh
