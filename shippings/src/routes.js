const { randomUUID } = require('node:crypto');

const database = [
    {
        id: randomUUID(),
        payment_id: randomUUID(),
        status: 'shipped',
    },
];

const routes = [
    {
        method: 'GET',
        path: '/shippings',
        handler: (req, res) => {
            return res.end(JSON.stringify(database));
        },
    },
    {
        method: 'POST',
        path: '/shippings',
        handler: (req, res) => {
            const { payment_id } = req.body;
            const shipping = {
                id: randomUUID(),
                payment_id,
                staus: 'pending',
            };
            database.push(shipping);
            return res.writeHead(201).end(JSON.stringify(shipping));
        },
    }
];

module.exports = routes;