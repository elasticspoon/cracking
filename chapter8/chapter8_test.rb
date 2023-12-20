require "minitest/autorun"
require_relative "chapter8"

class TestChapter8 < Minitest::Test
  Maze = Struct.new(:rows, :cols, :walls) do
    def invalid?(row, col)
      walls.include?([row, col])
    end
  end

  def test_find_path_no_obstacles
    maze = Maze.new(3, 3, {})
    path = find_path(maze)
    expected_path = [[0, 0], [1, 0], [2, 0], [2, 1], [2, 2]]

    assert_equal expected_path, path
  end

  def test_find_path
    maze = Maze.new(3, 3, {[0, 2] => true, [2, 1] => true})
    path = find_path(maze)
    expected_path = [[0, 0], [1, 0], [1, 1], [1, 2], [2, 2]]

    assert_equal expected_path, path
  end

  def test_find_path_many_obstacles
    maze = Maze.new(4, 4, {[1, 1] => true, [1, 2] => true, [2, 1] => true, [2, 2] => true, [3, 2] => true})
    path = find_path(maze)
    expected_path = [[0, 0], [0, 1], [0, 2], [0, 3], [1, 3], [2, 3], [3, 3]]

    assert_equal expected_path, path
  end

  def test_subsets_1
    array = [1, 2, 3]
    expected = [[1], [2], [3], [1, 2], [1, 3], [2, 3], [1, 2, 3]].sort
    got = subsets(array).sort

    assert_equal expected, got
  end

  def test_subsets_2
    array = [1, 2]
    expected = [[1], [2], [1, 2]].sort
    got = subsets(array).sort

    assert_equal expected, got
  end

  def test_subsets_cheat
    array = [1, 2, 3]
    expected = [[1], [2], [3], [1, 2], [1, 3], [2, 3], [1, 2, 3]].sort
    got = other_subsets(array).sort

    assert_equal expected, got
  end
end

class TestTowerOfHanoi < Minitest::Test
  def test_tower_size_1
    tower = TowerOfHanoi.new(1)
    expected = [[], [], [1]]

    move_tower(tower)

    got = tower.towers
    assert_equal expected, got
  end

  def test_tower_size_2
    tower = TowerOfHanoi.new(2)
    expected = [[], [], [2, 1]]

    move_tower(tower)

    got = tower.towers
    assert_equal expected, got
  end

  def test_tower_size_10
    tower = TowerOfHanoi.new(10)
    expected = [[], [], [10, 9, 8, 7, 6, 5, 4, 3, 2, 1]]

    move_tower(tower)

    got = tower.towers
    assert_equal expected, got
  end
end

class TestPermutations < Minitest::Test
  def test_permutations_1_length
    string = "a"
    expected = ["a"]

    got = perms(string)

    assert_equal expected, got
  end

  def test_permutations_2_length
    string = "ab"
    expected = ["ab", "ba"].sort

    got = perms(string).sort

    assert_equal expected, got
  end

  def test_permutations_3_length
    string = "abc"
    expected = ["abc", "acb", "bac", "bca", "cab", "cba"].sort

    got = perms(string).sort

    assert_equal expected, got
  end

  def test_permutations_with_duplicates
    string = "abba"
    expected = ["abba", "abab", "aabb", "baba", "baab", "bbaa"].sort

    got = perms(string).sort

    assert_equal expected, got
  end
end

class TestParenCombinations < Minitest::Test
  def test_length_1
    expected = ["()"]
    got = parens(1)
    assert_equal expected, got
  end

  def test_length_2
    expected = ["()()", "(())"].sort
    got = parens(2).sort
    assert_equal expected, got
  end

  def test_length_3
    expected = ["()()()", "()(())", "(())()", "(()())", "((()))"].sort
    got = parens(3).sort
    assert_equal expected, got
  end
end
