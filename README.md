### Goals

Learn Go while having fun.

### Experience with Go

One little proprietary application. Nothing fancy.

### What to expect

Random commits touching on random ideas or experiments I'm inspired to run. Might or might not evolve into something useful someday.

### Where to track these ideas

I'll do my best to list them here.

### Ideas touched upon so far

- [Grasp the language basics again](https://github.com/mk360/go-sandbox/commit/b8973b9ca4050a7d9b4e6ead2f381b6b2d8aab0f)
- [Split the main module into packages](https://github.com/mk360/go-sandbox/commit/b8973b9ca4050a7d9b4e6ead2f381b6b2d8aab0f)
- [Public and private functions](https://github.com/mk360/go-sandbox/commit/b8973b9ca4050a7d9b4e6ead2f381b6b2d8aab0f)
- [How to express a function's return type, declaring vs initializing a variable, remembering that we can return more than one value](https://github.com/mk360/go-sandbox/commit/b8973b9ca4050a7d9b4e6ead2f381b6b2d8aab0f) (and usually it's an error)
- [First HTTP GET request, first struct](https://github.com/mk360/go-sandbox/commit/b8973b9ca4050a7d9b4e6ead2f381b6b2d8aab0f)
- `p o i n t e r s`
- What happens if I dereference a referenced pointer? (like `int something = *&otherthing`)
- What happens if I initialize a struct through a pointer to another struct, then mess with both of them separately?
- [What happens if I fetch a floating-point value from an API and assign it to an int?](https://github.com/mk360/go-sandbox/commit/0df478a00acbe3cbfc5305c7ebb81912f9f8fd0a) (like, seriously, any API that returns a float) (damnit the program actually crashed))
- Or what if I fetched an int and assigned it to a float?
- Can I dynamically call functions simply by manipulating a pointer?
- [Can I create a stack overflow, and how long will it take me to?](https://github.com/mk360/go-sandbox/commit/a87416518067ace10634c5aeb6337618105636d7)
- [Can I bypass int8 size limitations if, instead of increasing a variable's value directly, I increased a pointer's referenced variable's value?](https://github.com/mk360/go-sandbox/commit/7309301d64223aad79751f8e076db3dc92efc384)