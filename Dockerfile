FROM registry.cn-shenzhen.aliyuncs.com/dolphindong/alpine:3.21.1
WORKDIR /app
COPY drone-ci-demo .
RUN chmod +x drone-ci-demo
ENTRYPOINT ["./drone-ci-demo"]