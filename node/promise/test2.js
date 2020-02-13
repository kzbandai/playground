function fetchURL(URL) {
    return new Promise((resolve, reject) => {
        const req = new XMLHttpRequest();
        req.open("GET", URL, true);
        req.onload = () => {
            if (200 <= req.status && req.status < 300) {
                resolve(req.responseText);
            } else {
                reject(new Error(req.statusText));
            }
        };
        req.onerror = () => {
            reject(new Error(req.statusText));
        };
        req.send();
    });
}

// 実行例
const URL = "https://httpbin.org/get";
fetchURL(URL).then(function onFulfilled(value) {
    console.log(value);
}).catch(function onRejected(error) {
    console.error(error);
});

// new Promise()の引数はfn
// fnは非同期処理を含んでいる
// fnの第一引数はresolve, 第二はreject
// rejectの中ではnew Error()がよく使われる

// Promise().then()でresolveのときの処理を書く
// then(fn), fnの第一引数はvalue
// onFulfilledは通名

// Promise().then().catch()でrejectのときの処理を書く
// catch(fn), fnの第一引数はerror
// onRejectedは通名