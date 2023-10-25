import { TreeNode, InsertIntoTree, inorderSuccessor, isBST, Graph, Node, hasPath, treeToLists, isBalanced, isBalancedB } from "./chapter4";

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

test("isBST: valid balanced", () => {
  let tree = new TreeNode(5);

  tree.left = new TreeNode(2);
  tree.right = new TreeNode(7);

  tree.left.left = new TreeNode(1);
  tree.left.right = new TreeNode(4);

  tree.right.left = new TreeNode(6);
  tree.right.right = new TreeNode(8);


  expect(isBST(tree)).toBe(true);
});

test("isBST: valid unbalanced", () => {
  let tree = new TreeNode(5);
  tree.left = new TreeNode(4);
  tree.left.left = new TreeNode(3);
  tree.left.left.left = new TreeNode(2);
  tree.left.left.left.left = new TreeNode(1);

  expect(isBST(tree)).toBe(true);
});

test("isBST: invalid", () => {
  let tree = new TreeNode(5);
  tree.left = new TreeNode(2);
  tree.right = new TreeNode(7);

  tree.left.left = new TreeNode(1);
  tree.left.right = new TreeNode(6);

  tree.right.left = new TreeNode(6);
  tree.right.right = new TreeNode(8);


  expect(isBST(tree)).toBe(false);
});

test("isBST: invalid tree", () => {
  let tree = null;
  expect(isBST(tree)).toBe(false);
});

describe("inorderSuccessor", () => {
  let tree = bstTree();

  test("inorderSuccessor: 25-> null", () => {
    let node = tree.right!;

    expect(inorderSuccessor(node)).toBe(null);
  });

  test("inorderSuccessor: 20-> 25", () => {
    expect(inorderSuccessor(tree)!.value).toBe(25);
  });

  test("inorderSuccessor: 14-> 20", () => {
    let node = tree.left!.right!.right!

    expect(inorderSuccessor(node)!.value).toBe(20);
  });
});




function bstTree(): TreeNode {
  let tree = new TreeNode(20);

  tree.left = new TreeNode(5);
  tree.left.parent = tree;
  tree.right = new TreeNode(25);
  tree.right.parent = tree;

  tree.left.left = new TreeNode(3);
  tree.left.left.parent = tree.left;
  tree.left.right = new TreeNode(12);
  tree.left.right.parent = tree.left;


  tree.left.right.left = new TreeNode(1);
  tree.left.right.left.parent = tree.left.right;
  tree.left.right.right = new TreeNode(4);
  tree.left.right.right.parent = tree.left.right;

  tree.left.right.left = new TreeNode(6);
  tree.left.right.left.parent = tree.left.right;
  tree.left.right.right = new TreeNode(14);
  tree.left.right.right.parent = tree.left.right;
  return tree
}

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

