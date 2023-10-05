class LLNode<T> {
  value: T;
  next: LLNode<T> | null;

  constructor(value: T) {
    this.value = value;
    this.next = null;
  }
}

function removeDupsNested(head: LLNode<number>): LLNode<number> {
  let val = head;

  while (val) {
    let current: LLNode<number> | null = val.next;
    let previous = val;
    while (current) {
      if (val.value === current.value) {
        previous.next = current.next;
      } else {
        previous = current;
      }
      current = current.next;
    }
    val = val.next!;
  }

  return head;
}
function removeDupsBuffer<T>(head: LLNode<T>): LLNode<T> {
  const nodeSet = new Set<T>();
  let current: LLNode<T> | null = head;
  let previous = head;

  while (current) {
    if (nodeSet.has(current.value)) {
      previous.next = current.next;
    } else {
      nodeSet.add(current.value);
      previous = current;
    }
    current = current.next;
  }

  return head;
}

function kthValue<T>(head: LLNode<T>, k: number): T {
  let len = 0;
  let current = head;
  while (current) {
    len++;
    current = current.next!;
  }

  let target = len - k;
  current = head;
  let index = 1;

  while (index != target) {
    current = current.next!;
    index++;
  }

  return current.value;
}

function deleteMiddleNode<T>(node: LLNode<T>): LLNode<T> | null {
  let previous = node;
  let current = node.next;

  while (current) {
    previous.value = current.value;
    if (!current.next) {
      previous.next = null;
      return node;
    }
    previous = current;
    current = current.next;
  }
  return null;
}

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

function testDeleteMiddleNode(): void {
  console.log("Testing deleteMiddleNode");
  let list = buildList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
  let middle = list.next!.next!;
  let res = printList(deleteMiddleNode(middle)!);
  console.log(res == "4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> null");

  list = buildList([1, 2]);
  middle = list;
  res = printList(deleteMiddleNode(middle)!);
  console.log(res == "2 -> null");
}

function testRemoveDups(): void {
  console.log("Testing removeDups");
  let list = buildList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
  let res = removeDupsBuffer(list);
  console.log(
    printList(res) == "1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> null"
  );
  res = removeDupsNested(list);
  console.log(
    printList(res) == "1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> null"
  );

  list = buildList([1, 1, 1, 1, 1, 1, 1, 1, 1, 1]);
  res = removeDupsBuffer(list);
  console.log(printList(res) == "1 -> null");
  res = removeDupsNested(list);
  console.log(printList(res) == "1 -> null");

  list = buildList([1]);
  res = removeDupsBuffer(list);
  console.log(printList(res) == "1 -> null");
  res = removeDupsNested(list);
  console.log(printList(res) == "1 -> null");
}

function testKthValue(): void {
  console.log("Testing kthValue");
  let list = buildList([1, 2, 3, 4, 5, 6, 7, 8, 9, 10]);
  let res = kthValue(list, 0);
  console.log(res == 10);
  res = kthValue(list, 4);
  console.log(res == 6);
  res = kthValue(list, 9);
  console.log(res == 1);
}

testRemoveDups();
testKthValue();
testDeleteMiddleNode();
