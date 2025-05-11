# https://hub.docker.com/_/node
FROM node:23-alpine AS builder

WORKDIR /app

# Copy package files first to leverage Docker cache for dependencies
COPY package*.json ./

RUN npm ci

COPY . .

RUN npm run build

# Create the final image with only runtime dependencies
FROM node:23-alpine AS runner

WORKDIR /app

# Copy the necessary files from the build stage
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

EXPOSE 3000

CMD ["npm", "start"]