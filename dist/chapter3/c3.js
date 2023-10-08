"use strict";
class TripleStack {
    array = [];
    lenA = 0;
    lenB = 0;
    lenC = 0;
    pushA(value) {
        this.array.unshift(value);
        this.lenA++;
    }
    pushB(value) {
        this.array.splice(this.lenA, 0, value);
        this.lenB++;
    }
    pushC(value) {
        this.array.push(value);
        this.lenC++;
    }
    popA() {
        if (this.lenA === 0) {
            return undefined;
        }
        this.lenA--;
        return this.array.shift();
    }
    popB() {
        if (this.lenB === 0) {
            return undefined;
        }
        this.lenB--;
        return this.array.splice(this.lenA, 1)[0];
    }
    popC() {
        if (this.lenC === 0) {
            return undefined;
        }
        this.lenC--;
        return this.array.pop();
    }
}
function testTripleStack() {
    console.log("testTripleStack");
    const ts = new TripleStack();
    for (let i = 0; i < 3; i++) {
        ts.pushA(i);
        ts.pushB(i);
        ts.pushC(i);
    }
    for (let i = 2; i >= 0; i--) {
        console.log(ts.popA() == i);
    }
    for (let i = 2; i >= 0; i--) {
        console.log(ts.popC() == i);
    }
    for (let i = 2; i >= 0; i--) {
        console.log(ts.popB() == i);
    }
}
testTripleStack();
