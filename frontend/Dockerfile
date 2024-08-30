# Use the official Node.js image to build the Angular app
FROM node:20.17.0 AS builder

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install --force

# Copy the source code
COPY . .

# Build the Angular app
RUN npm run build --prod

# Use the Nginx image to serve the Angular app
FROM nginx:alpine

COPY nginx.conf /etc/nginx/nginx.conf
# Copy the built Angular app from the previous stage
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port 80
EXPOSE 80
