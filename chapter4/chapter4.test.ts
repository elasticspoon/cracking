import { TreeNode, InsertIntoTree, isSubtreeByVal, isSubtreeByRef, inorderSuccessor, isBST, possibleBSTInsertOrder, Graph, Node, hasPath, treeToLists, commonA, commonA2, commonA3, isBalanced, isBalancedB } from "./chapter4";

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

describe("commonA3", () => {
  let tree = bstTree();

  test("commonA: 3, node not tree-> null", () => {
    let node1 = tree.left!.left!;
    let node2 = new TreeNode(11);

    expect(commonA3(tree, node1, node2)).toBeNull();
  });
  test("commonA: 3, 14 -> 5", () => {
    let node1 = tree.left!.left!;
    let node2 = tree.left!.right!.right!;

    expect(commonA3(tree, node1, node2)!.value).toBe(5);
  });

  test("commonA: 1, 25 -> 20", () => {
    let node1 = tree.left!.left!.left!;
    let node2 = tree.right!;

    expect(commonA3(tree, node1, node2)!.value).toBe(20)
  });

  test("commonA: 14, 20 -> 20", () => {
    let node1 = tree.left!.right!.right!;
    let node2 = tree;

    expect(commonA3(tree, node1, node2)!.value).toBe(20)
  });
});
describe("commonA2", () => {
  let tree = bstTree();

  test("commonA: 3, node not tree-> null", () => {
    let node1 = tree.left!.left!;
    let node2 = new TreeNode(11);

    expect(commonA2(node1, node2)).toBeNull();
  });
  test("commonA: 3, 14 -> 5", () => {
    let node1 = tree.left!.left!;
    let node2 = tree.left!.right!.right!;

    expect(commonA2(node1, node2)!.value).toBe(5);
  });

  test("commonA: 1, 25 -> 20", () => {
    let node1 = tree.left!.left!.left!;
    let node2 = tree.right!;

    expect(commonA2(node1, node2)!.value).toBe(20)
  });

  test("commonA: 14, 20 -> 20", () => {
    let node1 = tree.left!.right!.right!;
    let node2 = tree;

    expect(commonA2(node1, node2)!.value).toBe(20)
  });
});
describe("commonA", () => {
  let tree = bstTree();

  test("commonA: 3, node not tree-> null", () => {
    let node1 = tree.left!.left!;
    let node2 = new TreeNode(11);

    expect(commonA(node1, node2)).toBeNull();
  });

  test("commonA: 3, 14 -> 5", () => {
    let node1 = tree.left!.left!;
    let node2 = tree.left!.right!.right!;

    expect(commonA(node1, node2)!.value).toBe(5);
  });

  test("commonA: 1, 25 -> 20", () => {
    let node1 = tree.left!.left!.left!;
    let node2 = tree.right!;

    expect(commonA(node1, node2)!.value).toBe(20)
  });

  test("commonA: 14, 20 -> 20", () => {
    let node1 = tree.left!.right!.right!;
    let node2 = tree;

    expect(commonA(node1, node2)!.value).toBe(20)
  });
});

describe("possible BST insert order", () => {
  test("5 node tree: 3, 2, 5, 1,8", () => {
    let tree = new TreeNode(3);

    tree.left = new TreeNode(2);
    tree.right = new TreeNode(5);

    tree.left.left = new TreeNode(1);
    tree.right.right = new TreeNode(8);

    let expected = [[3, 2, 5, 1, 8], [3, 2, 5, 8, 1], [3, 5, 2, 1, 8], [3, 5, 2, 8, 1]]
    let result = possibleBSTInsertOrder(tree);

    expect(result).toEqual(expected);
  });
  test("3 node tree: 3, 2, 5", () => {
    let tree = new TreeNode(3);

    tree.left = new TreeNode(2);
    tree.right = new TreeNode(5);

    let expected = [[3, 2, 5], [3, 5, 2]]

    expect(possibleBSTInsertOrder(tree)).toEqual(expected);

  });
});

describe("isSubtreeByValue", () => {
  let tree = bstTree()
  let otherTree = bstTree();

  test("returns true when given the same tree", () => {
    let result = isSubtreeByVal(tree, otherTree)
    expect(result).toBe(true)
  })

  test("returns true when given a subtree", () => {
    let subtree = otherTree.left!.right!
    let result = isSubtreeByVal(tree, subtree)
    expect(result).toBe(true)
  });

  test("returns false when given a node not in tree", () => {
    let otherTree = new TreeNode(99);

    let result = isSubtreeByRef(tree, otherTree);
    expect(result).toBe(false)
  });
});
describe("isSubtreeByRef", () => {
  let tree = bstTree()
  let otherTree = bstTree();

  test("returns true when given the same tree", () => {
    let result = isSubtreeByRef(tree, tree)
    expect(result).toBe(true)
  })

  test("returns true when given a subtree", () => {
    let subtree = tree.left!.right!
    let result = isSubtreeByRef(tree, subtree)
    expect(result).toBe(true)
  });

  test("returns false when given different node by ref", () => {
    let subtree = otherTree.left!.right!
    let result = isSubtreeByRef(tree, subtree)

    expect(result).toBe(false)
  });
  test("returns false when given a node not in tree", () => {
    let subtree = new TreeNode(11)
    subtree.left = new TreeNode(12)

    let result = isSubtreeByRef(tree, subtree)
    expect(result).toBe(false)
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


  tree.left.left.left = new TreeNode(1);
  tree.left.left.left.parent = tree.left.left;
  tree.left.left.right = new TreeNode(4);
  tree.left.left.right.parent = tree.left.left;

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

