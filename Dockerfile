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
    software-properties-common \
    vim \
    gcc \
    default-jre
RUN cd /home/rjar \ 
    && curl -O https://storage.googleapis.com/golang/go1.13.5.linux-amd64.tar.gz \ 
    && tar -xvf go1.13.5.linux-amd64.tar.gz \ 
    && mv go /usr/local
RUN mv /home/rjar/resources/.profile ~/.profile
RUN mkdir /home/rjar/kafka \
    && cd /home/rjar/kafka \
    && curl "https://downloads.apache.org/kafka/2.6.0/kafka_2.12-2.6.0.tgz" -o ./kafka.tgz \
    && tar -xvzf kafka.tgz
ADD ./resources/broker1.properties /home/rjar/kafka/kafka_2.12-2.6.0/config/broker1.properties
ADD ./resources/broker2.properties /home/rjar/kafka/kafka_2.12-2.6.0/config/broker2.properties
ADD ./resources/broker3.properties /home/rjar/kafka/kafka_2.12-2.6.0/config/broker3.properties
RUN export GOPATH=$HOME/work
RUN export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
RUN export DEBIAN_FRONTEND=noninteractive
RUN export DEBCONF_NONINTERACTIVE_SEEN=true
## preesed tzdata
RUN echo "tzdata tzdata/Areas select Europe" > /tmp/preseed.txt; \
    echo "tzdata tzdata/Zones/Europe select Madrid" >> /tmp/preseed.txt; \
    debconf-set-selections /tmp/preseed.txt && \
    apt-get update && \
    apt-get install -y tzdata
RUN apt -y -qq install apache2 php libapache2-mod-php php-curl
RUN mv -f /home/rjar/web/ports.conf /etc/apache2/ports.conf
RUN mv -f /home/rjar/web/apache2.conf /etc/apache2/apache2.conf
RUN mv -f /home/rjar/web/php/* /var/www/html
RUN mv /var/www/html/index.html /var/www/html/index.html.bck
EXPOSE 4000
EXPOSE 8080
CMD ./home/rjar/start-postback-in-image.sh
