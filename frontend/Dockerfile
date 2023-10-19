FROM node:16-buster-slim

RUN mkdir /app
WORKDIR /app
COPY . .
RUN rm -rf .vscode .git .git .gitignore .sample-env .env \
     && npm install

RUN echo "#!/bin/bash" > run.sh && chmod +x run.sh
RUN echo "node app.js" >> run.sh

CMD ./run.sh
