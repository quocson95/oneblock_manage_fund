# Use the official Node.js image to build the Angular app
FROM node:18 AS builder

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY fe/ngx-admin/package*.json ./

# Install dependencies
RUN npm install

# Copy the source code
COPY fe/ngx-admin/ .

# Build the Angular app
RUN npm run build --prod

# Use the Nginx image to serve the Angular app
FROM nginx:alpine

# Copy the built Angular app from the previous stage
COPY --from=builder /app/dist/frontend /usr/share/nginx/html

# Expose port 80
EXPOSE 80