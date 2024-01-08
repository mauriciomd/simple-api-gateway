async function json(req, res) {
    setContentTypeHeader(res);
    await createPropertyBodyInReq(req);
}

function setContentTypeHeader(res) {
    res.setHeader('Content-Type', 'application/json')
}

async function createPropertyBodyInReq(req) {
    const buffer = [];
    for await (var chunk of req) {
        buffer.push(chunk);
    }

    if (!buffer.length) return;
    req.body = JSON.parse(Buffer.concat(buffer).toString());
}


module.exports = json;