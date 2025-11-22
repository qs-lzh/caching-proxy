# caching-proxy

**caching-proxy** is a command-line tool that implements an HTTP proxy server with Redis-based caching. It forwards incoming requests to an origin server while caching responses to improve performance and reduce unnecessary upstream traffic.

## 📦 Usage

Start the caching proxy server:

```
caching-proxy --port 3000 --origin http://dummyjson.com
```

After starting the server, you can send requests such as:

```
http://localhost:3000/products
```

The proxy server will:

1. Map the path `/products` to

   ```
   http://dummyjson.com/products
   ```

2. Check Redis for a cached response

3. Return the cached response if available

4. Otherwise fetch from the origin server and store the new response in Redis

## 📚 About This Project

This project was built to complete the **“Caching Server”** assignment on Roadmap.sh:

🔗 https://roadmap.sh/projects/caching-server

Developed by **qs-lzh**

## 👍 Support the Project

If you find this project useful, feel free to support it with an upvote on the Roadmap project page:

👉 *(Insert your Roadmap project link here)*
