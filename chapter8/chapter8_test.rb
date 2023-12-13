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
