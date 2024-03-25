# Stage 1: Install dependencies
FROM node:18 AS dependencies

# Set the working directory
WORKDIR /opt/app

# Copy package.json and yarn.lock to leverage Docker cache
COPY package.json yarn.lock* ./

# Install dependencies
RUN yarn

# Stage 2: Setup the development environment
FROM node:18 AS development

# Set the working directory
WORKDIR /opt/app

# Copy the project files
COPY . .

# Copy the node_modules from the dependencies stage
COPY --from=dependencies /opt/app/node_modules /opt/app/node_modules

# install angular cli
RUN npm install -g @angular/cli@16.1

# Expose the Angular serve port
EXPOSE 4200

# Run ng serve when the container launches
CMD ["ng", "serve", "--host", "0.0.0.0"]
