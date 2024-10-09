# Stage 2: Serve the built React app with a simple static server
FROM nginx:alpine

COPY static/* /usr/share/nginx/html

# Expose port 80 to serve the app
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
