const { randomUUID } = require('node:crypto');

const database = [
    {
        id: randomUUID(),
        user_id: randomUUID(),
        amount: 119.90
    },
];

const routes = [
    {
        method: 'GET',
        path: '/payments',
        handler: (req, res) => {
            return res.end(JSON.stringify(database));
        },
    },
    {
        method: 'POST',
        path: '/payments',
        handler: (req, res) => {
            const { user_id, amount } = req.body;
            const payment = {
                id: randomUUID(),
                user_id,
                amount,
            };
            database.push(payment);
            return res.writeHead(201).end(JSON.stringify(payment));
        },
    }
];

module.exports = routes;