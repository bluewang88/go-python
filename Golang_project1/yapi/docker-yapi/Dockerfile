FROM node:12-alpine
COPY repositories /etc/apk/repositories

# RUN npm install -g yapi-cli --registry https://registry.npmmirror.com

# 修改为官方 npm 源
RUN npm config set registry https://registry.npmjs.org/
RUN npm install -g fs-extra
RUN npm install -g yapi-cli

EXPOSE 3000 9090




