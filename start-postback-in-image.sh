cd /home/rjar/
service apache2 start
echo "Starting Zookeeper..."
/home/rjar/kafka/kafka_2.12-2.6.0/bin/zookeeper-server-start.sh /home/rjar/kafka/kafka_2.12-2.6.0/config/zookeeper.properties > /home/rjar/kafka/kafka_2.12-2.6.0/zookeper.log 2>&1&
echo "Zookeeper started."
echo "Starting Kafka Broker01..."
/home/rjar/kafka/kafka_2.12-2.6.0/bin/kafka-server-start.sh /home/rjar/kafka/kafka_2.12-2.6.0/config/broker1.properties > /home/rjar/kafka/kafka_2.12-2.6.0/kafka.log 2>&1&
echo "Kafka Broker01 started."
echo "Starting Kafka Broker02..."
/home/rjar/kafka/kafka_2.12-2.6.0/bin/kafka-server-start.sh /home/rjar/kafka/kafka_2.12-2.6.0/config/broker2.properties > /home/rjar/kafka/kafka_2.12-2.6.0/kafka.log 2>&1&
echo "Kafka Broker02 started."
echo "Starting Kafka Broker03..."
/home/rjar/kafka/kafka_2.12-2.6.0/bin/kafka-server-start.sh /home/rjar/kafka/kafka_2.12-2.6.0/config/broker3.properties > /home/rjar/kafka/kafka_2.12-2.6.0/kafka.log 2>&1&
echo "Kafka Broker03 started."
/usr/local/go/bin/go run github.com/rjar2020/post-delivery
