const fetch = require("node-fetch");

function fetchBookTitle() {
    // Fetch APIは指定URLのリソースを取得しPromiseを返す関数
    return fetch("https://azu.github.io/promises-book/json/book.json").then((res) => {
        return res.json(); // レスポンスをJSON形式としてパースする
    }).then((json) => {
        return json.title; // JSONからtitleプロパティを取り出す
    });
}

function main() {
    // `fetchBookTitle`関数は、取得したJSONの`title`プロパティでresolveされる
    fetchBookTitle().then((title) => {
        console.log(title); // => "JavaScript Promiseの本"
    });
}

main();
