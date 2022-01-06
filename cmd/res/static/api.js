const pagesURI = '/pages'
const pageURI = '/page?path=example.md'

async function getPage() {
    return await get(pageURI)
}


function get(uri, token) {
    return new Promise(function (resolve, reject) {
        let xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function () {
            if (xhr.readyState !== 4) {
                return;
            }
            if (xhr.status >= 200 && xhr.status < 300) {
                if (xhr.getResponseHeader('content-type').indexOf('application/json') !== -1) {
                    let resp = JSON.parse(xhr.responseText);
                    resolve(resp);
                } else {
                    resolve(xhr.responseText);
                }
            } else {
                if (xhr.getResponseHeader('content-type').indexOf('application/json') !== -1) {
                    let resp = JSON.parse(xhr.responseText);
                    reject(resp)
                } else {
                    reject(xhr.responseText)
                }
            }
        };

        xhr.open('GET', uri, true);
        xhr.setRequestHeader('content-type', 'application/json');
        xhr.setRequestHeader('Authorization', 'Bearer ' + token);
        xhr.send();
    });
}

function post(uri, token, data) {
    return new Promise(function (resolve, reject) {
        let xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function () {
            if (xhr.readyState !== 4) {
                return;
            }
            if (xhr.status >= 200 && xhr.status < 300) {
                if (xhr.getResponseHeader('content-type').indexOf('application/json') !== -1) {
                    let resp = JSON.parse(xhr.responseText);
                    resolve(resp);
                } else {
                    resolve(xhr.responseText);
                }
            } else {
                if (xhr.getResponseHeader('content-type').indexOf('application/json') !== -1) {
                    let resp = JSON.parse(xhr.responseText);
                    reject(resp)
                } else {
                    reject(xhr.responseText)
                }
            }
        };

        xhr.open('POST', uri, true);
        xhr.setRequestHeader('content-type', 'application/json');
        xhr.setRequestHeader('Authorization', 'Bearer ' + token);
        xhr.send(JSON.stringify(data));
    });
}