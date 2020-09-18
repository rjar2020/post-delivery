cd /home/rjar/
service apache2 start
/home/rjar/kafka/kafka_2.12-2.6.0/bin/zookeeper-server-start.sh /home/rjar/kafka/kafka_2.12-2.6.0/config/zookeeper.properties > /home/rjar/kafka/kafka_2.12-2.6.0/zookeper.log 2>&1&
/home/rjar/kafka/kafka_2.12-2.6.0/bin/kafka-server-start.sh /home/rjar/kafka/kafka_2.12-2.6.0/config/server.properties > /home/rjar/kafka/kafka_2.12-2.6.0/kafka.log 2>&1&
/usr/local/go/bin/go run github.com/rjar2020/post-delivery
