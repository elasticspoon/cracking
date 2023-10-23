import { InsertIntoTree, Graph, hasPath } from "./chapter4";

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

