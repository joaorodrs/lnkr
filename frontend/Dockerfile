# Etapa 1: Build da aplicação com Node.js
FROM node:18-alpine AS builder

WORKDIR /app

# Copia package.json e package-lock.json / yarn.lock para cache de dependências
COPY package*.json ./
# ou COPY yarn.lock ./

RUN npm install
# ou RUN yarn install

# Copia todo o código fonte
COPY . .

# Build da aplicação Vite
RUN npm run build
# ou RUN yarn build

# Etapa 2: Servir os arquivos estáticos com Nginx
FROM nginx:stable-alpine

# Copia build estático gerado para a pasta que o nginx serve
COPY --from=builder /app/dist /usr/share/nginx/html

# Copia configuração customizada do nginx (opcional)
# COPY nginx.conf /etc/nginx/nginx.conf

# Expor porta padrão do nginx
EXPOSE 80

# Comando padrão
CMD ["nginx", "-g", "daemon off;"]
