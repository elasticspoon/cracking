import { findIntersection, kthValue } from "./chapter-2";
import { LLNode } from "./chapter-2";

test("findIntersection: null", () => {
  let a = buildList([1, 2, 3, 4, 5]);
  let b = buildList([6, 7, 8, 9, 10]);
  let res = findIntersection(a, b);
  expect(res).toBeNull();
});

test("findIntersection at 3", () => {
  let a = buildList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
  let b = buildList([11, 12, 13, 14, 15, 16, 17, 18, 19, 20]);
  b.next!.next!.next!.next!.next = a.next!.next;
  let res = findIntersection(a, b);
  expect(res!.value).toBe(3)
});


test("kth Value: middle", () => {
  let list = buildList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
  let res = kthValue(list, 4);
  expect(res).toBe(6);
});
test("kth Value: last", () => {
  let list = buildList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
  let res = kthValue(list, 9);
  expect(res).toBe(1);
});
test("kth Value: first", () => {
  let list = buildList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
  let res = kthValue(list, 0);
  expect(res).toBe(10);
});


function buildList(arr: number[]): LLNode<number> {
  let head = new LLNode(arr[0]);
  let current = head;
  for (let i = 1; i < arr.length; i++) {
    current.next = new LLNode(arr[i]);
    current = current.next;
  }
  return head;
}

function printList<T>(head: LLNode<T>): string {
  let current = head;
  let out = "";
  while (current) {
    out += `${current.value} -> `;
    current = current.next!;
  }
  out += "null";
  return out;
}

