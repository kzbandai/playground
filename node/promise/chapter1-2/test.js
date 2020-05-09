function asyncFunction() {
    return new Promise((resolve) => {
        setTimeout(() => {
            resolve("Async Hello world");
        }, 16);
    });
}

asyncFunction().then((value) => {
    console.log(value); // => 'Async Hello world'
}).catch((error) => {
    console.error(error);
});
