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
RUN curl --insecure https://download.docker.com/linux/ubuntu/gpg | apt-key add -
RUN apt-get update
RUN add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu focal test"
RUN curl -L "https://github.com/docker/compose/releases/download/1.27.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
RUN chmod +x /usr/local/bin/docker-compose
RUN apt-get -y install vim
RUN apt-get -y install gcc
RUN cd /home/rjar \ 
    && curl -O https://storage.googleapis.com/golang/go1.13.5.linux-amd64.tar.gz \ 
    && tar -xvf go1.13.5.linux-amd64.tar.gz \ 
    && mv go /usr/local
RUN mv /home/rjar/resources/.profile ~/.profile
CMD ["ubuntu:20.10"]