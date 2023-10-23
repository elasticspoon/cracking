export class TreeNode {
  public value: number;
  public height: number;
  public left: TreeNode | null
  public right: TreeNode | null
  constructor(value: number) {
    this.value = value;
    this.left = null;
    this.right = null;
    this.height = 1;
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
