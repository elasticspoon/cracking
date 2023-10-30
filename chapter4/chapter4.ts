export class TreeNode {
  public value: number;
  public height: number;
  public left: TreeNode | null
  public right: TreeNode | null
  public parent: TreeNode | null
  constructor(value: number) {
    this.value = value;
    this.left = null;
    this.right = null;
    this.height = 1;
    this.parent = null;
  }

  public insert(value: number): void {
    if (value < this.value) {
      if (this.left) {
        this.left.insert(value);
      } else {
        this.left = new TreeNode(value);
      }
      this.height = Math.max(this.left.height, this.right ? this.right.height : 0) + 1;
    } else if (value > this.value) {
      if (this.right) {
        this.right.insert(value);
      } else {
        this.right = new TreeNode(value);
      }
      this.height = Math.max(this.right.height, this.left ? this.left.height : 0) + 1;
    }

  }


}

export function InsertIntoTree(array: number[]): TreeNode | null {
  if (array.length === 0) { return null }

  let queue: [number, number][] = [[0, array.length - 1]];
  let tree: TreeNode | null = null;

  while (queue.length > 0) {
    let [left, right]: [number, number] = queue.shift()!;
    let mid = (left + right) >> 1;

    if (tree === null) {
      tree = new TreeNode(array[mid]);
    } else {
      tree.insert(array[mid]);
    }

    if (left <= mid - 1) {
      queue.push([left, mid - 1]);
    }
    if (mid + 1 <= right) {
      queue.push([mid + 1, right]);
    }
  }

  return tree
}

export type Graph = { [key: number]: number[] };

export function hasPath(graph: Graph, start: number, end: number): boolean {
  const queueS: number[] = [start];
  const queueE: number[] = [end];
  const visitedS = new Set();
  const visitedE = new Set();

  while (queueS.length > 0 || queueE.length > 0) {
    if (queueS.length > 0) {
      let node: number = queueS.shift()!;
      visitedS.add(node);
      if (visitedE.has(node)) return true;
      for (let child of graph[node]) {
        if (visitedS.has(child)) continue;
        queueS.push(child);
      }
    }
    if (queueE.length > 0) {
      let node: number = queueE.shift()!;
      visitedE.add(node);
      if (visitedS.has(node)) return true;
      for (let child of graph[node]) {
        if (visitedE.has(child)) continue;
        queueE.push(child);
      }
    }
  }
  return false
}

export class Node {
  public value: any;
  public next: Node | null;
  constructor(value: any) {
    this.value = value;
    this.next = null;
  }
}

export function treeToLists(tree: TreeNode): Node[] {
  let nodes: Node[] = [];
  let queue: [TreeNode, number][] = [[tree, 1]];

  while (queue.length > 0) {
    let [tNode, level]: [TreeNode, number] = queue.shift()!;
    let node = new Node(tNode.value);

    if (nodes[level - 1]) {
      node.next = nodes[level - 1];
    }
    nodes[level - 1] = node;

    if (tNode.left) {
      queue.push([tNode.left, level + 1]);
    }
    if (tNode.right) {
      queue.push([tNode.right, level + 1]);
    }
  }
  return nodes
}

export function isBalanced(tree: TreeNode | null): boolean {
  if (tree === null) { return false }
  let depthSizes: number[] = [];
  let queue: [TreeNode, number][] = [[tree, 0]];

  while (queue.length > 0) {
    let [tNode, level]: [TreeNode, number] = queue.shift()!;
    if (depthSizes[level - 1] && depthSizes[level - 1] < 2 ** (level - 1)) {
      return false
    }

    depthSizes[level] = (depthSizes[level] ?? 0) + 1;
    if (tNode.left) {
      queue.push([tNode.left, level + 1]);
    }
    if (tNode.right) {
      queue.push([tNode.right, level + 1]);
    }
  }

  return true
}

function checkHeight(tree: TreeNode | null): number {
  if (tree === null) { return -1 }

  let leftHeight = checkHeight(tree.left);
  if (leftHeight === -Infinity) { return -Infinity }

  let rightHeight = checkHeight(tree.right);
  if (rightHeight === -Infinity) { return -Infinity }

  let heightDiff = Math.abs(leftHeight - rightHeight);
  if (heightDiff > 1) {
    return -Infinity
  } else {
    return Math.max(leftHeight, rightHeight) + 1;
  }
}

export function isBalancedB(tree: TreeNode | null): boolean {
  return checkHeight(tree) !== -Infinity
}

export function isBST(tree: TreeNode | null): boolean {
  if (tree === null) { return false }

  let stack: (TreeNode | number)[] = [tree];
  let result: number[] = [];

  while (stack.length > 0) {
    let value = stack.pop()!;
    if (typeof value === "number") {
      result.push(value);
    } else {
      if (value.right) {
        stack.push(value.right);
      }
      stack.push(value.value);
      if (value.left) {
        stack.push(value.left);
      }
    }

    for (let i = 1; i < result.length; i++) {
      if (result[i] < result[i - 1]) {
        return false
      }
    }
  }
  return true
}

export function inorderSuccessor(node: TreeNode): TreeNode | null {
  if (node.right) { return subtreeMin(node.right) }
  if (node.parent && isLeftChild(node, node.parent)) { return node.parent }

  while (node.parent) {
    if (isLeftChild(node, node.parent)) {
      return node.parent
    }
    node = node.parent;
  }


  return null
}


function subtreeMin(tree: TreeNode): TreeNode {
  while (tree.left) {
    tree = tree.left;
  }
  return tree
}

function isLeftChild(node: TreeNode, parent: TreeNode): boolean {
  return node === parent.left
}

export function commonA3(root: TreeNode, nodeA: TreeNode, nodeB: TreeNode): TreeNode | null {
  // check that both nodes are in the tree
  if (!contains(root, nodeA) || !contains(root, nodeB)) { return null }
  return recCommonA(root, nodeA, nodeB)
}

// if subtreeBeingChecked includes a and not b return a
// if subtreeBeingChecked includes b and not a return b
// if subtreeBeingChecked has neither a nor b return null
// else return commonA of a and b
function recCommonA(subtreeBeingChecked: TreeNode | null, nodeA: TreeNode, nodeB: TreeNode): TreeNode | null {
  // subtree had neither a or b. propogate it up
  if (subtreeBeingChecked === null) { return null }
  // subtree was the commonA. propagate it up
  if (subtreeBeingChecked === nodeA || subtreeBeingChecked === nodeB) { return subtreeBeingChecked }

  // checking left subtree for a and b
  let left = recCommonA(subtreeBeingChecked.left, nodeA, nodeB);
  // if it returned a value and that value is not null, a or b then that value is commonA. propogate it up
  if (left && left !== nodeA && left !== nodeB) { return left }


  // checking right subtree for a and b
  let right = recCommonA(subtreeBeingChecked.right, nodeA, nodeB);
  // if it returned a value and that value is not null, a or b then that value is commonA. propogate it up
  if (right && right !== nodeA && right !== nodeB) { return right }

  // Base case

  // both left and right returned values that are not null, a or b
  // this means current root is common anc
  if (left && right) {
    return subtreeBeingChecked
  } else if (subtreeBeingChecked === nodeA || subtreeBeingChecked === nodeB) {
    // if a or b is the root then it is the commonA, return it
    return subtreeBeingChecked
  } else if (left || right) {
    // either left or right was null
    // return the non null value
    return left ? left : right
  } else {
    return null
  }

}

function contains(tree: TreeNode | null, node: TreeNode): boolean {
  if (tree === null) { return false }
  if (tree === node) { return true }
  return contains(tree.left, node) || contains(tree.right, node)
}


export function commonA2(treeA: TreeNode, treeB: TreeNode): TreeNode | null {
  let [currA, currB]: [TreeNode | null, TreeNode | null] = [treeA, treeB]
  let [depthA, depthB]: [number, number] = [0, 0];

  while (currA) {
    currA = currA.parent;
    depthA++
  }
  while (currB) {
    currB = currB.parent;
    depthB++
  }

  let diff = Math.abs(depthA - depthB);

  if (depthA > depthB) {
    [currA, currB] = [treeA, treeB]
  } else {
    [currA, currB] = [treeB, treeA]
  }

  for (let i = 0; i < diff; i++) {
    currA = currA!.parent;
  }

  while (currA !== currB) {
    currA = currA!.parent;
    currB = currB!.parent;
    depthA--
  }


  return currB
}
export function commonA(treeA: TreeNode, treeB: TreeNode): TreeNode | null {
  let [pathA, pathB]: [string, string] = ["", ""]
  let [currA, currB]: [TreeNode, TreeNode] = [treeA, treeB]

  while (currA.parent) {
    if (isLeftChild(currA, currA.parent)) {
      pathA += "L"
    } else {
      pathA += "R"
    }
    currA = currA.parent;
  }

  while (currB.parent) {
    if (isLeftChild(currB, currB.parent)) {
      pathB += "L"
    } else {
      pathB += "R"
    }
    currB = currB.parent;
  }

  if (currA !== currB) { return null }

  let i = 1;
  while (pathA.at(-i) && pathB.at(-i) && pathA.at(-i) === pathB.at(-i)) {
    i++
  }

  let common = pathA.slice(pathA.length - (i - 1))

  for (let j = 0; j < common.length; j++) {
    if (common[j] === "L") {
      currA = currA.left!;
    } else {
      currA = currA.right!;
    }
  }
  return currA
}


export function possibleBSTInsertOrder(tree: TreeNode): number[][] {
  let queue: [number, TreeNode][] = [[0, tree]];
  let levels: number[][] = [];
  let result: number[][] = [];

  while (queue.length > 0) {
    let [level, tNode]: [number, TreeNode] = queue.shift()!;
    levels[level] = (levels[level] ?? []).concat(tNode.value);
    if (tNode.left) {
      queue.push([level + 1, tNode.left]);
    }
    if (tNode.right) {
      queue.push([level + 1, tNode.right]);
    }
  }

  let totalPoss = levels.reduce((acc, curr) => acc * factorial(curr.length), 1);

  for (let i = 0; i < totalPoss; i++) {
    let divisor = totalPoss;
    let temp: number[] = [];
    for (let arr of levels) {
      divisor = divisor / factorial(arr.length);
      let index = Math.floor(i / divisor) % arr.length;
      for (let j = 0; j < arr.length; j++) {
        temp.push(arr[(j + index) % arr.length]);
      }
    }
    result.push(temp);
  }


  return result
}

function factorial(n: number): number {
  if (n === 0) { return 1 }
  return n * factorial(n - 1)
}

