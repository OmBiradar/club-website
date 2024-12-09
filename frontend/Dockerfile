# Dockerfile for frontend (serving HTML, CSS, JS)
FROM alpine:latest

# Install Nginx to serve static files
RUN apk --no-cache add nginx

# Copy frontend files to the Nginx directory
COPY ./ /usr/share/nginx/html/

# Expose the port that Nginx will use
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
