require 'set'

def analyze(input=[])
  return :stable if input.length < 2

  seen = Set.new
  last = input[0]

  (1..input.length-1).each do |i|
    return :stable if input[i] == last

    if input[i] > last
      seen.add(:increasing)
    else
      seen.add(:decreasing)
    end

    last = input[i]
  end

  return :stable if seen.include?(:increasing) && seen.include?(:decreasing)

  return :increasing if seen.include?(:increasing)

  :decreasing
end

def print(name, input)
  puts "name:\t#{name}"
  puts "input:\t#{input.join(' ')}"
  puts "output:\t#{analyze(input)}"
  puts '=' * 80
end

print('Increasing', [1, 2, 3, 4, 5])
print('Decreasing', [5, 4, 3, 2, 1, 0, -1, -2, -3])
print('Stable', [1, 2, 3, 2, 1, 4, 5])
print('Empty', [])
print('Single', [1])
print('EqualWhenIncreasing', [1, 2, 2])
print('EqualWhenDecreasing', [5, 4, 4])
