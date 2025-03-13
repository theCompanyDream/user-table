# Stage 1: Install dependencies
FROM node:22-alpine AS dependencies

# Set the working directory
WORKDIR /opt/app

RUN npm install -g pnpm

# Copy package.json and yarn.lock to leverage Docker cache
COPY package.json pnpm-lock.yaml* ./

# Install dependencies
RUN pnpm install

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
