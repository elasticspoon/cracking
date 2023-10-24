import { TreeNode, InsertIntoTree, Graph, Node, hasPath, treeToLists, isBalanced, isBalancedB } from "./chapter4";

test("InsertIntoTree: height 4", () => {
  let array = [1, 2, 3, 4, 5, 6, 7, 8, 9];
  let tree = InsertIntoTree(array);

  expect(tree!.height).toBe(4);
});

test("InsertIntoTree: height 3", () => {
  let array = [1, 2, 3, 4, 5, 6, 7];
  let tree = InsertIntoTree(array);

  expect(tree!.height).toBe(3);
});
test("InsertIntoTree: height 2", () => {
  let array = [5, 6, 7];
  let tree = InsertIntoTree(array);

  expect(tree!.height).toBe(2);
});
test("InsertIntoTree: height 1", () => {
  let array = [9];
  let tree = InsertIntoTree(array);

  expect(tree!.height).toBe(1);
});
test("InsertIntoTree: empty array", () => {
  let array: number[] = [];
  let tree = InsertIntoTree(array);

  expect(tree).toBe(null);
});

let graph: Graph = { 0: [1, 2, 8], 1: [7, 8, 10], 2: [10], 10: [], 14: [], 8: [], 7: [14, 4], 4: [8] }

test("hasPath: 0 -> 10", () => {
  expect(hasPath(graph, 0, 10)).toBe(true)
});
test("hasPath: 10 -> 0", () => {
  expect(hasPath(graph, 10, 1)).toBe(true)
});
test("hasPath: 10 -> 4", () => {
  expect(hasPath(graph, 10, 4)).toBe(false)
});

test("treeToLists: unbalanced", () => {
  let tree = buildTree([1, 2, 3, 4, 5])!;
  let expected = [[1], [2], [3], [4], [5]];

  for (let list of treeToLists(tree)) {
    expect(nodeListToArray(list)).toEqual(expected.shift());
  }
});

test("treeToLists: balanced", () => {
  let tree = buildTree([5, 2, 7, 1, 3, 6, 8])!;
  let expected = [[5], [7, 2], [8, 6, 3, 1]];

  for (let list of treeToLists(tree)) {
    expect(nodeListToArray(list)).toEqual(expected.shift());
  }
});
test("treeToLists: unbalanced", () => {
  let tree = buildTree([1, 2, 3, 4, 5])!;
  let expected = [[1], [2], [3], [4], [5]];

  for (let list of treeToLists(tree)) {
    expect(nodeListToArray(list)).toEqual(expected.shift());
  }
});

test("treeToLists: balanced", () => {
  let tree = buildTree([5, 2, 7, 1, 3, 6, 8])!;
  let expected = [[5], [7, 2], [8, 6, 3, 1]];

  for (let list of treeToLists(tree)) {
    expect(nodeListToArray(list)).toEqual(expected.shift());
  }
});

test("isBalanced: balanced", () => {
  let tree = buildTree([5, 2, 7, 1, 3, 6, 8])!;
  expect(isBalanced(tree)).toBe(true);
});

test("isBalanced: balanced barely", () => {
  let tree = buildTree([1, 2])!;
  expect(isBalanced(tree)).toBe(true);
});
test("isBalanced: unbalanced", () => {
  let tree = buildTree([1, 2, 3, 4, 5])!;
  expect(isBalanced(tree)).toBe(false);
});

test("isBalancedBookV: balanced", () => {
  let tree = buildTree([5, 2, 7, 1, 3, 6, 8])!;
  expect(isBalancedB(tree)).toBe(true);
});

test("isBalancedBookV: balanced barely", () => {
  let tree = buildTree([1, 2])!;
  expect(isBalancedB(tree)).toBe(true);
});
test("isBalancedBookV: unbalanced", () => {
  let tree = buildTree([1, 2, 3, 4, 5])!;
  expect(isBalancedB(tree)).toBe(false);
});

function buildTree(array: number[]): TreeNode | null {
  let tree: TreeNode | null = null;
  for (let value of array) {
    if (tree === null) {
      tree = new TreeNode(value);
    } else {
      tree.insert(value);
    }
  }
  return tree
}

function nodeListToArray(node: Node | null): any[] {
  let array: any[] = [];
  while (node !== null) {
    array.push(node.value);
    node = node.next;
  }
  return array;
}
