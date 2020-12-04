const promise = new Promise((resolve) => {
    console.log("inner promise"); // 1
    resolve(42);
});
promise.then((value) => {
    console.log(value); // 3
});

console.log("outer promise"); // 2