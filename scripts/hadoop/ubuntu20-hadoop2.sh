#!/bin/bash

apt-get update -y 
apt-get install openjdk-8-jdk -y
java -version
apt-get install software-properties-common -y 
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64/jre/
mkdir /opt/hadoop
wget https://dlcdn.apache.org/hadoop/common/hadoop-2.10.1/hadoop-2.10.1.tar.gz
tar xvf hadoop-2.10.1.tar.gz -C /opt/hadoop
rm hadoop-2.10.1.tar.gz
export HADOOP_PREFIX="/opt/hadoop/hadoop-2.10.1"
export PATH="$PATH:$HADOOP_PREFIX/bin:$HADOOP_PREFIX/sbin"
export HADOOP_MAPRED_HOME=${HADOOP_PREFIX}
export HADOOP_COMMON_HOME=${HADOOP_PREFIX}
export HADOOP_HDFS_HOME=${HADOOP_PREFIX}
export YARN_HOME=${HADOOP_PREFIX}
tee -a ~/.bashrc > /dev/null <<EOT
export HADOOP_PREFIX="/opt/hadoop/hadoop-2.10.1"
export PATH="$PATH:\$HADOOP_PREFIX/bin:\$HADOOP_PREFIX/sbin"
export HADOOP_MAPRED_HOME=\${HADOOP_PREFIX}
export HADOOP_COMMON_HOME=\${HADOOP_PREFIX}
export HADOOP_HDFS_HOME=\${HADOOP_PREFIX}
export YARN_HOME=\${HADOOP_PREFIX}
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64/jre/
EOT
tee -a /home/ubuntu/.bashrc > /dev/null <<EOT
export HADOOP_PREFIX="/opt/hadoop/hadoop-2.10.1"
export PATH="$PATH:\$HADOOP_PREFIX/bin:\$HADOOP_PREFIX/sbin"
export HADOOP_MAPRED_HOME=\${HADOOP_PREFIX}
export HADOOP_COMMON_HOME=\${HADOOP_PREFIX}
export HADOOP_HDFS_HOME=\${HADOOP_PREFIX}
export YARN_HOME=\${HADOOP_PREFIX}
export JAVA_HOME=/usr/lib/jvm/java-8-openjdk-amd64/jre/
EOT
chown ubuntu:ubuntu -R /opt/hadoop
mkdir /hadoop 
chown -R ubuntu:ubuntu /hadoop 
wget http://archive.apache.org/dist/ambari/ambari-2.7.6/apache-ambari-2.7.6-src.tar.gz