#使用了镜像大小体积只有5MB的alpine镜像
FROM alpine:latest
#作者信息
MAINTAINER silver fox
#设置工作路径
RUN mkdir -p /data/service
WORKDIR /data/service
#把上文编译好的api文件添加到镜像里
ADD api /data/service
#暴露容器内部端口
EXPOSE 8080
#入口
ENTRYPOINT ["./api"]
