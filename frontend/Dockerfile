# Frontend Dockerfile
FROM node:18 as builder
WORKDIR /app
COPY package.json tsconfig.json jest.config.js index.html vite.config.ts ./
COPY src ./src
RUN npm install && npm run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]
