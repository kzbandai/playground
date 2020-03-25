const fetch = require("node-fetch");

// `async`をつけて`fetchBookTitle`関数をAsync Functionとして定義
async function fetchBookTitle() {
    // リクエストしてリソースを取得する
    const res = await fetch("https://azu.github.io/promises-book/json/book.json");
    // レスポンスをJSON形式としてパースする
    const json = await res.json();
    // JSONからtitleプロパティを取り出す
    return json.title;
}

// `async`をつけて`main`関数をAsync Functionとして定義
async function main() {
    // `await`式で`fetchBookTitle`の非同期処理が完了するまで待つ
    // `fetchBookTitle`がresolveした値が返り値になる
    const title = await fetchBookTitle();
    console.log(title); // => "JavaScript Promiseの本"
}

main();
