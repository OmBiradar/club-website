# Use Alpine Linux as base image
FROM alpine:latest

# Install Nginx to serve static files
RUN apk --no-cache add nginx

# Copy frontend files to the Nginx directory
COPY frontend/ /usr/share/nginx/html/

# Copy a custom Nginx configuration
COPY frontend/nginx.conf /etc/nginx/nginx.conf

# Expose port 8081 for the frontend
EXPOSE 8081

# Start the Nginx server
CMD ["nginx", "-g", "daemon off;"]
