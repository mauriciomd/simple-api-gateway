const http = require('node:http')
const routes = require('./routes')
const jsonMiddleware = require('./json_middleware')

const api = http.createServer(async (req, res) => {
    
    const { method, url } = req;
    await jsonMiddleware(req, res);

    const route = routes.find(r => r.method === method && r.path === url);
    if (!route) {
        return res.writeHead(404);
    }
    
    return route.handler(req, res);
});

api.listen(3000);