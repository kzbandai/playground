Promise.resolve("成功").finally(() => {
    console.log("成功時に実行される");
});
Promise.reject(new Error("失敗")).finally(() => {
    console.log("失敗時に実行される");
});