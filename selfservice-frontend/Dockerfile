FROM node:6.11.0

# The base node image sets a very verbose log level.
ENV NPM_CONFIG_LOGLEVEL warn
COPY . .
RUN npm install
RUN npm run build --production

# Install `serve` to run the application.
RUN npm install -g serve

# Set the command to start the node server.
CMD serve -s build

# Tell Docker about the port we'll run on.
EXPOSE 5000
