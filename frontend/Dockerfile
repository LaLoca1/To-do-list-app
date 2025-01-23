# Use a Node image to build the frontend app 
FROM node:16-alpine as builder 

# Set the working directory 
WORKDIR /app 

# Copy package.json and package-lock.json 
COPY package.json package-lock.json ./ 

# Install dependencies 
RUN npm install 

# Copy the source code 
COPY . .

# Build the app
RUN npm run build 

# Use an Nginx image to servce the built app
FROM nginx:alpine 

# Copy the built app from the builder stage 
COPY --from=builder /app/build /usr/share/nginx/html 

# Expose port 80 
EXPOSE 80 

# Start nginx 
CMD ["nginx", "-g", "daemon off;"] 

