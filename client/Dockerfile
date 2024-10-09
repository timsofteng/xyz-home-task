FROM node:alpine AS builder

WORKDIR /app
RUN npm install -g pnpm

COPY package.json pnpm-lock.yaml ./
RUN pnpm install

# Copy the rest of the application source files
COPY . .
RUN pnpm lint
RUN pnpm test

ARG HOST
ARG SERVER_PORT
ARG API_PORT

# Build the application
RUN VITE_HOST=$HOST \
    VITE_SERVER_PORT=$SERVER_PORT \
    VITE_API_PORT=$API_PORT pnpm build

FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
